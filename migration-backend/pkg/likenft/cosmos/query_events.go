package cosmos

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/util/httputil"
)

var ErrQueryNFTEvents = errors.New("err querying nft events")

type queryNFTEventsRequest struct {
	ClassId      string   `url:"class_id,omitempty"`
	NFTId        string   `url:"nft_id,omitempty"`
	ActionTypes  []string `url:"action_type,omitempty"`
	IgnoreToList string   `url:"ignore_to_list,omitempty"`
}

func (c *LikeNFTCosmosClient) MakeQueryNFTEventsRequest(
	classId string,
	nftId string,
) queryNFTEventsRequest {
	return queryNFTEventsRequest{
		ClassId: classId,
		NFTId:   nftId,
		ActionTypes: []string{
			"/cosmos.nft.v1beta1.MsgSend",
			"buy_nft",
			"sell_nft",
			"new_class",
		},
		IgnoreToList: c.nftEventsIgnoreToList,
	}
}

type QueryEventsPageRequest struct {
	Key        int    `url:"pagination.key,omitempty"`
	Offset     uint64 `url:"pagination.offset,omitempty"`
	Limit      uint64 `url:"pagination.limit,omitempty"`
	CountTotal bool   `url:"pagination.countTotal,omitempty"`
	Reverse    bool   `url:"reverse,omitempty"`
}

type QueryEventsResponse struct {
	Events     []model.Event           `json:"events,omitempty"`
	Pagination QueryEventsPageResponse `json:"pagination,omitempty"`
}

type QueryEventsPageResponse struct {
	NextKey int    `json:"next_key,omitempty"`
	Count   uint64 `json:"count,omitempty"`
}

func (c *LikeNFTCosmosClient) queryNFTEvents(
	requestParams queryNFTEventsRequest,
	pageParams QueryEventsPageRequest,
) (*QueryEventsResponse, error) {
	u, err := url.Parse("/likechain/likenft/v1/event")

	if err != nil {
		return nil, errors.Join(ErrQueryNFTEvents, fmt.Errorf("url.Parse %s", "/likechain/likenft/v1/event"), err)
	}

	requestQuery, err := query.Values(requestParams)

	if err != nil {
		return nil, errors.Join(ErrQueryNFTEvents, fmt.Errorf("query.Values requestParams"), err)
	}

	pageQuery, err := query.Values(pageParams)

	if err != nil {
		return nil, errors.Join(ErrQueryNFTEvents, fmt.Errorf("query.Values pageParams"), err)
	}

	u.RawQuery = fmt.Sprintf("%s&%s", requestQuery.Encode(), pageQuery.Encode())

	base, err := url.Parse(c.NodeURL)

	if err != nil {
		return nil, errors.Join(ErrQueryNFTEvents, fmt.Errorf("url.Parse %s", c.NodeURL), err)
	}

	resp, err := c.HTTPClient.Get(base.ResolveReference(u).String())

	if err != nil {
		return nil, errors.Join(ErrQueryNFTEvents, fmt.Errorf("c.HTTPClient.Get"), err)
	}
	if err = httputil.HandleResponseStatus(resp); err != nil {
		return nil, errors.Join(ErrQueryNFTEvents, err)
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var response QueryEventsResponse
	err = decoder.Decode(&response)
	if err != nil {
		return nil, errors.Join(ErrQueryNFTEvents, fmt.Errorf("decoder.Decode"), err)
	}
	return &response, nil
}

var (
	ErrQueryAllNFTEvents = errors.New("err querying all nft events")
)

func (c *LikeNFTCosmosClient) QueryAllNFTEvents(
	request queryNFTEventsRequest,
) ([]model.Event, error) {
	limit := uint64(100)
	p := QueryEventsPageRequest{
		Limit:      limit,
		Reverse:    false,
		CountTotal: true,
	}

	events := make([]model.Event, 0)

	resp, err := c.queryNFTEvents(request, p)

	if err != nil {
		return nil, errors.Join(ErrQueryAllNFTEvents, fmt.Errorf("c.queryNFTEvents"), err)
	}

	events = append(events, resp.Events...)

	page := uint64(math.Ceil(float64(resp.Pagination.Count) / float64(limit)))

	key := resp.Pagination.NextKey
	for p := uint64(2); p <= page; p++ {
		_p := QueryEventsPageRequest{
			Limit:   limit,
			Key:     key,
			Reverse: false,
		}

		resp, err = c.queryNFTEvents(request, _p)

		if err != nil {
			return nil, errors.Join(ErrQueryAllNFTEvents, fmt.Errorf("c.queryNFTEvents"), err)
		}

		events = append(events, resp.Events...)
		key = resp.Pagination.NextKey
	}

	return events, nil
}
