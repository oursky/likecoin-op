package db

import "github.com/likecoin/like-migration-backend/pkg/model"

func QueryNFTMigrationByCosmosAddress(
	tx TxLike,
	cosmosAddress string,
) (*model.NFTMigration, error) {
	row := tx.QueryRow(
		`SELECT
	id,
	cosmos_address,
	eth_address,
	status
FROM nft_migration WHERE cosmos_address = $1`,
		cosmosAddress,
	)

	nftMigration := &model.NFTMigration{}

	err := row.Scan(
		&nftMigration.Id,
		&nftMigration.CosmosAddress,
		&nftMigration.EthAddress,
		&nftMigration.Status,
	)

	if err != nil {
		return nil, err
	}

	return nftMigration, nil
}

func InsertNFTMigration(
	tx TxLike,
	nftMigration *model.NFTMigration,
) (*model.NFTMigration, error) {
	row := tx.QueryRow(
		`INSERT INTO nft_migration (
	cosmos_address,
	eth_address,
	status
) VALUES ($1, $2, $3) RETURNING id`,
		nftMigration.CosmosAddress,
		nftMigration.EthAddress,
		nftMigration.Status,
	)

	lastInsertId := 0
	err := row.Scan(&lastInsertId)

	if err != nil {
		return nil, err
	}

	n := *nftMigration
	n.Id = uint64(lastInsertId)

	return &n, nil
}
