package db

import (
	"fmt"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

func QueryNFTMigrationClassesByNFTMigrationId(
	tx TxLike,
	nftMigrationId uint64,
) ([]model.NFTMigrationClass, error) {
	rows, err := tx.Query(
		`SELECT
	id,
	nft_migration_id,
	created_at,
	cosmos_class_id,
	name,
	likerland_url,
	image,
	status,
	enqueue_time,
	evm_tx_hash,
	failed_reason
FROM nft_migration_class WHERE nft_migration_id = $1`,
		nftMigrationId,
	)

	if err != nil {
		return nil, err
	}

	nftMigrationClasses := []model.NFTMigrationClass{}
	for rows.Next() {
		nftMigrationClass := &model.NFTMigrationClass{}

		err := rows.Scan(
			&nftMigrationClass.Id,
			&nftMigrationClass.NFTMigrationId,
			&nftMigrationClass.CreatedAt,
			&nftMigrationClass.CosmosClassId,
			&nftMigrationClass.Name,
			&nftMigrationClass.LikerlandUrl,
			&nftMigrationClass.Image,
			&nftMigrationClass.Status,
			&nftMigrationClass.EnqueueTime,
			&nftMigrationClass.EvmTxHash,
			&nftMigrationClass.FailedReason,
		)

		if err != nil {
			return nil, err
		}

		nftMigrationClasses = append(nftMigrationClasses, *nftMigrationClass)
	}

	return nftMigrationClasses, nil
}

func InsertNFTMigrationClasses(
	tx TxLike,
	nftMigrationClasses []model.NFTMigrationClass,
) error {
	if len(nftMigrationClasses) == 0 {
		return nil
	}
	valueStrings := make([]string, 0, len(nftMigrationClasses))
	valueArgs := make([]interface{}, 0, len(nftMigrationClasses)*9)

	for i, class := range nftMigrationClasses {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*9+1, i*9+2, i*9+3, i*9+4, i*9+5, i*9+6, i*9+7, i*9+8, i*9+9))
		valueArgs = append(valueArgs, class.NFTMigrationId)
		valueArgs = append(valueArgs, class.CosmosClassId)
		valueArgs = append(valueArgs, class.Name)
		valueArgs = append(valueArgs, class.LikerlandUrl)
		valueArgs = append(valueArgs, class.Image)
		valueArgs = append(valueArgs, class.Status)
		valueArgs = append(valueArgs, class.EnqueueTime)
		valueArgs = append(valueArgs, class.EvmTxHash)
		valueArgs = append(valueArgs, class.FailedReason)
	}

	stmt := fmt.Sprintf(`INSERT INTO nft_migration_class (
	nft_migration_id,
	cosmos_class_id,
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
