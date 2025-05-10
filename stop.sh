#!/bin/bash

set -e

echo "🧹 Cleaning up Docker for this project..."

# Bring down docker-compose and remove volumes + orphans
docker-compose down --volumes --remove-orphans

# Dynamically get project name (Docker image prefixes might vary)
PROJECT_NAME=$(basename "$PWD" | tr '[:upper:]' '[:lower:]' | tr -cd '[:alnum:]-')

echo "🔥 Attempting to remove images built by this project: ${PROJECT_NAME}*"
docker images --format "{{.Repository}}:{{.Tag}} {{.ID}}" | grep "^${PROJECT_NAME}" | awk '{print $2}' | xargs -r docker rmi -f

# Remove unnamed/dangling volumes
echo "🗑️ Removing dangling volumes..."
docker volume prune -f

# Remove build cache
echo "♻️ Pruning build cache..."
docker builder prune -f

echo "🚀 Docker environment is clean and ready for a fresh start!"
