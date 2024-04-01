
# Start with a base image
FROM golang:1.22-alpine

# Define the default values for the build arguments
ARG PORT=1337
ARG REDIS_HOST="redis:6379"

# Set the working directory inside the container
WORKDIR /backend

# Copy the Go module files
COPY go.mod go.sum ./

# Copy the source code into the container
COPY . .

# Download the Go module dependencies
RUN go mod download

# Build the Go application
RUN go build .

ENV ENV_PORT=$PORT
ENV ENV_REDIS_HOST=$REDIS_HOST

# Deliberately using shell form of entry point to allow for environment variable substitution
ENTRYPOINT ./backend --port $ENV_PORT --redis-address $ENV_REDIS_HOST