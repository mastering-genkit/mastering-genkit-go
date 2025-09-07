# Chapter 15: Client Integration Patterns - Cooking Battle App

## Overview

This example demonstrates how to integrate Genkit Go applications with various client frameworks (Next.js, Angular, Flutter) using a unified API contract. The application theme is a "Cooking Battle" where users can compete by creating dishes based on given ingredients and constraints.

## Architecture

```mermaid
graph TB
    subgraph "Client Applications"
        NC[Next.js Client]
        AC[Angular Client]
        FC[Flutter Client]
    end
    
    subgraph "API Gateway"
        CR[Cloud Run / Local Server<br/>:9090]
    end
    
    subgraph "Security Layer"
        FA[Firebase Authentication<br/>Anonymous Auth]
        CORS[CORS Middleware]
    end
    
    subgraph "Monitoring"
        OT[OpenTelemetry<br/>Traces & Metrics]
        GCM[Google Cloud Monitoring<br/>Logs & Dashboards]
    end
    
    subgraph "Genkit Flows"
        CBChat[cookingBattleChat<br/>DefineStreamingFlow<br/>SSE]
        CBAction[cookingBattleAction<br/>DefineFlow<br/>REST]
    end
    
    subgraph "Tools (Chapter 8)"
        T1[searchRecipeDatabase]
        T2[checkIngredientStock]
        T3[calculateNutrition]
    end
    
    subgraph "Data Layer"
        FS[Firestore<br/>recipes collection]
    end
    
    subgraph "AI Models"
        Gemini[Gemini-2.5-flash<br/>Text Generation & Analysis]
        Imagen[Imagen4<br/>Image Generation]
        GeminiImg[Gemini-2.5-flash-image<br/>Image Generation]
    end
    
    NC -->|POST /{flowName}<br/>Bearer Token| CR
    AC -->|POST /{flowName}<br/>Bearer Token| CR
    FC -->|genkit package<br/>Bearer Token| CR
    
    CR --> FA
    FA -->|Verify Anonymous Token| CORS
    CORS --> OT
    OT --> CBChat
    OT --> CBAction
    
    CBChat -->|Stream Response| Gemini
    CBChat -.->|Tool Calling| T1
    CBChat -.->|Tool Calling| T2
    CBChat -.->|Tool Calling| T3
    
    CBAction -->|Generate Image| Imagen
    CBAction -->|Generate Image| GeminiImg
    CBAction -->|Analyze Image| Gemini
    CBAction -.->|Tool Calling| T1
    CBAction -.->|Tool Calling| T2
    CBAction -.->|Tool Calling| T3
    
    T1 -->|Query| FS
    T2 -->|Query| FS
    
    OT -->|Export| GCM
    
    style CBChat fill:#e1f5fe
    style CBAction fill:#fff3e0
    style T1 fill:#f3e5f5
    style T2 fill:#f3e5f5
    style T3 fill:#f3e5f5
    style FS fill:#fff8e1
    style OT fill:#e8f5e9
```

## Application Flow

### 1. Authentication Flow

- Clients obtain anonymous authentication token from Firebase Authentication
- All API requests include `Authorization: Bearer <token>` header
- Server validates `sign_in_provider == "anonymous"`

### 2. Cooking Battle Flow

1. **Ingredient Selection**: Users choose ingredients through chat interaction
2. **Recipe Generation**: AI generates recipes based on constraints
3. **Cooking Simulation**: Step-by-step guidance via streaming chat
4. **Dish Completion**: Image generation using Imagen4 or Gemini-2.5-flash-image
5. **Evaluation**: Image analysis using Gemini-2.5-flash
6. **Battle Result**: Winner determination based on evaluation

### 3. Response Types

- **Streaming (SSE)**: `cookingBattleChat` returns `text/event-stream`
- **REST (JSON)**: `cookingBattleAction` returns `application/json`
- Client automatically handles response based on `Content-Type` header

