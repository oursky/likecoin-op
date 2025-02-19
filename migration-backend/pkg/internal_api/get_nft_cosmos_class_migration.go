package internal_api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type GetCosmosClassMigrationResponseBody struct {
	NFTCosmosClassMigration *model.NFTCosmosClassMigration `json:"nft_cosmos_class_migration,omitempty"`
	ErrorDescription        string                         `json:"error_description,omitempty"`
}

func (a *InternalAPI) GetNFTCosmosClassMigration(cosmosClassId string) (*model.NFTCosmosClassMigration, error) {
	request, err := http.NewRequest("GET",
		fmt.Sprintf("%s/internal/nft-cosmos-class-migration/%s", a.APIUrlBase, cosmosClassId), nil)
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
	var response GetCosmosClassMigrationResponseBody
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return response.NFTCosmosClassMigration, nil
}
