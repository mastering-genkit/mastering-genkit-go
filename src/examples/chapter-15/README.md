# Chapter 15: Client Integration Patterns - Recipe Quest App

## Overview

This example demonstrates how to integrate Genkit Go applications with various client frameworks (Next.js, Angular, Flutter) using a unified API contract. The application theme is a "Recipe Quest" where users can discover new recipes by getting random ingredients and letting AI guide them through recipe creation, visualization, and evaluation with achievements and chef titles.

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
        CreateRecipe[createRecipe<br/>DefineStreamingFlow<br/>SSE]
        CreateImage[createImage<br/>DefineFlow<br/>REST]
        CookingEvaluate[cookingEvaluate<br/>DefineFlow<br/>REST with Titles]
    end
    
    subgraph "Educational Tools (Chapter 8)"
        T1[checkIngredientCompatibility<br/>Firestore Tool]
        T2[estimateCookingDifficulty<br/>Calculation Tool]
    end
    
    subgraph "Data Layer"
        FS[Firestore<br/>ingredient_combinations collection]
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
    OT --> CreateRecipe
    OT --> CreateImage
    OT --> CookingEvaluate
    
    CreateRecipe -->|Stream Response| Gemini
    CreateRecipe -.->|Tool Calling| T1
    CreateRecipe -.->|Tool Calling| T2
    
    CreateImage -->|Generate Image| Imagen
    CreateImage -->|Generate Image| GeminiImg
    
    CookingEvaluate -->|Analyze Recipe| Gemini
    
    T1 -->|Query| FS
    
    OT -->|Export| GCM
    
    style CreateRecipe fill:#e1f5fe
    style CreateImage fill:#fff3e0
    style CookingEvaluate fill:#f3e5f5
    style T1 fill:#f3e5f5
    style T2 fill:#f3e5f5
    style FS fill:#fff8e1
    style OT fill:#e8f5e9
```

## Application Flow

### 1. Authentication Flow

- Clients obtain anonymous authentication token from Firebase Authentication
- All API requests include `Authorization: Bearer <token>` header
- Server validates `sign_in_provider == "anonymous"`

### 2. Recipe Quest Flow

1. **Random Ingredient Challenge**: System randomly selects 3-4 ingredients from predefined pools
2. **Recipe Generation** (Streaming): AI creates custom recipe using ingredient compatibility and difficulty tools
3. **Dish Visualization**: Image generation using Imagen3 (googleai/imagen-3.0-generate-002)
4. **AI Judge Evaluation**: Recipe analysis with detailed feedback and scoring
5. **Quest Complete**: Final score, chef title, and achievements awarded
   - Titles: "🏆 Legendary Quest Master", "⭐ Elite Recipe Explorer", "📚 Recipe Student", etc.
   - Achievements: "Innovation Master", "Technique Virtuoso", "Triple Crown Winner"

### 3. Response Types

- **Streaming (SSE)**: `createRecipe` returns `text/event-stream` with progressive recipe generation
- **REST (JSON)**: `createImage` and `cookingEvaluate` return `application/json`
- Client automatically handles response based on `Content-Type` header

## Directory Structure

```text
chapter-15/
├── server/                      # Genkit Go Server
│   ├── internal/
│   │   ├── flows/              # Flow definitions
│   │   │   ├── generate_recipe.go  # Streaming recipe generation flow
│   │   │   ├── create_image.go     # Image generation flow
│   │   │   └── evaluate.go         # Evaluation flow (with titles & achievements)
│   │   ├── tools/              # Educational tool implementations
│   │   │   ├── compatibility.go    # Ingredient compatibility checker (Firestore)
│   │   │   └── difficulty.go       # Cooking difficulty estimator (Calculation)
│   │   └── structs/            # Data structures
│   │       ├── client/
│   │       │   ├── recipe_input.go     # Recipe flow request DTOs
│   │       │   ├── recipe_output.go    # Recipe flow response DTOs
│   │       │   ├── image_input.go      # Image flow request DTOs
│   │       │   ├── image_output.go     # Image flow response DTOs
│   │       │   ├── evaluate_input.go   # Evaluate flow request DTOs
│   │       │   └── evaluate_output.go  # Evaluate flow response DTOs (with titles)
│   │       ├── tools/
│   │       │   ├── compatibility.go    # Compatibility tool input/output structs
│   │       │   └── difficulty.go       # Difficulty tool input/output structs
│   │       └── error.go        # Error responses
│   ├── firestore-data/         # Firestore configuration
│   │   ├── local/              # Local emulator data
│   │   └── remote/
│   │       └── firestore.tf    # Terraform Firestore configuration
│   ├── main.go                 # Entry point
│   ├── go.mod
│   └── go.sum
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

