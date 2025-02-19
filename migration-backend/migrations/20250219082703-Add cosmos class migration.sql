
-- +migrate Up
CREATE TABLE nft_cosmos_class_migration
(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT now(),
  cosmos_class_id text NOT NULL,
  evm_class_id text NULL,
  evm_tx_hash text NULL,
  status text NOT NULL,
  failed_reason text NULL,
  UNIQUE (cosmos_class_id)
);

-- +migrate Down

DROP TABLE nft_cosmos_class_migration;
