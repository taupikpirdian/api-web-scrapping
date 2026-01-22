package handlers

import (
	"api-web-scrapping/internal/application/dto"
	"api-web-scrapping/internal/application/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MarketDataHandler struct {
	useCase *usecases.MarketDataUseCase
}

func NewMarketDataHandler(useCase *usecases.MarketDataUseCase) *MarketDataHandler {
	return &MarketDataHandler{
		useCase: useCase,
	}
}

// GetAll handles GET /api/v1/market-data
// Retrieves all market data from v_latest_market_data view
func (h *MarketDataHandler) GetAll(c *gin.Context) {
	marketDataList, err := h.useCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "internal_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, marketDataList)
}

// GetByEmiten handles GET /api/v1/market-data/emiten/:emiten
// Retrieves market data for a specific emiten
func (h *MarketDataHandler) GetByEmiten(c *gin.Context) {
	emiten := c.Param("emiten")

	marketDataList, err := h.useCase.GetByEmiten(emiten)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error:   "not_found",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, marketDataList)
}

// GetLatestByEmiten handles GET /api/v1/market-data/emiten/:emiten/latest
// Retrieves the latest market data for a specific emiten
func (h *MarketDataHandler) GetLatestByEmiten(c *gin.Context) {
	emiten := c.Param("emiten")

	marketData, err := h.useCase.GetLatestByEmiten(emiten)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error:   "not_found",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, marketData)
}

// GetLatestByAllEmiten handles GET /api/v1/market-data/latest
// Retrieves the latest market data for all emitens
func (h *MarketDataHandler) GetLatestByAllEmiten(c *gin.Context) {
	marketDataList, err := h.useCase.GetLatestByAllEmiten()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "internal_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, marketDataList)
}
