package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type EvmMigrateBookRequest struct {
	LikeClassID string `json:"like_class_id"`
	EvmClassID  string `json:"evm_class_id"`
}

type EvmMigrateBookResponse struct {
	MigratedClassIds      []string `json:"migratedClassIds"`
	MigratedCollectionIds []string `json:"migratedCollectionIds"`
	Error                 string   `json:"error"`
}

func (a *LikecoinAPI) SubmitEvmBookMigrated(request *EvmMigrateBookRequest) (*EvmMigrateBookResponse, error) {
	return a.EvmMigrateBook(request)
}

func (a *LikecoinAPI) EvmMigrateBook(request *EvmMigrateBookRequest) (*EvmMigrateBookResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/wallet/evm/migrate/book", a.LikecoinAPIUrlBase),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var response EvmMigrateBookResponse
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
