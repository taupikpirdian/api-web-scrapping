package entities

import (
	"time"
)

// StockPriceSummary represents daily OHLC stock price data
type StockPriceSummary struct {
	ID         int64     `json:"id" db:"id"`
	EmitenID   int64     `json:"emiten_id" db:"emiten_id"`
	Date       time.Time `json:"date" db:"date"`
	OpenPrice  float64   `json:"open_price" db:"open_price"`
	HighPrice  float64   `json:"high_price" db:"high_price"`
	LowPrice   float64   `json:"low_price" db:"low_price"`
	ClosePrice float64   `json:"close_price" db:"close_price"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// StockPriceSummaryDetail represents stock price summary with emiten information
type StockPriceSummaryDetail struct {
	ID          int64     `json:"id" db:"id"`
	EmitenID    int64     `json:"emiten_id" db:"emiten_id"`
	Symbol      string    `json:"symbol" db:"symbol"`
	Name        string    `json:"name" db:"name"`
	Sector      string    `json:"sector" db:"sector"`
	Date        time.Time `json:"date" db:"date"`
	OpenPrice   float64   `json:"open_price" db:"open_price"`
	HighPrice   float64   `json:"high_price" db:"high_price"`
	LowPrice    float64   `json:"low_price" db:"low_price"`
	ClosePrice  float64   `json:"close_price" db:"close_price"`
	PriceChange float64   `json:"price_change" db:"price_change"`
	ChangePct   float64   `json:"change_percent" db:"change_percent"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
