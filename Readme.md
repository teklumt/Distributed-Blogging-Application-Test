# Distributed Microservices Project

A distributed system built with microservices to manage API routing, authentication, blogging, notifications, and user data. The system is containerized using Docker, orchestrated by Kubernetes, and connected through RabbitMQ for real-time inter-service communication.

## 📂 Project Structure

The project is organized into distinct services with dependencies, configurations, and Kubernetes deployment files for each:

```plaintext
api-gateway/
├── .env                    # Environment variables
├── config.yaml             # API Gateway configuration
├── Dockerfile              # API Gateway containerization
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── kubernetes/
│   ├── deployment.yaml     # Kubernetes deployment configuration
│   └── service.yaml        # Kubernetes service configuration
├── main.go                 # API Gateway entry point
└── middleware/             # Middleware components for API Gateway

Auth_service/
├── cmd/                    # CLI code for Auth service
├── config/                 # Service configurations
├── delivery/
│   └── routers/
│       └── routers.go      # Auth service routing
├── Dockerfile              # Auth service containerization
├── domain/                 # Domain models and interfaces
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── infrastructure/         # Infrastructure components
├── messaging/              # Messaging components
└── repository/             # Repository implementations

Blog_service/
├── cmd/                    # CLI code for Blog service
├── config/                 # Service configurations
├── delivery/
│   └── routers/
│       └── routers.go      # Blog service routing
├── Dockerfile              # Blog service containerization
├── domain/                 # Domain models and interfaces
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── infrastructure/         # Database components
├── messaging/              # Messaging components
└── repository/             # Repository implementations

Notification_service/
├── cmd/                    # CLI code for Notification service
├── config/                 # Service configurations
├── delivery/
│   └── routers/
│       └── routers.go      # Notification service routing setup
├── Dockerfile              # Dockerfile for Notification service containerization
├── domain/                 # Domain models and interfaces
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── infrastructure/         # Infrastructure code, such as messaging connections
├── Makefile                # Build and deployment commands
├── messaging/              # RabbitMQ message-handling components
├── repository/             # Repository implementations
└── usecase/                # Use case implementations


User_service/
├── cmd/                    # CLI code for User service
├── config/                 # Service configurations
├── delivery/
│   └── routers/
│       └── routers.go      # User service routing setup
├── Dockerfile              # Dockerfile for User service containerization
├── domain/                 # Domain models and interfaces
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── infrastructure/         # Database and infrastructure components
├── Makefile                # Build and deployment commands
├── messaging/              # Messaging components
├── repository/             # Repository implementations
└── usecase/                # Use case implementations

database/
├── config/                 # Database configuration
├── deployment/             # Deployment scripts
└── schema/                 # Database schema files

k8s/
├── api-gateway/
├── Auth_service/
├── Blog_service/
├── Notification/
├── RabbitMQ/
└── User_services/
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

## 📝 License

This project is licensed under the MIT License.

---

Happy deploying and managing your microservices architecture!

```

```
