package schema

import (
	"slices"

	"likenft-indexer/ent/schema/typeutil"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// EVMEvent holds the schema definition for the EVMEvent entity.
type EVMEventProcessedBlockHeight struct {
	ent.Schema
}

// Fields of the EVMEvent.
func (EVMEventProcessedBlockHeight) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("contract_type").Values(
			"book_nft",
			"like_protocol",
		).Comment("For reference only"),
		field.String("contract_address").NotEmpty(),
		field.Enum("event").Values(
			"ContractURIUpdated",
			"NewBookNFT",
			"OwnershipTransferred",
			"TransferWithMemo",
			"Transfer",
		),
		field.Uint64("block_height").GoType(typeutil.Uint64(0)).
			SchemaType(typeutil.Uint64SchemaType).
			ValueScanner(typeutil.Uint64ValueScanner),
	}
}

// Edges of the EVMEventTransferWithMemo.
func (EVMEventProcessedBlockHeight) Edges() []ent.Edge {
	return nil
}

func (EVMEventProcessedBlockHeight) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("contract_type", "contract_address", "event").Unique(),
	}
}

func (EVMEventProcessedBlockHeight) Annotations() []schema.Annotation {
	return slices.Concat(
		typeutil.Uint64Annotations("block_height"),
	)
}
