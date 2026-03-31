#!/bin/bash
export DATABASE_URL="postgresql://futstats:password@localhost:5432/futstats"
export PORT="8080"
export ENVIRONMENT="development"
export SECRET_KEY="local-dev-secret-key-futstats"
export LOG_LEVEL="info"
export LOG_FORMAT="console"
export ALLOWED_ORIGINS="http://localhost:3000"

echo "Starting FutStats server with local Docker PostgreSQL..."
echo "Database: postgresql://futstats:password@localhost:5432/futstats"
echo "Server will run on: http://localhost:8080"
echo "Frontend URL: http://localhost:3000"
echo ""

go run cmd/main.go
