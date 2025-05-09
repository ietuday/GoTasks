# Use the official Go image
FROM golang:1.24.2-bookworm

# Set the working directory
WORKDIR /app

# Install air for hot reload
RUN go install github.com/air-verse/air@v1.61.7

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy everything else
COPY . .

# Expose port 8080 for the app
EXPOSE 8080

# Run using air (auto-reloads on changes)
CMD ["air"]
