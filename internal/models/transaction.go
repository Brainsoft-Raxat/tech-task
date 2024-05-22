package models

import "github.com/google/uuid"

type Transaction struct {
	ID         uuid.UUID `db:"id" json:"id"`
	Value      float64   `db:"value" json:"value"`
	AccountID  uuid.UUID `db:"account_id" json:"account_id"`
	GroupType  string    `db:"group_type" json:"group_type"`
	Account2ID uuid.UUID `db:"account2_id,omitempty" json:"account2_id,omitempty"`
	CreatedAt  string    `db:"created_at" json:"created_at"`
	UpdatedAt  string    `db:"updated_at" json:"updated_at"`
}

const (
	GroupTypeIncome   = "income"
	GroupTypeOutcome  = "outcome"
	GroupTypeTransfer = "transfer"
)
