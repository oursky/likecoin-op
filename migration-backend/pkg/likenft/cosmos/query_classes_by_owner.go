package cosmos

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
)

type QueryNFTClassesByOwnerRequest struct {
	QueryNFTClassesByOwnerPageRequest
	ISCNOwner string `url:"iscn_owner"`
}

type QueryNFTClassesByOwnerPageRequest struct {
	Key        *int    `url:"pagination.key"`
	Offset     *uint64 `url:"pagination.offset"`
	Limit      *uint64 `url:"pagination.limit"`
	CountTotal *bool   `url:"pagination.countTotal"`
	Reverse    *bool   `url:"reverse"`
}

type QueryNFTClassesByOwnerResponse struct {
	Classes    []model.ClassListItem              `json:"classes,omitempty"`
	Pagination QueryNFTClassesByOwnerPageResponse `json:"pagination,omitempty"`
}

type QueryNFTClassesByOwnerPageResponse struct {
	NextKey int    `json:"next_key,omitempty"`
	Count   uint64 `json:"count,omitempty"`
}

func (c *LikeNFTCosmosClient) QueryNFTClassesByOwner(request QueryNFTClassesByOwnerRequest) (*QueryNFTClassesByOwnerResponse, error) {
	v, _ := query.Values(request)
	url := fmt.Sprintf("%s/likechain/likenft/v1/class?%v", c.NodeURL, v.Encode())
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		fmt.Printf("c.HTTPClient.Get error %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var res QueryNFTClassesByOwnerResponse
	err = decoder.Decode(&res)
	if err != nil {
		fmt.Printf("decoder.Decode error %v\n", err)
		return nil, err
	}
	return &res, nil
}

type QueryAllNFTClassesByOwnerRequest struct {
	ISCNOwner string `url:"iscn_owner"`
}

type QueryAllNFTClasssesByOwnerResponse struct {
	Classes []model.ClassListItem `json:"classes,omitempty"`
}

func (c *LikeNFTCosmosClient) QueryAllNFTClassesByOwner(ctx context.Context, request QueryAllNFTClassesByOwnerRequest) (*QueryAllNFTClasssesByOwnerResponse, error) {
	successChan := make(chan []model.ClassListItem, 1)
	errChan := make(chan error, 1)

	go func() {
		classes := make([]model.ClassListItem, 0)
		var nextKey *int = nil

		for {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
			}
			queryNFTClassesByOwnerResponse, err := c.QueryNFTClassesByOwner(QueryNFTClassesByOwnerRequest{
				ISCNOwner: request.ISCNOwner,
				QueryNFTClassesByOwnerPageRequest: QueryNFTClassesByOwnerPageRequest{
					Key: nextKey,
				},
			})
			if err != nil {
				errChan <- err
				return
			}
			if queryNFTClassesByOwnerResponse.Pagination.Count == 0 {
				successChan <- classes
				return
			}
			classes = append(classes, queryNFTClassesByOwnerResponse.Classes...)
			nextKey = &queryNFTClassesByOwnerResponse.Pagination.NextKey
		}
	}()

	select {
	case err := <-errChan:
		return nil, err
	case classes := <-successChan:
		return &QueryAllNFTClasssesByOwnerResponse{
			Classes: classes,
		}, nil
	}
}
