package persistence

import (
	"api-web-scrapping/internal/domain/entities"
	"api-web-scrapping/internal/domain/repositories"
	"database/sql"
)

type marketDataRepositoryImpl struct {
	db *sql.DB
}

// NewMarketDataRepository creates a new market data repository
func NewMarketDataRepository(db *sql.DB) repositories.MarketDataRepository {
	return &marketDataRepositoryImpl{db: db}
}

// GetAll retrieves all market data from the view
func (r *marketDataRepositoryImpl) GetAll() ([]entities.MarketData, error) {
	query := `
		SELECT id, emiten, open_price, high_price, low_price, close_price,
		       volume, value, frequency, date, created_at, updated_at, deleted_at
		FROM v_latest_market_data
		ORDER BY emiten
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var marketDataList []entities.MarketData
	for rows.Next() {
		var md entities.MarketData
		err := rows.Scan(
			&md.ID,
			&md.Emiten,
			&md.OpenPrice,
			&md.HighPrice,
			&md.LowPrice,
			&md.ClosePrice,
			&md.Volume,
			&md.Value,
			&md.Frequency,
			&md.Date,
			&md.CreatedAt,
			&md.UpdatedAt,
			&md.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		marketDataList = append(marketDataList, md)
	}

	return marketDataList, nil
}

// GetByEmiten retrieves market data for a specific emiten
func (r *marketDataRepositoryImpl) GetByEmiten(emiten string) ([]entities.MarketData, error) {
	query := `
		SELECT id, emiten, open_price, high_price, low_price, close_price,
		       volume, value, frequency, date, created_at, updated_at, deleted_at
		FROM v_latest_market_data
		WHERE emiten = $1
		ORDER BY date DESC
	`

	rows, err := r.db.Query(query, emiten)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var marketDataList []entities.MarketData
	for rows.Next() {
		var md entities.MarketData
		err := rows.Scan(
			&md.ID,
			&md.Emiten,
			&md.OpenPrice,
			&md.HighPrice,
			&md.LowPrice,
			&md.ClosePrice,
			&md.Volume,
			&md.Value,
			&md.Frequency,
			&md.Date,
			&md.CreatedAt,
			&md.UpdatedAt,
			&md.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		marketDataList = append(marketDataList, md)
	}

	return marketDataList, nil
}

// GetLatestByEmiten retrieves the latest market data for a specific emiten
func (r *marketDataRepositoryImpl) GetLatestByEmiten(emiten string) (*entities.MarketData, error) {
	query := `
		SELECT id, emiten, open_price, high_price, low_price, close_price,
		       volume, value, frequency, date, created_at, updated_at, deleted_at
		FROM v_latest_market_data
		WHERE emiten = $1
		LIMIT 1
	`

	var md entities.MarketData
	err := r.db.QueryRow(query, emiten).Scan(
		&md.ID,
		&md.Emiten,
		&md.OpenPrice,
		&md.HighPrice,
		&md.LowPrice,
		&md.ClosePrice,
		&md.Volume,
		&md.Value,
		&md.Frequency,
		&md.Date,
		&md.CreatedAt,
		&md.UpdatedAt,
		&md.DeletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &md, nil
}

// GetLatestByAllEmiten retrieves the latest market data for all emitens
func (r *marketDataRepositoryImpl) GetLatestByAllEmiten() ([]entities.MarketData, error) {
	query := `
		SELECT id, emiten, open_price, high_price, low_price, close_price,
		       volume, value, frequency, date, created_at, updated_at, deleted_at
		FROM v_latest_market_data
		ORDER BY emiten
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var marketDataList []entities.MarketData
	for rows.Next() {
		var md entities.MarketData
		err := rows.Scan(
			&md.ID,
			&md.Emiten,
			&md.OpenPrice,
			&md.HighPrice,
			&md.LowPrice,
			&md.ClosePrice,
			&md.Volume,
			&md.Value,
			&md.Frequency,
			&md.Date,
			&md.CreatedAt,
			&md.UpdatedAt,
			&md.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		marketDataList = append(marketDataList, md)
	}

	return marketDataList, nil
}
