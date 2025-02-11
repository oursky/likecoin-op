
-- +migrate Up

CREATE TABLE nft_migration
(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT now(),
  cosmos_address text NOT NULL,
  eth_address text NOT NULL,
  status text NOT NULL,
  UNIQUE (cosmos_address)
);

CREATE TABLE nft_migration_class
(
  id SERIAL PRIMARY KEY,
  nft_migration_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  cosmos_class_id text NOT NULL,
  name text NOT NULL,
  likerland_url text NOT NULL,
  image text NOT NULL,
  status text NOT NULL,
  enqueue_time TIMESTAMP NOT NULL,
  evm_tx_hash text NULL,
  failed_reason text NULL,
  CONSTRAINT fk_nft_migration
    FOREIGN KEY(nft_migration_id)
      REFERENCES nft_migration(id),
  UNIQUE (nft_migration_id, cosmos_class_id)
);

CREATE TABLE nft_migration_nft
(
  id SERIAL PRIMARY KEY,
  nft_migration_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  cosmos_class_id text NOT NULL,
  cosmos_nft_id text NOT NULL,
  name text NOT NULL,
  likerland_url text NOT NULL,
  image text NOT NULL,
  status text NOT NULL,
  enqueue_time TIMESTAMP NOT NULL,
  evm_tx_hash text NULL,
  failed_reason text NULL,
  CONSTRAINT fk_nft_migration
    FOREIGN KEY(nft_migration_id)
      REFERENCES nft_migration(id),
  UNIQUE (nft_migration_id, cosmos_class_id, cosmos_nft_id)
);

-- +migrate Down

DROP TABLE nft_migration_nft;
DROP TABLE nft_migration_class;
DROP TABLE nft_migration;
