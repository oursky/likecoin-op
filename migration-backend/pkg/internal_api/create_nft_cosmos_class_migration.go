package internal_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type CreateCosmosClassMigrationRequestBody struct {
	NFTCosmosClassMigration *model.NFTCosmosClassMigration `json:"nft_cosmos_class_migration,omitempty"`
}

type CreateCosmosClassMigrationResponseBody struct {
	NFTCosmosClassMigration *model.NFTCosmosClassMigration `json:"nft_cosmos_class_migration,omitempty"`
	ErrorDescription        string                         `json:"error_description,omitempty"`
}

func (a *InternalAPI) CreateNFTCosmosClassMigration(m *CreateCosmosClassMigrationRequestBody) (*model.NFTCosmosClassMigration, error) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/internal/nft-cosmos-class-migration", a.APIUrlBase), buf)
	if err != nil {
		return nil, err
	}
	a.attachAPIKey(request)
	resp, err := a.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var response CreateCosmosClassMigrationResponseBody
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return response.NFTCosmosClassMigration, nil
}
