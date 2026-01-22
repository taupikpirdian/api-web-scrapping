# API Documentation

## Base URL
`http://localhost:8080/api/v1`

## Authentication

### Login
**POST** `/auth/login`

Authenticate a user and return a JWT token.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "token": "jwt_token_here",
  "user": {
    "id": "user_uuid",
    "email": "user@example.com",
    "full_name": "John Doe"
  }
}
```

## Market Data

### Get All Market Data
**GET** `/market-data`

Get all market data from view.

**Response:**
```json
[
  {
    "id": 1,
    "emiten": "BBCA",
    "open_price": 9200,
    "high_price": 9300,
    "low_price": 9150,
    "close_price": 9250,
    "volume": 1000000,
    "value": 9250000000,
    "frequency": 5000,
    "date": "2026-01-19T00:00:00Z",
    "created_at": "...",
    "updated_at": "..."
  }
]
```

### Get Latest Market Data for All Emitens
**GET** `/market-data/latest`

Get the latest market data for all emitens.

### Get Market Data by Emiten
**GET** `/market-data/emiten/:emiten`

Get market data for a specific emiten.

**URL Parameters:**
- `emiten`: Stock symbol (e.g., BBCA)

### Get Latest Market Data by Emiten
**GET** `/market-data/emiten/:emiten/latest`

Get the latest market data for a specific emiten.
