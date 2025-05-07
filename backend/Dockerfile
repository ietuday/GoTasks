# Use the official Go image with version 1.24.2 based on Debian Bookworm (stable & secure)
FROM golang:1.24.2-bookworm

# Set the working directory inside the container to /app
WORKDIR /app

# Copy Go module files first to leverage Docker cache for faster builds on rebuild
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application source code into the container
COPY . .

# Build the Go application and output an executable named "main"
RUN go build -o main .

# Tell Docker that the application will listen on port 8080
EXPOSE 8080

# Command to run the app when the container starts
CMD ["./main"]
