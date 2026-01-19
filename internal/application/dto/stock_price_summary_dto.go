package dto

import "time"

// StockPriceSummaryResponse represents the response for stock price summary
type StockPriceSummaryResponse struct {
	ID         int64     `json:"id"`
	EmitenID   int64     `json:"emiten_id"`
	Date       time.Time `json:"date"`
	OpenPrice  float64   `json:"open_price"`
	HighPrice  float64   `json:"high_price"`
	LowPrice   float64   `json:"low_price"`
	ClosePrice float64   `json:"close_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// StockPriceSummaryDetailResponse represents the detailed response with emiten info
type StockPriceSummaryDetailResponse struct {
	ID          int64     `json:"id"`
	EmitenID    int64     `json:"emiten_id"`
	Symbol      string    `json:"symbol"`
	Name        string    `json:"name"`
	Sector      string    `json:"sector"`
	Date        time.Time `json:"date"`
	OpenPrice   float64   `json:"open_price"`
	HighPrice   float64   `json:"high_price"`
	LowPrice    float64   `json:"low_price"`
	ClosePrice  float64   `json:"close_price"`
	PriceChange float64   `json:"price_change"`
	ChangePct   float64   `json:"change_percent"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GetStockPriceSummaryRequest represents the request to get stock price summary
type GetStockPriceSummaryRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

// GetStockPriceBySymbolRequest represents the request to get stock price by symbol
type GetStockPriceBySymbolRequest struct {
	Symbol string `uri:"symbol" binding:"required"`
}

// GetStockPriceByDateRequest represents the request to get stock price by date
type GetStockPriceByDateRequest struct {
	Date string `uri:"date" binding:"required"`
}

// GetStockPriceByRangeRequest represents the request to get stock price by date range
type GetStockPriceByRangeRequest struct {
	StartDate string `form:"start_date" binding:"required"`
	EndDate   string `form:"end_date" binding:"required"`
	Page      int    `form:"page" binding:"min=1"`
	PageSize  int    `form:"page_size" binding:"min=1,max=100"`
}

// GetAllStockPricesRequest represents the request to get all stock prices with pagination
type GetAllStockPricesRequest struct {
	Page     int `form:"page" binding:"min=1"`
	PageSize int `form:"page_size" binding:"min=1,max=100"`
}

// CreateStockPriceSummaryRequest represents the request to create stock price summary
type CreateStockPriceSummaryRequest struct {
	EmitenID   int64     `json:"emiten_id" binding:"required"`
	Date       time.Time `json:"date" binding:"required"`
	OpenPrice  float64   `json:"open_price" binding:"required"`
	HighPrice  float64   `json:"high_price" binding:"required"`
	LowPrice   float64   `json:"low_price" binding:"required"`
	ClosePrice float64   `json:"close_price" binding:"required"`
}

// UpdateStockPriceSummaryRequest represents the request to update stock price summary
type UpdateStockPriceSummaryRequest struct {
	EmitenID   int64     `json:"emiten_id" binding:"required"`
	Date       time.Time `json:"date" binding:"required"`
	OpenPrice  float64   `json:"open_price" binding:"required"`
	HighPrice  float64   `json:"high_price" binding:"required"`
	LowPrice   float64   `json:"low_price" binding:"required"`
	ClosePrice float64   `json:"close_price" binding:"required"`
}

// StockPriceListResponse represents the paginated list response
type StockPriceListResponse struct {
	Data       []*StockPriceSummaryDetailResponse `json:"data"`
	Pagination *PaginationResponse                 `json:"pagination,omitempty"`
}

// TopMoversResponse represents the response for top gainers/losers
type TopMoversResponse struct {
	Date     time.Time                          `json:"date"`
	Gainers  []*StockPriceSummaryDetailResponse `json:"gainers,omitempty"`
	Losers   []*StockPriceSummaryDetailResponse `json:"losers,omitempty"`
}

// ErrorResponse represents error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// PaginationResponse represents pagination metadata
type PaginationResponse struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalCount int64 `json:"total_count"`
	TotalPages int   `json:"total_pages"`
}