### Collection: `ingredient_combinations`

Recipe Quest uses a simplified educational schema focused on ingredient compatibility analysis:

```json
{
  "collection": "ingredient_combinations",
  "documents": [
    {
      "document_id": "combo_001",
      "fields": {
        "ingredients": ["chicken", "rice"],
        "compatibility_score": 9,
        "flavor_profile": "savory",
        "cuisine_style": "asian",
        "tips": "Perfect for comfort food and one-pot dishes",
        "difficulty_bonus": 1
      }
    },
    {
      "document_id": "combo_002",
      "fields": {
        "ingredients": ["salmon", "lemon"],
        "compatibility_score": 10,
        "flavor_profile": "fresh",
        "cuisine_style": "mediterranean",
        "tips": "Classic pairing with bright citrus notes",
        "difficulty_bonus": 0
      }
    },
    {
      "document_id": "combo_003",
      "fields": {
        "ingredients": ["pasta", "garlic"],
        "compatibility_score": 8,
        "flavor_profile": "aromatic",
        "cuisine_style": "italian",
        "tips": "Foundation of many Italian classics",
        "difficulty_bonus": 0
      }
    }
  ]
}
```

### Field Descriptions

| Field | Type | Description |
|-------|------|-------------|
| `ingredients` | array | Ingredient combination (2-4 ingredients) |
| `compatibility_score` | number | Compatibility rating (1-10) |
| `flavor_profile` | string | Flavor characteristic (savory, fresh, aromatic, etc.) |
| `cuisine_style` | string | Cuisine style (asian, mediterranean, italian, etc.) |
| `tips` | string | Cooking tips and recommendations |
| `difficulty_bonus` | number | Bonus points for difficulty estimation (0-3) |

### Terraform Configuration

The Firestore data is provisioned using Terraform. The configuration uses the following resources:

- [`google_firestore_document`](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/firestore_document) - Creates Firestore documents
- [`google_firestore_field`](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/firestore_field) - Configures field-level settings (if needed)

See `server/firestore-data/remote/firestore.tf` for the complete configuration. To apply:

```bash
cd server/firestore-data/remote/
terraform init
terraform plan -var="project_id=your-project-id"
terraform apply -var="project_id=your-project-id"
```

## API Contract

### Endpoints

Recipe Quest uses dedicated endpoints for each flow:

```text
POST /generateRecipe      # Streaming recipe generation
POST /createImage         # Image generation
POST /evaluateDish        # Evaluation with titles & achievements
```

### Request Headers

```text
Authorization: Bearer <ID_TOKEN>
Content-Type: application/json
```

### Request Bodies

#### Recipe Generation (Streaming)
```json
{
  "ingredients": ["chicken", "rice", "garlic", "soy sauce"],
  "constraints": {"difficulty": "medium", "time": "quick"}
}
```

#### Image Generation
```json
{
  "dishName": "One-Pan Garlic-Soy Chicken with Rice",
  "description": "A delicious one-pan dish perfect for weeknight dinner"
}
```

#### Recipe Evaluation
```json
{
  "dishName": "One-Pan Garlic-Soy Chicken with Rice",
  "description": "A delicious one-pan dish...",
  "constraints": {"difficulty": "medium", "time": "quick"}
}
```

### Response Types

#### Streaming (SSE) - Recipe Generation

```text
Content-Type: text/event-stream

data: {"type": "content", "content": "🍳 Starting recipe creation..."}
data: {"type": "content", "content": "👨‍🍳 Analyzing ingredients..."}
data: {"type": "content", "content": "Recipe name: One-Pan Garlic-Soy..."}
data: {"type": "done"}
```

#### Non-Streaming (JSON) - Image Generation

```json
{
  "success": true,
  "imageUrl": "data:image/png;base64,iVBORw0KG...",
  "dishName": "One-Pan Garlic-Soy Chicken with Rice",
  "error": ""
}
```

#### Non-Streaming (JSON) - Recipe Evaluation

