package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/internal_api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/logic"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/spf13/cobra"
)

var mintNFTClassCmd = &cobra.Command{
	Use:   "mint-nft-class cosmos-class-id to-eth-address",
	Short: "Mint NFT Class",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			cmd.Usage()
			return
		}

		envCfg := cmd.Context().Value(config.ContextKey).(*config.EnvConfig)

		likenftClient := &cosmos.LikeNFTCosmosClient{
			HTTPClient: &http.Client{
				Timeout: 5 * time.Second,
			},
			NodeURL: envCfg.CosmosNodeUrl,
		}

		ethClient, err := ethclient.Dial(envCfg.EthNetworkPublicRPCURL)
		if err != nil {
			panic(err)
		}
		privateKey, err := crypto.HexToECDSA(envCfg.EthWalletPrivateKey)
		if err != nil {
			panic(err)
		}
		contractAddress := common.HexToAddress(envCfg.EthLikeNFTContractAddress)

		evmLikeNFT := evm.NewLikeNFT(
			ethClient,
			privateKey,
			envCfg.EthChainId,
			contractAddress,
		)

		internalAPI := &internal_api.InternalAPI{
			HTTPClient: &http.Client{
				Timeout: 5 * time.Second,
			},
			APIUrlBase: "http://localhost:8091",
		}

		cosmosClassId := args[0]
		ethAddress := args[1]

		m := &model.NFTCosmosClassMigration{
			CosmosClassId: cosmosClassId,
			Status:        model.NFTCosmosClassMigrationStatusInit,
		}
		_, err = internalAPI.CreateNFTCosmosClassMigration(&internal_api.CreateCosmosClassMigrationRequestBody{
			NFTCosmosClassMigration: m,
		})

		if err != nil {
			panic(err)
		}

		var resolveChan = make(chan *model.NFTCosmosClassMigration)
		var rejectChan = make(chan error)
		go func() {
			cc, err := logic.MintNFTClass(likenftClient, &evmLikeNFT, internalAPI, cosmosClassId, ethAddress)
			if err != nil {
				rejectChan <- err
				return
			}
			resolveChan <- cc
		}()

		select {
		case cc := <-resolveChan:
			fmt.Printf("Mint nft class completed from %v to %v", cosmosClassId, *cc.EvmClassId)
		case err = <-rejectChan:
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(mintNFTClassCmd)
}
