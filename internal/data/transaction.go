package data

import (
	"github.com/Brainsoft-Raxat/tech-task/internal/models"
)

type CreateTransactionRequest struct {
	Value      float64 `json:"value" validate:"required,gt=0"`
	AccountID  string  `json:"account_id" validate:"required,uuid4"`
	GroupType  string  `json:"group_type" validate:"required,oneof=income outcome transfer"`
	Account2ID string  `json:"account2_id,omitempty" validate:"omitempty,uuid4"`
}

type CreateTransactionResponse struct {
	Transaction models.Transaction `json:"transaction"`
}

type GetAllTransactionsByAccountIDRequest struct {
	AccountID string `json:"account_id" validate:"required,uuid4"`
}

type GetAllTransactionsByAccountIDResponse struct {
	Transactions []models.Transaction `json:"transactions"`
}

type GetTransactionByIDRequest struct {
	ID string `json:"id" validate:"required,uuid4"`
}

type GetTransactionByIDResponse struct {
	Transaction models.Transaction `json:"transaction"`
}

// type UpdateTransactionRequest struct {
// 	ID         string  `json:"id"`
// 	Value      float64 `json:"value"`
// 	AccountID  string  `json:"account_id"`
// 	GroupType  string  `json:"group_type"`
// 	Account2ID string  `json:"account2_id,omitempty"`
// }

// type UpdateTransactionResponse struct {
// 	Transaction models.Transaction `json:"transaction"`
// }

type DeleteTransactionRequest struct {
	ID string `json:"id" validate:"required,uuid4"`
}

type DeleteTransactionResponse struct {
	// Transaction models.Transaction `json:"transaction"`
}
