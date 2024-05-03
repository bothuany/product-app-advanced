#!/bin/bash

docker run --name postgres-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 6432:5432 -d postgres:latest

echo "Postgresql starting..."
sleep 3

docker exec -it postgres-test psql -U postgres -d postgres -c "CREATE DATABASE product-app-advanced;"
sleep 3
echo "Database product-app-advanced created"

