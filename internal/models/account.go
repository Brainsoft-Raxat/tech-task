package models

import "github.com/google/uuid"

type Account struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string  `db:"name" json:"name"`
	Balance   float64 `db:"balance" json:"balance"`
	CreatedAt string  `db:"created_at" json:"created_at"`
	UpdatedAt string  `db:"updated_at" json:"updated_at"`
}
