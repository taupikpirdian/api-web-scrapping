# Docker Deployment Guide

## Prerequisites
- Docker installed
- Docker Compose installed (for docker-compose deployment)

## Quick Start (Development)

### 1. Build and Run
```bash
# Build the Docker image
make docker-build

# Start the container
make docker-up

# View logs
make docker-logs

# Stop the container
make docker-down
```

### 2. Manual Commands
```bash
# Build
docker-compose build

# Start
docker-compose up -d

# View logs
docker-compose logs -f

# Stop
docker-compose down
```

## Production Deployment

### 1. Prepare Environment
```bash
# Copy environment file
cp .env.example .env

# Edit .env and set your JWT_SECRET
nano .env
```

### 2. Build Production Image
```bash
# Build production image
docker build -t api-web-scrapping:latest .
```

### 3. Deploy with Docker Compose
```bash
# Deploy production stack
docker-compose -f deploy.yml up -d

# View logs
docker-compose -f deploy.yml logs -f

# Stop
docker-compose -f deploy.yml down
```

### 4. Deploy with Docker Run
```bash
docker run -d \
  --name api-web-scrapping \
  -p 8080:8080 \
  -e AUTH_JWT_SECRET=your-production-secret \
  -e AUTH_TOKEN_DURATION=24h \
  --restart unless-stopped \
  api-web-scrapping:latest
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVER_PORT` | Server port | `8080` |
| `AUTH_JWT_SECRET` | JWT secret key | **Required in production** |
| `AUTH_TOKEN_DURATION` | Token expiration time | `24h` |

## Health Check

The application includes a health check endpoint:
```bash
curl http://localhost:8080/health
```

## Nginx Reverse Proxy (Optional)

For production, you can use Nginx as a reverse proxy:

1. Configure SSL certificates in `./ssl/` directory
2. Update `nginx.conf` with your domain
3. Uncomment HTTPS section in nginx.conf
4. The deploy.yml includes Nginx service

## Security Best Practices

1. **Change JWT Secret**: Always set a strong `JWT_SECRET` in production
2. **Use HTTPS**: Configure SSL certificates for production
3. **Run as Non-Root**: The Dockerfile uses non-root user
4. **Read-Only Filesystem**: Production deploy uses read-only filesystem
5. **Resource Limits**: CPU and memory limits are configured
6. **Rate Limiting**: Nginx includes rate limiting (10 req/s)

## Monitoring

### View Logs
```bash
# All logs
docker-compose logs -f

# Specific service
docker-compose logs -f api

# Last 100 lines
docker-compose logs --tail=100 api
```

### Container Stats
```bash
docker stats api-web-scrapping
```

## Troubleshooting

### Container won't start
```bash
# Check logs
docker-compose logs api

# Check container status
docker ps -a
```

### Port already in use
```bash
# Find process using port 8080
lsof -i :8080

# Change port in docker-compose.yml
ports:
  - "8081:8080"
```

### Build issues
```bash
# Clean build
docker-compose down -v
docker system prune -f
docker-compose build --no-cache
```

## Updating the Application

```bash
# Pull latest code
git pull

# Rebuild
docker-compose build

# Restart with new image
docker-compose up -d
```

## Cleanup

```bash
# Stop and remove containers
docker-compose down

# Remove volumes
docker-compose down -v

# Remove images
docker rmi api-web-scrapping:latest
```
