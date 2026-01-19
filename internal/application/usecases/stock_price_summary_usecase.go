package usecases

import (
	"api-web-scrapping/internal/application/dto"
	"api-web-scrapping/internal/domain/entities"
	"api-web-scrapping/internal/domain/repositories"
	"errors"
	"fmt"
	"time"
)

type StockPriceSummaryUseCase struct {
	summaryRepo repositories.StockPriceSummaryRepository
}

func NewStockPriceSummaryUseCase(summaryRepo repositories.StockPriceSummaryRepository) *StockPriceSummaryUseCase {
	return &StockPriceSummaryUseCase{
		summaryRepo: summaryRepo,
	}
}

// GetByID retrieves a stock price summary by ID
func (uc *StockPriceSummaryUseCase) GetByID(id int64) (*dto.StockPriceSummaryResponse, error) {
	summary, err := uc.summaryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if summary == nil {
		return nil, errors.New("stock price summary not found")
	}

	return uc.entityToResponse(summary), nil
}

// GetBySymbolAndDate retrieves a stock price summary by symbol and date
func (uc *StockPriceSummaryUseCase) GetBySymbolAndDate(symbol string, dateStr string) (*dto.StockPriceSummaryDetailResponse, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}

	summary, err := uc.summaryRepo.GetBySymbolAndDate(symbol, date)
	if err != nil {
		return nil, err
	}
	if summary == nil {
		return nil, fmt.Errorf("stock price summary not found for symbol %s on %s", symbol, dateStr)
	}

	return uc.entityToDetailResponse(summary), nil
}

// GetBySymbol retrieves all stock price summaries for a symbol
func (uc *StockPriceSummaryUseCase) GetBySymbol(symbol string, page, pageSize int) (*dto.StockPriceListResponse, error) {
	offset := (page - 1) * pageSize

	summaries, err := uc.summaryRepo.GetBySymbol(symbol, pageSize, offset)
	if err != nil {
		return nil, err
	}

	totalCount, err := uc.summaryRepo.CountByEmitenID(summaries[0].EmitenID)
	if err != nil {
		return nil, err
	}

	response := &dto.StockPriceListResponse{
		Data: entitiesToDetailResponses(summaries),
		Pagination: &dto.PaginationResponse{
			Page:       page,
			PageSize:   pageSize,
			TotalCount: totalCount,
			TotalPages: int(totalCount) / pageSize,
		},
	}

	return response, nil
}

// GetByDateRange retrieves summaries within a date range
func (uc *StockPriceSummaryUseCase) GetByDateRange(startDate, endDate string, page, pageSize int) (*dto.StockPriceListResponse, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %w", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %w", err)
	}

	offset := (page - 1) * pageSize

	summaries, err := uc.summaryRepo.GetByDateRange(start, end, pageSize, offset)
	if err != nil {
		return nil, err
	}

	totalCount, err := uc.summaryRepo.Count()
	if err != nil {
		return nil, err
	}

	response := &dto.StockPriceListResponse{
		Data: entitiesToDetailResponses(summaries),
		Pagination: &dto.PaginationResponse{
			Page:       page,
			PageSize:   pageSize,
			TotalCount: totalCount,
			TotalPages: int(totalCount) / pageSize,
		},
	}

	return response, nil
}

// GetBySymbolAndDateRange retrieves summaries for a symbol within a date range
func (uc *StockPriceSummaryUseCase) GetBySymbolAndDateRange(symbol, startDate, endDate string, page, pageSize int) (*dto.StockPriceListResponse, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %w", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %w", err)
	}

	offset := (page - 1) * pageSize

	summaries, err := uc.summaryRepo.GetBySymbolAndDateRange(symbol, start, end, pageSize, offset)
	if err != nil {
		return nil, err
	}

	response := &dto.StockPriceListResponse{
		Data: entitiesToDetailResponses(summaries),
		Pagination: &dto.PaginationResponse{
			Page:       page,
			PageSize:   pageSize,
			TotalCount: int64(len(summaries)),
			TotalPages: 1,
		},
	}

	return response, nil
}

// GetLatestBySymbol retrieves the latest summary for a symbol
func (uc *StockPriceSummaryUseCase) GetLatestBySymbol(symbol string) (*dto.StockPriceSummaryDetailResponse, error) {
	summary, err := uc.summaryRepo.GetLatestBySymbol(symbol)
	if err != nil {
		return nil, err
	}
	if summary == nil {
		return nil, fmt.Errorf("no stock price summary found for symbol %s", symbol)
	}

	return uc.entityToDetailResponse(summary), nil
}

