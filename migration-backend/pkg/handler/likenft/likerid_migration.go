package likenft

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
)

type MigrateLikerIDWithEthAddressRequestBody struct {
	CosmosAddress        string `json:"cosmos_address,omitempty"`
	CosmosPubKey         string `json:"cosmos_pub_key,omitempty"`
	EthAddress           string `json:"eth_address,omitempty"`
	CosmosSignature      string `json:"cosmos_signature,omitempty"`
	CosmosSigningMessage string `json:"cosmos_signing_message,omitempty"`
	EthSignature         string `json:"eth_signature,omitempty"`
	EthSigningMessage    string `json:"eth_signing_message,omitempty"`
}

type MigrateLikerIDWithEthAddressResponseBody struct {
	Message          string `json:"message,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type LikerIDMigrationHandler struct {
	LikecoinAPI *likecoin_api.LikecoinAPI
}

func (p *LikerIDMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data MigrateLikerIDWithEthAddressRequestBody
	err := decoder.Decode(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, &CreateSigningMessageResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	err = p.handle(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, &CreateSigningMessageResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &CreateSigningMessageResponseBody{
		Message: "OK",
	})
}

func (p *LikerIDMigrationHandler) handle(data *MigrateLikerIDWithEthAddressRequestBody) error {
	recoveredAddr, err := ethereum.RecoverAddress(data.EthSignature, []byte(data.EthSigningMessage))
	if err != nil {
		return err
	}

	// The address from request data maybe in lower case
	// while the recovered address maybe in uppercase
	if !strings.EqualFold(recoveredAddr.Hex(), data.EthAddress) {
		return fmt.Errorf("ethereum signature not verified")
	}

	// The cosmos signature will be verified by the likecoin api
	response, err := p.LikecoinAPI.MigrateUserEVMWallet(&likecoin_api.MigrateUserEVMWalletRequest{
		CosmosAddress:          data.CosmosAddress,
		CosmosSignature:        data.CosmosSignature,
		CosmosPublicKey:        data.CosmosPubKey,
		CosmosSignatureContent: data.CosmosSigningMessage,
		SignMethod:             likecoin_api.MigrateUserEVMWalletRequestSignMethodMemo,
	})

	if err != nil {
		return err
	}

	if response.MigrateLikerLandError != "" {
		return fmt.Errorf("migrate liker land error: %s", response.MigrateLikerLandError)
	}

	if response.MigrateLikerIdError != "" {
		return fmt.Errorf("migrate liker id error: %s", response.MigrateLikerIdError)
	}

	return nil
}
