#!/bin/bash

echo "🧹 Cleaning up Docker environment..."

# Stop and remove all containers defined in docker-compose
docker-compose down --volumes --remove-orphans

# Remove dangling images
echo "🗑️ Removing dangling Docker images..."
docker image prune -f

# Remove unused volumes
echo "🧼 Removing unused Docker volumes..."
docker volume prune -f

# Optional full wipe: Uncomment to remove **all** images
# echo "🔥 Removing all Docker images (use with caution!)"
# docker rmi -f $(docker images -q)

# Optional: Remove build cache to ensure clean rebuild
echo "🧽 Cleaning up Docker build cache..."
docker builder prune -af

# Rebuild and restart
echo "📦 Building and starting fresh containers..."
docker-compose up --build --force-recreate

echo "🚀 Docker environment is clean and running!"
