# 📜 Task CRUD App (Go + React + MongoDB)

This project is a full-stack Task Management application built using:

* ⚙️ **Go (Golang)** for backend API
* ⚛️ **React** for frontend UI
* 🍃 **MongoDB** as the database
* 🐳 **Docker + Docker Compose** for containerized development

---

## 📦 Project Structure

```
├── backend/         # Go backend API
├── frontend/        # React frontend
├── docker-compose.yml
├── start-clean.sh   # Clean and start Docker containers
└── stop.sh          # Stop and fully clean Docker containers/images/volumes
```

---

## ✨ Getting Started

### 📋 Prerequisites

Make sure you have the following installed:

* Docker
* Docker Compose

### 🔧 Run the Project

To run everything fresh (clean, build, up):

```bash
./start-clean.sh
```

To stop everything and remove images, volumes, and cache:

```bash
./stop.sh
```

### 🌐 Access App

* Frontend: [http://localhost:3000](http://localhost:3000)
* Backend: [http://localhost:8080](http://localhost:8080)
* MongoDB: localhost:27017

---

## 🐳 Docker Services

### `mongo`

* Uses the official `mongo:6` image
* Persists data using a Docker named volume `mongo-data`

### `backend`

* Golang API container
* Auto-restarts and uses volume mount for live code reload

### `frontend`

* React development server using `npm start`
* Mounts code for hot reload in development

---

## 💠 Project Scripts

### 🛉 `stop.sh` – Full Cleanup

This script **removes everything** related to this project:

* Shuts down containers
* Deletes associated volumes
* Removes Docker images built by this project
* Prunes build cache

```bash
./stop.sh
```

```bash
#!/bin/bash

echo "🚭 Stopping and cleaning Docker for this project..."

# Stop and remove containers/volumes/networks
docker-compose down --volumes --remove-orphans

# Get project directory name to target images
PROJECT_NAME=$(basename "$PWD")

echo "🔥 Removing images built by this project: $PROJECT_NAME*"
docker images "$PROJECT_NAME*" -q | xargs -r docker rmi -f

echo "🧼 Removing dangling volumes..."
docker volume prune -f

echo "♻️ Pruning build cache..."
docker builder prune -af

echo "✅ Project cleanup complete."
```

---

### ✨ `start-clean.sh` – Fresh Start

This script runs the full cleanup **and then immediately starts fresh containers**.

```bash
./start-clean.sh
```

```bash
#!/bin/bash

echo "🚹 Cleaning up Docker environment..."

# Stop and remove all containers and volumes
docker-compose down --volumes --remove-orphans

# Remove dangling images
docker image prune -f

# Remove unused volumes
docker volume prune -f

# Clean build cache
docker builder prune -af

# Start containers fresh
echo "📦 Rebuilding and starting fresh containers..."
docker-compose up --build --force-recreate

echo "🚀 Docker environment is clean and running!"
```

---

## 📁 Volumes Used

### `mongo-data`

Defined in `docker-compose.yml` to persist MongoDB data.

```yaml
volumes:
  mongo-data:
```

---

## 💡 Notes

* React supports hot reloading via mounted volume and `npm start` in Docker
* Backend Go container watches for changes via mounted volume (optionally use `air` or `reflex` for live reload)

---

## 🤝 Contributing

PRs welcome. Fork and fire away 🔥

---

## 📜 License

MIT License
