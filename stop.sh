#!/bin/bash

echo "🧹 Cleaning up Docker for this project..."

# Bring down docker-compose and remove volumes + orphans
docker-compose down --volumes --remove-orphans

# Remove project-related images (assuming they start with your directory name)
PROJECT_NAME=$(basename "$PWD")

echo "🔥 Removing images built by this project: $PROJECT_NAME*"
docker images "$PROJECT_NAME*" -q | xargs -r docker rmi -f

# Remove unnamed/dangling volumes
echo "🗑️ Removing dangling volumes..."
docker volume prune -f

# Optionally remove build cache (safe and useful)
echo "♻️ Pruning build cache..."
docker builder prune -f

echo "🚀 Docker environment is clean and ready for a fresh start!"