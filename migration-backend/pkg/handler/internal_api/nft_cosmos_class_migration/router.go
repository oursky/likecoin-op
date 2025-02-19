package nft_cosmos_class_migration

import (
	"database/sql"
	"net/http"
)

type NFTCosmosClassMigrationRouter struct {
	Db *sql.DB
}

func (h *NFTCosmosClassMigrationRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("POST /nft-cosmos-class-migration", &CreateNFTCosmosClassMigrationHandler{
		Db: h.Db,
	})
	router.Handle("GET /nft-cosmos-class-migration/", &GetNFTCosmosClassMigrationHandler{
		Db: h.Db,
	})
	router.Handle("PUT /nft-cosmos-class-migration/", &UpdateNFTCosmosClassMigrationHandler{
		Db: h.Db,
	})

	return router
}
