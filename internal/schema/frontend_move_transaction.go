package schema

const (
	FrontendMoveStatusPending   = "pending"
	FrontendMoveStatusConfirmed = "confirmed"
	FrontendMoveStatusFailed    = "failed"
)

// FrontendMoveTransaction records move transaction information initiated from frontend
type FrontendMoveTransaction struct {
	Base
	GameContract   string `json:"game_contract" gorm:"index:idx_game_contract"`             // Dispute Game contract address
	TxHash         string `json:"tx_hash" gorm:"type:varchar(128);uniqueIndex:idx_tx_hash"` // Transaction hash
	Claimant       string `json:"claimant" gorm:"type:varchar(128)"`                        // Initiator address
	ParentIndex    string `json:"parent_index"`                                             // Parent index
	Claim          string `json:"claim"`                                                    // Claim data
	IsAttack       bool   `json:"is_attack"`                                                // Whether it's an attack
	ChallengeIndex string `json:"challenge_index"`                                          // Challenge index
	DisputedClaim  string `json:"disputed_claim"`                                           // Disputed claim
	BlockNumber    int64  `json:"block_number"`                                             // Block number
	Status         string `json:"status" gorm:"type:varchar(20);default:pending"`           // Status
	ErrorMessage   string `json:"error_message,omitempty"`                                  // Error message (if any)
	SubmittedAt    int64  `json:"submitted_at"`                                             // Submission timestamp
	ConfirmedAt    int64  `json:"confirmed_at,omitempty"`                                   // Confirmation timestamp
}

func (FrontendMoveTransaction) TableName() string {
	return "frontend_move_transactions"
}
