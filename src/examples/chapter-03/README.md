# Chapter 3: Setting Up Your Development Environment - Getting Started with Genkit Go

This example demonstrates the very first Genkit Go application to verify your development environment setup. It implements a simple greeting flow that accepts a user's name and generates personalized welcome messages using AI.

## Features

- Simple greeting flow using Google AI (Gemini 2.5 Flash)
- System prompt configuration for consistent AI behavior
- HTTP API endpoint for easy testing
- Minimal setup perfect for verifying installation
- Error handling and logging

## API Endpoints

### 1. Greeting Flow
**Endpoint:** `POST /greetingFlow`

Generates personalized greetings for users based on their name input.

## Running the Example

1. Set up your Google AI API key:
   ```bash
   export GEMINI_API_KEY="your-api-key-here"
   ```

2. Run the application with Developer UI:
   ```bash
   genkit start -- go run .
   ```

4. The server will start and you can make POST requests to:
   - `POST /greetingFlow` - Generate personalized greetings

## Example Usage

### Generate Greeting
Send a POST request with a user's name:
- `"Alice"`
- `"Bob"`
- `"Charlie"`

## cURL Examples

Here are practical examples of calling the flow using cURL commands:

### Generate Greeting
```bash
curl -X POST http://127.0.0.1:9090/greetingFlow \
  -H "Content-Type: application/json" \
  -d '{"data":"Alice"}'
```

**Example Response:**
```json
{
  "result": "Hello Alice! Welcome to our application! It's wonderful to have you here. I hope you have a fantastic experience exploring everything we have to offer. Thank you for joining us!"
}
```

## Running with Genkit CLI

With your application running via `genkit start -- go run .`, open another terminal and use the Genkit CLI:

```bash
genkit flow:run greetingFlow '"Alice"'
```

## Project Structure

- `main.go` - Entry point with greeting flow definition and HTTP server setup

This simple example serves as the foundation for understanding Genkit Go basics and verifying your development environment is properly configured.
