# Chapter 13 Examples: Monitoring and Observability

This directory contains examples demonstrating different monitoring and observability approaches with Genkit Go.

## Examples Overview

### 1. Google Cloud Monitoring (`cloud/`)
Shows integration with Google Cloud's monitoring services:
- Google Cloud plugin configuration
- Cloud Trace, Logging, and Monitoring integration
- Required API setup and authentication

### 2. OpenTelemetry Integration (`otel/`)
Demonstrates using OpenTelemetry for flexible monitoring:
- OTLP exporter configuration
- Integration with various backends (Jaeger, Prometheus)
- Custom metric and trace configuration

## Running the Examples

### Prerequisites

1. **AWS Credentials**: All examples use AWS Bedrock, so ensure you have:
   ```bash
   export AWS_ACCESS_KEY_ID=your_access_key
   export AWS_SECRET_ACCESS_KEY=your_secret_key
   export AWS_REGION=us-east-1
   ```

2. **Google Cloud (for cloud example)**:
   ```bash
   export GOOGLE_CLOUD_PROJECT=your-project-id
   # Ensure you have authenticated with gcloud
   gcloud auth application-default login
   ```

### Google Cloud Example

```bash
cd cloud/
go mod tidy
# Set your Google Cloud project ID
export GOOGLE_CLOUD_PROJECT=your-project-id
go run main.go
```

View traces and metrics in the Google Cloud Console.

### OpenTelemetry Example

```bash
cd otel/
go mod tidy
go run main.go
```

This example exports telemetry data using OTLP. You can configure it to work with:
- Jaeger for distributed tracing
- Prometheus for metrics
- Any OTLP-compatible backend

## Testing the Examples

You can test any example by sending a POST request:

```bash
curl -X POST http://localhost:3000/chatFlow \
  -H "Content-Type: application/json" \
  -d '{"message": "Hello, how are you?"}'
```

## Configuration Options

### OpenTelemetry Presets

The OpenTelemetry example demonstrates different presets:

- `PresetOTLP`: For OpenTelemetry Protocol exporters
- `PresetJaeger`: For Jaeger backend integration
- `PresetPrometheus`: For Prometheus metrics
- `PresetZipkin`: For Zipkin tracing

### Google Cloud Configuration

The Google Cloud example shows:

- Project ID configuration
- Force export for development
- Integration with Cloud Trace, Logging, and Monitoring
  