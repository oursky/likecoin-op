package internal_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type UpdateNFTCosmosClassMigrationRequestBody struct {
	NFTCosmosClassMigration *model.NFTCosmosClassMigration `json:"nft_cosmos_class_migration,omitempty"`
}

type UpdateNFTCosmosClassMigrationResponseBody struct {
	NFTCosmosClassMigration *model.NFTCosmosClassMigration `json:"nft_cosmos_class_migration,omitempty"`
	ErrorDescription        string                         `json:"error_description,omitempty"`
}

func (a *InternalAPI) UpdateNFTCosmosClassMigration(m *UpdateNFTCosmosClassMigrationRequestBody) (*model.NFTCosmosClassMigration, error) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("PUT", fmt.Sprintf("%s/internal/nft-cosmos-class-migration/%s", a.APIUrlBase, m.NFTCosmosClassMigration.CosmosClassId), buf)
	if err != nil {
		return nil, err
	}
	resp, err := a.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var response UpdateNFTCosmosClassMigrationResponseBody
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return response.NFTCosmosClassMigration, nil
}
