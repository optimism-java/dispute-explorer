CREATE DATABASE dispute_explorer WITH ENCODING ='UTF8';
-- Switch to the newly created database
\c dispute_explorer;

-- Create sync_blocks table
DROP TABLE if exists sync_blocks;
CREATE TABLE IF NOT EXISTS sync_blocks
(
    id           SERIAL PRIMARY KEY,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    blockchain   VARCHAR(32) NOT NULL,
    miner        VARCHAR(42) NOT NULL,
    block_time   BIGINT      NOT NULL,
    block_number BIGINT      NOT NULL,
    block_hash   VARCHAR(66) NOT NULL,
    tx_count     BIGINT      NOT NULL,
    event_count  BIGINT      NOT NULL,
    parent_hash  VARCHAR(66) NOT NULL,
    status       VARCHAR(32) NOT NULL,
    check_count  BIGINT      NOT NULL
);
CREATE INDEX if not exists status_index ON sync_blocks (status);
CREATE INDEX if not exists tx_count_index ON sync_blocks (tx_count);
CREATE INDEX if not exists check_count_index ON sync_blocks (check_count);

-- Create sync_events table
DROP TABLE if exists sync_events;
CREATE TABLE IF NOT EXISTS sync_events
(
    id                SERIAL PRIMARY KEY,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    sync_block_id     BIGINT      NOT NULL,
    blockchain        VARCHAR(32) NOT NULL,
    block_time        BIGINT      NOT NULL,
    block_number      BIGINT      NOT NULL,
    block_hash        VARCHAR(66) NOT NULL,
    block_log_indexed BIGINT      NOT NULL,
    tx_index          BIGINT      NOT NULL,
    tx_hash           VARCHAR(66) NOT NULL,
    event_name        VARCHAR(32) NOT NULL,
    event_hash        VARCHAR(66) NOT NULL,
    contract_address  VARCHAR(42) NOT NULL,
    data              JSONB       NOT NULL,
    status            VARCHAR(32) NOT NULL,
    retry_count       BIGINT               DEFAULT 0
);

-- ----------------------------
-- Table structure for dispute_game
-- ----------------------------
Drop table if exists dispute_game;
CREATE TABLE IF NOT EXISTS dispute_game
(
    id                SERIAL PRIMARY KEY,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    sync_block_id     BIGINT      NOT NULL,
    blockchain        VARCHAR(32) NOT NULL,
    block_time        BIGINT      NOT NULL,
    block_number      BIGINT      NOT NULL,
    block_hash        VARCHAR(66) NOT NULL,
    block_log_indexed BIGINT      NOT NULL,
    tx_index          BIGINT      NOT NULL,
    tx_hash           VARCHAR(66) NOT NULL,
    event_name        VARCHAR(32) NOT NULL,
    event_hash        VARCHAR(66) NOT NULL,
    contract_address  VARCHAR(42) NOT NULL,
    game_contract     varchar(42) NOT NULL,
    game_type         int         NOT NULL,
    l2_block_number   bigint      NOT NULL,
    status            int         NOT NULL,
    computed          boolean     default false
);
CREATE INDEX if not exists dispute_game_index ON dispute_game (contract_address, game_contract);

-- ----------------------------
-- Table structure for game_claim_data
-- ----------------------------
DROP TABLE if exists game_claim_data;
CREATE TABLE IF NOT EXISTS game_claim_data
(
    id                SERIAL PRIMARY KEY,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    game_contract     varchar(42)  NOT NULL,
    data_index        int          NOT NULL,
    parent_index      bigint       NOT NULL,
    countered_by      varchar(42)  NOT NULL,
    claimant          varchar(64)  NOT NULL,
    bond              bigint       NOT NULL,
    claim             varchar(64)  NOT NULL,
    position          bigint       NOT NULL,
    clock             bigint       NOT NULL,
    output_block      bigint       NOT NUll,
);
CREATE INDEX if not exists dispute_game_data_index ON game_claim_data (game_contract, data_index);

-- Table structure for game_credit
-- ----------------------------
DROP TABLE if exists game_credit;
CREATE TABLE IF NOT EXISTS game_credit
(
    id                SERIAL PRIMARY KEY,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    game_contract     varchar(42)  NOT NULL,
    address           varchar(64)  NOT NULL,
    credit            bigint       NOT NULL
);

