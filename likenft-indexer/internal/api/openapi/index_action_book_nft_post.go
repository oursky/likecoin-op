package openapi

import (
	"context"
	"errors"
	"time"

	"likenft-indexer/cmd/worker/task"
	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"

	"github.com/hibiken/asynq"
)

func (h *OpenAPIHandler) IndexActionBookNftBooknftIDPost(
	ctx context.Context,
	params api.IndexActionBookNftBooknftIDPostParams,
) (*api.IndexActionBookNftBooknftIDPostOK, error) {
	task, err := task.NewCheckBookNFTToLatestBlockNumberTask(params.BooknftID)
	if err != nil {
		return nil, err
	}
	taskInfo, err := h.asynqClient.Enqueue(
		task,
		asynq.MaxRetry(0),
		asynq.Timeout(10*time.Minute),
		asynq.Unique(10*time.Minute),
	)
	if err != nil {
		if errors.Is(err, asynq.ErrDuplicateTask) {
			return &api.IndexActionBookNftBooknftIDPostOK{
				Message: "Already requested",
			}, nil
		}
		return nil, err
	}
	return &api.IndexActionBookNftBooknftIDPostOK{
		Message: "OK",
		TaskID:  model.MakeOptString(&taskInfo.ID),
	}, nil
}
