#!/bin/bash

# Simple Docker Build Script for futStats
# This script builds and optionally pushes your Docker image

set -e

# Configuration
IMAGE_NAME="futstats-api"
REGISTRY="bermanbenjamin"
TAG="latest"
FULL_IMAGE_NAME="${REGISTRY}/${IMAGE_NAME}:${TAG}"

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

# Build Docker image
build_image() {
    log_info "Building Docker image: ${FULL_IMAGE_NAME}"
    
    # Build the image
    docker build -t ${IMAGE_NAME} -t ${FULL_IMAGE_NAME} .
    
    log_info "Docker image built successfully!"
    
    # Show image info
    docker images ${IMAGE_NAME}
}

# Push to Docker Hub
push_image() {
    log_info "Pushing Docker image to Docker Hub: ${FULL_IMAGE_NAME}"
    
    # Check if logged in to Docker Hub
    if ! docker info | grep -q "Username"; then
        log_warn "Please login to Docker Hub first:"
        docker login
    fi
    
    # Push the image
    docker push ${FULL_IMAGE_NAME}
    
    log_info "Docker image pushed successfully!"
    log_info "Image available at: https://hub.docker.com/r/${REGISTRY}/${IMAGE_NAME}"
}

# Test image locally
test_image() {
    log_info "Testing Docker image locally..."
    
    # Stop any existing containers
    docker stop futstats-test 2>/dev/null || true
    docker rm futstats-test 2>/dev/null || true
    
    # Run the container
    docker run -d --name futstats-test \
        -p 8080:8080 \
        -e DATABASE_URL="sqlite:///tmp/test.db" \
        -e SECRET_KEY="test-key" \
        -e ENVIRONMENT="test" \
        ${FULL_IMAGE_NAME}
    
    # Wait for container to start
    sleep 5
    
    # Test health endpoint
    if curl -f http://localhost:8080/health; then
        log_info "Health check passed!"
        log_info "API is running at: http://localhost:8080"
    else
        log_error "Health check failed!"
        docker logs futstats-test
        exit 1
    fi
    
    # Cleanup
    docker stop futstats-test
    docker rm futstats-test
    
    log_info "Local test completed successfully!"
}

# Main script logic
case "${1:-build}" in
    "build")
        build_image
        ;;
    "push")
        build_image
        push_image
        ;;
    "test")
        build_image
        test_image
        ;;
    "all")
        build_image
        test_image
        push_image
        ;;
    *)
        echo "Usage: $0 [build|push|test|all]"
        echo ""
        echo "Commands:"
        echo "  build - Build Docker image"
        echo "  push  - Build and push to Docker Hub"
        echo "  test  - Build and test locally"
        echo "  all   - Build, test, and push"
        echo ""
        echo "Prerequisites:"
        echo "  - Docker installed and running"
        echo "  - Docker Hub account (for push)"
        exit 1
        ;;
esac