```json
{
  "success": true,
  "score": 60,
  "feedback": "Detailed professional feedback...",
  "creativityScore": 5,
  "techniqueScore": 6,
  "appealScore": 7,
  "title": "📚 Recipe Student",
  "achievement": "Completed Recipe Quest",
  "error": ""
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
# Streaming recipe generation (note: -s flag for streaming)
genkit flow:run createRecipe -s --input='{"ingredients":["chicken","rice","garlic","soy sauce"],"constraints":{"difficulty":"medium","time":"quick"}}'

# Generate dish image
genkit flow:run createImage --input='{"dishName":"One-Pan Garlic-Soy Chicken with Rice","description":"A delicious one-pan dish perfect for weeknight dinner"}'

# Evaluate recipe with chef titles and achievements
genkit flow:run cookingEvaluate --input='{"dishName":"One-Pan Garlic-Soy Chicken with Rice","description":"A delicious one-pan dish with chicken, rice, garlic and soy sauce. Perfect for quick weeknight dinner.","constraints":{"difficulty":"medium","time":"quick"}}'
```

## Educational Tools Implementation

Recipe Quest features two carefully designed tools that serve as educational examples:

### 1. checkIngredientCompatibility (Firestore Tool)

**Purpose**: Analyzes ingredient combinations using Firestore database
**Educational Value**: Demonstrates Firestore queries and complex data operations

**Features**:
- Searches `ingredient_combinations` collection for matching pairs
- Returns compatibility scores (1-10), flavor profiles, and cuisine styles
- Provides cooking tips and difficulty bonuses
- Handles partial matches and fallback estimates

**Example Response**:
```json
{
  "ingredients": ["chicken", "rice"],
  "compatibilityScore": 9,
  "flavorProfile": "savory", 
  "cuisineStyle": "asian",
  "tips": "Perfect for comfort food and one-pot dishes",
  "difficultyBonus": 1,
  "overallRating": "Perfect Match"
}
```

### 2. estimateCookingDifficulty (Calculation Tool)

**Purpose**: Calculates cooking difficulty based on multiple factors
**Educational Value**: Shows complex business logic and conditional calculations

**Features**:
- Analyzes ingredient count, cooking steps, and cooking methods
- Estimates preparation time and required equipment
- Determines skill requirements and provides helpful tips
- Returns structured difficulty assessment

**Example Response**:
```json
{
  "level": "Medium",
  "score": 6,
  "reasoning": "Medium difficulty (score 6/10): Using 4 ingredients, 5 cooking steps required, Involves fry, simmer",
  "timeEstimate": 35,
  "skillsRequired": ["basic knife skills", "heat control"],
  "equipmentRequired": ["cutting board", "knife", "frying pan", "pot"],
  "tips": "Good challenge for developing skills. Prep all ingredients before cooking and watch your timing."
}
```

## Security Considerations

- API keys are managed server-side only (never exposed to clients)
- Anonymous authentication prevents abuse while maintaining easy access
- CORS configured for specific origins in production
- Rate limiting can be added via middleware

## Connecting to Production Firestore

While this chapter focuses on local development with the Firestore emulator, you can deploy the ingredient compatibility data to a production Firestore instance using the provided Terraform configuration.

### Prerequisites

- A Google Cloud Project with Firebase enabled
- Terraform 1.0 or later installed
- Google Cloud CLI (`gcloud`) authenticated

### Deploying Master Data with Terraform

Navigate to the Terraform configuration directory:

```bash
cd src/examples/chapter-15/server/firestore-data/remote
```

Initialize and apply the Terraform configuration:

```bash
# Initialize Terraform
terraform init

# Plan the deployment (replace with your project ID)
terraform plan -var="project_id=your-project-id"

# Apply the configuration
terraform apply -var="project_id=your-project-id"
```

This will create the `ingredient_combinations` collection with all the master data for ingredient compatibility.

### Updating Your Go Server

To connect to the production Firestore instead of the emulator, update the Firestore client initialization in `main.go`:

```go
// Change from:
firestoreClient, err := firestore.NewClient(ctx, "local-emulator")

// To:
firestoreClient, err := firestore.NewClient(ctx, "your-project-id")
```

Also, ensure the `FIRESTORE_EMULATOR_HOST` environment variable is NOT set when running in production mode.

## Resources

- [Genkit Documentation](https://genkit.dev)
- [Firebase Authentication](https://firebase.google.com/docs/auth)
- [Cloud Run Documentation](https://cloud.google.com/run/docs)
- [Terraform Google Provider](https://registry.terraform.io/providers/hashicorp/google/latest/docs)
- [OpenTelemetry Go](https://opentelemetry.io/docs/languages/go/)
