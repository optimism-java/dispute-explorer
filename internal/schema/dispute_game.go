package schema

// 0-In progress 1- Challenger wins 2- Defender wins

const (
	DisputeGameStatusInProgress    = 0
	DisputeGameStatusChallengerWin = 1
	DisputeGameStatusDefenderWin   = 2

	DisputeGameOnChainStatusValid    = "valid"
	DisputeGameOnChainStatusRollBack = "rollback"
)

type DisputeGame struct {
	Base
	SyncBlockID     int64  `json:"sync_block_id"`
	Blockchain      string `json:"blockchain"`
	BlockTime       int64  `json:"block_time"`
	BlockNumber     int64  `json:"block_number"`
	BlockHash       string `json:"block_hash"`
	BlockLogIndexed int64  `json:"block_log_indexed"`
	TxIndex         int64  `json:"tx_index"`
	TxHash          string `json:"tx_hash"`
	EventName       string `json:"event_name"`
	EventHash       string `json:"event_hash"`
	ContractAddress string `json:"contract_address"`
	GameContract    string `json:"game_contract"`
	GameType        uint32 `json:"game_type"`
	L2BlockNumber   int64  `json:"l_2_block_number"`
	Status          uint8  `json:"status"`
	Computed        bool   `json:"computed"`
	CalculateLost   bool   `json:"calculate_lost"`
	OnChainStatus   string `json:"on_chain_status"`
	ClaimDataLen    int64  `json:"claim_data_len"`
	GetLenStatus    bool   `json:"get_len_status"`
	HasFrontendMove bool   `json:"has_frontend_move" gorm:"default:false"` // Whether contains frontend-initiated move transactions
}

func (DisputeGame) TableName() string {
	return "dispute_game"
}
