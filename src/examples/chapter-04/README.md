# Chapter 4: Mastering AI Generation

This example demonstrates comprehensive AI generation capabilities in Genkit Go, showcasing text generation, prompt templates, image analysis, audio transcription, and image generation. The application implements five distinct flows to explore different aspects of AI-powered content generation.

## Features

- Multiple AI generation patterns using Google AI (Gemini 2.5 Flash)
- Dotprompt template system for reusable prompts
- Multimodal AI capabilities (text, image, audio)
- Image generation using AI models
- Middleware integration for logging and monitoring
- Professional cooking instructor theme
- Structured flow architecture

## Flows Available

### 1. Basic Generation Flow
Simple text generation with inline prompts for cooking instruction explanations.

### 2. Dotprompt Flow
Template-based generation using external prompt files for consistent cooking guidance.

### 3. Image Analysis Flow
Analyzes images from URLs and provides detailed descriptions of visual content.

### 4. Audio Transcription Flow
Transcribes audio content from URLs into accurate text format.

### 5. Image Generation Flow
Generates images based on text descriptions using AI image models.

## API Endpoints

### 1. Basic Generation Flow
**Endpoint:** `POST /basicGenerationFlow`

Generates cooking explanations using inline prompts.

### 2. Dotprompt Flow
**Endpoint:** `POST /dotpromptFlow`

Uses template-based prompts from the `prompts/` directory for consistent responses.

### 3. Image Analysis Flow
**Endpoint:** `POST /imageAnalysisFlow`

Analyzes images from provided URLs and describes their content in detail.

### 4. Audio Transcription Flow
**Endpoint:** `POST /audioTranscriptionFlow`

Transcribes audio files from URLs into text format.

### 5. Image Generation Flow
**Endpoint:** `POST /imageGenerationFlow`

Generates images based on text descriptions.

## Running the Example

1. Set up your Google AI API key:
   ```bash
   export GEMINI_API_KEY="your-api-key-here"
   ```

2. Run the application with Developer UI:
   ```bash
   genkit start -- go run .
   ```

3. The server will start and you can make POST requests to any of the available endpoints:
   - `POST /basicGenerationFlow` - Basic cooking instruction explanations
   - `POST /dotpromptFlow` - Template-based cooking guidance
   - `POST /imageAnalysisFlow` - Image content analysis
   - `POST /audioTranscriptionFlow` - Audio transcription
   - `POST /imageGenerationFlow` - AI image generation

## Example Usage

### Basic Text Generation
Send requests like:
- `"how to make pasta"`
- `"knife safety techniques"`
- `"understanding cooking temperatures"`

### Dotprompt Template Usage
Send cooking topic requests like:
- `"seasoning basics"`
- `"food storage tips"`
- `"kitchen equipment guide"`

### Image Analysis
Send image URLs for analysis:
- `"https://example.com/dish-photo.jpg"`
- `"https://example.com/ingredients.png"`

### Audio Transcription
Send audio URLs for transcription:
- `"https://example.com/cooking-tutorial.mp3"`
- `"https://example.com/recipe-instructions.wav"`

### Image Generation
Send text descriptions for image creation:
- `"a beautiful plated pasta dish with herbs"`
- `"fresh vegetables arranged on a cutting board"`

## cURL Examples

Here are practical examples of calling each flow using cURL commands:

### Basic Generation
```bash
curl -X POST http://127.0.0.1:9090/basicGenerationFlow \
  -H "Content-Type: application/json" \
  -d '{"data":"how to properly dice onions"}'
```

### Dotprompt Template
```bash
curl -X POST http://127.0.0.1:9090/dotpromptFlow \
  -H "Content-Type: application/json" \
  -d '{"data":"knife safety techniques"}'
```

### Image Analysis
```bash
curl -X POST http://127.0.0.1:9090/imageAnalysisFlow \
  -H "Content-Type: application/json" \
  -d '{"data":"https://example.com/food-image.jpg"}'
```

### Audio Transcription
```bash
curl -X POST http://127.0.0.1:9090/audioTranscriptionFlow \
  -H "Content-Type: application/json" \
  -d '{"data":"https://example.com/cooking-audio.mp3"}'
```

### Image Generation
```bash
curl -X POST http://127.0.0.1:9090/imageGenerationFlow \
  -H "Content-Type: application/json" \
  -d '{"data":"a perfectly cooked steak with roasted vegetables"}'
```

## Running with Genkit CLI

With your application running via `genkit start -- go run .`, open another terminal and use the Genkit CLI:

```bash
# Basic generation
genkit flow:run basicGenerationFlow '"how to properly dice onions"'

# Dotprompt template
genkit flow:run dotpromptFlow '"knife safety techniques"'

# Image analysis
genkit flow:run imageAnalysisFlow '"https://example.com/food-image.jpg"'

# Audio transcription
genkit flow:run audioTranscriptionFlow '"https://example.com/cooking-audio.mp3"'

# Image generation
genkit flow:run imageGenerationFlow '"a perfectly cooked steak with roasted vegetables"'
```

## Project Structure

- `main.go` - Entry point with flow registration and HTTP server setup
- `internal/flows/` - Flow implementations
  - `basic-generation.go` - Simple text generation flow
  - `dotprompt.go` - Template-based generation flow
  - `image-analysis.go` - Image analysis flow
  - `audio-transcription.go` - Audio transcription flow
  - `image-generation.go` - Image generation flow
  - `middleware.go` - Logging middleware for all flows
- `prompts/` - Dotprompt template files
  - `cooking-instructor.prompt` - Template for cooking instruction generation

## Template System

The dotprompt flow uses external template files stored in the `prompts/` directory. The `cooking-instructor.prompt` file defines:

- Model configuration (googleai/gemini-2.5-flash)
- Input schema (topic: string)
- Template content with variable substitution

This approach promotes reusability and maintainability of prompt content across different parts of your application.

## Middleware Integration

All flows use logging middleware to provide consistent request/response monitoring. The middleware tracks:

- Request timing
- Input parameters
- Response status
- Error conditions

This demonstrates how to integrate cross-cutting concerns into Genkit Go flows for production readiness.
