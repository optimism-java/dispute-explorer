package schema

type GameLostBond struct {
	Base
	GameContract string `json:"game_contract"`
	Address      string `json:"address"`
	Bond         string `json:"bond"`
}

func (GameLostBond) TableName() string {
	return "game_lost_bonds"
}
