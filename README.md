# API Web Scrapping - Stock Price Summary API

A RESTful API for managing and querying stock price summary data (OHLC) built with Go and PostgreSQL.

## Features

- ✅ Clean Architecture (Domain, Application, Infrastructure, Presentation layers)
- ✅ Stock price summary CRUD operations
- ✅ Query by symbol, date range, pagination
- ✅ Top gainers & losers endpoint
- ✅ PostgreSQL database with migrations
- ✅ Docker support for easy deployment
- ✅ Comprehensive API documentation
- ✅ JWT Authentication

## Tech Stack

- **Go 1.21** - Programming language
- **Gin** - HTTP web framework
- **PostgreSQL** - Database
- **lib/pq** - PostgreSQL driver
- **Docker** - Containerization

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Docker and Docker Compose (optional, for containerized database)

## Quick Start

### 1. Setup Database

```bash
# Start PostgreSQL container
make db-up

# Run migrations
make db-migrate

# Verify database is running
make db-logs
```

### 2. Configure Environment

```bash
# Copy environment file
cp .env.example .env

# Edit .env if needed (defaults work for local development)
nano .env
```

### 3. Install Dependencies

```bash
go mod download
go mod tidy
```

### 4. Run the Application

```bash
# Run directly
go run cmd/api/main.go

# Or build and run
go build -o main cmd/api/main.go
./main
```

The API will start on `http://localhost:8080`

### 5. Test the API

```bash
# Health check
curl http://localhost:8080/health

# Get all stock prices
curl http://localhost:8080/api/v1/stock-prices

# Get latest price for BBCA
curl http://localhost:8080/api/v1/stock-prices/symbol/BBCA/latest

# Get top movers for today
curl http://localhost:8080/api/v1/stock-prices/movers/2026-01-19
```

## Available Endpoints

### Stock Price Summary

- `GET /api/v1/stock-prices` - Get all stock prices with pagination
- `GET /api/v1/stock-prices/:id` - Get stock price by ID
- `GET /api/v1/stock-prices/symbol/:symbol` - Get prices by symbol
- `GET /api/v1/stock-prices/symbol/:symbol/latest` - Get latest price by symbol
- `GET /api/v1/stock-prices/symbol/:symbol/date/:date` - Get price by symbol and date
- `GET /api/v1/stock-prices/range` - Get prices by date range
- `GET /api/v1/stock-prices/symbol/:symbol/range` - Get symbol prices by date range
- `GET /api/v1/stock-prices/movers/:date` - Get top gainers & losers
- `POST /api/v1/stock-prices` - Create stock price summary
- `PUT /api/v1/stock-prices/:id` - Update stock price summary
- `DELETE /api/v1/stock-prices/:id` - Delete stock price summary

### Authentication

- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/register` - User registration

### Health Check

- `GET /health` - API health status

See [API.md](API.md) for detailed API documentation.

## Project Structure

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── application/
│   │   ├── dto/                 # Data Transfer Objects
│   │   │   └── auth_dto.go
│   │   └── usecases/            # Business logic layer
│   │       └── auth_usecase.go
│   ├── domain/
│   │   ├── entities/            # Domain entities
│   │   │   └── user.go
│   │   └── repositories/        # Repository interfaces
│   │       └── user_repository.go
│   ├── infrastructure/
│   │   ├── auth/                # Authentication implementation
│   │   ├── config/              # Configuration
│   │   │   └── config.go
│   │   └── persistence/         # Repository implementations
│   │       └── user_repository_impl.go
│   └── presentation/
│       ├── handlers/            # HTTP handlers
│       │   └── auth_handler.go
│       └── routes/              # Route definitions
│           └── routes.go
├── pkg/
│   └── auth/                    # JWT authentication
│       └── jwt.go
├── tests/
│   └── unit/
│       ├── domain/              # Domain layer tests
│       │   └── user_test.go
│       └── usecase/             # Use case layer tests
│           └── auth_usecase_test.go
├── go.mod
└── README.md
```

## Architecture Layers

### 1. Domain Layer (Core Business Logic)
- **Entities**: Core business objects (User)
- **Repository Interfaces**: Contracts for data access
- No dependencies on outer layers

### 2. Application Layer (Use Cases)
- **Use Cases**: Business logic orchestration
- **DTOs**: Data transfer objects for API
- Depends on Domain layer

### 3. Infrastructure Layer (External Concerns)
- **Persistence**: Repository implementations (in-memory, database)
- **Auth**: JWT implementation
- **Config**: Configuration management
- Implements interfaces from Domain layer

### 4. Presentation Layer (API)
- **Handlers**: HTTP request handlers
- **Routes**: API route definitions
- Depends on Application layer

## Features

- JWT Authentication
- User Registration
- User Login
- Clean Architecture with DDD
- Comprehensive unit tests
- In-memory repository (easily replaceable with real database)

## API Endpoints

### Health Check
```
GET /health
```

### Authentication
```
POST /api/v1/auth/login
POST /api/v1/auth/register
```

## Running the Application

1. Install dependencies:
```bash
go mod download
```

2. Run the application:
```bash
go run cmd/api/main.go
```

The server will start on port 8080.

## Running Tests

Run all tests:
```bash
go test ./tests/unit/... -v
```

Run specific test:
```bash
go test ./tests/unit/usecase -v
go test ./tests/unit/domain -v
```

Run with coverage:
```bash
go test ./tests/unit/... -cover
```

Run benchmark tests:
```bash
go test ./tests/unit/... -bench=.
```

## Example Usage

### Register a new user:
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "full_name": "John Doe"
  }'
```

### Login:
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

## Configuration

Configuration is managed in `internal/infrastructure/config/config.go`. For production, consider using environment variables or a configuration file.

## Dependencies

- Gin - Web framework
- JWT - Authentication
- UUID - Unique identifiers
- Bcrypt - Password hashing
- Testify - Testing framework

## Next Steps

1. Replace in-memory repository with real database (PostgreSQL, MySQL, etc.)
2. Add middleware for authentication
3. Implement refresh token flow
4. Add more comprehensive validation
5. Add integration tests
6. Add database migrations
7. Add API documentation (Swagger)
8. Implement logging middleware
