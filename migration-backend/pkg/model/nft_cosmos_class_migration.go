package model

import "time"

type NFTCosmosClassMigrationStatus string

const (
	NFTCosmosClassMigrationStatusInit       NFTCosmosClassMigrationStatus = "init"
	NFTCosmosClassMigrationStatusInProgress NFTCosmosClassMigrationStatus = "in_progress"
	NFTCosmosClassMigrationStatusCompleted  NFTCosmosClassMigrationStatus = "completed"
	NFTCosmosClassMigrationStatusFailed     NFTCosmosClassMigrationStatus = "failed"
)

type NFTCosmosClassMigration struct {
	Id            uint64                        `json:"id"`
	CreatedAt     time.Time                     `json:"created_at"`
	CosmosClassId string                        `json:"cosmos_class_id"`
	EvmClassId    *string                       `json:"evm_class_id"`
	EvmTxHash     *string                       `json:"evm_tx_hash"`
	Status        NFTCosmosClassMigrationStatus `json:"status"`
	FailedReason  *string                       `json:"failed_reason"`
}
