package nft_cosmos_class_migration

import (
	"database/sql"
	"encoding/json"
	"net/http"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

type CreateCosmosClassMigrationRequestBody struct {
	NFTCosmosClassMigration *model.NFTCosmosClassMigration `json:"nft_cosmos_class_migration,omitempty"`
}

type CreateCosmosClassMigrationResponseBody struct {
	NFTCosmosClassMigration *model.NFTCosmosClassMigration `json:"nft_cosmos_class_migration,omitempty"`
	ErrorDescription        string                         `json:"error_description,omitempty"`
}

type CreateNFTCosmosClassMigrationHandler struct {
	Db *sql.DB
}

func (p *CreateNFTCosmosClassMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data CreateCosmosClassMigrationRequestBody
	err := decoder.Decode(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, &CreateCosmosClassMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	c, err := p.handle(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, &GetCosmosClassMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &GetCosmosClassMigrationResponseBody{
		NFTCosmosClassMigration: c,
	})
}

func (p *CreateNFTCosmosClassMigrationHandler) handle(req *CreateCosmosClassMigrationRequestBody) (*model.NFTCosmosClassMigration, error) {
	return appdb.InsertNFTCosmosClassMigration(p.Db, req.NFTCosmosClassMigration)
}
