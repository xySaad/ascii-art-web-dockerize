#!/bin/bash

# Clean up unused Docker containers, images, and system resources
docker system prune -f

# Define the image name
IMAGE_NAME="ascii-art-web_image"

# Build the Docker image
docker build -t $IMAGE_NAME .

# Check if a container with the same name is running and stop it
CONTAINER_ID=$(docker ps -q -f ancestor=$IMAGE_NAME)
echo container id is $CONTAINER_ID
if [ -n "$CONTAINER_ID" ]; then
  docker stop $CONTAINER_ID
  docker rm $CONTAINER_ID
fi

# Run the Docker container with the specified name and port mapping
docker run -d -p 8080:8080 --name ascii-app $IMAGE_NAME

docker system prune -f