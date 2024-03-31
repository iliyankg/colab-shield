# Start with a base image
FROM golang:1.22-alpine

ARG PORT=1337
ARG REDIS_HOST=redis

# Set the working directory inside the container
WORKDIR /backend

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

#TODO: Only for dev purposes and should be removed.
COPY ../protos ../protos 

# Build the Go application
RUN go build -o backend

ENV ENV_PORT=$PORT
ENV ENV_REDIS_HOST=$REDIS_HOST

# Set the entry point for the container
CMD ["./backend", "--port", $ENV_PORT, "--redis-address", $ENV_REDIS_HOST]
