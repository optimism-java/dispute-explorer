package v0

import "github.com/optimism-java/dispute-explorer/migration"

type SyncBlock struct {
	migration.Base
	Blockchain  string `json:"blockchain" gorm:"type:varchar(32);notnull"`
	Miner       string `json:"miner" gorm:"type:varchar(42);notnull"`
	BlockTime   int64  `json:"block_time" gorm:"type:bigint;notnull"`
	BlockNumber int64  `json:"block_number" gorm:"type:bigint;notnull"`
	BlockHash   string `json:"block_hash" gorm:"type:varchar(66);notnull"`
	TxCount     int64  `json:"tx_count" gorm:"type:bigint;notnull;index:tx_count"`
	EventCount  int64  `json:"event_count" gorm:"type:bigint;notnull"`
	ParentHash  string `json:"parent_hash" gorm:"type:varchar(66);notnull"`
	Status      string `json:"status" gorm:"type:varchar(32);notnull;index:status_index"`
	CheckCount  int64  `json:"check_count" gorm:"type:bigint;notnull;index:check_count"`
}

func (SyncBlock) TableName() string {
	return "sync_blocks"
}

type SyncEvent struct {
	migration.Base
	SyncBlockID     int64  `json:"sync_block_id" gorm:"type:bigint;notnull"`
	Blockchain      string `json:"blockchain" gorm:"type:varchar(32);notnull"`
	BlockTime       int64  `json:"block_time" gorm:"type:bigint;notnull"`
	BlockNumber     int64  `json:"block_number" gorm:"type:bigint;notnull"`
	BlockHash       string `json:"block_hash" gorm:"type:varchar(66);notnull"`
	BlockLogIndexed int64  `json:"block_log_indexed" gorm:"type:bigint;notnull"`
	TxIndex         int64  `json:"tx_index" gorm:"type:bigint;notnull"`
	TxHash          string `json:"tx_hash" gorm:"type:varchar(66);notnull"`
	EventName       string `json:"event_name" gorm:"type:varchar(32);notnull"`
	EventHash       string `json:"event_hash" gorm:"type:varchar(66);notnull"`
	ContractAddress string `json:"contract_address" gorm:"type:varchar(42);notnull"`
	Data            string `json:"data" gorm:"type:json;notnull"`
	Status          string `json:"status" gorm:"type:varchar(32);notnull;index:status_index"`
	RetryCount      int64  `json:"retry_count" gorm:"type:bigint;notnull"`
}

func (SyncEvent) TableName() string {
	return "sync_events"
}

type DisputeGame struct {
	migration.Base
	SyncBlockID     int64  `json:"sync_block_id" gorm:"type:bigint;notnull"`
	Blockchain      string `json:"blockchain" gorm:"type:varchar(32);notnull"`
	BlockTime       int64  `json:"block_time" gorm:"type:bigint;notnull"`
	BlockNumber     int64  `json:"block_number" gorm:"type:bigint;notnull"`
	BlockHash       string `json:"block_hash" gorm:"type:varchar(66);notnull"`
	BlockLogIndexed int64  `json:"block_log_indexed" gorm:"type:bigint;notnull"`
	TxIndex         int64  `json:"tx_index" gorm:"type:bigint;notnull"`
	TxHash          string `json:"tx_hash" gorm:"type:varchar(66);notnull"`
	EventName       string `json:"event_name" gorm:"type:varchar(32);notnull"`
	EventHash       string `json:"event_hash" gorm:"type:varchar(66);notnull"`
	ContractAddress string `json:"contract_address" gorm:"type:varchar(42);notnull;index:dispute_game_index"`
	GameContract    string `json:"game_contract" gorm:"type:varchar(42);notnull;index:dispute_game_index"`
	GameType        uint32 `json:"game_type" gorm:"type:int;notnull"`
	L2BlockNumber   int64  `json:"l_2_block_number" gorm:"type:bigint;notnull"`
	Status          uint8  `json:"status" gorm:"type:int;notnull;index:status_index"`
	Computed        bool   `json:"computed" gorm:"type:tinyint(1);notnull;default:0"`
}

func (DisputeGame) TableName() string {
	return "dispute_game"
}

type GameClaimData struct {
	migration.Base
	GameContract string `json:"game_contract" gorm:"type:varchar(42);notnull"`
	DataIndex    int64  `json:"data_index" gorm:"type:int;notnull"`
	ParentIndex  uint32 `json:"parent_index" gorm:"type:bigint;notnull"`
	CounteredBy  string `json:"countered_by" gorm:"type:varchar(42);notnull"`
	Claimant     string `json:"claimant" gorm:"type:varchar(64);notnull"`
	Bond         uint64 `json:"bond" gorm:"type:bigint;notnull"`
	Claim        string `json:"claim" gorm:"type:varchar(64);notnull"`
	Position     uint64 `json:"position" gorm:"type:bigint;notnull"`
	Clock        int64  `json:"clock" gorm:"type:bigint;notnull"`
	OutputBlock  uint64 `json:"output_block" gorm:"type:bigint;notnull"`
	EventID      int64  `json:"event_id" gorm:"type:bigint;notnull"`
}

func (GameClaimData) TableName() string {
	return "game_claim_data"
}

type GameCredit struct {
	migration.Base
	GameContract string `json:"game_contract" gorm:"type:varchar(42);notnull"`
	Address      string `json:"address" gorm:"type:varchar(64);notnull"`
	Credit       string `json:"credit" gorm:"type:bigint;notnull"`
}

func (GameCredit) TableName() string {
	return "game_credit"
}

var ModelSchemaList = []interface{}{SyncBlock{}, SyncEvent{}, DisputeGame{}, GameClaimData{}, GameCredit{}}
