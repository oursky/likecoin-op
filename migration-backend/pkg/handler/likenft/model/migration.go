package model

import "time"

type Migration struct {
	Status  MigrationStatus `json:"status"`
	Classes []Class         `json:"classes"`
	NFTs    []NFT           `json:"nfts"`
}

type MigrationStatus string

const (
	MigrationStatusPreview    MigrationStatus = "preview"
	MigrationStatusPending    MigrationStatus = "pending"
	MigrationStatusInProgress MigrationStatus = "in_progress"
	MigrationStatusSuccess    MigrationStatus = "success"
	MigrationStatusFailed     MigrationStatus = "failed"
)

type Class struct {
	Id           string               `json:"id"`
	Name         string               `json:"name"`
	LikerlandUrl string               `json:"likerland_url"`
	Image        string               `json:"image"`
	Status       ClassMigrationStatus `json:"status"`
	EnqueuedTime time.Time            `json:"enqueued_time,omitempty"`
	EVMTxHash    string               `json:"evm_tx_hash,omitempty"`
	FailedReason string               `json:"failed_reason,omitempty"`
}

type ClassMigrationStatus string

const (
	ClassMigrationStatusPending    ClassMigrationStatus = "pending"
	ClassMigrationStatusScheduled  ClassMigrationStatus = "scheduled"
	ClassMigrationStatusInProgress ClassMigrationStatus = "in_progress"
	ClassMigrationStatusSuccess    ClassMigrationStatus = "success"
	ClassMigrationStatusFailed     ClassMigrationStatus = "failed"
)

type NFT struct {
	ClassId      string             `json:"class_id"`
	Id           string             `json:"id"`
	Name         string             `json:"name"`
	LikerlandUrl string             `json:"likerland_url"`
	Image        string             `json:"image"`
	Status       NFTMigrationStatus `json:"status"`
	EnqueuedTime time.Time          `json:"enqueued_time,omitempty"`
	EVMTxHash    string             `json:"evm_tx_hash,omitempty"`
	FailedReason string             `json:"failed_reason,omitempty"`
}

type NFTMigrationStatus string

const (
	NFTMigrationStatusPending    NFTMigrationStatus = "pending"
	NFTMigrationStatusScheduled  NFTMigrationStatus = "scheduled"
	NFTMigrationStatusInProgress NFTMigrationStatus = "in_progress"
	NFTMigrationStatusSuccess    NFTMigrationStatus = "success"
	NFTMigrationStatusFailed     NFTMigrationStatus = "failed"
)
