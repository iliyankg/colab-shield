# ----------------------------------------------------------
# Multi-stage Dockerfile for building the Go application
# ----------------------------------------------------------
FROM golang:1.22-alpine AS builder

# Define the default values for the build arguments
ARG ARG_GRPC_PORT=1337
ARG ARG_HTTP_PORT=1338
ARG ARG_REDIS_HOST="redis"
ARG ARG_REDIS_PORT=6379
ARG ARG_REDIS_PASSWORD="redis"
ARG ARG_GIN_MODE="release"

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

# ----------------------------------------------------------
# Multi-stage Dockerfile for running the Go application
# ----------------------------------------------------------
FROM golang:1.22-alpine

RUN apk add --no-cache curl

WORKDIR /app

# Copy the built Go application from the builder stage
COPY --from=builder /backend/backend /backend

# Set up environment variables
ENV COLABSHIELD_GRPC_PORT=$ARG_GRPC_PORT
ENV COLABSHIELD_HTTP_PORT=$ARG_HTTP_PORT
ENV REDIS_HOST=$ARG_REDIS_HOST
ENV REDIS_PORT=$ARG_REDIS_PORT
ENV REDIS_PASSWORD=$ARG_REDIS_PASSWORD
ENV GIN_MODE=$ARG_GIN_MODE

# Healthcheck 
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD curl --fail http://localhost:$COLABSHIELD_HTTP_PORT/health || exit 1 

EXPOSE $ARG_GRPC_PORT
EXPOSE $ARG_HTTP_PORT

# Deliberately using shell form of entry point to allow for environment variable substitution
ENTRYPOINT ./backend