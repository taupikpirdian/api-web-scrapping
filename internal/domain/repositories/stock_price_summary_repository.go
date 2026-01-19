package repositories

import (
	"api-web-scrapping/internal/domain/entities"
	"time"
)

// StockPriceSummaryRepository defines the interface for stock price summary data operations
type StockPriceSummaryRepository interface {
	// GetByID retrieves a stock price summary by ID
	GetByID(id int64) (*entities.StockPriceSummary, error)

	// GetByEmitenIDAndDate retrieves a stock price summary by emiten ID and date
	GetByEmitenIDAndDate(emitenID int64, date time.Time) (*entities.StockPriceSummary, error)

	// GetBySymbolAndDate retrieves a stock price summary by symbol and date
	GetBySymbolAndDate(symbol string, date time.Time) (*entities.StockPriceSummaryDetail, error)

	// GetByEmitenID retrieves all stock price summaries for an emiten
	GetByEmitenID(emitenID int64, limit, offset int) ([]*entities.StockPriceSummary, error)

	// GetBySymbol retrieves all stock price summaries for a symbol with details
	GetBySymbol(symbol string, limit, offset int) ([]*entities.StockPriceSummaryDetail, error)

	// GetByDateRange retrieves summaries within a date range
	GetByDateRange(startDate, endDate time.Time, limit, offset int) ([]*entities.StockPriceSummaryDetail, error)

	// GetBySymbolAndDateRange retrieves summaries for a symbol within a date range
	GetBySymbolAndDateRange(symbol string, startDate, endDate time.Time, limit, offset int) ([]*entities.StockPriceSummaryDetail, error)

	// GetLatestBySymbol retrieves the latest summary for a symbol
	GetLatestBySymbol(symbol string) (*entities.StockPriceSummaryDetail, error)

	// GetLatestByEmitenID retrieves the latest summary for an emiten ID
	GetLatestByEmitenID(emitenID int64) (*entities.StockPriceSummary, error)

	// GetAll retrieves all stock price summaries with pagination
	GetAll(limit, offset int) ([]*entities.StockPriceSummaryDetail, error)

	// GetTopGainers retrieves top gainers for a specific date
	GetTopGainers(date time.Time, limit int) ([]*entities.StockPriceSummaryDetail, error)

	// GetTopLosers retrieves top losers for a specific date
	GetTopLosers(date time.Time, limit int) ([]*entities.StockPriceSummaryDetail, error)

	// Create creates a new stock price summary
	Create(summary *entities.StockPriceSummary) error

	// Update updates an existing stock price summary
	Update(summary *entities.StockPriceSummary) error

	// Upsert inserts or updates a stock price summary
	Upsert(summary *entities.StockPriceSummary) error

	// Delete deletes a stock price summary
	Delete(id int64) error

	// Count returns the total count of stock price summaries
	Count() (int64, error)

	// CountByEmitenID returns the count of summaries for an emiten
	CountByEmitenID(emitenID int64) (int64, error)
}
