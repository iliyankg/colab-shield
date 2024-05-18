#!/bin/bash

# Check if Docker Compose is already up
if docker-compose -f ../docker-compose.yml ps | grep -q "Up"; then
    echo "Docker Compose is already up."

    # Stop the backend service
    docker-compose -f ../docker-compose.yml stop backend

    # Remove the backend service before rebuilding
    yes Y | docker-compose -f ../docker-compose.yml rm backend

    # Rebuild the backend service
    docker-compose -f ../docker-compose.yml build backend || { echo "Failed to build the backend service." ; exit 1; }

    # Start the backend service
    docker-compose -f ../docker-compose.yml up backend

    # Stop the backend service
    docker-compose -f ../docker-compose.yml stop backend
else
    echo "Docker Compose is not up."

    # Start the Docker Compose
    docker-compose -f ../docker-compose.yml up --build -d  ||  { echo "Failed to build the backend service." ; exit 1; }

    # Tail the logs of the backend service
    docker-compose -f ../docker-compose.yml logs -f backend

    # Stop the backend service
    docker-compose -f ../docker-compose.yml stop backend
fi