
```markdown
# Distributed Microservices Project

A robust distributed system featuring microservices for API management, authentication, blogging, notifications, and user management. This project is containerized with Docker and deployed on Kubernetes, utilizing RabbitMQ for efficient inter-service messaging.

## 📂 Project Structure

The repository is organized into distinct services, each with its own set of dependencies, configurations, and Kubernetes deployment scripts:

```plaintext
api-gateway/
├── .env                    # Environment variables
├── config.yaml             # Configuration file for the API Gateway
├── Dockerfile              # Dockerfile for API Gateway containerization
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── kubernetes/
│   ├── deployment.yaml     # Kubernetes deployment configuration
│   └── service.yaml        # Kubernetes service configuration
├── main.go                 # Main entry point for the API Gateway
├── Makefile                # Build and deployment commands
└── middleware/             # API Gateway middleware components

Auth_service/
├── cmd/                    # CLI code
├── config/                 # Service configurations
├── delivery/
│   └── routers/
│       └── routers.go      # Auth service routing setup
├── Dockerfile              # Dockerfile for Auth service containerization
├── domain/                 # Domain models and interfaces
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── infrastructure/         # Database and infrastructure components
├── Makefile                # Build and deployment commands
├── messaging/              # Messaging components
├── repository/             # Repository implementations
└── usecase/                # Use case implementations

Blog_service/
├── cmd/                    # CLI code
├── config/                 # Service configurations
├── delivery/
│   └── routers/
│       └── routers.go      # Blog service routing setup
├── Dockerfile              # Dockerfile for Blog service containerization
├── domain/                 # Domain models and interfaces
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── infrastructure/         # Database and infrastructure components
├── Makefile                # Build and deployment commands
├── messaging/              # Messaging components
├── repository/
│   └── blog_repository.go  # Repository implementation
└── usecase/
    └── blog_usecase.go     # Use case implementations

database/
├── config/                 # Database configuration files
├── deployment/             # Database deployment scripts
└── schema/                 # Database schema files

k8s/
├── api-gateway/
│   ├── api-gateway-deployment.yaml  # Deployment configuration for API Gateway
│   └── api-gateway-service.yaml     # Service configuration for API Gateway
├── Auth_service/
│   └── auth-service-deployment.yaml # Deployment configuration for Auth service
├── Blog_service/
│   └── blog-deployment.yaml         # Deployment configuration for Blog service
├── Notification/
│   └── notification-deployment.yaml # Deployment configuration for Notification service
├── RabbitMQ/
│   └── rabbitmq-deployment.yaml     # RabbitMQ deployment configuration
└── User_services/
    └── user-service-deployment.yaml # Deployment configuration for User service

Makefile                # Project-wide build and deployment commands
Notification_service/   # Notification service code and configurations
User_service/           # User service code and configurations
```

## 🚀 Microservices Overview

- **API Gateway**: Central entry point that manages routing, request validation, and load balancing.
- **Auth Service**: Responsible for user authentication, session management, and JWT token issuance.
- **Blog Service**: Handles blog content management including creation, updating, retrieval, and deletion.
- **Notification Service**: Manages notifications via RabbitMQ, delivering real-time messages to users.
- **User Service**: Manages user data, including profiles, updates, and retrieval.

## 🛠️ Database Setup

The `database` folder contains essential configurations and deployment scripts, including schema definitions, for setting up and initializing the database for each microservice.

## ☸️ Kubernetes Deployment

Each service has dedicated Kubernetes manifests located in the `k8s/` folder. These manifests include deployment and service configurations to manage container orchestration and network access.

## 📬 Messaging System

RabbitMQ enables seamless message-based communication across services. Deployment details are found in `k8s/RabbitMQ/rabbitmq-deployment.yaml`.

## 🏗️ Building and Deploying

1. **Clone the repository** and navigate to the project root.
2. **Build the services** by running:
   ```bash
   make build
   ```
3. **Deploy to Kubernetes**:
   ```bash
   make deploy
   ```
4. **Access** the application through the API Gateway.

## 📝 License

This project is licensed under the MIT License.

---

Enjoy deploying your distributed system!
```