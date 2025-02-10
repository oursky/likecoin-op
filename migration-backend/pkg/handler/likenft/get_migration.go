package likenft

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/handler/likenft/model"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	cosmos_model "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
)

type GetMigrationPreviewRequestBody struct {
	CosmosAddress string
}

type GetMigrationPreviewResponseBody struct {
	Migration        *model.Migration `json:"migration,omitempty"`
	ErrorDescription string           `json:"error_description,omitempty"`
}

type GetMigrationHandler struct {
	LikerlandUrlBase    string
	LikeNFTCosmosClient *cosmos.LikeNFTCosmosClient
}

func (h *GetMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cosmosAddress := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	m, err := h.handle(GetMigrationPreviewRequestBody{
		CosmosAddress: cosmosAddress,
	})

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, &CreateSigningMessageResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &GetMigrationPreviewResponseBody{
		Migration: m,
	})
}

func (h *GetMigrationHandler) handle(request GetMigrationPreviewRequestBody) (*model.Migration, error) {
	// TODO handle db
	return h.handlePreview(request)
}

func (h *GetMigrationHandler) handlePreview(request GetMigrationPreviewRequestBody) (*model.Migration, error) {
	classesResponse, err := h.LikeNFTCosmosClient.QueryNFTClassesByOwner(cosmos.QueryNFTClassesByOwnerRequest{
		Owner: request.CosmosAddress,
		QueryNFTClassesByOwnerPageRequest: cosmos.QueryNFTClassesByOwnerPageRequest{
			Limit:      100,
			CountTotal: true,
		},
	})
	if err != nil {
		return nil, err
	}
	nftsResponse, err := h.LikeNFTCosmosClient.QueryNFTsByOwner(cosmos.QueryNFTsByOwnerRequest{
		Owner: request.CosmosAddress,
		QueryNFTsByOwnerPageRequest: cosmos.QueryNFTsByOwnerPageRequest{
			Limit:      100,
			CountTotal: true,
		},
	})
	if err != nil {
		return nil, err
	}

	classes := make([]model.Class, len(classesResponse.Classes))
	for i := range classesResponse.Classes {
		c := *h.mapClass(classesResponse.Classes[i])
		c.Status = model.ClassMigrationStatusPending
		classes[i] = c
	}
	nfts := make([]model.NFT, len(nftsResponse.NFTs))
	for i := range nftsResponse.NFTs {
		nfts[i] = *h.mapNFT(nftsResponse.NFTs[i])
	}

	return &model.Migration{
		Status:  model.MigrationStatusPreview,
		Classes: classes,
		NFTs:    nfts,
	}, nil
}

func (h *GetMigrationHandler) mapClass(c cosmos_model.Class) *model.Class {
	return &model.Class{
		Id:           c.Id,
		Name:         c.Name,
		LikerlandUrl: fmt.Sprintf("%v/nft/class/%v", h.LikerlandUrlBase, c.Id),
		Image:        c.Metadata.Image,
		Status:       model.ClassMigrationStatusPending,
	}
}

func (h *GetMigrationHandler) mapNFT(n cosmos_model.NFT) *model.NFT {
	return &model.NFT{
		ClassId:      n.ClassId,
		Id:           n.Id,
		Name:         n.Data.Metadata.Name,
		LikerlandUrl: fmt.Sprintf("%v/nft/class/%v/%v", h.LikerlandUrlBase, n.ClassId, n.Id),
		Image:        n.Data.Metadata.Image,
		Status:       model.NFTMigrationStatusPending,
	}
}
