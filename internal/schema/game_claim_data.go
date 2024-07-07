package schema

type GameClaimData struct {
	Base
	GameContract string `json:"game_contract"`
	DataIndex    int64  `json:"data_index"`
	ParentIndex  uint32 `json:"parent_index"`
	CounteredBy  string `json:"countered_by"`
	Claimant     string `json:"claimant"`
	Bond         uint64 `json:"bond"`
	Claim        string `json:"claim"`
	Position     uint64 `json:"position"`
	Clock        int64  `json:"clock"`
	OutputBlock  uint64 `json:"output_block"`
	EventID      int64  `json:"event_id"`
}

func (GameClaimData) TableName() string {
	return "game_claim_data"
}
