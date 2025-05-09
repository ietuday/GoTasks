#!/bin/bash

echo "🧹 Cleaning up Docker environment..."

# Stop and remove all containers defined in docker-compose
docker-compose down --volumes --remove-orphans

# Remove dangling images
echo "🗑️ Removing dangling Docker images..."
docker image prune -f

# Remove all unused volumes
echo "🧼 Removing unused Docker volumes..."
docker volume prune -f

# Optionally remove all images built (uncomment if needed)
# echo "🔥 Removing all Docker images (use with caution!)"
# docker rmi $(docker images -q) -f

echo "📦 Building and starting fresh containers..."
docker-compose up --build --force-recreate
echo "🚀 Docker environment is clean and running!"