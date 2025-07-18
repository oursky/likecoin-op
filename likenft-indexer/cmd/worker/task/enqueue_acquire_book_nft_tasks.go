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

const TypeEnqueueAcquireBookNFTTasksPayload = "enqueue-acquire-book-nft-tasks"

type EnqueueAcquireBookNFTTasksPayload struct {
}

func NewEnqueueAcquireBookNFTTasksTask() (*asynq.Task, error) {
	payload, err := json.Marshal(EnqueueAcquireBookNFTTasksPayload{})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(
		TypeEnqueueAcquireBookNFTTasksPayload,
		payload,
		asynq.Queue(TypeEnqueueAcquireBookNFTTasksPayload),
	), nil
}

func HandleEnqueueAcquireBookNFTTasks(ctx context.Context, t *asynq.Task) error {
	logger := appcontext.LoggerFromContext(ctx)
	config := appcontext.ConfigFromContext(ctx)
	inspector := appcontext.AsynqInspectorFromContext(ctx)

	mylogger := logger.WithGroup("HandleEnqueueAcquireBookNFTTasks")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var p EnqueueAcquireBookNFTTasksPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal: %v", err)
	}

	queueInfo, err := inspector.GetQueueInfo(TypeAcquireBookNFTEventsTaskPayload)
	if err != nil {
		return fmt.Errorf("inspector.GetQueueInfo: %v", err)
	}
	currentLength := queueInfo.Pending + queueInfo.Active
	numberOfItemsToBeFetched := config.TaskAcquireBookNFTEventsMaxQueueLength - currentLength

	dbService := database.New()
	nftClassAcquireBookNFTEventsRepository := database.MakeNFTClassAcquireBookNFTEventsRepository(dbService)

	nftClasses, err := nftClassAcquireBookNFTEventsRepository.
		RequestForEnqueue(ctx, numberOfItemsToBeFetched)

	if err != nil {
		return fmt.Errorf("nftClassAcquireBookNFTEventsRepository.RequestForEnqueue: %v", err)
	}

	myLogger := logger.With("batchSize", len(nftClasses))
	mylogger.Info("nft classes found")

	lifecycleObjects := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycles(
		nftClassAcquireBookNFTEventsRepository,
		nftClasses,
		nftclassacquirebooknftevent.MakeCalculateScoreFn(
			config.TaskAcquireBookNFTEventsScoreBlockHeightContribution,
			config.TaskAcquireBookNFTEventsScoreWeight0Constant,
			config.TaskAcquireBookNFTEventsScoreWeight1Constant,
			config.TaskAcquireBookNFTEventsScoreWeightContribution,
		),
	)

	myLogger.Info("Enqueueing EnqueueAcquireBookNFTTask tasks...")
	for _, lifecycleObject := range lifecycleObjects {
		if _, err := lifecycleObject.WithEnqueueing(ctx, func(nftClass *ent.NFTClass) (*asynq.TaskInfo, error) {
			return EnqueueAcquireBookNFTTask(ctx, nftClass.Address)
		}); err != nil {
			myLogger.ErrorContext(ctx, "lifecycleObject.WithEnqueueing", "err", err)
		}
	}

	return nil
}

func EnqueueAcquireBookNFTTask(ctx context.Context, contractAddress string) (*asynq.TaskInfo, error) {
	asynqClient := appcontext.AsynqClientFromContext(ctx)

	t, err := NewEnqueueAcquireBookNFTTasksAcquireBookNFTTask(contractAddress)
	if err != nil {
		return nil, err
	}

	taskInfo, err := asynqClient.Enqueue(t, asynq.MaxRetry(0))
	if err != nil {
		return nil, err
	}
	return taskInfo, nil
}

func init() {
	t := Tasks.Register(task.DefineTask(
		TypeEnqueueAcquireBookNFTTasksPayload,
		HandleEnqueueAcquireBookNFTTasks,
	))
	PeriodicTasks.Register(task.DefinePeriodicTask(
		t,
		NewEnqueueAcquireBookNFTTasksTask,
	))
}
