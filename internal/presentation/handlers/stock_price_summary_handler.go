package handlers

import (
	"api-web-scrapping/internal/application/dto"
	"api-web-scrapping/internal/application/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StockPriceSummaryHandler struct {
	useCase *usecases.StockPriceSummaryUseCase
}

func NewStockPriceSummaryHandler(useCase *usecases.StockPriceSummaryUseCase) *StockPriceSummaryHandler {
	return &StockPriceSummaryHandler{
		useCase: useCase,
	}
}

// GetByID handles GET /api/v1/stock-prices/:id
func (h *StockPriceSummaryHandler) GetByID(c *gin.Context) {
	var req dto.GetStockPriceSummaryRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "bad_request",
			Message: err.Error(),
		})
		return
	}

	summary, err := h.useCase.GetByID(req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error:   "not_found",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, summary)
}

// GetBySymbolAndDate handles GET /api/v1/stock-prices/symbol/:symbol/date/:date
func (h *StockPriceSummaryHandler) GetBySymbolAndDate(c *gin.Context) {
	symbol := c.Param("symbol")
	dateStr := c.Param("date")

	summary, err := h.useCase.GetBySymbolAndDate(symbol, dateStr)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error:   "not_found",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, summary)
}

// GetBySymbol handles GET /api/v1/stock-prices/symbol/:symbol
func (h *StockPriceSummaryHandler) GetBySymbol(c *gin.Context) {
	symbol := c.Param("symbol")

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeInt = 10
	}

	summaries, err := h.useCase.GetBySymbol(symbol, pageInt, pageSizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "internal_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, summaries)
}

// GetByDateRange handles GET /api/v1/stock-prices/range
func (h *StockPriceSummaryHandler) GetByDateRange(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	if startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "bad_request",
			Message: "start_date and end_date are required",
		})
		return
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeInt = 10
	}

	summaries, err := h.useCase.GetByDateRange(startDate, endDate, pageInt, pageSizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "internal_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, summaries)
}

// GetBySymbolAndDateRange handles GET /api/v1/stock-prices/symbol/:symbol/range
func (h *StockPriceSummaryHandler) GetBySymbolAndDateRange(c *gin.Context) {
	symbol := c.Param("symbol")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	if startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "bad_request",
			Message: "start_date and end_date are required",
		})
		return
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeInt = 10
	}

	summaries, err := h.useCase.GetBySymbolAndDateRange(symbol, startDate, endDate, pageInt, pageSizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "internal_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, summaries)
}

// GetLatestBySymbol handles GET /api/v1/stock-prices/symbol/:symbol/latest
func (h *StockPriceSummaryHandler) GetLatestBySymbol(c *gin.Context) {
	symbol := c.Param("symbol")

	summary, err := h.useCase.GetLatestBySymbol(symbol)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error:   "not_found",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, summary)
}

// GetAll handles GET /api/v1/stock-prices
func (h *StockPriceSummaryHandler) GetAll(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeInt = 10
	}

	summaries, err := h.useCase.GetAll(pageInt, pageSizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "internal_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, summaries)
}

// GetTopMovers handles GET /api/v1/stock-prices/movers/:date
func (h *StockPriceSummaryHandler) GetTopMovers(c *gin.Context) {
	dateStr := c.Param("date")
	limit := c.DefaultQuery("limit", "10")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	movers, err := h.useCase.GetTopMovers(dateStr, limitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "internal_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, movers)
}

// Create handles POST /api/v1/stock-prices
func (h *StockPriceSummaryHandler) Create(c *gin.Context) {
	var req dto.CreateStockPriceSummaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "bad_request",
			Message: err.Error(),
		})
		return
	}

	summary, err := h.useCase.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "internal_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, summary)
}

// Update handles PUT /api/v1/stock-prices/:id
func (h *StockPriceSummaryHandler) Update(c *gin.Context) {
	var req dto.UpdateStockPriceSummaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "bad_request",
			Message: err.Error(),
		})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "bad_request",
			Message: "invalid id",
		})
		return
	}

	summary, err := h.useCase.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "internal_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, summary)
}

// Delete handles DELETE /api/v1/stock-prices/:id
func (h *StockPriceSummaryHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "bad_request",
			Message: "invalid id",
		})
		return
	}

	err = h.useCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "internal_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Stock price summary deleted successfully",
	})
}
