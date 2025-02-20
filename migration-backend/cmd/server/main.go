package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"database/sql"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	_ "github.com/lib/pq"

	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/handler/internal_api"
	"github.com/likecoin/like-migration-backend/pkg/handler/likenft"
	"github.com/likecoin/like-migration-backend/pkg/handler/user"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
)

func main() {
	logger := slog.New(slog.Default().Handler())
	err := godotenv.Load()
	if errors.Is(err, os.ErrNotExist) {
		logger.Warn("skip loading .env as it is absent")
	} else if err != nil {
		panic(err)
	}
	envCfg, err := LoadEnvConfigFromEnv()
	if err != nil {
		panic(err)
	}

	routePrefixMux := http.NewServeMux()
	mainMux := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	db, err := sql.Open("postgres", envCfg.DbConnectionStr)
	if err != nil {
		panic(err)
	}

	client, err := ethclient.Dial(envCfg.EthNetworkPublicRPCURL)

	if err != nil {
		panic(err)
	}

	cosmosAPI := &api.CosmosAPI{
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		NodeURL: envCfg.CosmosNodeUrl,
	}
	likenftCosmosClient := &cosmos.LikeNFTCosmosClient{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		NodeURL: envCfg.CosmosNodeUrl,
	}
	likecoinAPI := &likecoin_api.LikecoinAPI{
		LikecoinAPIUrlBase: envCfg.LikecoinAPIUrlBase,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	mainMux.Handle("/healthz", &handler.HealthzHandler{})
	mainMux.Handle("/init_likecoin_migration_from_cosmos", &handler.InitLikeCoinMigrationFromCosmosHandler{
		Db:                     db,
		EthClient:              client,
		CosmosAPI:              cosmosAPI,
		EthWalletPrivateKey:    envCfg.EthWalletPrivateKey,
		EthNetworkPublicRPCURL: envCfg.EthNetworkPublicRPCURL,
		EthTokenAddress:        envCfg.EthTokenAddress,
	})
	mainMux.Handle("/migration_record/", &handler.GetLikeCoinMigrationRecordHandler{
		Db:        db,
		EthClient: client,
	})
	likeNFTRouter := likenft.LikeNFTRouter{
		Db:                  db,
		CosmosAPI:           cosmosAPI,
		LikeNFTCosmosClient: likenftCosmosClient,
	}
	mainMux.Handle("/likenft/", http.StripPrefix("/likenft", likeNFTRouter.Router()))

	userRouter := user.UserRouter{
		Db:          db,
		LikecoinAPI: likecoinAPI,
	}
	mainMux.Handle("/user/", http.StripPrefix("/user", userRouter.Router()))

	internalRouter := internal_api.InternalRouter{
		Db: db,
	}
	mainMux.Handle("/internal/", internal_api.InternalAPIMiddleware(envCfg.InternalAPIKey, http.StripPrefix("/internal", internalRouter.Router())))

	routePrefixMux.Handle(fmt.Sprintf("%s/", envCfg.RoutePrefix), http.StripPrefix(envCfg.RoutePrefix, mainMux))

	server := &http.Server{
		Addr:              envCfg.ListenAddr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           c.Handler(routePrefixMux),
	}

	logger.Info("listening",
		"addr", envCfg.ListenAddr,
	)

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
