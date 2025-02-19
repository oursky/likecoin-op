package nft_cosmos_class_migration

import (
	"database/sql"
	"net/http"
	"strings"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

type GetCosmosClassMigrationResponseBody struct {
	NFTCosmosClassMigration *model.NFTCosmosClassMigration `json:"nft_cosmos_class_migration,omitempty"`
	ErrorDescription        string                         `json:"error_description,omitempty"`
}

type GetNFTCosmosClassMigrationHandler struct {
	Db *sql.DB
}

func (p *GetNFTCosmosClassMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cosmosClassId := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	c, err := p.handle(cosmosClassId)

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

func (p *GetNFTCosmosClassMigrationHandler) handle(cosmosClassId string) (*model.NFTCosmosClassMigration, error) {
	return appdb.QueryNFTCosmosClassMigrationByCosmosClassId(p.Db, cosmosClassId)
}
