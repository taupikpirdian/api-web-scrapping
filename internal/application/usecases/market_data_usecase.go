package usecases

import (
	"api-web-scrapping/internal/application/dto"
	"api-web-scrapping/internal/domain/entities"
	"api-web-scrapping/internal/domain/repositories"
	"fmt"
)

type MarketDataUseCase struct {
	marketDataRepo repositories.MarketDataRepository
}

func NewMarketDataUseCase(marketDataRepo repositories.MarketDataRepository) *MarketDataUseCase {
	return &MarketDataUseCase{
		marketDataRepo: marketDataRepo,
	}
}

// GetAll retrieves all market data from the view
func (uc *MarketDataUseCase) GetAll() (*dto.MarketDataListResponse, error) {
	marketDataList, err := uc.marketDataRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return &dto.MarketDataListResponse{
		Data: entitiesToResponses(marketDataList),
	}, nil
}

// GetByEmiten retrieves market data for a specific emiten
func (uc *MarketDataUseCase) GetByEmiten(emiten string) (*dto.MarketDataListResponse, error) {
	marketDataList, err := uc.marketDataRepo.GetByEmiten(emiten)
	if err != nil {
		return nil, err
	}

	if len(marketDataList) == 0 {
		return nil, fmt.Errorf("no market data found for emiten %s", emiten)
	}

	return &dto.MarketDataListResponse{
		Data: entitiesToResponses(marketDataList),
	}, nil
}

// GetLatestByEmiten retrieves the latest market data for a specific emiten
func (uc *MarketDataUseCase) GetLatestByEmiten(emiten string) (*dto.MarketDataResponse, error) {
	marketData, err := uc.marketDataRepo.GetLatestByEmiten(emiten)
	if err != nil {
		return nil, err
	}

	if marketData == nil {
		return nil, fmt.Errorf("no market data found for emiten %s", emiten)
	}

	return entityToResponse(marketData), nil
}

// GetLatestByAllEmiten retrieves the latest market data for all emitens
func (uc *MarketDataUseCase) GetLatestByAllEmiten() (*dto.MarketDataListResponse, error) {
	marketDataList, err := uc.marketDataRepo.GetLatestByAllEmiten()
	if err != nil {
		return nil, err
	}

	return &dto.MarketDataListResponse{
		Data: entitiesToResponses(marketDataList),
	}, nil
}

// Helper functions

func entityToResponse(entity *entities.MarketData) *dto.MarketDataResponse {
	return &dto.MarketDataResponse{
		ID:         entity.ID,
		Emiten:     entity.Emiten,
		OpenPrice:  entity.OpenPrice,
		HighPrice:  entity.HighPrice,
		LowPrice:   entity.LowPrice,
		ClosePrice: entity.ClosePrice,
		Date:       entity.Date,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
	}
}

func entitiesToResponses(entities []entities.MarketData) []dto.MarketDataResponse {
	responses := make([]dto.MarketDataResponse, len(entities))
	for i, entity := range entities {
		responses[i] = dto.MarketDataResponse{
			ID:         entity.ID,
			Emiten:     entity.Emiten,
			OpenPrice:  entity.OpenPrice,
			HighPrice:  entity.HighPrice,
			LowPrice:   entity.LowPrice,
			ClosePrice: entity.ClosePrice,
			Date:       entity.Date,
			CreatedAt:  entity.CreatedAt,
			UpdatedAt:  entity.UpdatedAt,
		}
	}
	return responses
}
