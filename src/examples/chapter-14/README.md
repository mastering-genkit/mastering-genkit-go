# Chapter 14 Examples: Deploying Genkit Applications

This directory contains examples showing how to deploy Genkit Go applications to various platforms and environments.

## Example Overview

This example showcases a production-ready Genkit Go application with:

- **Docker Support**: Containerized deployment with multi-stage builds
- **Kubernetes Manifests**: Complete K8s deployment with HPA and service configuration
- **Health Checks**: Built-in `/health` and `/ready` endpoints for container orchestration
- **Production Structure**: Organized codebase with proper separation of concerns
- **Environment Configuration**: Configurable port and runtime settings

### Application Structure

```
chapter-14/
├── main.go              # Main HTTP server with Genkit integration
├── internal/
│   ├── flows/           # Genkit flows and business logic
│   └── handlers/        # HTTP handlers for health checks
├── k8s/                 # Kubernetes deployment manifests
│   ├── deployment.yml   # Application deployment
│   ├── svc.yml         # Service configuration
│   └── hpa.yml         # Horizontal Pod Autoscaler
├── Dockerfile           # Multi-stage Docker build
└── .dockerignore        # Docker build optimization

```

## Running the Examples

### Prerequisites

1. **Go 1.21+** installed
2. **Docker** (for containerized deployment)
3. **Kubernetes cluster** (for K8s deployment) you can use Minikube or Kind for local testing

### Local Development

1. **Navigate to the example directory:**
   ```bash
   cd src/examples/chapter-14
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the application locally:**
   ```bash
   go run .
   ```

4. **Test the endpoints:**
   ```bash
   # Health check
   curl http://localhost:9091/health
   
   # Readiness check
   curl http://localhost:9091/ready
   
   # Simple flow endpoint
   curl -X POST http://localhost:9091/simpleFlow \
     -H "Content-Type: application/json" \
     -d '{"data": {"input": "Hello, world!"}}'
   ```

### Docker Deployment

1. **Build the Docker image:**
   ```bash
   docker build -t genkit-deployment-example .
   ```

2. **Run the container:**
   ```bash
   docker run -p 8080:8080 genkit-deployment-example
   ```

3. **Test the containerized application:**
   ```bash
   curl http://localhost:8080/health
   ```

### Kubernetes Deployment

1. **Apply the Kubernetes manifests:**
   ```bash
   kubectl apply -f k8s/
   ```

2. **Check the deployment status:**
   ```bash
   kubectl get pods
   kubectl get services
   ```

3. **Access the application:**
   ```bash
   # Get the service endpoint
   kubectl get svc genkit-deployment-service
   
   # Port forward for local testing
   kubectl port-forward svc/genkit-deployment-service 8080:80
   ```

## Health Checks

The application includes standard health check endpoints:

- `GET /health`: Basic health check
- `GET /ready`: Readiness probe for Kubernetes

These endpoints are essential for container orchestration platforms to manage application lifecycle.
