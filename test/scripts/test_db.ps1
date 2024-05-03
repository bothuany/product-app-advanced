# Start the PostgreSQL container
docker run --name postgres-go-advanced -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 7432:5432 -d postgres:latest

Write-Host "Postgresql starting..."
Start-Sleep -Seconds 3

# Create the database
docker exec -it postgres-go-advanced psql -U postgres -d postgres -c "CREATE DATABASE productappadvanced;"
Start-Sleep -Seconds 3
Write-Host "Database productappadvanced created"
