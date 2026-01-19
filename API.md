# Stock Price Summary API Documentation

## Base URL
```
http://localhost:8080/api/v1
```

## Overview
This API provides endpoints to manage and query stock price summary data (OHLC - Open, High, Low, Close).

## Authentication
Currently, endpoints are open. In production, admin endpoints (POST, PUT, DELETE) should require authentication.

---

## Endpoints

### 1. Get All Stock Prices
Get all stock price summaries with pagination.

**Endpoint:** `GET /stock-prices`

**Query Parameters:**
- `page` (optional, default: 1) - Page number
- `page_size` (optional, default: 10, max: 100) - Items per page

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/stock-prices?page=1&page_size=10"
```

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "emiten_id": 1,
      "symbol": "BBCA",
      "name": "Bank Central Asia Tbk",
      "sector": "Finance",
      "date": "2026-01-19T00:00:00Z",
      "open_price": 9200,
      "high_price": 9300,
      "low_price": 9150,
      "close_price": 9250,
      "price_change": 50,
      "change_percent": 0.54,
      "created_at": "2026-01-19T10:00:00Z",
      "updated_at": "2026-01-19T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total_count": 100,
    "total_pages": 10
  }
}
```

---

### 2. Get Stock Price by ID
Get a specific stock price summary by ID.

**Endpoint:** `GET /stock-prices/:id`

**URL Parameters:**
- `id` (required) - Stock price summary ID

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/stock-prices/1"
```

**Response:**
```json
{
  "id": 1,
  "emiten_id": 1,
  "date": "2026-01-19T00:00:00Z",
  "open_price": 9200,
  "high_price": 9300,
  "low_price": 9150,
  "close_price": 9250,
  "created_at": "2026-01-19T10:00:00Z",
  "updated_at": "2026-01-19T10:00:00Z"
}
```

---

### 3. Get Stock Prices by Symbol
Get all stock price summaries for a specific symbol with pagination.

**Endpoint:** `GET /stock-prices/symbol/:symbol`

**URL Parameters:**
- `symbol` (required) - Stock symbol (e.g., BBCA, TLKM)

**Query Parameters:**
- `page` (optional, default: 1) - Page number
- `page_size` (optional, default: 10, max: 100) - Items per page

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/stock-prices/symbol/BBCA?page=1&page_size=30"
```

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "emiten_id": 1,
      "symbol": "BBCA",
      "name": "Bank Central Asia Tbk",
      "sector": "Finance",
      "date": "2026-01-19T00:00:00Z",
      "open_price": 9200,
      "high_price": 9300,
      "low_price": 9150,
      "close_price": 9250,
      "price_change": 50,
      "change_percent": 0.54,
      "created_at": "2026-01-19T10:00:00Z",
      "updated_at": "2026-01-19T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "page_size": 30,
    "total_count": 250,
    "total_pages": 9
  }
}
```

---

### 4. Get Latest Stock Price by Symbol
Get the most recent stock price for a symbol.

**Endpoint:** `GET /stock-prices/symbol/:symbol/latest`

**URL Parameters:**
- `symbol` (required) - Stock symbol

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/stock-prices/symbol/BBCA/latest"
```

**Response:**
```json
{
  "id": 1,
  "emiten_id": 1,
  "symbol": "BBCA",
  "name": "Bank Central Asia Tbk",
  "sector": "Finance",
  "date": "2026-01-19T00:00:00Z",
  "open_price": 9200,
  "high_price": 9300,
  "low_price": 9150,
  "close_price": 9250,
  "price_change": 50,
  "change_percent": 0.54,
  "created_at": "2026-01-19T10:00:00Z",
  "updated_at": "2026-01-19T10:00:00Z"
}
```

---

### 5. Get Stock Price by Symbol and Date
Get a stock price summary for a specific symbol and date.

**Endpoint:** `GET /stock-prices/symbol/:symbol/date/:date`

