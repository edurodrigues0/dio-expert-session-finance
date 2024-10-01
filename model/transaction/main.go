package transaction

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type Transactions []Transaction
