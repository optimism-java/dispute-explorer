Create Database If Not Exists dispute_explorer Character Set UTF8;
USE dispute_explorer;


SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
DROP TABLE IF EXISTS `dispute_game`;
CREATE TABLE `dispute_game`
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
    `game_contract`     varchar(42) NOT NULL COMMENT ' game_contract',
    `game_type`         int         NOT NULL COMMENT ' game_type',
    `l2_block_number`   bigint      NOT NULL COMMENT ' l2_block_number',
    `status`            varchar(32) NOT NULL COMMENT ' status -1-initial  0-In progress 1- Challenger wins 2- Defender wins',
    PRIMARY KEY (`id`),
    KEY `status_index` (`status`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1011299
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for game_claim_data
-- ----------------------------
DROP TABLE IF EXISTS `game_claim_data`;
CREATE TABLE `game_claim_data`
(
    `id`                bigint       NOT NULL AUTO_INCREMENT,
    `created_at`        datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `game_contract`     varchar(42)  NOT NULL COMMENT ' game_contract',
    `data_index`        int          NOT NULL COMMENT ' data_index',
    `parent_index`      bigint       NOT NULL COMMENT ' parent_index',
    `countered_by`      varchar(42)  NOT NULL COMMENT ' countered_by',
    `claimant`          varchar(64)  NOT NULL COMMENT ' claimant',
    `bond`              bigint       NOT NULL COMMENT ' bond',
    `claim`             varchar(64)  NOT NULL COMMENT ' claim',
    `position`          bigint       NOT NULL COMMENT ' position',
    `clock`             bigint       NOT NULL COMMENT ' clock',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1011299
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

