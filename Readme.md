# Distributed Microservices Project
![Uploading design.png…]()

A distributed system built with microservices to manage API routing, authentication, blogging, notifications, and user data. The system is containerized using Docker, orchestrated by Kubernetes, and connected through RabbitMQ for real-time inter-service communication.

## 📂 Project Structure

The project is organized into distinct services with dependencies, configurations, and Kubernetes deployment files for each:

```plaintext
api-gateway/
├── .env                    # Environment variables for API Gateway
├── config.yaml             # API Gateway configuration file
├── Dockerfile              # API Gateway containerization
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── kubernetes/
│   ├── deployment.yaml     # Kubernetes deployment config for API Gateway
│   └── service.yaml        # Kubernetes service config for API Gateway
├── main.go                 # Entry point for the API Gateway
└── middleware/             # Middleware components for API Gateway

Auth_service/
├── cmd/                    # CLI code and entry point for Auth service
├── config/                 # Service-specific configurations
├── delivery/
│   └── routers/
│       └── routers.go      # Auth service routing configuration
├── Dockerfile              # Auth service containerization
├── domain/                 # Domain models and interfaces
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── infrastructure/         # Infrastructure components like DB connections
├── messaging/              # Messaging components for inter-service communication
└── repository/             # Repository implementations for data handling

Blog_service/
├── cmd/                    # CLI code and entry point for Blog service
├── config/                 # Service-specific configurations
├── delivery/
│   └── routers/
│       └── routers.go      # Blog service routing configuration
├── Dockerfile              # Blog service containerization
├── domain/                 # Domain models and interfaces
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── infrastructure/         # Database infrastructure components
├── messaging/              # Messaging components
└── repository/             # Repository implementations

Notification_service/
├── cmd/                    # CLI code and entry point for Notification service
├── config/                 # Service-specific configurations
├── delivery/
│   └── routers/
│       └── routers.go      # Notification service routing configuration
├── Dockerfile              # Containerization for Notification service
├── domain/                 # Domain models and interfaces
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── infrastructure/         # Messaging and database infrastructure
├── Makefile                # Build and deployment commands
├── messaging/              # RabbitMQ or other message handling components
├── repository/             # Repository implementations for data handling
└── usecase/                # Use case business logic implementations

User_service/
├── cmd/                    # CLI code and entry point for User service
├── config/                 # Service-specific configurations
├── delivery/
│   └── routers/
│       └── routers.go      # User service routing configuration
├── Dockerfile              # Containerization for User service
├── domain/                 # Domain models and interfaces
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── infrastructure/         # Database and infrastructure components
├── Makefile                # Build and deployment commands
├── messaging/              # Messaging components
├── repository/             # Repository implementations
└── usecase/                # Use case business logic implementations

database/
├── config/                 # Database configuration files
├── deployment/             # Deployment scripts for database
└── schema/                 # Database schema files

k8s/
├── api-gateway/
│   ├── api-gateway-deployment.yaml  # Deployment config for API Gateway
│   └── api-gateway-service.yaml     # Service config for API Gateway
├── Auth_service/
│   └── auth-service-deployment.yaml # Deployment config for Auth service
├── Blog_service/
│   └── blog-deployment.yaml         # Deployment config for Blog service
├── Notification/
│   └── notification-deployment.yaml # Deployment config for Notification service
├── RabbitMQ/
│   └── rabbitmq-deployment.yaml     # Deployment config for RabbitMQ
├── monitoring/
│   ├── grafana/
│   │   ├── grafana-deployment.yaml  # Deployment config for Grafana
│   │   └── grafana-service.yaml     # Service config for Grafana
│   ├── prometheus/
│   │   ├── prometheus-deployment.yaml  # Deployment config for Prometheus
│   │   └── prometheus-service.yaml     # Service config for Prometheus
│   └── alertmanager/
│       ├── alertmanager-deployment.yaml  # Deployment config for Alertmanager
│       └── alertmanager-service.yaml     # Service config for Alertmanager
└── User_service/
    └── user-service-deployment.yaml # Deployment config for User service
```

## 🚀 Microservices Architecture

The microservices are interconnected as follows:

**API Gateway** ←→ **Auth Service** ←→ **User Service**  
 ↓ ↓  
**Blog Service** ←→ **Notification Service** ←→ **RabbitMQ**

### Microservice Descriptions:

- **API Gateway**: Routes requests to appropriate services and handles request validation.
- **Auth Service**: Manages authentication, session, and JWT issuance.
- **Blog Service**: Handles blog content management.
- **Notification Service**: Sends notifications through RabbitMQ, delivering messages to users.
- **User Service**: Manages user data, including profile updates and retrieval.

## 🛠️ Manual Deployment

1. **Set Up Database**:

   - Configure and deploy the database by following scripts in the `database/` folder.

2. **Build Docker Containers**:

   - For each service (e.g., `api-gateway`, `Auth_service`), navigate to its directory and run:
     ```bash
     docker build -t <service-name> .
     ```

3. **Deploy Services on Kubernetes**:

   - Apply Kubernetes configurations for each service:
     ```bash
     kubectl apply -f k8s/<service-directory>/
     ```
   - Example:
     ```bash
     kubectl apply -f k8s/api-gateway/
     kubectl apply -f k8s/Auth_service/
     # Repeat for other services
     ```

4. **Run RabbitMQ for Messaging**:

   - Deploy RabbitMQ from `k8s/RabbitMQ/rabbitmq-deployment.yaml`:
     ```bash
     kubectl apply -f k8s/RabbitMQ/rabbitmq-deployment.yaml
     ```

5. **Access the System**:
   - The API Gateway serves as the main entry point. Use its external IP to connect to the system.

## 📬 Messaging System

RabbitMQ is deployed to enable seamless message-based communication among services. The Notification Service uses it to trigger real-time notifications based on event messages.

```

```
