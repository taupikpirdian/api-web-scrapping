.PHONY: build docker-build docker-up docker-down docker-logs docker-clean db-up db-down db-migrate db-rollback db-logs db-shell

# Build the application
build:
	go build -o main ./cmd/api

# Docker build
docker-build:
	docker-compose build

# Docker up
docker-up:
	docker-compose up -d

# Docker down
docker-down:
	docker-compose down

# Docker logs
docker-logs:
	docker-compose logs -f

# Docker clean
docker-clean:
	docker-compose down -v
	docker system prune -f

# Production build
docker-build-prod:
	docker build -t api-web-scrapping:latest .

# Production deploy
docker-deploy-prod:
	docker-compose -f deploy.yml up -d

# View logs production
docker-logs-prod:
	docker-compose -f deploy.yml logs -f

# Database - Start PostgreSQL
db-up:
	docker-compose -f migrations/docker-compose-db.yml up -d

# Database - Stop PostgreSQL
db-down:
	docker-compose -f migrations/docker-compose-db.yml down

# Database - View logs
db-logs:
	docker-compose -f migrations/docker-compose-db.yml logs -f postgres

# Database - Run migrations
db-migrate:
	@echo "Running database migrations..."
	./scripts/run-migrations.sh

# Database - Rollback last migration
db-rollback:
	@echo "Rolling back last migration..."
	./scripts/rollback-migration.sh

# Database - Connect to PostgreSQL shell
db-shell:
	docker exec -it api-web-scrapping-db psql -U ${POSTGRES_USER:-admin} -d ${POSTGRES_DB:-api_web_scrapping}

# Database - Backup database
db-backup:
	docker exec api-web-scrapping-db pg_dump -U ${POSTGRES_USER:-admin} ${POSTGRES_DB:-api_web_scrapping} > backup_$$(date +%Y%m%d_%H%M%S).sql

# Database - Restore database
db-restore:
	@read -p "Enter backup file name: " file; \
	docker exec -i api-web-scrapping-db psql -U ${POSTGRES_USER:-admin} ${POSTGRES_DB:-api_web_scrapping} < $$file

# Database - Reset database (WARNING: Deletes all data)
db-reset:
	docker-compose -f migrations/docker-compose-db.yml down -v
	docker-compose -f migrations/docker-compose-db.yml up -d
	@echo "Waiting for database to be ready..."
	sleep 5
	make db-migrate
