# Chapter 6: Building Flows - Music Instrument Recommendation

This example demonstrates Flow orchestration in Genkit Go by building a music instrument recommendation system. It showcases how to compose multiple Flows to create complex, multi-step workflows.

## Overview

The system recommends musical instruments based on genre preferences through a series of orchestrated Flows:

1. **Genre Analysis**: Categorizes music genres as "acoustic" or "electronic"
2. **Instrument Selection**: Recommends appropriate instruments based on the category
3. **Detail Generation**: Creates comprehensive recommendations with explanations

## Architecture

```text
recommendationFlow (Orchestrator)
├── analyzeGenreFlow → Categorizes genre
├── acousticInstrumentFlow → For acoustic genres
├── electronicInstrumentFlow → For electronic genres
└── Generate Details → Creates final recommendation
```

## Project Structure

```text
chapter-06/
├── main.go                          # HTTP server and Flow registration
├── internal/
│   ├── flows/
│   │   ├── analyze-genre.go        # Genre categorization Flow
│   │   ├── acoustic.go              # Acoustic instrument Flow
│   │   ├── electronic.go            # Electronic instrument Flow
│   │   └── recommendation.go        # Orchestrator Flow
│   └── structs/
│       ├── recommendation-input.go  # Client request structure
│       └── recommendation-output.go # Server response structure
├── prompts/                         # Dotprompt templates
│   ├── analyze-genre.prompt
│   ├── acoustic-instrument.prompt
│   ├── electronic-instrument.prompt
│   └── recommendation-details.prompt
├── go.mod
└── go.sum
```

## Prerequisites

- Go 1.25.0 or later
- Google AI API key (Gemini 2.5 Flash)
- Genkit CLI

## Setup

1. Set your Google AI API key:

```bash
export GOOGLE_GENAI_API_KEY="your-api-key-here"
```

2. Install dependencies:

```bash
go mod download
```

## Running the Example

### Option 1: With Developer UI (Recommended)

```bash
genkit start -- go run .
```

This starts both the application server (port 9090) and Developer UI (port 4000).

### Option 2: Standalone Server

```bash
go run .
```

The server will start on `http://localhost:9090`.

## Testing the Flows

### Using cURL

Test the orchestrator Flow:

```bash
# Jazz recommendation (acoustic path)
curl -X POST http://localhost:9090/recommendationFlow \
  -H "Content-Type: application/json" \
  -d '{
    "data": {
      "genre": "classical",
      "experience": "beginner"
    }
  }'

# EDM recommendation (electronic path)
curl -X POST http://localhost:9090/recommendationFlow \
  -H "Content-Type: application/json" \
  -d '{
    "data": {
      "genre": "edm",
      "experience": "expert"
    }
  }'
```

### Using Genkit CLI

```bash
# Test the complete orchestration
genkit flow:run recommendationFlow '{"genre":"rock","experience":"beginner"}'

# Test individual Flows
genkit flow:run analyzeGenreFlow '"techno"'
genkit flow:run acousticInstrumentFlow '"blues"'
genkit flow:run electronicInstrumentFlow '"house"'
```

### Using Developer UI

1. Open <http://localhost:4000>
2. Navigate to Flows → recommendationFlow
3. Enter input JSON:

```json
{
  "genre": "classical",
  "experience": "beginner"
}
```

4. Click Run to see the orchestration in action

## Key Concepts Demonstrated

### 1. Flow Orchestration

- **Flow-to-Flow Calls**: Using `Run()` to execute child Flows
- **Conditional Branching**: Choosing different Flows based on intermediate results
- **Data Transformation**: Passing data between Flows

### 2. Dotprompt Integration

- **Structured Prompts**: Managing prompts in `.prompt` files
- **Temperature Control**: Different temperature settings for each task
- **Schema Definition**: Using Picoschema for input/output validation

### 3. Type Safety

- **Structured Input/Output**: Using Go structs for type-safe data handling
- **JSON Marshaling**: Automatic conversion between structs and JSON

## Example Response

```json
{
  "instrument": "Acoustic Guitar",
  "why": "The acoustic guitar is perfect for jazz beginners because it provides a solid foundation for understanding chord progressions and melody. Its versatility allows you to explore various jazz styles from smooth to bebop.",
  "starter_items": [
    "Guitar picks (various thicknesses)",
    "Digital tuner",
    "Padded gig bag",
    "Jazz chord chart book",
    "Guitar strap"
  ]
}
```
