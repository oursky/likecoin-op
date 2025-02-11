package model

type NFTMigrationStatus string

const (
	NFTMigrationStatusPreview    NFTMigrationStatus = "preview"
	NFTMigrationStatusPending    NFTMigrationStatus = "pending"
	NFTMigrationStatusScheduled  NFTMigrationStatus = "scheduled"
	NFTMigrationStatusInProgress NFTMigrationStatus = "in_progress"
	NFTMigrationStatusSuccess    NFTMigrationStatus = "success"
	NFTMigrationStatusFailed     NFTMigrationStatus = "failed"
)

type NFTMigration struct {
	Id            uint64
	CosmosAddress string
	EthAddress    string
	Status        NFTMigrationStatus
}