**URL Parameters:**
- `symbol` (required) - Stock symbol
- `date` (required) - Date in YYYY-MM-DD format

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/stock-prices/symbol/BBCA/date/2026-01-19"
```

**Response:**
```json
{
  "id": 1,
  "emiten_id": 1,
  "symbol": "BBCA",
  "name": "Bank Central Asia Tbk",
  "sector": "Finance",
  "date": "2026-01-19T00:00:00Z",
  "open_price": 9200,
  "high_price": 9300,
  "low_price": 9150,
  "close_price": 9250,
  "price_change": 50,
  "change_percent": 0.54,
  "created_at": "2026-01-19T10:00:00Z",
  "updated_at": "2026-01-19T10:00:00Z"
}
```

---

### 6. Get Stock Prices by Date Range
Get all stock price summaries within a date range.

**Endpoint:** `GET /stock-prices/range`

**Query Parameters:**
- `start_date` (required) - Start date in YYYY-MM-DD format
- `end_date` (required) - End date in YYYY-MM-DD format
- `page` (optional, default: 1) - Page number
- `page_size` (optional, default: 10, max: 100) - Items per page

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/stock-prices/range?start_date=2026-01-01&end_date=2026-01-19&page=1&page_size=50"
```

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "emiten_id": 1,
      "symbol": "BBCA",
      "name": "Bank Central Asia Tbk",
      "sector": "Finance",
      "date": "2026-01-19T00:00:00Z",
      "open_price": 9200,
      "high_price": 9300,
      "low_price": 9150,
      "close_price": 9250,
      "price_change": 50,
      "change_percent": 0.54,
      "created_at": "2026-01-19T10:00:00Z",
      "updated_at": "2026-01-19T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "page_size": 50,
    "total_count": 500,
    "total_pages": 10
  }
}
```

---

### 7. Get Stock Prices by Symbol and Date Range
Get stock price summaries for a specific symbol within a date range.

**Endpoint:** `GET /stock-prices/symbol/:symbol/range`

**URL Parameters:**
- `symbol` (required) - Stock symbol

**Query Parameters:**
- `start_date` (required) - Start date in YYYY-MM-DD format
- `end_date` (required) - End date in YYYY-MM-DD format
- `page` (optional, default: 1) - Page number
- `page_size` (optional, default: 10, max: 100) - Items per page

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/stock-prices/symbol/BBCA/range?start_date=2026-01-01&end_date=2026-01-19"
```

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "emiten_id": 1,
      "symbol": "BBCA",
      "name": "Bank Central Asia Tbk",
      "sector": "Finance",
      "date": "2026-01-19T00:00:00Z",
      "open_price": 9200,
      "high_price": 9300,
      "low_price": 9150,
      "close_price": 9250,
      "price_change": 50,
      "change_percent": 0.54,
      "created_at": "2026-01-19T10:00:00Z",
      "updated_at": "2026-01-19T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total_count": 19,
    "total_pages": 2
  }
}
```

---

### 8. Get Top Movers (Gainers & Losers)
Get top gainers and losers for a specific date.

**Endpoint:** `GET /stock-prices/movers/:date`

**URL Parameters:**
- `date` (required) - Date in YYYY-MM-DD format

**Query Parameters:**
- `limit` (optional, default: 10) - Number of top gainers/losers to return

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/stock-prices/movers/2026-01-19?limit=20"
```

**Response:**
```json
{
  "date": "2026-01-19T00:00:00Z",
  "gainers": [
    {
      "id": 1,
      "emiten_id": 1,
      "symbol": "BBCA",
      "name": "Bank Central Asia Tbk",
      "sector": "Finance",
      "date": "2026-01-19T00:00:00Z",
      "open_price": 9200,
      "high_price": 9300,
      "low_price": 9150,
      "close_price": 9250,
      "price_change": 50,
      "change_percent": 0.54,
      "created_at": "2026-01-19T10:00:00Z",
      "updated_at": "2026-01-19T10:00:00Z"
    }
  ],
  "losers": [
    {
      "id": 2,
      "emiten_id": 2,
      "symbol": "TLKM",
      "name": "Telkom Indonesia Tbk",
      "sector": "Technology",
      "date": "2026-01-19T00:00:00Z",
      "open_price": 3500,
      "high_price": 3550,
      "low_price": 3450,
      "close_price": 3475,
      "price_change": -25,
      "change_percent": -0.71,
      "created_at": "2026-01-19T10:00:00Z",
      "updated_at": "2026-01-19T10:00:00Z"
    }
  ]
}
```

