# Recipe Quest - Next Client

A Next.js 15 client application for the Recipe Quest cooking battle game, communicating with a Genkit Go server. This app implements an Onion Architecture pattern for clean, maintainable code.

## 🎯 Features

- **Interactive Cooking Game**: Select ingredients, generate recipes, create dish images, and get AI evaluations
- **Real-time Streaming**: Watch recipes being generated in real-time with streaming responses
- **Clean Architecture**: Built with Onion Architecture principles for maintainability
- **Tech Stack**: Next.js 15, React 19, TypeScript, and Tailwind CSS
- **AI-Powered**: Integration with Genkit Go flows for recipe generation, image creation, and evaluation

## 🏗️ Architecture

### Onion Architecture

This project follows Onion Architecture principles with clear separation of concerns:

```text
src/
├── domain/                    # Domain Layer (Core)
│   ├── models/               # Domain Models
│   │   ├── game/            # Game state models
│   │   ├── recipe/          # Recipe-related models
│   │   ├── image/           # Image generation models
│   │   ├── evaluate/        # Evaluation models
│   │   └── error/           # Error handling models
│   └── repositories.ts       # Repository Interfaces
│
├── usecases/                 # Use Case Layer
│   ├── generate-recipe.ts    # Recipe generation use case
│   ├── create-image.ts       # Image creation use case
│   └── evaluate-dish.ts      # Dish evaluation use case
│
├── infrastructure/           # Infrastructure Layer
│   ├── http/                # HTTP Communication
│   │   ├── dto/            # Data Transfer Objects
│   │   ├── mappers/        # DTO ↔ Domain mapping
│   │   ├── client/         # HTTP client utilities
│   │   ├── config/         # HTTP configuration
│   │   └── repository/     # Repository implementations
│   └── auth/               # Authentication
│       └── firebase.ts     # Firebase setup (if needed)
│
├── components/              # UI Components
│   ├── GameProgress.tsx    # Game progress indicator
│   ├── GameResult.tsx      # Game results display
│   ├── IngredientCards.tsx # Ingredient selection
│   ├── RecipeDisplay.tsx   # Recipe presentation
│   └── ImageDisplay.tsx    # Generated image display
│
└── app/                    # Next.js App Router
    ├── page.tsx           # Home page
    ├── quest/             # Game quest page
    ├── hooks/             # Custom React hooks
    └── composition.ts     # Dependency injection setup
```

### Layer Responsibilities

1. **Domain Layer (Core)**
   - Business logic and domain models
   - No external dependencies - pure TypeScript
   - Repository interface definitions

2. **Use Case Layer**
   - Application business logic
   - Depends only on domain layer
   - Orchestrates domain models and repositories

3. **Infrastructure Layer**
   - External system implementations
   - HTTP clients, authentication, DTO mapping
   - Repository interface implementations

4. **Presentation Layer**
   - UI components and pages
   - React/Next.js specific implementations
   - Connects UI to use cases

## 🚀 Quick Start

### Prerequisites

- **Node.js** 18.17 or later
- **npm** or **yarn**
- **Genkit Go Server** running on `http://127.0.0.1:9090`

### Installation

```bash
# Install dependencies
npm install
```

### Development Server

```bash
# Start the development server
npm run build
npm run dev

# Open http://localhost:3000 in your browser
```

## 🔧 Development Guide

### Adding New Features

Follow these steps to add new functionality while maintaining clean architecture:

1. **Define Domain Models**

   ```typescript
   // src/domain/models/[feature]/[model].ts
   export interface NewFeature {
     id: string;
     name: string;
     // ... define domain properties
   }
   ```

2. **Create Repository Interface**

   ```typescript
   // src/domain/repositories.ts
   export interface NewFeatureRepository {
     getFeature(id: string): Promise<NewFeature>;
     createFeature(data: CreateFeatureRequest): Promise<NewFeature>;
   }
   ```

3. **Implement Use Case**

   ```typescript
   // src/usecases/new-feature.ts
   export class GetNewFeatureUseCase {
     constructor(private repository: NewFeatureRepository) {}
     
     async execute(id: string): Promise<NewFeature> {
       return await this.repository.getFeature(id);
     }
   }
   ```

4. **Implement Infrastructure**

   ```typescript
   // src/infrastructure/http/repository/new-feature-repo.ts
   export class HttpNewFeatureRepository implements NewFeatureRepository {
     constructor(private httpClient: HttpClient) {}
     
     async getFeature(id: string): Promise<NewFeature> {
       // HTTP implementation
       const response = await this.httpClient.get(`/feature/${id}`);
       return mapDTOToDomain(response.data);
     }
   }
   ```

5. **Wire Dependencies**

   ```typescript
   // app/composition.ts
   const newFeatureRepo = new HttpNewFeatureRepository(httpClient);
   export const getNewFeature = new GetNewFeatureUseCase(newFeatureRepo);
   ```

### Genkit Flow Communication

This app communicates with Genkit Go flows using specific patterns:

#### Request Format

```typescript
// All requests are wrapped
{ data: <your-payload> }
```

#### Response Formats

```typescript
// Regular responses
{ result: <response-data> }

// Streaming responses (SSE)
{ message: <streaming-chunk> }  // For intermediate chunks
{ result: <final-result> }       // For completion
```

#### Example Usage

```typescript
// Streaming recipe generation
const response = await fetch('/generateRecipe', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'text/event-stream',
  },
  body: JSON.stringify({
    data: { ingredients: ['tomato', 'basil', 'mozzarella'] }
  })
});
```
