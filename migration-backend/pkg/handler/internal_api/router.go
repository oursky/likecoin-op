package internal_api

import (
	"database/sql"
	"net/http"

	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/handler/internal_api/nft_cosmos_class_migration"
)

func InternalAPIMiddleware(apiKey string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKeyHeader := r.Header.Get("X-API-KEY")
		if apiKeyHeader == apiKey {
			next.ServeHTTP(w, r)
		} else {
			handler.SendJSON(w, http.StatusUnauthorized, nil)
		}
	})
}

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
