package persistence

import (
	"api-web-scrapping/internal/domain/entities"
	"api-web-scrapping/internal/domain/repositories"
	"database/sql"
	"fmt"
	"time"
)

type stockPriceSummaryRepositoryImpl struct {
	db *sql.DB
}

// NewStockPriceSummaryRepository creates a new stock price summary repository
func NewStockPriceSummaryRepository(db *sql.DB) repositories.StockPriceSummaryRepository {
	return &stockPriceSummaryRepositoryImpl{db: db}
}

func (r *stockPriceSummaryRepositoryImpl) GetByID(id int64) (*entities.StockPriceSummary, error) {
	query := `
		SELECT id, emiten_id, date, open_price, high_price, low_price, close_price, created_at, updated_at
		FROM stock_price_summary
		WHERE id = $1
	`

	var summary entities.StockPriceSummary
	err := r.db.QueryRow(query, id).Scan(
		&summary.ID,
		&summary.EmitenID,
		&summary.Date,
		&summary.OpenPrice,
		&summary.HighPrice,
		&summary.LowPrice,
		&summary.ClosePrice,
		&summary.CreatedAt,
		&summary.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *stockPriceSummaryRepositoryImpl) GetByEmitenIDAndDate(emitenID int64, date time.Time) (*entities.StockPriceSummary, error) {
	query := `
		SELECT id, emiten_id, date, open_price, high_price, low_price, close_price, created_at, updated_at
		FROM stock_price_summary
		WHERE emiten_id = $1 AND date = $2
	`

	var summary entities.StockPriceSummary
	err := r.db.QueryRow(query, emitenID, date).Scan(
		&summary.ID,
		&summary.EmitenID,
		&summary.Date,
		&summary.OpenPrice,
		&summary.HighPrice,
		&summary.LowPrice,
		&summary.ClosePrice,
		&summary.CreatedAt,
		&summary.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *stockPriceSummaryRepositoryImpl) GetBySymbolAndDate(symbol string, date time.Time) (*entities.StockPriceSummaryDetail, error) {
	query := `
		SELECT sps.id, sps.emiten_id, e.symbol, e.name, e.sector,
		       sps.date, sps.open_price, sps.high_price, sps.low_price, sps.close_price,
		       (sps.close_price - sps.open_price) as price_change,
		       CASE WHEN sps.open_price > 0
		            THEN ROUND(((sps.close_price - sps.open_price) / sps.open_price * 100)::NUMERIC, 2)
		            ELSE 0
		       END as change_percent,
		       sps.created_at, sps.updated_at
		FROM stock_price_summary sps
		INNER JOIN emitens e ON sps.emiten_id = e.id
		WHERE e.symbol = $1 AND sps.date = $2
	`

	var summary entities.StockPriceSummaryDetail
	err := r.db.QueryRow(query, symbol, date).Scan(
		&summary.ID,
		&summary.EmitenID,
		&summary.Symbol,
		&summary.Name,
		&summary.Sector,
		&summary.Date,
		&summary.OpenPrice,
		&summary.HighPrice,
		&summary.LowPrice,
		&summary.ClosePrice,
		&summary.PriceChange,
		&summary.ChangePct,
		&summary.CreatedAt,
		&summary.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *stockPriceSummaryRepositoryImpl) GetByEmitenID(emitenID int64, limit, offset int) ([]*entities.StockPriceSummary, error) {
	query := `
		SELECT id, emiten_id, date, open_price, high_price, low_price, close_price, created_at, updated_at
		FROM stock_price_summary
		WHERE emiten_id = $1
		ORDER BY date DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.Query(query, emitenID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summaries []*entities.StockPriceSummary
	for rows.Next() {
		var summary entities.StockPriceSummary
		err := rows.Scan(
			&summary.ID,
			&summary.EmitenID,
			&summary.Date,
			&summary.OpenPrice,
			&summary.HighPrice,
			&summary.LowPrice,
			&summary.ClosePrice,
			&summary.CreatedAt,
			&summary.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, &summary)
	}

	return summaries, nil
}

func (r *stockPriceSummaryRepositoryImpl) GetBySymbol(symbol string, limit, offset int) ([]*entities.StockPriceSummaryDetail, error) {
	query := `
		SELECT sps.id, sps.emiten_id, e.symbol, e.name, e.sector,
		       sps.date, sps.open_price, sps.high_price, sps.low_price, sps.close_price,
		       (sps.close_price - sps.open_price) as price_change,
		       CASE WHEN sps.open_price > 0
		            THEN ROUND(((sps.close_price - sps.open_price) / sps.open_price * 100)::NUMERIC, 2)
		            ELSE 0
		       END as change_percent,
		       sps.created_at, sps.updated_at
		FROM stock_price_summary sps
		INNER JOIN emitens e ON sps.emiten_id = e.id
		WHERE e.symbol = $1
		ORDER BY sps.date DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.Query(query, symbol, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summaries []*entities.StockPriceSummaryDetail
	for rows.Next() {
		var summary entities.StockPriceSummaryDetail
		err := rows.Scan(
			&summary.ID,
			&summary.EmitenID,
			&summary.Symbol,
			&summary.Name,
			&summary.Sector,
			&summary.Date,
			&summary.OpenPrice,
			&summary.HighPrice,
			&summary.LowPrice,
			&summary.ClosePrice,
			&summary.PriceChange,
			&summary.ChangePct,
			&summary.CreatedAt,
			&summary.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, &summary)
	}

	return summaries, nil
}

func (r *stockPriceSummaryRepositoryImpl) GetByDateRange(startDate, endDate time.Time, limit, offset int) ([]*entities.StockPriceSummaryDetail, error) {
	query := `
		SELECT sps.id, sps.emiten_id, e.symbol, e.name, e.sector,
		       sps.date, sps.open_price, sps.high_price, sps.low_price, sps.close_price,
		       (sps.close_price - sps.open_price) as price_change,
		       CASE WHEN sps.open_price > 0
		            THEN ROUND(((sps.close_price - sps.open_price) / sps.open_price * 100)::NUMERIC, 2)
		            ELSE 0
		       END as change_percent,
		       sps.created_at, sps.updated_at
		FROM stock_price_summary sps
		INNER JOIN emitens e ON sps.emiten_id = e.id
		WHERE sps.date >= $1 AND sps.date <= $2
		ORDER BY sps.date DESC, e.symbol
		LIMIT $3 OFFSET $4
	`

	rows, err := r.db.Query(query, startDate, endDate, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summaries []*entities.StockPriceSummaryDetail
	for rows.Next() {
		var summary entities.StockPriceSummaryDetail
		err := rows.Scan(
			&summary.ID,
			&summary.EmitenID,
			&summary.Symbol,
			&summary.Name,
			&summary.Sector,
			&summary.Date,
			&summary.OpenPrice,
			&summary.HighPrice,
			&summary.LowPrice,
			&summary.ClosePrice,
			&summary.PriceChange,
			&summary.ChangePct,
			&summary.CreatedAt,
			&summary.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, &summary)
	}

	return summaries, nil
}

func (r *stockPriceSummaryRepositoryImpl) GetBySymbolAndDateRange(symbol string, startDate, endDate time.Time, limit, offset int) ([]*entities.StockPriceSummaryDetail, error) {
	query := `
		SELECT sps.id, sps.emiten_id, e.symbol, e.name, e.sector,
		       sps.date, sps.open_price, sps.high_price, sps.low_price, sps.close_price,
		       (sps.close_price - sps.open_price) as price_change,
		       CASE WHEN sps.open_price > 0
		            THEN ROUND(((sps.close_price - sps.open_price) / sps.open_price * 100)::NUMERIC, 2)
		            ELSE 0
		       END as change_percent,
		       sps.created_at, sps.updated_at
		FROM stock_price_summary sps
		INNER JOIN emitens e ON sps.emiten_id = e.id
		WHERE e.symbol = $1 AND sps.date >= $2 AND sps.date <= $3
		ORDER BY sps.date DESC
		LIMIT $4 OFFSET $5
	`

	rows, err := r.db.Query(query, symbol, startDate, endDate, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summaries []*entities.StockPriceSummaryDetail
	for rows.Next() {
		var summary entities.StockPriceSummaryDetail
		err := rows.Scan(
			&summary.ID,
			&summary.EmitenID,
			&summary.Symbol,
			&summary.Name,
			&summary.Sector,
			&summary.Date,
			&summary.OpenPrice,
			&summary.HighPrice,
			&summary.LowPrice,
			&summary.ClosePrice,
			&summary.PriceChange,
			&summary.ChangePct,
			&summary.CreatedAt,
			&summary.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, &summary)
	}

	return summaries, nil
}

func (r *stockPriceSummaryRepositoryImpl) GetLatestBySymbol(symbol string) (*entities.StockPriceSummaryDetail, error) {
	query := `
		SELECT sps.id, sps.emiten_id, e.symbol, e.name, e.sector,
		       sps.date, sps.open_price, sps.high_price, sps.low_price, sps.close_price,
		       (sps.close_price - sps.open_price) as price_change,
		       CASE WHEN sps.open_price > 0
		            THEN ROUND(((sps.close_price - sps.open_price) / sps.open_price * 100)::NUMERIC, 2)
		            ELSE 0
		       END as change_percent,
		       sps.created_at, sps.updated_at
		FROM stock_price_summary sps
		INNER JOIN emitens e ON sps.emiten_id = e.id
		WHERE e.symbol = $1
		ORDER BY sps.date DESC
		LIMIT 1
	`

	var summary entities.StockPriceSummaryDetail
	err := r.db.QueryRow(query, symbol).Scan(
		&summary.ID,
		&summary.EmitenID,
		&summary.Symbol,
		&summary.Name,
		&summary.Sector,
		&summary.Date,
		&summary.OpenPrice,
		&summary.HighPrice,
		&summary.LowPrice,
		&summary.ClosePrice,
		&summary.PriceChange,
		&summary.ChangePct,
		&summary.CreatedAt,
		&summary.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *stockPriceSummaryRepositoryImpl) GetLatestByEmitenID(emitenID int64) (*entities.StockPriceSummary, error) {
	query := `
		SELECT id, emiten_id, date, open_price, high_price, low_price, close_price, created_at, updated_at
		FROM stock_price_summary
		WHERE emiten_id = $1
		ORDER BY date DESC
		LIMIT 1
	`

	var summary entities.StockPriceSummary
	err := r.db.QueryRow(query, emitenID).Scan(
		&summary.ID,
		&summary.EmitenID,
		&summary.Date,
		&summary.OpenPrice,
		&summary.HighPrice,
		&summary.LowPrice,
		&summary.ClosePrice,
		&summary.CreatedAt,
		&summary.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *stockPriceSummaryRepositoryImpl) GetAll(limit, offset int) ([]*entities.StockPriceSummaryDetail, error) {
	query := `
		SELECT sps.id, sps.emiten_id, e.symbol, e.name, e.sector,
		       sps.date, sps.open_price, sps.high_price, sps.low_price, sps.close_price,
		       (sps.close_price - sps.open_price) as price_change,
		       CASE WHEN sps.open_price > 0
		            THEN ROUND(((sps.close_price - sps.open_price) / sps.open_price * 100)::NUMERIC, 2)
		            ELSE 0
		       END as change_percent,
		       sps.created_at, sps.updated_at
		FROM stock_price_summary sps
		INNER JOIN emitens e ON sps.emiten_id = e.id
		ORDER BY sps.date DESC, e.symbol
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summaries []*entities.StockPriceSummaryDetail
	for rows.Next() {
		var summary entities.StockPriceSummaryDetail
		err := rows.Scan(
			&summary.ID,
			&summary.EmitenID,
			&summary.Symbol,
			&summary.Name,
			&summary.Sector,
			&summary.Date,
			&summary.OpenPrice,
			&summary.HighPrice,
			&summary.LowPrice,
			&summary.ClosePrice,
			&summary.PriceChange,
			&summary.ChangePct,
			&summary.CreatedAt,
			&summary.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, &summary)
	}

	return summaries, nil
}

func (r *stockPriceSummaryRepositoryImpl) GetTopGainers(date time.Time, limit int) ([]*entities.StockPriceSummaryDetail, error) {
	query := `
		SELECT sps.id, sps.emiten_id, e.symbol, e.name, e.sector,
		       sps.date, sps.open_price, sps.high_price, sps.low_price, sps.close_price,
		       (sps.close_price - sps.open_price) as price_change,
		       CASE WHEN sps.open_price > 0
		            THEN ROUND(((sps.close_price - sps.open_price) / sps.open_price * 100)::NUMERIC, 2)
		            ELSE 0
		       END as change_percent,
		       sps.created_at, sps.updated_at
		FROM stock_price_summary sps
		INNER JOIN emitens e ON sps.emiten_id = e.id
		WHERE sps.date = $1
		ORDER BY change_percent DESC
		LIMIT $2
	`

	rows, err := r.db.Query(query, date, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summaries []*entities.StockPriceSummaryDetail
	for rows.Next() {
		var summary entities.StockPriceSummaryDetail
		err := rows.Scan(
			&summary.ID,
			&summary.EmitenID,
			&summary.Symbol,
			&summary.Name,
			&summary.Sector,
			&summary.Date,
			&summary.OpenPrice,
			&summary.HighPrice,
			&summary.LowPrice,
			&summary.ClosePrice,
			&summary.PriceChange,
			&summary.ChangePct,
			&summary.CreatedAt,
			&summary.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, &summary)
	}

	return summaries, nil
}

func (r *stockPriceSummaryRepositoryImpl) GetTopLosers(date time.Time, limit int) ([]*entities.StockPriceSummaryDetail, error) {
	query := `
		SELECT sps.id, sps.emiten_id, e.symbol, e.name, e.sector,
		       sps.date, sps.open_price, sps.high_price, sps.low_price, sps.close_price,
		       (sps.close_price - sps.open_price) as price_change,
		       CASE WHEN sps.open_price > 0
		            THEN ROUND(((sps.close_price - sps.open_price) / sps.open_price * 100)::NUMERIC, 2)
		            ELSE 0
		       END as change_percent,
		       sps.created_at, sps.updated_at
		FROM stock_price_summary sps
		INNER JOIN emitens e ON sps.emiten_id = e.id
		WHERE sps.date = $1
		ORDER BY change_percent ASC
		LIMIT $2
	`

	rows, err := r.db.Query(query, date, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summaries []*entities.StockPriceSummaryDetail
	for rows.Next() {
		var summary entities.StockPriceSummaryDetail
		err := rows.Scan(
			&summary.ID,
			&summary.EmitenID,
			&summary.Symbol,
			&summary.Name,
			&summary.Sector,
			&summary.Date,
			&summary.OpenPrice,
			&summary.HighPrice,
			&summary.LowPrice,
			&summary.ClosePrice,
			&summary.PriceChange,
			&summary.ChangePct,
			&summary.CreatedAt,
			&summary.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, &summary)
	}

	return summaries, nil
}

func (r *stockPriceSummaryRepositoryImpl) Create(summary *entities.StockPriceSummary) error {
	query := `
		INSERT INTO stock_price_summary (emiten_id, date, open_price, high_price, low_price, close_price)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRow(
		query,
		summary.EmitenID,
		summary.Date,
		summary.OpenPrice,
		summary.HighPrice,
		summary.LowPrice,
		summary.ClosePrice,
	).Scan(&summary.ID, &summary.CreatedAt, &summary.UpdatedAt)
}

func (r *stockPriceSummaryRepositoryImpl) Update(summary *entities.StockPriceSummary) error {
	query := `
		UPDATE stock_price_summary
		SET emiten_id = $1, date = $2, open_price = $3, high_price = $4, low_price = $5, close_price = $6, updated_at = NOW()
		WHERE id = $7
	`

	result, err := r.db.Exec(
		query,
		summary.EmitenID,
		summary.Date,
		summary.OpenPrice,
		summary.HighPrice,
		summary.LowPrice,
		summary.ClosePrice,
		summary.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func (r *stockPriceSummaryRepositoryImpl) Upsert(summary *entities.StockPriceSummary) error {
	query := `
		INSERT INTO stock_price_summary (emiten_id, date, open_price, high_price, low_price, close_price)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (emiten_id, date)
		DO UPDATE SET
			open_price = EXCLUDED.open_price,
			high_price = EXCLUDED.high_price,
			low_price = EXCLUDED.low_price,
			close_price = EXCLUDED.close_price,
			updated_at = NOW()
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRow(
		query,
		summary.EmitenID,
		summary.Date,
		summary.OpenPrice,
		summary.HighPrice,
		summary.LowPrice,
		summary.ClosePrice,
	).Scan(&summary.ID, &summary.CreatedAt, &summary.UpdatedAt)
}

func (r *stockPriceSummaryRepositoryImpl) Delete(id int64) error {
	query := `DELETE FROM stock_price_summary WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func (r *stockPriceSummaryRepositoryImpl) Count() (int64, error) {
	var count int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM stock_price_summary").Scan(&count)
	return count, err
}

func (r *stockPriceSummaryRepositoryImpl) CountByEmitenID(emitenID int64) (int64, error) {
	var count int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM stock_price_summary WHERE emiten_id = $1", emitenID).Scan(&count)
	return count, err
}
