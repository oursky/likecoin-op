package model

import "time"

type NFTMigrationNFTStatus string

const (
	NFTMigrationNFTStatusPreview    NFTMigrationNFTStatus = "preview"
	NFTMigrationNFTStatusPending    NFTMigrationNFTStatus = "pending"
	NFTMigrationNFTStatusScheduled  NFTMigrationNFTStatus = "scheduled"
	NFTMigrationNFTStatusInProgress NFTMigrationNFTStatus = "in_progress"
	NFTMigrationNFTStatusSuccess    NFTMigrationNFTStatus = "success"
	NFTMigrationNFTStatusFailed     NFTMigrationNFTStatus = "failed"
)

type NFTMigrationNFT struct {
	Id             uint64
	NFTMigrationId uint64
	CreatedAt      time.Time
	CosmosClassId  string
	CosmosNFTId    string
	Name           string
	LikerlandUrl   string
	Image          string
	Status         NFTMigrationNFTStatus
	EnqueueTime    *time.Time
	EvmTxHash      *string
	FailedReason   *string
}
