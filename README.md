# Microservices Architecture

This project demonstrates a microservices architecture using Go, gRPC, Redis, and Docker. It consists of three main services:

1. Auth Service (gRPC)
2. Order Service (Redis)
3. Product Service (HTTP)

## Services Overview

### Auth Service (auth-grpc-microservice)
- Handles user authentication and authorization
- Implements gRPC for communication
- Port: 8000

### Order Service (order-redis-microservice)
- Manages order processing and storage
- Uses Redis for data persistence
- Port: 8081

### Product Service (product-service)
- Provides product information and handles product-related operations
- Acts as an API gateway, communicating with both Auth and Order services
- Port: 8082

## Prerequisites

- Docker
- Docker Compose
- Go 1.21 or later (for local development)

## Getting Started

1. Clone the repository:
   ```
   git clone https://github.com/caiotheodoro/microservices-project.git
   cd microservices-project
   ```

2. Build and run the services using Docker Compose:
   ```
   docker-compose up --build
   ```

3. The services will be available at:
   - Auth Service: localhost:8000 (gRPC)
   - Order Service: http://localhost:8081
   - Product Service: http://localhost:8082

## API Endpoints

### Product Service

- `GET /products`: Retrieve a list of products
- `POST /order`: Create a new order
- `GET /health`: Check the health of the service

## Development

To run the services locally for development:

1. Install dependencies for each service:
   ```
   cd auth-grpc-microservice && go mod download
   cd ../order-redis-microservice && go mod download
   cd ../product-service && go mod download
   ```

2. Run each service individually:
   ```
   cd auth-grpc-microservice && go run cmd/main.go
   cd ../order-redis-microservice && go run main.go
   cd ../product-service && go run cmd/main.go
   ```
