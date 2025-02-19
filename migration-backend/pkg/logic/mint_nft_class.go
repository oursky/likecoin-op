package logic

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/likecoin/like-migration-backend/pkg/internal_api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func MintNFTClass(c *cosmos.LikeNFTCosmosClient, l *evm.LikeNFT, i *internal_api.InternalAPI, cosmosClassId string, evmAddress string) (*model.NFTCosmosClassMigration, error) {
	m, err := i.GetNFTCosmosClassMigration(cosmosClassId)
	if err != nil {
		return nil, err
	}
	if m.Status != model.NFTCosmosClassMigrationStatusInit {
		return nil, fmt.Errorf("error the cosmos class migration status is not init")
	}
	m.Status = model.NFTCosmosClassMigrationStatusInProgress
	_, err = i.UpdateNFTCosmosClassMigration(&internal_api.UpdateNFTCosmosClassMigrationRequestBody{
		NFTCosmosClassMigration: m,
	})
	if err != nil {
		return nil, err
	}

	cosmosQueryClassResponse, err := c.QueryClassByClassId(cosmos.QueryClassByClassIdRequest{ClassId: cosmosClassId})
	if err != nil {
		return nil, mintNFTClassFailed(i, m, err)
	}

	class := cosmosQueryClassResponse.Class

	// (evm class) <- mint evm class(cosmos class)
	classId, txHash, err := l.NewClass(like_protocol.MsgNewClass{
		Creator: common.HexToAddress(evmAddress),
		Input: like_protocol.ClassInput{
			Name:     class.Name,
			Symbol:   class.Symbol,
			Metadata: "",
			Config: like_protocol.ClassConfig{
				MaxSupply: 0,
			},
		},
	})
	if err != nil {
		return nil, mintNFTClassFailed(i, m, err)
	}

	evmTxhash := hexutil.Encode(txHash.Bytes())
	m.EvmTxHash = &evmTxhash
	evmClassId := hexutil.Encode(classId.Bytes())
	m.EvmClassId = &evmClassId
	m.Status = model.NFTCosmosClassMigrationStatusCompleted

	// transfer ownership (evm class, mint to evm address)
	// Should already be transfered (by creator)

	// (maybe evm class) <- get evm class(class id)
	i.UpdateNFTCosmosClassMigration(&internal_api.UpdateNFTCosmosClassMigrationRequestBody{
		NFTCosmosClassMigration: m,
	})
	if err != nil {
		return nil, err
	}

	// Minted
	return m, nil
}

func mintNFTClassFailed(i *internal_api.InternalAPI, m *model.NFTCosmosClassMigration, reason error) error {
	m.Status = model.NFTCosmosClassMigrationStatusFailed
	failedReason := reason.Error()
	m.FailedReason = &failedReason
	_, err := i.UpdateNFTCosmosClassMigration(&internal_api.UpdateNFTCosmosClassMigrationRequestBody{
		NFTCosmosClassMigration: m,
	})
	if err != nil {
		return err
	}
	return reason
}