// GetAll retrieves all stock price summaries with pagination
func (uc *StockPriceSummaryUseCase) GetAll(page, pageSize int) (*dto.StockPriceListResponse, error) {
	offset := (page - 1) * pageSize

	summaries, err := uc.summaryRepo.GetAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	totalCount, err := uc.summaryRepo.Count()
	if err != nil {
		return nil, err
	}

	response := &dto.StockPriceListResponse{
		Data: entitiesToDetailResponses(summaries),
		Pagination: &dto.PaginationResponse{
			Page:       page,
			PageSize:   pageSize,
			TotalCount: totalCount,
			TotalPages: int(totalCount) / pageSize,
		},
	}

	return response, nil
}

// GetTopMovers retrieves top gainers and losers for a specific date
func (uc *StockPriceSummaryUseCase) GetTopMovers(dateStr string, limit int) (*dto.TopMoversResponse, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}

	gainers, err := uc.summaryRepo.GetTopGainers(date, limit)
	if err != nil {
		return nil, err
	}

	losers, err := uc.summaryRepo.GetTopLosers(date, limit)
	if err != nil {
		return nil, err
	}

	response := &dto.TopMoversResponse{
		Date:    date,
		Gainers: entitiesToDetailResponses(gainers),
		Losers:  entitiesToDetailResponses(losers),
	}

	return response, nil
}

// Create creates a new stock price summary
func (uc *StockPriceSummaryUseCase) Create(req *dto.CreateStockPriceSummaryRequest) (*dto.StockPriceSummaryResponse, error) {
	summary := &entities.StockPriceSummary{
		EmitenID:   req.EmitenID,
		Date:       req.Date,
		OpenPrice:  req.OpenPrice,
		HighPrice:  req.HighPrice,
		LowPrice:   req.LowPrice,
		ClosePrice: req.ClosePrice,
	}

	err := uc.summaryRepo.Create(summary)
	if err != nil {
		return nil, err
	}

	return uc.entityToResponse(summary), nil
}

// Update updates an existing stock price summary
func (uc *StockPriceSummaryUseCase) Update(id int64, req *dto.UpdateStockPriceSummaryRequest) (*dto.StockPriceSummaryResponse, error) {
	summary, err := uc.summaryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if summary == nil {
		return nil, errors.New("stock price summary not found")
	}

	summary.EmitenID = req.EmitenID
	summary.Date = req.Date
	summary.OpenPrice = req.OpenPrice
	summary.HighPrice = req.HighPrice
	summary.LowPrice = req.LowPrice
	summary.ClosePrice = req.ClosePrice

	err = uc.summaryRepo.Update(summary)
	if err != nil {
		return nil, err
	}

	return uc.entityToResponse(summary), nil
}

// Delete deletes a stock price summary
func (uc *StockPriceSummaryUseCase) Delete(id int64) error {
	_, err := uc.summaryRepo.GetByID(id)
	if err != nil {
		return err
	}

	return uc.summaryRepo.Delete(id)
}

// Helper functions

func (uc *StockPriceSummaryUseCase) entityToResponse(entity *entities.StockPriceSummary) *dto.StockPriceSummaryResponse {
	return &dto.StockPriceSummaryResponse{
		ID:         entity.ID,
		EmitenID:   entity.EmitenID,
		Date:       entity.Date,
		OpenPrice:  entity.OpenPrice,
		HighPrice:  entity.HighPrice,
		LowPrice:   entity.LowPrice,
		ClosePrice: entity.ClosePrice,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
	}
}

func (uc *StockPriceSummaryUseCase) entityToDetailResponse(entity *entities.StockPriceSummaryDetail) *dto.StockPriceSummaryDetailResponse {
	return &dto.StockPriceSummaryDetailResponse{
		ID:          entity.ID,
		EmitenID:    entity.EmitenID,
		Symbol:      entity.Symbol,
		Name:        entity.Name,
		Sector:      entity.Sector,
		Date:        entity.Date,
		OpenPrice:   entity.OpenPrice,
		HighPrice:   entity.HighPrice,
		LowPrice:    entity.LowPrice,
		ClosePrice:  entity.ClosePrice,
		PriceChange: entity.PriceChange,
		ChangePct:   entity.ChangePct,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func entitiesToDetailResponses(entities []*entities.StockPriceSummaryDetail) []*dto.StockPriceSummaryDetailResponse {
	responses := make([]*dto.StockPriceSummaryDetailResponse, len(entities))
	for i, entity := range entities {
		responses[i] = &dto.StockPriceSummaryDetailResponse{
			ID:          entity.ID,
			EmitenID:    entity.EmitenID,
			Symbol:      entity.Symbol,
			Name:        entity.Name,
			Sector:      entity.Sector,
			Date:        entity.Date,
			OpenPrice:   entity.OpenPrice,
			HighPrice:   entity.HighPrice,
			LowPrice:    entity.LowPrice,
			ClosePrice:  entity.ClosePrice,
			PriceChange: entity.PriceChange,
			ChangePct:   entity.ChangePct,
			CreatedAt:   entity.CreatedAt,
			UpdatedAt:   entity.UpdatedAt,
		}
	}
	return responses
}