---

### 9. Create Stock Price Summary
Create a new stock price summary (Admin only).

**Endpoint:** `POST /stock-prices`

**Request Body:**
```json
{
  "emiten_id": 1,
  "date": "2026-01-19T00:00:00Z",
  "open_price": 9200,
  "high_price": 9300,
  "low_price": 9150,
  "close_price": 9250
}
```

**Example Request:**
```bash
curl -X POST "http://localhost:8080/api/v1/stock-prices" \
  -H "Content-Type: application/json" \
  -d '{
    "emiten_id": 1,
    "date": "2026-01-19T00:00:00Z",
    "open_price": 9200,
    "high_price": 9300,
    "low_price": 9150,
    "close_price": 9250
  }'
```

**Response:**
```json
{
  "id": 1,
  "emiten_id": 1,
  "date": "2026-01-19T00:00:00Z",
  "open_price": 9200,
  "high_price": 9300,
  "low_price": 9150,
  "close_price": 9250,
  "created_at": "2026-01-19T10:00:00Z",
  "updated_at": "2026-01-19T10:00:00Z"
}
```

---

### 10. Update Stock Price Summary
Update an existing stock price summary (Admin only).

**Endpoint:** `PUT /stock-prices/:id`

**URL Parameters:**
- `id` (required) - Stock price summary ID

**Request Body:**
```json
{
  "emiten_id": 1,
  "date": "2026-01-19T00:00:00Z",
  "open_price": 9250,
  "high_price": 9350,
  "low_price": 9200,
  "close_price": 9300
}
```

**Example Request:**
```bash
curl -X PUT "http://localhost:8080/api/v1/stock-prices/1" \
  -H "Content-Type: application/json" \
  -d '{
    "emiten_id": 1,
    "date": "2026-01-19T00:00:00Z",
    "open_price": 9250,
    "high_price": 9350,
    "low_price": 9200,
    "close_price": 9300
  }'
```

**Response:**
```json
{
  "id": 1,
  "emiten_id": 1,
  "date": "2026-01-19T00:00:00Z",
  "open_price": 9250,
  "high_price": 9350,
  "low_price": 9200,
  "close_price": 9300,
  "created_at": "2026-01-19T10:00:00Z",
  "updated_at": "2026-01-19T11:00:00Z"
}
```

---

### 11. Delete Stock Price Summary
Delete a stock price summary (Admin only).

**Endpoint:** `DELETE /stock-prices/:id`

**URL Parameters:**
- `id` (required) - Stock price summary ID

**Example Request:**
```bash
curl -X DELETE "http://localhost:8080/api/v1/stock-prices/1"
```

**Response:**
```json
{
  "message": "Stock price summary deleted successfully"
}
```

---

## Error Responses

All endpoints may return error responses in the following format:

```json
{
  "error": "error_code",
  "message": "Human-readable error message"
}
```

### Common Error Codes:
- `bad_request` - Invalid request parameters
- `not_found` - Resource not found
- `internal_error` - Internal server error

---

## Status Codes
- `200 OK` - Request successful
- `201 Created` - Resource created successfully
- `400 Bad Request` - Invalid request parameters
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

---

## Examples with curl

### Get latest price for BBCA
```bash
curl "http://localhost:8080/api/v1/stock-prices/symbol/BBCA/latest"
```

### Get BBCA prices for last 30 days
```bash
curl "http://localhost:8080/api/v1/stock-prices/symbol/BBCA/range?start_date=2025-12-20&end_date=2026-01-19"
```

### Get top 20 gainers and losers today
```bash
curl "http://localhost:8080/api/v1/stock-prices/movers/2026-01-19?limit=20"
```

### Create new stock price
```bash
curl -X POST "http://localhost:8080/api/v1/stock-prices" \
  -H "Content-Type: application/json" \
  -d '{
    "emiten_id": 1,
    "date": "2026-01-19T00:00:00Z",
    "open_price": 9200,
    "high_price": 9300,
    "low_price": 9150,
    "close_price": 9250
  }'
```

---

## Next Steps

1. ✅ Start database: `make db-up`
2. ✅ Run migrations: `make db-migrate`
3. ✅ Start API server
4. ✅ Test endpoints with curl or Postman
