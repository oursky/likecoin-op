package likenft

import (
	"database/sql"
	"net/http"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
)

type LikeNFTRouter struct {
	LikerlandUrlBase    string
	Db                  *sql.DB
	LikeNFTCosmosClient *cosmos.LikeNFTCosmosClient
}

func (h *LikeNFTRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("POST /signing_message", &CreateSigningMessageHandler{
		Db: h.Db,
	})
	router.Handle("POST /likerid/migration", &LikerIDMigrationHandler{})
	router.Handle("GET /migration/", &GetMigrationHandler{
		Db:                  h.Db,
		LikerlandUrlBase:    h.LikerlandUrlBase,
		LikeNFTCosmosClient: h.LikeNFTCosmosClient,
	})
	router.Handle("POST /migration", &CreateMigrationHandler{
		LikerlandUrlBase:    h.LikerlandUrlBase,
		Db:                  h.Db,
		LikeNFTCosmosClient: h.LikeNFTCosmosClient,
	})

	return router
}
