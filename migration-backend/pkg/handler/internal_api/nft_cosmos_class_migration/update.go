package nft_cosmos_class_migration

import (
	"database/sql"
	"encoding/json"
	"net/http"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

type UpdateNFTCosmosClassMigrationRequestBody struct {
	NFTCosmosClassMigration *model.NFTCosmosClassMigration `json:"nft_cosmos_class_migration,omitempty"`
}

type UpdateNFTCosmosClassMigrationResponseBody struct {
	NFTCosmosClassMigration *model.NFTCosmosClassMigration `json:"nft_cosmos_class_migration,omitempty"`
	ErrorDescription        string                         `json:"error_description,omitempty"`
}

type UpdateNFTCosmosClassMigrationHandler struct {
	Db *sql.DB
}

func (p *UpdateNFTCosmosClassMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data UpdateNFTCosmosClassMigrationRequestBody
	err := decoder.Decode(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, &UpdateNFTCosmosClassMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	err = p.handle(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, &GetCosmosClassMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &GetCosmosClassMigrationResponseBody{
		NFTCosmosClassMigration: data.NFTCosmosClassMigration,
	})
}

func (p *UpdateNFTCosmosClassMigrationHandler) handle(req *UpdateNFTCosmosClassMigrationRequestBody) error {
	return appdb.UpdateNFTCosmosClassMigrationByCosmosClassId(p.Db, req.NFTCosmosClassMigration)
}
