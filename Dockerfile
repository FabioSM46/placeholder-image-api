# Start with a minimal base image for Go
FROM golang:1.22-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first, for dependency caching
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Build the Go app
RUN go build -o main

# Expose the port your app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
