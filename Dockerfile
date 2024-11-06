# Start from a Go base image
FROM golang:1.20-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Expose port 8081 for Service2
EXPOSE 8081

# Command to run the application
CMD ["go", "run", "main.go"]
