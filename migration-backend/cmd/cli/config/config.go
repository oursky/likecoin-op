package config

import (
	"github.com/kelseyhightower/envconfig"
)

type contextKey struct{}

var ContextKey = &contextKey{}

type EnvConfig struct {
	CosmosNodeUrl                   string `envconfig:"COSMOS_NODE_URL"`
	CosmosNodeHTTPTimeoutSeconds    int    `envconfig:"COSMOS_NODE_HTTP_TIMEOUT_SECONDS" default:"10"`
	CosmosNftEventsIgnoreToList     string `envconfig:"COSMOS_NFT_EVENTS_IGNORE_TO_LIST"`
	CosmosLikeCoinNetworkConfigPath string `envconfig:"COSMOS_LIKECOIN_NETWORK_CONFIG_PATH"`
	EthSignerBaseUrl                string `envconfig:"ETH_SIGNER_BASE_URL"`
	EthSignerAPIKey                 string `envconfig:"ETH_SIGNER_API_KEY"`
	EthNetworkPublicRPCURL          string `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`
	EthTokenAddress                 string `envconfig:"ETH_TOKEN_ADDRESS"`
	EthLikeNFTContractAddress       string `envconfig:"ETH_LIKENFT_CONTRACT_ADDRESS"`
	DbConnectionStr                 string `envconfig:"DB_CONNECTION_STR"`
	LikecoinAPIUrlBase              string `envconfig:"LIKECOIN_API_URL_BASE"`
	LikecoinAPIHTTPTimeoutSeconds   int    `envconfig:"LIKECOIN_API_HTTP_TIMEOUT_SECONDS" default:"10"`

	InitialNewClassOwner      string `envconfig:"INITIAL_NEW_CLASS_OWNER"`
	InitialNewClassMinter     string `envconfig:"INITIAL_NEW_CLASS_MINTER"`
	InitialNewClassUpdater    string `envconfig:"INITIAL_NEW_CLASS_UPDATER"`
	InitialBatchMintNFTsOwner string `envconfig:"INITIAL_BATCH_MINT_NFTS_OWNER"`

	W3SpaceDID          string `envconfig:"W3_SPACE_DID"`
	W3UcanDIDPrivateKey string `envconfig:"W3_UCAN_DID_PRIVATE_KEY"`
	W3UcanProofPath     string `envconfig:"W3_UCAN_PROOF_PATH"`

	IpfsHTTPBaseURL    string `envconfig:"IPFS_HTTP_BASE_URL"`
	ArweaveHTTPBaseURL string `envconfig:"ARWEAVE_HTTP_BASE_URL"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