## Directory Structure

```text
chapter-15/
├── server/                      # Genkit Go Server
│   ├── internal/
│   │   ├── flows/              # Flow definitions
│   │   │   ├── chat.go         # Streaming chat flow
│   │   │   └── action.go       # Non-streaming action flow
│   │   ├── handlers/           # HTTP handlers
│   │   │   ├── flow.go         # Main flow handler
│   │   │   └── middleware.go   # Auth & CORS middleware
│   │   ├── tools/              # Tool implementations (Chapter 8)
│   │   │   ├── recipe.go       # Recipe database search
│   │   │   ├── ingredients.go  # Ingredient stock check
│   │   │   └── nutrition.go    # Nutrition calculator
│   │   └── structs/            # Data structures
│   │       ├── client/
│   │       │   ├── chat_input.go    # Chat flow request DTOs
│   │       │   ├── chat_output.go   # Chat flow response DTOs
│   │       │   ├── action_input.go  # Action flow request DTOs
│   │       │   └── action_output.go # Action flow response DTOs
│   │       ├── domain/
│   │       │   ├── recipe.go        # Recipe domain model
│   │       │   ├── stock.go         # Ingredient stock model
│   │       │   └── battle.go        # Battle result model
│   │       ├── tools/
│   │       │   ├── recipe.go        # Recipe tool input structs
│   │       │   ├── ingredients.go   # Ingredient tool input structs
│   │       │   └── nutrition.go     # Nutrition tool input structs
│   │       └── error.go        # Error responses
│   ├── prompts/                # Dotprompt templates
│   │   └── chat.prompt
│   ├── main.go                 # Entry point
│   ├── go.mod
│   ├── go.sum
│   └── README.md               # Server setup guide
│
├── terraform/                   # Infrastructure as Code
│   ├── firestore.tf            # Firestore document definitions
│   ├── variables.tf            # Terraform variables
│   └── README.md               # Terraform setup guide
│
├── client/                      # Client Applications
│   ├── next/                   # Next.js implementation
│   │   ├── src/
│   │   │   └── lib/
│   │   │       └── genkit-client.ts
│   │   ├── package.json
│   │   └── README.md
│   ├── angular/                # Angular implementation
│   │   ├── src/
│   │   │   └── app/
│   │   │       └── services/
│   │   │           └── genkit.service.ts
│   │   ├── package.json
│   │   └── README.md
│   └── flutter/                # Flutter implementation
│       ├── lib/
│       │   └── services/
│       │       └── genkit_service.dart
│       ├── pubspec.yaml
│       └── README.md
│
└── README.md                   # This file
```

## Technology Stack

### Server Side

- **Framework**: Genkit Go
- **Runtime**: Cloud Run (production) / Local server (development)
- **Authentication**: Firebase Authentication (anonymous auth)
- **Database**: Firestore (recipe storage)
- **Monitoring**: OpenTelemetry, Google Cloud Monitoring
- **AI Models**:
  - Gemini-2.5-flash (text generation & image analysis)
  - Imagen4 or Gemini-2.5-flash-image (image generation)

### Client Side

- **Next.js**: TypeScript, fetch API, eventsource-parser
- **Angular**: TypeScript, HttpClient, RxJS
- **Flutter**: Dart, genkit package

## Firestore Schema

### Collection: `recipes`

The application uses a single Firestore collection to store recipe data and ingredient information:

