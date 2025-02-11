package likenft

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	api_model "github.com/likecoin/like-migration-backend/pkg/handler/likenft/model"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	cosmos_model "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

type GetMigrationPreviewRequestBody struct {
	CosmosAddress string
}

type GetMigrationPreviewResponseBody struct {
	Migration        *api_model.Migration `json:"migration,omitempty"`
	ErrorDescription string               `json:"error_description,omitempty"`
}

type GetMigrationHandler struct {
	Db                  *sql.DB
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

func (h *GetMigrationHandler) handle(request GetMigrationPreviewRequestBody) (*api_model.Migration, error) {
	nftMigration, err := appdb.QueryNFTMigrationByCosmosAddress(h.Db, request.CosmosAddress)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return h.handlePreview(request)
		}
		return nil, err
	}

	return h.handleDb(nftMigration)
}

func (h *GetMigrationHandler) handlePreview(request GetMigrationPreviewRequestBody) (*api_model.Migration, error) {
	classesResponse, err := h.LikeNFTCosmosClient.QueryAllNFTClassesByOwner(cosmos.QueryAllNFTClassesByOwnerRequest{
		Owner: request.CosmosAddress,
	})
	if err != nil {
		return nil, err
	}
	nftsResponse, err := h.LikeNFTCosmosClient.QueryAllNFTsByOwner(cosmos.QueryAllNFTsByOwnerRequest{
		Owner: request.CosmosAddress,
	})
	if err != nil {
		return nil, err
	}

	classes := make([]api_model.Class, len(classesResponse.Classes))
	for i := range classesResponse.Classes {
		c := *h.mapCosmosClass(classesResponse.Classes[i])
		c.Status = api_model.ClassMigrationStatusPending
		classes[i] = c
	}
	nfts := make([]api_model.NFT, len(nftsResponse.NFTs))
	for i := range nftsResponse.NFTs {
		nfts[i] = *h.mapCosmosNFT(nftsResponse.NFTs[i])
	}

	return &api_model.Migration{
		Status:  api_model.MigrationStatusPreview,
		Classes: classes,
		NFTs:    nfts,
	}, nil
}

func (h *GetMigrationHandler) handleDb(nftMigration *model.NFTMigration) (*api_model.Migration, error) {
	nftMigrationClasses, err := appdb.QueryNFTMigrationClassesByNFTMigrationId(h.Db, nftMigration.Id)

	if err != nil {
		return nil, err
	}

	nftMigrationNFTs, err := appdb.QueryNFTMigrationNFTsByNFTMigrationId(h.Db, nftMigration.Id)

	if err != nil {
		return nil, err
	}

	classes := make([]api_model.Class, len(nftMigrationClasses))
	for i := range nftMigrationClasses {
		c := *h.mapDbClass(nftMigrationClasses[i])
		classes[i] = c
	}
	nfts := make([]api_model.NFT, len(nftMigrationNFTs))
	for i := range nftMigrationNFTs {
		nfts[i] = *h.mapDbNFT(nftMigrationNFTs[i])
	}

	return &api_model.Migration{
		Status:  api_model.MigrationStatus(nftMigration.Status),
		Classes: classes,
		NFTs:    nfts,
	}, nil
}

func (h *GetMigrationHandler) mapCosmosClass(c cosmos_model.Class) *api_model.Class {
	return &api_model.Class{
		Id:           c.Id,
		Name:         c.Name,
		LikerlandUrl: fmt.Sprintf("%v/nft/class/%v", h.LikerlandUrlBase, c.Id),
		Image:        c.Metadata.Image,
		Status:       api_model.ClassMigrationStatusPending,
	}
}

func (h *GetMigrationHandler) mapCosmosNFT(n cosmos_model.NFT) *api_model.NFT {
	return &api_model.NFT{
		ClassId:      n.ClassId,
		Id:           n.Id,
		Name:         n.Data.Metadata.Name,
		LikerlandUrl: fmt.Sprintf("%v/nft/class/%v/%v", h.LikerlandUrlBase, n.ClassId, n.Id),
		Image:        n.Data.Metadata.Image,
		Status:       api_model.NFTMigrationStatusPending,
	}
}

func (h *GetMigrationHandler) mapDbClass(c model.NFTMigrationClass) *api_model.Class {
	return &api_model.Class{
		Id:           c.CosmosClassId,
		Name:         c.Name,
		LikerlandUrl: c.LikerlandUrl,
		Image:        c.Image,
		Status:       api_model.ClassMigrationStatus(c.Status),
		EnqueuedTime: c.EnqueueTime,
		EVMTxHash:    c.EvmTxHash,
		FailedReason: c.FailedReason,
	}
}

func (h *GetMigrationHandler) mapDbNFT(n model.NFTMigrationNFT) *api_model.NFT {
	return &api_model.NFT{
		ClassId:      n.CosmosClassId,
		Id:           n.CosmosNFTId,
		Name:         n.Name,
		LikerlandUrl: n.LikerlandUrl,
		Image:        n.Image,
		Status:       api_model.NFTMigrationStatus(n.Status),
		EnqueuedTime: n.EnqueueTime,
		EVMTxHash:    n.EvmTxHash,
		FailedReason: n.FailedReason,
	}
}
