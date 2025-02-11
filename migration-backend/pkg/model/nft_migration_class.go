package model

import "time"

type NFTMigrationClassStatus string

const (
	NFTMigrationClassStatusPreview    NFTMigrationClassStatus = "preview"
	NFTMigrationClassStatusPending    NFTMigrationClassStatus = "pending"
	NFTMigrationClassStatusScheduled  NFTMigrationClassStatus = "scheduled"
	NFTMigrationClassStatusInProgress NFTMigrationClassStatus = "in_progress"
	NFTMigrationClassStatusSuccess    NFTMigrationClassStatus = "success"
	NFTMigrationClassStatusFailed     NFTMigrationClassStatus = "failed"
)

type NFTMigrationClass struct {
	Id             uint64
	NFTMigrationId uint64
	CreatedAt      *time.Time
	CosmosClassId  string
	Name           string
	LikerlandUrl   string
	Image          string
	Status         NFTMigrationClassStatus
	EnqueueTime    *time.Time
	EvmTxHash      *string
	FailedReason   *string
}