```json
{
  "collection": "recipes",
  "documents": [
    {
      "document_id": "recipe_001",
      "fields": {
        "name": "Classic Tomato Pasta",
        "category": "Italian",
        "difficulty": "easy",
        "prepTime": 15,
        "cookTime": 20,
        "ingredients": [
          {
            "name": "pasta",
            "amount": 200,
            "unit": "g"
          },
          {
            "name": "tomato",
            "amount": 3,
            "unit": "pieces"
          },
          {
            "name": "garlic",
            "amount": 2,
            "unit": "cloves"
          }
        ],
        "nutrition": {
          "calories": 450,
          "protein": 15,
          "carbs": 65,
          "fat": 12
        },
        "tags": ["vegetarian", "quick", "budget-friendly"]
      }
    },
    {
      "document_id": "recipe_002",
      "fields": {
        "name": "Grilled Chicken Teriyaki",
        "category": "Japanese",
        "difficulty": "medium",
        "prepTime": 30,
        "cookTime": 25,
        "ingredients": [
          {
            "name": "chicken",
            "amount": 500,
            "unit": "g"
          },
          {
            "name": "soy_sauce",
            "amount": 50,
            "unit": "ml"
          },
          {
            "name": "mirin",
            "amount": 30,
            "unit": "ml"
          }
        ],
        "nutrition": {
          "calories": 380,
          "protein": 42,
          "carbs": 18,
          "fat": 15
        },
        "tags": ["protein-rich", "japanese", "grilled"]
      }
    }
  ]
}
```

### Field Descriptions

| Field | Type | Description |
|-------|------|-------------|
| `name` | string | Recipe name |
| `category` | string | Cuisine category (Italian, Japanese, etc.) |
| `difficulty` | string | Difficulty level (easy, medium, hard) |
| `prepTime` | number | Preparation time in minutes |
| `cookTime` | number | Cooking time in minutes |
| `ingredients` | array | List of ingredients with amount and unit |
| `nutrition` | object | Nutritional information per serving |
| `tags` | array | Search tags for filtering |

### Terraform Configuration

The Firestore data is provisioned using Terraform. The configuration uses the following resources:

- [`google_firestore_document`](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/firestore_document) - Creates Firestore documents
- [`google_firestore_field`](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/firestore_field) - Configures field-level settings (if needed)

See `terraform/firestore.tf` for the complete configuration. To apply:

```bash
cd terraform/
terraform init
terraform plan -var="project_id=your-project-id"
terraform apply -var="project_id=your-project-id"
```

## API Contract

### Endpoint

```text
POST /{flowName}
```

### Request Headers

```text
Authorization: Bearer <ID_TOKEN>
Content-Type: application/json
```

### Request Body

```json
{
  "messages": [...],
  "ingredients": [...],
  "constraints": {...},
  "action": "generate_recipe|create_image|evaluate"
}
```

### Response Types

#### Streaming (SSE)

```text
Content-Type: text/event-stream

data: {"type": "content", "content": "..."}
data: {"type": "done"}
```

#### Non-Streaming (JSON)

```text
Content-Type: application/json

{
  "result": {...},
  "metadata": {...}
}
```

## Prerequisites

### Required Software

- **Go 1.22+** - For Genkit Go server development
- **Node.js 20+** - For Genkit CLI and client applications
- **Terraform CLI** - For provisioning Firestore data
- **Google Cloud CLI** (`gcloud`) - For Cloud Run deployment
- **Firebase CLI** - For Firebase project management

### Firebase Project Setup

1. **Create Firebase Project**
   ```bash
   firebase projects:create your-project-id
   firebase projects:use your-project-id
   ```

2. **Enable Firebase Authentication**
   - Go to Firebase Console > Authentication
   - Enable "Anonymous" sign-in method
   - This allows users to authenticate without credentials

3. **Initialize Firestore Database**
   - Go to Firebase Console > Firestore Database
   - Create database in production mode
   - Select location: `asia-northeast1` (or your preferred region)
   - Database will be populated via Terraform

### Google Cloud Setup

1. **Enable Required APIs**
   ```bash
   gcloud services enable \
     cloudrun.googleapis.com \
     firestore.googleapis.com \
     generativelanguage.googleapis.com \
     firebase.googleapis.com
   ```

2. **Obtain API Keys**
   - Create Google AI API key: https://aistudio.google.com/apikey
   - Store securely (will be used as `GOOGLE_GENAI_API_KEY`)

