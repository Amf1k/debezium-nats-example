package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID       uuid.UUID       `json:"id" db:"id"`
	Name     string          `json:"name" db:"name"`
	UPC      string          `json:"upc" db:"upc"`
	Price    decimal.Decimal `json:"price" db:"price"`
	Quantity uint            `json:"quantity" db:"quantity"`
}
