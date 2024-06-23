package schema

type GameCredit struct {
	Base
	GameContract string `json:"game_contract"`
	Address      string `json:"address"`
	Credit       string `json:"credit"`
}

func (GameCredit) TableName() string {
	return "game_credit"
}