### Terraform Setup

1. **Install Terraform CLI**
   ```bash
   # macOS
   brew install terraform
   
   # Or download from: https://www.terraform.io/downloads
   ```

2. **Authenticate with Google Cloud**
   ```bash
   gcloud auth application-default login
   ```

## Quick Start

### Server Setup

```bash
cd server/
go mod init chapter-15/server
go get github.com/firebase/genkit/go@latest

# Development mode
genkit start -- go run .

# Production deployment
gcloud run deploy --source . --port 9090
```

### Client Setup

See individual README files in each client directory:

- [Next.js Setup](client/next/README.md)
- [Angular Setup](client/angular/README.md)
- [Flutter Setup](client/flutter/README.md)

## Environment Variables

### Server

```bash
PROJECT_ID=firebase-genkit-sample
OPENAI_API_KEY=your-openai-api-key
GEMINI_API_KEY=your-gemini-api-key
PORT=9090
```

### Client

```bash
NEXT_PUBLIC_API_URL=http://localhost:9090  # or Cloud Run URL
NEXT_PUBLIC_FIREBASE_CONFIG='{...}'        # Firebase config JSON
```

## Development Workflow

1. Start the server in development mode with Genkit Developer UI
2. Test flows using Developer UI at http://localhost:4000
3. Run client applications pointing to local server
4. Deploy to Cloud Run for production testing

### Local Development with Firestore Emulator

For Go applications, you need to run the Firestore Emulator and Genkit server separately:

**Terminal 1: Firestore Emulator (with seed data)**
```bash
cd src/examples/chapter-15/server
firebase emulators:start --only firestore --import=./firestore-data/local
```

**Terminal 2: Genkit Go Server (development mode)**
```bash
cd src/examples/chapter-15/server
export PROJECT_ID=firebase-genkit-sample
export FIRESTORE_EMULATOR_HOST=127.0.0.1:8080
# Add your API keys if needed:
# export GEMINI_API_KEY=your-gemini-api-key
# export OPENAI_API_KEY=your-openai-api-key
genkit start -- go run .
```

**Note**: Unlike JavaScript/Node.js applications, Go applications cannot use the combined command `genkit start -- firebase emulators:start` as they run separate HTTP servers. The Go application automatically connects to the Firestore Emulator when `FIRESTORE_EMULATOR_HOST` is set.

### Run flows from CLI (flow:run)

```bash
# Streaming chat
genkit flow:run cookingBattleChat --input='{"messages":[{"role":"user","content":"What can I cook with tomato and pasta?"}],"ingredients":["tomato","pasta"],"constraints":{"diet":"vegetarian"}}'

# Generate recipe (non-streaming)
genkit flow:run cookingBattleAction --input='{"action":"generate_recipe","ingredients":["tomato","pasta"]}'

# Create image (optional)
genkit flow:run cookingBattleAction --input='{"action":"create_image","dishName":"Tomato Pasta","description":"Light pasta with fresh tomatoes and basil"}'

# Evaluate dish (optional)
genkit flow:run cookingBattleAction --input='{"action":"evaluate","dishName":"Tomato Pasta","description":"Contest plating","imageUrl":"data:image/png;base64,...."}'
```

## Security Considerations

- API keys are managed server-side only (never exposed to clients)
- Anonymous authentication prevents abuse while maintaining easy access
- CORS configured for specific origins in production
- Rate limiting can be added via middleware

## Resources

- [Genkit Documentation](https://genkit.dev)
- [Firebase Authentication](https://firebase.google.com/docs/auth)
- [Cloud Run Documentation](https://cloud.google.com/run/docs)
- [Terraform Google Provider](https://registry.terraform.io/providers/hashicorp/google/latest/docs)
- [OpenTelemetry Go](https://opentelemetry.io/docs/languages/go/)
