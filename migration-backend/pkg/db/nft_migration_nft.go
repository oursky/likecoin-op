package db

import (
	"fmt"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

func QueryNFTMigrationNFTsByNFTMigrationId(
	tx TxLike,
	nftMigrationId uint64,
) ([]model.NFTMigrationNFT, error) {
	rows, err := tx.Query(
		`SELECT
	id,
	nft_migration_id,
	created_at,
	cosmos_class_id,
	cosmos_nft_id,
	name,
	likerland_url,
	image,
	status,
	enqueue_time,
	evm_tx_hash,
	failed_reason
FROM nft_migration_nft WHERE nft_migration_id = $1`,
		nftMigrationId,
	)

	if err != nil {
		return nil, err
	}

	nftMigrationNFTs := []model.NFTMigrationNFT{}
	for rows.Next() {
		nftMigrationNFT := &model.NFTMigrationNFT{}

		err := rows.Scan(
			&nftMigrationNFT.Id,
			&nftMigrationNFT.NFTMigrationId,
			&nftMigrationNFT.CreatedAt,
			&nftMigrationNFT.CosmosClassId,
			&nftMigrationNFT.CosmosNFTId,
			&nftMigrationNFT.Name,
			&nftMigrationNFT.LikerlandUrl,
			&nftMigrationNFT.Image,
			&nftMigrationNFT.Status,
			&nftMigrationNFT.EnqueueTime,
			&nftMigrationNFT.EvmTxHash,
			&nftMigrationNFT.FailedReason,
		)

		if err != nil {
			return nil, err
		}

		nftMigrationNFTs = append(nftMigrationNFTs, *nftMigrationNFT)
	}

	return nftMigrationNFTs, nil
}

func InsertNFTMigrationNFTs(
	tx TxLike,
	nftMigrationNFTs []model.NFTMigrationNFT,
) error {
	if len(nftMigrationNFTs) == 0 {
		return nil
	}

	valueStrings := make([]string, 0, len(nftMigrationNFTs))
	valueArgs := make([]interface{}, 0, len(nftMigrationNFTs)*10)

	for i, class := range nftMigrationNFTs {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*10+1, i*10+2, i*10+3, i*10+4, i*10+5, i*10+6, i*10+7, i*10+8, i*10+9, i*10+10))
		valueArgs = append(valueArgs, class.NFTMigrationId)
		valueArgs = append(valueArgs, class.CosmosClassId)
		valueArgs = append(valueArgs, class.CosmosNFTId)
		valueArgs = append(valueArgs, class.Name)
		valueArgs = append(valueArgs, class.LikerlandUrl)
		valueArgs = append(valueArgs, class.Image)
		valueArgs = append(valueArgs, class.Status)
		valueArgs = append(valueArgs, class.EnqueueTime)
		valueArgs = append(valueArgs, class.EvmTxHash)
		valueArgs = append(valueArgs, class.FailedReason)
	}

	stmt := fmt.Sprintf(`INSERT INTO nft_migration_nft (
	nft_migration_id,
	cosmos_class_id,
	cosmos_nft_id,
	name,
	likerland_url,
	image,
	status,
	enqueue_time,
	evm_tx_hash,
	failed_reason
) VALUES %s`, strings.Join(valueStrings, ","))

	_, err := tx.Exec(stmt, valueArgs...)
	return err
}
