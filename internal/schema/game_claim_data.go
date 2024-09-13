package schema

const (
	GameClaimDataOnChainStatusValid    = "valid"
	GameClaimDataOnChainStatusRollBack = "rollback"
)

type GameClaimData struct {
	Base
	GameContract  string `json:"game_contract"`
	DataIndex     int64  `json:"data_index"`
	ParentIndex   uint32 `json:"parent_index"`
	CounteredBy   string `json:"countered_by"`
	Claimant      string `json:"claimant"`
	Bond          string `json:"bond"`
	Claim         string `json:"claim"`
	Position      string `json:"position"`
	Clock         string `json:"clock"`
	OutputBlock   uint64 `json:"output_block"`
	EventID       int64  `json:"event_id"`
	OnChainStatus string `json:"on_chain_status"`
}

func (GameClaimData) TableName() string {
	return "game_claim_data"
}
