package likenft

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	api_model "github.com/likecoin/like-migration-backend/pkg/handler/likenft/model"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

type CreateMigrationRequestBody struct {
	CosmosAddress string `json:"cosmos_address,omitempty"`
	EthAddress    string `json:"eth_address,omitempty"`
}

type CreateMigrationResponseBody struct {
	Migration        *api_model.Migration `json:"migration"`
	ErrorDescription string               `json:"error_description,omitempty"`
}

type CreateMigrationHandler struct {
	LikerlandUrlBase    string
	Db                  *sql.DB
	LikeNFTCosmosClient *cosmos.LikeNFTCosmosClient
}

func (h *CreateMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data CreateMigrationRequestBody
	err := decoder.Decode(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, &CreateMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	migration, err := h.handle(data)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, &CreateMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &CreateMigrationResponseBody{
		Migration: migration,
	})
}

func (h *CreateMigrationHandler) handle(request CreateMigrationRequestBody) (*api_model.Migration, error) {
	nftSigningMessage, err := appdb.QueryNFTSigningMessageByCosmosAddressAndEthAddress(h.Db, request.CosmosAddress, request.EthAddress)

	if err != nil {
		return nil, err
	}

	// TODO check if cosmos address and eth address pair matches that on likerland

	classesResponse, err := h.LikeNFTCosmosClient.QueryAllNFTClassesByOwner(cosmos.QueryAllNFTClassesByOwnerRequest{
		Owner: nftSigningMessage.CosmosAddress,
	})
	if err != nil {
		return nil, err
	}

	nftsResponse, err := h.LikeNFTCosmosClient.QueryAllNFTsByOwner(cosmos.QueryAllNFTsByOwnerRequest{
		Owner: nftSigningMessage.CosmosAddress,
	})
	if err != nil {
		return nil, err
	}

	nftMigration := &model.NFTMigration{
		CosmosAddress: nftSigningMessage.CosmosAddress,
		EthAddress:    nftSigningMessage.EthAddress,
		Status:        model.NFTMigrationStatusPending,
	}

	tx, err := h.Db.Begin()
	defer tx.Commit()

	if err != nil {
		return nil, err
	}

	nftMigration, err = appdb.InsertNFTMigration(tx, nftMigration)

	if err != nil {
		return nil, err
	}

	nftMigrationClasses := make([]model.NFTMigrationClass, 0, len(classesResponse.Classes))

	for _, cosmosClass := range classesResponse.Classes {
		now := time.Now().UTC()
		nowPtr := &now
		nftMigrationClasses = append(nftMigrationClasses,
			model.NFTMigrationClass{
				NFTMigrationId: nftMigration.Id,
				CosmosClassId:  cosmosClass.Id,
				Name:           cosmosClass.Name,
				LikerlandUrl:   fmt.Sprintf("%v/nft/class/%v", h.LikerlandUrlBase, cosmosClass.Id),
				Image:          cosmosClass.Metadata.Image,
				Status:         model.NFTMigrationClassStatusPending,
				EnqueueTime:    nowPtr,
			})
	}

	nftMigrationNFTs := make([]model.NFTMigrationNFT, 0, len(nftsResponse.NFTs))

	for _, cosmosNFT := range nftsResponse.NFTs {
		now := time.Now().UTC()
		nowPtr := &now
		nftMigrationNFTs = append(nftMigrationNFTs, model.NFTMigrationNFT{
			NFTMigrationId: nftMigration.Id,
			CosmosClassId:  cosmosNFT.ClassId,
			CosmosNFTId:    cosmosNFT.Id,
			Name:           cosmosNFT.Data.Metadata.Name,
			LikerlandUrl:   fmt.Sprintf("%v/nft/class/%v/%v", h.LikerlandUrlBase, cosmosNFT.ClassId, cosmosNFT.Id),
			Image:          cosmosNFT.Data.Metadata.Image,
			Status:         model.NFTMigrationNFTStatusPending,
			EnqueueTime:    nowPtr,
		})
	}

	err = appdb.InsertNFTMigrationClasses(tx, nftMigrationClasses)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = appdb.InsertNFTMigrationNFTs(tx, nftMigrationNFTs)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	classes := make([]api_model.Class, len(nftMigrationClasses))
	for i := range nftMigrationClasses {
		c := *h.mapClass(nftMigrationClasses[i])
		c.Status = api_model.ClassMigrationStatusPending
		classes[i] = c
	}
	nfts := make([]api_model.NFT, len(nftMigrationNFTs))
	for i := range nftMigrationNFTs {
		nfts[i] = *h.mapNFT(nftMigrationNFTs[i])
	}

	return &api_model.Migration{
		Status:  api_model.MigrationStatus(nftMigration.Status),
		Classes: classes,
		NFTs:    nfts,
	}, nil
}

func (h *CreateMigrationHandler) mapClass(c model.NFTMigrationClass) *api_model.Class {
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

func (h *CreateMigrationHandler) mapNFT(n model.NFTMigrationNFT) *api_model.NFT {
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
