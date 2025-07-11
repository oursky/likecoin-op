package db

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type QueryLikeNFTMigrationActionNewClassFilter struct {
	CosmosClassId *string
	EvmClassId    *string
}

func makeQueryLikeNFTMigrationActionNewClassWhereClauseFromFilter(f QueryLikeNFTMigrationActionNewClassFilter) (string, []any) {
	wheres := make([]string, 0)
	attributes := make([]any, 0)

	if f.CosmosClassId != nil {
		wheres = append(wheres, fmt.Sprintf("cosmos_class_id = $%d", len(wheres)+1))
		attributes = append(attributes, f.CosmosClassId)
	}

	if f.EvmClassId != nil {
		wheres = append(wheres, fmt.Sprintf("evm_class_id = $%d", len(wheres)+1))
		attributes = append(attributes, f.EvmClassId)
	}

	whereClause := ""
	if len(wheres) > 0 {
		whereClause = fmt.Sprintf("%s %s", "WHERE", strings.Join(wheres, " AND "))
	}

	return whereClause, attributes
}

func QueryLikeNFTMigrationActionNewClass(
	tx TxLike,
	filter QueryLikeNFTMigrationActionNewClassFilter,
) (*model.LikeNFTMigrationActionNewClass, error) {
	whereClause, attributes := makeQueryLikeNFTMigrationActionNewClassWhereClauseFromFilter(filter)
	row := tx.QueryRow(
		fmt.Sprintf(`SELECT
    id,
    created_at,
    cosmos_class_id,
    initial_owner,
    initial_minter,
    initial_updater,
    default_royalty_fraction,
    should_premint_all_nfts,
    initial_batch_mint_owner,
    status,
    number_of_cosmos_nfts_found,
    evm_class_id,
    evm_tx_hash,
    failed_reason
FROM likenft_migration_action_new_class
%s
`, whereClause),
		attributes...,
	)

	var defaultRoyaltyFractionStr string
	var numberOfCosmosNFTsFoundStr *string
	m := &model.LikeNFTMigrationActionNewClass{}

	err := row.Scan(
		&m.Id,
		&m.CreatedAt,
		&m.CosmosClassId,
		&m.InitialOwner,
		&m.InitialMintersStr,
		&m.InitialUpdatersStr,
		&defaultRoyaltyFractionStr,
		&m.ShouldPremintAllNFTs,
		&m.InitialBatchMintOwner,
		&m.Status,
		&numberOfCosmosNFTsFoundStr,
		&m.EvmClassId,
		&m.EvmTxHash,
		&m.FailedReason,
	)

	defaultRoyaltyFraction, _ := new(big.Int).SetString(defaultRoyaltyFractionStr, 10)
	m.DefaultRoyaltyFraction = defaultRoyaltyFraction

	if err != nil {
		return nil, err
	}

	if numberOfCosmosNFTsFoundStr != nil {
		numberOfCosmosNFTsFound, err := strconv.ParseUint(*numberOfCosmosNFTsFoundStr, 10, 64)
		if err != nil {
			return nil, err
		}
		m.NumberOfCosmosNFTsFound = &numberOfCosmosNFTsFound
	}

	return m, nil
}

func InsertLikeNFTMigrationActionNewClass(
	tx TxLike,
	a *model.LikeNFTMigrationActionNewClass,
) (*model.LikeNFTMigrationActionNewClass, error) {
	var numberOfCosmosNFTsFoundStr *string
	if a.NumberOfCosmosNFTsFound != nil {
		str := strconv.FormatUint(*a.NumberOfCosmosNFTsFound, 10)
		numberOfCosmosNFTsFoundStr = &str
	}
	row := tx.QueryRow(
		`INSERT INTO likenft_migration_action_new_class (
    cosmos_class_id,
    initial_owner,
    initial_minter,
    initial_updater,
    default_royalty_fraction,
    should_premint_all_nfts,
    initial_batch_mint_owner,
    status,
    number_of_cosmos_nfts_found,
    evm_class_id,
    evm_tx_hash,
    failed_reason
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING id`,
		a.CosmosClassId,
		a.InitialOwner,
		a.InitialMintersStr,
		a.InitialUpdatersStr,
		a.DefaultRoyaltyFraction.String(),
		a.ShouldPremintAllNFTs,
		a.InitialBatchMintOwner,
		a.Status,
		numberOfCosmosNFTsFoundStr,
		a.EvmClassId,
		a.EvmTxHash,
		a.FailedReason,
	)

	lastInsertId := 0
	err := row.Scan(&lastInsertId)

	if err != nil {
		return nil, err
	}

	a.Id = uint64(lastInsertId)

	return a, nil
}

func UpdateLikeNFTMigrationActionNewClass(
	tx TxLike,
	a *model.LikeNFTMigrationActionNewClass,
) error {
	var numberOfCosmosNFTsFoundStr *string
	if a.NumberOfCosmosNFTsFound != nil {
		str := strconv.FormatUint(*a.NumberOfCosmosNFTsFound, 10)
		numberOfCosmosNFTsFoundStr = &str
	}
	result, err := tx.Exec(
		`UPDATE likenft_migration_action_new_class SET
    cosmos_class_id = $1,
    initial_owner = $2,
    initial_minter = $3,
    initial_updater = $4,
    default_royalty_fraction = $5,
    should_premint_all_nfts = $6,
    initial_batch_mint_owner = $7,
    status = $8,
    number_of_cosmos_nfts_found = $9,
    evm_class_id = $10,
    evm_tx_hash = $11,
    failed_reason = $12
WHERE id = $13`,
		a.CosmosClassId,
		a.InitialOwner,
		a.InitialMintersStr,
		a.InitialUpdatersStr,
		a.DefaultRoyaltyFraction.String(),
		a.ShouldPremintAllNFTs,
		a.InitialBatchMintOwner,
		a.Status,
		numberOfCosmosNFTsFoundStr,
		a.EvmClassId,
		a.EvmTxHash,
		a.FailedReason,
		a.Id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("update affect no rows")
	}

	return nil
}
