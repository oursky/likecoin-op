package task

import (
	"context"
	"encoding/json"
	"fmt"

	appcontext "likenft-indexer/cmd/worker/context"
	"likenft-indexer/ent"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/logic/nftclassacquirebooknftevent"
	"likenft-indexer/internal/worker/task"

	"github.com/hibiken/asynq"
)

const TypeEnqueueAcquireBookNFTTasksAcquireBookNFTPayload = "enqueue-acquire-book-nft-tasks--acquire-book-nft"

type EnqueueAcquireBookNFTTasksAcquireBookNFTPayload struct {
	ContractAddress string
}

func NewEnqueueAcquireBookNFTTasksAcquireBookNFTTask(
	contractAddress string,
) (*asynq.Task, error) {
	payload, err := json.Marshal(EnqueueAcquireBookNFTTasksAcquireBookNFTPayload{
		contractAddress,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(
		TypeEnqueueAcquireBookNFTTasksAcquireBookNFTPayload,
		payload,
		asynq.Queue(TypeEnqueueAcquireBookNFTTasksAcquireBookNFTPayload),
	), nil
}

func HandleEnqueueAcquireBookNFTTasksAcquireBookNFTTask(ctx context.Context, t *asynq.Task) error {
	config := appcontext.ConfigFromContext(ctx)
	dbService := database.New()
	nftClassAcquireBookNFTEventsRepository := database.MakeNFTClassAcquireBookNFTEventsRepository(dbService)

	var p EnqueueAcquireBookNFTTasksAcquireBookNFTPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal: %v", err)
	}

	lifecycle, err := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycleFromAddress(
		ctx,
		nftClassAcquireBookNFTEventsRepository,
		p.ContractAddress,
		nftclassacquirebooknftevent.MakeCalculateScoreFn(
			config.TaskAcquireBookNFTEventsScoreBlockHeightContribution,
			config.TaskAcquireBookNFTEventsScoreWeight0Constant,
			config.TaskAcquireBookNFTEventsScoreWeight1Constant,
			config.TaskAcquireBookNFTEventsScoreWeightContribution,
		),
	)
	if err != nil {
		return fmt.Errorf("p.ToLifecycle: %v", err)
	}

	return lifecycle.WithEnqueued(ctx, func(nftClass *ent.NFTClass) error {
		return HandleAcquireBookNFTEvents(ctx, []string{nftClass.Address})
	})
}

func init() {
	Tasks.Register(task.DefineTask(
		TypeEnqueueAcquireBookNFTTasksAcquireBookNFTPayload,
		HandleEnqueueAcquireBookNFTTasksAcquireBookNFTTask,
	))
}
