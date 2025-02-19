package internal_api

import (
	"database/sql"
	"net/http"

	"github.com/likecoin/like-migration-backend/pkg/handler/internal_api/nft_cosmos_class_migration"
)

type InternalRouter struct {
	Db *sql.DB
}

func (h *InternalRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	nftCosmosClassMigrationHandler := &nft_cosmos_class_migration.NFTCosmosClassMigrationRouter{
		Db: h.Db,
	}
	router.Handle("/nft-cosmos-class-migration", nftCosmosClassMigrationHandler.Router())
	router.Handle("/nft-cosmos-class-migration/", nftCosmosClassMigrationHandler.Router())

	return router
}
