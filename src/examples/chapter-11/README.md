# Chapter 11 Example: Evaluations with Genkit Go

This example shows how to implement a evaluation systems for AI applications using Genkit Go. It showcases three types of evaluators: built-in evaluators, custom non-billed evaluators, and custom billed evaluators that use AI models for sophisticated assessment.

## Features

- **Built-in Evaluators**: Uses Genkit's pre-built deep equality evaluator
- **Custom Non-billed Evaluators**: Implements custom evaluation logic without AI model costs
- **Custom Billed Evaluators**: Sophisticated AI-powered evaluators using Anthropic's Claude
- **Chat Flow**: A simple conversational AI flow for demonstration
- **Dataset Management**: Includes test datasets for evaluation scenarios

## Project Structure

```
src/examples/chapter-11/
├── main.go                     # Main application entry point
├── go.mod                      # Go module dependencies
├── go.sum                      # Dependency checksums
├── README.md                   # This documentation
├── .gitignore                  # Git ignore patterns
└── internal/                   # Internal packages
    ├── flows/                  # Flow definitions
    │   └── chat.go             # Chat flow implementation
    ├── evals/                  # Evaluator implementations
    │   ├── genkitEvaluators.go # Built-in evaluator configurations
    │   ├── NonBilledEvaluators.go # Custom cost-free evaluators
    │   └── billedEvaluators.go # AI-powered evaluators
    └── datasets/               # Test datasets for evaluations
```

## API Endpoints

### 1. Chat Flow
**Endpoint:** `POST /chatFlow`

The main conversational flow for generating AI responses.

**Request Body:**
```json
{
  "message": "Your question or message here"
}
```

**Response:**
```json
{
  "message": "AI assistant response"
}
```

## Running the Example

### Prerequisites

1. **Anthropic API Key**: Required for the AI model and billed evaluators
   ```bash
   export ANTHROPIC_API_KEY="your-anthropic-api-key-here"
   ```

2. **Go 1.21+** installed

### Setup and Execution

1. **Navigate to the project directory:**
   ```bash
   cd src/examples/chapter-11
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the application:**
   ```bash
   go run .
   ```

4. **Access the application:**
   - Server runs on `http://localhost:9090`
   - Use the `/chatFlow` endpoint to interact with the AI
   - Evaluations are triggered automatically during development


```bash
curl -X POST http://localhost:9090/chatFlow \
  -H "Content-Type: application/json" \
  -d '{"data": {"message": "What are the benefits of using AI evaluations?"}}'
```

## Run evaluations using the Genkit CLI

```bash
genkit eval:flow chatFlow --input internal/datasets/chat-dataset.json
```