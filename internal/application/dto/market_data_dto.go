package dto

import "time"

// MarketDataResponse represents the response for market data
type MarketDataResponse struct {
	ID         int64      `json:"id"`
	Emiten     string     `json:"emiten"`
	OpenPrice  float64    `json:"open_price"`
	HighPrice  float64    `json:"high_price"`
	LowPrice   float64    `json:"low_price"`
	ClosePrice float64    `json:"close_price"`
	Volume     int64      `json:"volume"`
	Value      float64    `json:"value"`
	Frequency  int        `json:"frequency"`
	Date       time.Time  `json:"date"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// MarketDataListResponse represents the list response for market data
type MarketDataListResponse struct {
	Data []MarketDataResponse `json:"data"`
}
