package db

import (
	"time"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

func QueryNFTCosmosClassMigrationByCosmosClassId(tx TxLike, cosmosClassId string) (*model.NFTCosmosClassMigration, error) {
	row := tx.QueryRow(
		`SELECT
	id,
	created_at,
	cosmos_class_id,
	evm_class_id,
	evm_tx_hash,
	status,
	failed_reason
FROM nft_cosmos_class_migration WHERE cosmos_class_id = $1`,
		cosmosClassId,
	)

	c := &model.NFTCosmosClassMigration{}

	err := row.Scan(
		&c.Id,
		&c.CreatedAt,
		&c.CosmosClassId,
		&c.EvmClassId,
		&c.EvmTxHash,
		&c.Status,
		&c.FailedReason,
	)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func InsertNFTCosmosClassMigration(tx TxLike, c *model.NFTCosmosClassMigration) (*model.NFTCosmosClassMigration, error) {
	row := tx.QueryRow(
		`INSERT INTO nft_cosmos_class_migration (
	cosmos_class_id,
	evm_class_id,
	evm_tx_hash,
	status,
	failed_reason
) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`,
		c.CosmosClassId,
		c.EvmClassId,
		c.EvmTxHash,
		c.Status,
		c.FailedReason,
	)

	var lastInsertId uint64
	var createdAt time.Time
	err := row.Scan(&lastInsertId, &createdAt)

	if err != nil {
		return nil, err
	}

	m := *c
	m.Id = lastInsertId

	return &m, nil
}

func UpdateNFTCosmosClassMigrationByCosmosClassId(tx TxLike, c *model.NFTCosmosClassMigration) error {
	_, err := tx.Exec(
		`UPDATE nft_cosmos_class_migration SET
	evm_class_id = $2,
	evm_tx_hash = $3,
	status = $4,
	failed_reason = $5
WHERE cosmos_class_id = $1`,
		c.CosmosClassId,
		c.EvmClassId,
		c.EvmTxHash,
		c.Status,
		c.FailedReason,
	)

	return err
}
