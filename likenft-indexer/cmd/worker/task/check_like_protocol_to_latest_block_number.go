package task

import (
	"context"
	"encoding/json"
	"fmt"
	"os/signal"
	"syscall"

	appcontext "likenft-indexer/cmd/worker/context"
	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/logic/contractevmeventacquirer"

	"github.com/hibiken/asynq"
)

const TypeCheckLikeProtocolToLatestBlockNumberPayload = "check-like-protocol-to-latest-block-number"

type CheckLikeProtocolToLatestBlockNumberPayload struct {
	ContractAddress string
}

func NewCheckLikeProtocolToLatestBlockNumberTask(contractAddress string) (*asynq.Task, error) {
	payload, err := json.Marshal(CheckLikeProtocolToLatestBlockNumberPayload{
		ContractAddress: contractAddress,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCheckLikeProtocolToLatestBlockNumberPayload, payload), nil
}

func HandleCheckLikeProtocolToLatestBlockNumber(ctx context.Context, t *asynq.Task) error {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := appcontext.LoggerFromContext(ctx)
	envCfg := appcontext.ConfigFromContext(ctx)
	asynqClient := appcontext.AsynqClientFromContext(ctx)
	evmQueryClient := appcontext.EvmQueryClientFromContext(ctx)
	evmClient := appcontext.EvmClientFromContext(ctx)

	dbService := database.New()
	likeProtocolRepository := database.MakeLikeProtocolRepository(dbService)
	evmEventRepository := database.MakeEVMEventRepository(dbService)

	mylogger := logger.WithGroup("HandleCheckLikeProtocolToLatestBlockNumber")

	var p CheckLikeProtocolToLatestBlockNumberPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		mylogger.Error("json.Unmarshal CheckLikeProtocolToLatestBlockNumberPayload", "err", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	likeProtocol, err := likeProtocolRepository.GetLikeProtocol(ctx, p.ContractAddress)
	if err != nil {
		return err
	}

	latestBlockNumber, err := evmClient.BlockNumber(ctx)
	if err != nil {
		return err
	}

	mylogger = mylogger.With("latestBlockNumber", latestBlockNumber)

	blockStarts := make([]uint64, 0)

	for i := uint64(likeProtocol.LatestEventBlockNumber); i < latestBlockNumber; i = i + envCfg.EvmEventQueryNumberOfBlocksLimit {
		blockStarts = append(blockStarts, i)
	}

	acquirer := contractevmeventacquirer.NewContractEvmEventsAcquirer(
		evmQueryClient,
		evmEventRepository,
		evmQueryClient,
		evmClient,
		contractevmeventacquirer.ContractEvmEventsAcquirerContractTypeLikeProtocol,
		[]string{p.ContractAddress},
	)

	errorChan := make(chan error, 1)
	doneChan := make(chan struct{}, 1)
	go func() {
		for i, blockStart := range blockStarts {
			select {
			case <-ctx.Done():
				return
			default:
			}
			mylogger := mylogger.With(
				"partition",
				fmt.Sprintf("%d/%d", i, len(blockStarts)),
			)
			newBlockNumber, err := acquirer.Acquire(
				ctx,
				mylogger,
				blockStart,
				envCfg.EvmEventQueryNumberOfBlocksLimit,
			)
			if err != nil {
				errorChan <- err
				return
			}
			err = likeProtocolRepository.CreateOrUpdateLatestEventBlockHeight(
				ctx,
				p.ContractAddress,
				typeutil.Uint64(newBlockNumber),
			)
			if err != nil {
				errorChan <- err
				return
			}
		}
		doneChan <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return nil
	case <-doneChan:
		task, err := NewCheckReceivedEVMEventsTask()
		if err != nil {
			return err
		}

		_, err = asynqClient.Enqueue(task, asynq.MaxRetry(0))

		return err
	case err := <-errorChan:
		return err
	}
}
