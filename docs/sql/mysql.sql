Create Database If Not Exists dispute_explorer Character Set UTF8;
USE dispute_explorer;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

SET GLOBAL binlog_format = 'ROW';

-- ----------------------------
-- Table structure for sync_blocks
-- ----------------------------
DROP TABLE IF EXISTS `sync_blocks`;
CREATE TABLE `sync_blocks`
(
    `id`           bigint      NOT NULL AUTO_INCREMENT,
    `created_at`   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `blockchain`   varchar(32) NOT NULL COMMENT ' chain name',
    `miner`        varchar(42) NOT NULL COMMENT ' miner',
    `block_time`   bigint      NOT NULL COMMENT ' block_time',
    `block_number` bigint      NOT NULL COMMENT ' block_number',
    `block_hash`   varchar(66) NOT NULL COMMENT ' block hash',
    `tx_count`     bigint      NOT NULL COMMENT ' tx count',
    `event_count`  bigint      NOT NULL COMMENT ' event count',
    `parent_hash`  varchar(66) NOT NULL COMMENT ' parent hash',
    `status`       varchar(32) NOT NULL COMMENT ' status',
    `check_count`  bigint      NOT NULL COMMENT ' check count',
    PRIMARY KEY (`id`),
    KEY `status_index` (`status`),
    KEY `tx_count_index` (`tx_count`),
    KEY `check_count_index` (`check_count`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 2923365
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for sync_events
-- ----------------------------
DROP TABLE IF EXISTS `sync_events`;
CREATE TABLE `sync_events`
(
    `id`                bigint      NOT NULL AUTO_INCREMENT,
    `created_at`        datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `sync_block_id`     bigint      NOT NULL COMMENT ' sync_block_id',
    `blockchain`        varchar(32) NOT NULL COMMENT ' blockchain',
    `block_time`        bigint      NOT NULL COMMENT ' block_time',
    `block_number`      bigint      NOT NULL COMMENT ' block_number',
    `block_hash`        varchar(66) NOT NULL COMMENT ' block_hash',
    `block_log_indexed` bigint      NOT NULL COMMENT ' block_log_indexed',
    `tx_index`          bigint      NOT NULL COMMENT ' tx_index',
    `tx_hash`           varchar(66) NOT NULL COMMENT ' tx_hash',
    `event_name`        varchar(32) NOT NULL COMMENT ' event_name',
    `event_hash`        varchar(66) NOT NULL COMMENT ' event_hash',
    `contract_address`  varchar(42) NOT NULL COMMENT ' contract_address',
    `data`              json        NOT NULL COMMENT ' data',
    `status`            varchar(32) NOT NULL COMMENT ' status',
    `retry_count`       bigint               DEFAULT '0' COMMENT 'retry_count',
    PRIMARY KEY (`id`),
    KEY `status_index` (`status`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1011299
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;


-- ----------------------------
-- Table structure for dispute_game
-- ----------------------------
DROP TABLE IF EXISTS dispute_game;
CREATE TABLE IF NOT EXISTS dispute_game
(
    `id`                bigint      NOT NULL AUTO_INCREMENT,
    `created_at`        datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `sync_block_id`     bigint      NOT NULL COMMENT ' sync_block_id',
    `blockchain`        varchar(32) NOT NULL COMMENT ' blockchain',
    `block_time`        bigint      NOT NULL COMMENT ' block_time',
    `block_number`      bigint      NOT NULL COMMENT ' block_number',
    `block_hash`        varchar(66) NOT NULL COMMENT ' block_hash',
    `block_log_indexed` bigint      NOT NULL COMMENT ' block_log_indexed',
    `tx_index`          bigint      NOT NULL COMMENT ' tx_index',
    `tx_hash`           varchar(66) NOT NULL COMMENT ' tx_hash',
    `event_name`        varchar(32) NOT NULL COMMENT ' event_name',
    `event_hash`        varchar(66) NOT NULL COMMENT ' event_hash',
    `contract_address`  varchar(42) NOT NULL COMMENT ' contract_address',
    `game_contract`     varchar(42) NOT NULL,
    `game_type`         int         NOT NULL,
    `l2_block_number`   bigint      NOT NULL,
    `status`            int         NOT NULL,
    `computed`          tinyint(1)  NOT NULL DEFAULT 0 COMMENT ' 1-already get game credit 0- not yet',
    PRIMARY KEY (`id`),
    KEY `status_index` (`status`),
    KEY `dispute_game_index` (`contract_address`, `game_contract`)
);

-- ----------------------------
-- Table structure for game_claim_data
-- ----------------------------
DROP TABLE IF EXISTS game_claim_data;
CREATE TABLE IF NOT EXISTS game_claim_data
(
    `id`                bigint       NOT NULL AUTO_INCREMENT,
    `created_at`        datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `game_contract`     varchar(42)  NOT NULL,
    `data_index`        int          NOT NULL,
    `parent_index`      bigint       NOT NULL,
    `countered_by`      varchar(42)  NOT NULL,
    `claimant`          varchar(64)  NOT NULL,
    `bond`              bigint       NOT NULL,
    `claim`             varchar(64)  NOT NULL,
    `position`          bigint       NOT NULL,
    `clock`             bigint       NOT NULL,
    `output_block`      bigint       NOT NULL,
    `event_id`          bigint       NOT NULL,
    PRIMARY KEY (`id`),
    KEY `credit_index` (`game_contract`, `data_index`)
);

-- ----------------------------
-- Table structure for game_credit
-- ----------------------------
DROP TABLE IF EXISTS game_credit;
CREATE TABLE IF NOT EXISTS game_credit
(
    `id`                bigint      NOT NULL AUTO_INCREMENT,
    `created_at`        datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `game_contract`     varchar(42) NOT NULL,
    `address`           varchar(64) NOT NULL,
    `credit`            numeric     NOT NULL,
    PRIMARY KEY (`id`)
)
