# Use official Golang image as the base
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o /go-rest-api

# Expose the API port
EXPOSE 8080

# Run the Go app
CMD ["/go-rest-api"]
