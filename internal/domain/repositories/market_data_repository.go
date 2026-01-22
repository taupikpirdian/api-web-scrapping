package repositories

import (
	"api-web-scrapping/internal/domain/entities"
)

// MarketDataRepository defines the interface for market data operations
type MarketDataRepository interface {
	GetAll() ([]entities.MarketData, error)
	GetByEmiten(emiten string) ([]entities.MarketData, error)
	GetLatestByEmiten(emiten string) (*entities.MarketData, error)
	GetLatestByAllEmiten() ([]entities.MarketData, error)
}
