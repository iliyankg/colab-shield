# Start with a base image
FROM golang:1.22-alpine

# Define the default values for the build arguments
ARG ARG_GRPC_PORT=1337
ARG ARG_HTTP_PORT=1338
ARG ARG_REDIS_HOST="redis"
ARG ARG_REDIS_PORT=6379
ARG ARG_REDIS_PASSWORD="redis"

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

ENV COLABSHIELD_GRPC_PORT=$ARG_GRPC_PORT
ENV COLABSHIELD_HTTP_PORT=$ARG_HTTP_PORT
ENV REDIS_HOST=$ARG_REDIS_HOST
ENV REDIS_PORT=$ARG_REDIS_PORT
ENV REDIS_PASSWORD=$ARG_REDIS_PASSWORD

# Deliberately using shell form of entry point to allow for environment variable substitution
ENTRYPOINT ./backend