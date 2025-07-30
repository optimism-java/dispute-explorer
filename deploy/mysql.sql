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
    `created_at`   datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `blockchain`   varchar(32) NOT NULL,
    `miner`        varchar(42) NOT NULL,
    `block_time`   bigint      NOT NULL,
    `block_number` bigint      NOT NULL,
    `block_hash`   varchar(66) NOT NULL,
    `tx_count`     bigint      NOT NULL,
    `event_count`  bigint      NOT NULL,
    `parent_hash`  varchar(66) NOT NULL,
    `status`       varchar(32) NOT NULL,
    `check_count`  bigint      NOT NULL,
    PRIMARY KEY (`id`),
    KEY `tx_count` (`tx_count`),
    KEY `status_index` (`status`),
    KEY `check_count` (`check_count`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 3343515
  DEFAULT CHARSET = utf8mb3;
-- ----------------------------
-- Table structure for sync_events
-- ----------------------------
DROP TABLE IF EXISTS `sync_events`;
CREATE TABLE `sync_events`
(
    `id`                bigint       NOT NULL AUTO_INCREMENT,
    `created_at`        datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `sync_block_id`     bigint       NOT NULL,
    `blockchain`        varchar(32)  NOT NULL,
    `block_time`        bigint       NOT NULL,
    `block_number`      bigint       NOT NULL,
    `block_hash`        varchar(66)  NOT NULL,
    `block_log_indexed` bigint       NOT NULL,
    `tx_index`          bigint       NOT NULL,
    `tx_hash`           varchar(66)  NOT NULL,
    `event_name`        varchar(32)  NOT NULL,
    `event_hash`        varchar(66)  NOT NULL,
    `contract_address`  varchar(42)  NOT NULL,
    `data`              varchar(256) NOT NULL,
    `status`            varchar(32)  NOT NULL,
    `retry_count`       bigint       NOT NULL,
    PRIMARY KEY (`id`),
    KEY `status_index` (`status`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 71494
  DEFAULT CHARSET = utf8mb3;


-- ----------------------------
-- Table structure for dispute_game
-- ----------------------------
DROP TABLE IF EXISTS `dispute_game`;
CREATE TABLE `dispute_game`
(
    `id`                bigint      NOT NULL AUTO_INCREMENT,
    `created_at`        datetime             DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        datetime             DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `sync_block_id`     bigint      NOT NULL,
    `blockchain`        varchar(32) NOT NULL,
    `block_time`        bigint      NOT NULL,
    `block_number`      bigint      NOT NULL,
    `block_hash`        varchar(66) NOT NULL,
    `block_log_indexed` bigint      NOT NULL,
    `tx_index`          bigint      NOT NULL,
    `tx_hash`           varchar(66) NOT NULL,
    `event_name`        varchar(32) NOT NULL,
    `event_hash`        varchar(66) NOT NULL,
    `contract_address`  varchar(42) NOT NULL,
    `game_contract`     varchar(42) NOT NULL,
    `game_type`         int         NOT NULL,
    `l2_block_number`   bigint      NOT NULL,
    `status`            tinyint     NOT NULL,
    `computed`          tinyint(1)  NOT NULL DEFAULT '0',
    `calculate_lost`    tinyint(1)  NOT NULL DEFAULT '0',
    `on_chain_status`   varchar(32) NOT NULL DEFAULT 'valid',
    `claim_data_len`    bigint      NOT NULL DEFAULT '1',
    `get_len_status`    tinyint(1)  NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY `status_index` (`status`),
    KEY `dispute_game_index` (`contract_address`, `game_contract`),
    KEY `dispute_on_chain_status_index` (`on_chain_status`),
    KEY `dispute_claim_data_len_index` (`claim_data_len`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 35578
  DEFAULT CHARSET = utf8mb3;

-- ----------------------------
-- Table structure for game_claim_data
-- ----------------------------
DROP TABLE IF EXISTS `game_claim_data`;
CREATE TABLE `game_claim_data`
(
    `id`              bigint       NOT NULL AUTO_INCREMENT,
    `created_at`      datetime              DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      datetime              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `game_contract`   varchar(42)  NOT NULL,
    `data_index`      bigint       NOT NULL,
    `parent_index`    bigint       NOT NULL,
    `countered_by`    varchar(42)  NOT NULL,
    `claimant`        varchar(64)  NOT NULL,
    `bond`            varchar(128) NOT NULL,
    `claim`           varchar(64)  NOT NULL,
    `position`        varchar(128) NOT NULL,
    `clock`           varchar(128) NOT NULL,
    `output_block`    bigint       NOT NULL,
    `event_id`        bigint       NOT NULL,
    `on_chain_status` varchar(32)  NOT NULL DEFAULT 'valid',
    PRIMARY KEY (`id`),
    KEY `claim_on_chain_status_index` (`on_chain_status`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 36273
  DEFAULT CHARSET = utf8mb3;

-- ----------------------------
-- Table structure for game_credit
-- ----------------------------
DROP TABLE IF EXISTS `game_credit`;
CREATE TABLE `game_credit`
(
    `id`            bigint      NOT NULL AUTO_INCREMENT,
    `created_at`    datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `game_contract` varchar(42) NOT NULL,
    `address`       varchar(64) NOT NULL,
    `credit`        bigint      NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 36342
  DEFAULT CHARSET = utf8mb3;

-- ----------------------------
-- Table structure for game_credit
-- ----------------------------
DROP TABLE IF EXISTS `game_lost_bonds`;
CREATE TABLE `game_lost_bonds`
(
    `id`            bigint       NOT NULL AUTO_INCREMENT,
    `created_at`    datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `game_contract` varchar(42)  NOT NULL,
    `address`       varchar(64)  NOT NULL,
    `bond`          varchar(128) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 14080
  DEFAULT CHARSET = utf8mb3;
