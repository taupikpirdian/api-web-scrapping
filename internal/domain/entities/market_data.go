package entities

import (
	"time"
)

// MarketData represents market data from v_latest_market_data view
type MarketData struct {
	ID             int64      `json:"id" db:"id"`
	Emiten         string     `json:"emiten" db:"emiten"`
	OpenPrice      float64    `json:"open_price" db:"open_price"`
	HighPrice      float64    `json:"high_price" db:"high_price"`
	LowPrice       float64    `json:"low_price" db:"low_price"`
	ClosePrice     float64    `json:"close_price" db:"last_price"`
	Date           time.Time  `json:"date" db:"date_time_scraping"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
