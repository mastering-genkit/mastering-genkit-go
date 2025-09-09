# Recipe Quest - Angular Client

An Angular 20 client application for the Recipe Quest cooking battle game, communicating with a Genkit Go server. This app implements an Onion Architecture pattern for clean, maintainable code.

## 🎯 Features

- **Interactive Cooking Game**: Select ingredients, generate recipes, create dish images, and get AI evaluations
- **Real-time Streaming**: Watch recipes being generated in real-time with streaming responses
- **Clean Architecture**: Built with Onion Architecture principles for maintainability
- **Modern Tech Stack**: Angular 20.2.2, TypeScript 5.8, RxJS 7.8, and standalone components
- **AI-Powered**: Integration with Genkit Go flows for recipe generation, image creation, and evaluation
- **Enhanced Performance**: Optimized build process and runtime performance with Angular 20
- **Zero Vulnerabilities**: Security-focused dependency management

## 🆕 Angular 20 Highlights

This project leverages the latest Angular 20.2.2 features:

- **Improved Bundle Optimization**: Faster build times and smaller bundle sizes
- **Enhanced Type Safety**: Better TypeScript integration with stricter type checking
- **Latest Dependencies**: Up-to-date ecosystem with security patches
- **Modern Development Tools**: Advanced CLI and build tooling
- **Better Performance**: Optimized change detection and rendering

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
│   └── auth/               # Authentication (future)
│
└── app/                    # Angular App
    ├── quest/              # Quest component
    ├── services/           # Angular services
    ├── composition/        # Dependency injection setup
    └── app.component.ts    # Root component
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
   - HTTP clients, DTO mapping
   - Repository interface implementations

4. **Presentation Layer**
   - Angular components and services
   - RxJS observables for reactive programming
   - Connects UI to use cases through services

## 🚀 Quick Start

### Prerequisites

- **Node.js** 18.19 or later
- **npm** or **yarn**
- **Angular CLI** 20.x
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

# Open http://localhost:4200 in your browser
```

## 🔧 Development Guide

### Angular Architecture Patterns

This app uses Angular-specific patterns while maintaining clean architecture:

#### Services for State Management

```typescript
// Game state management service
@Injectable({ providedIn: 'root' })
export class GameStateService {
  private stateSubject = new BehaviorSubject<GameState>(initialState);
  public state$ = this.stateSubject.asObservable();
  
  // Actions that modify state
  startGame(): void { /* ... */ }
  addIngredient(ingredient: string): void { /* ... */ }
}
```

#### Reactive Programming with RxJS

```typescript
// Service for streaming recipe generation
@Injectable({ providedIn: 'root' })
export class GenerateRecipeService {
  private recipeStreamSubject = new Subject<RecipeResponseDomain>();
  public recipeStream$ = this.recipeStreamSubject.asObservable();
  
  async generateRecipe(request: RecipeRequestDomain): Promise<void> {
    // Async generator for streaming responses
    for await (const response of useCase.execute(request)) {
      this.recipeStreamSubject.next(response);
    }
  }
}
```

#### Component Integration

```typescript
@Component({
  selector: 'app-quest',
  standalone: true,
  templateUrl: './quest.component.html'
})
export class QuestComponent implements OnInit, OnDestroy {
  state$ = this.gameStateService.state$;
  
  constructor(
    private gameStateService: GameStateService,
    private generateRecipeService: GenerateRecipeService
  ) {}
  
  ngOnInit() {
    this.subscriptions.add(
      this.generateRecipeService.recipeStream$.subscribe(response => {
        // Handle streaming responses
      })
    );
  }
}
```

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
     async getFeature(id: string): Promise<NewFeature> {
       const response = await httpFetch(`/feature/${id}`);
       return mapDTOToDomain(response.data);
     }
   }
   ```

5. **Create Angular Service**

   ```typescript
   // src/app/services/new-feature.service.ts
   @Injectable({ providedIn: 'root' })
   export class NewFeatureService {
     constructor(private compositionService: CompositionService) {}
     
     async getFeature(id: string): Promise<NewFeature> {
       const useCase = this.compositionService.getNewFeatureUseCase();
       return await useCase.execute(id);
     }
   }
   ```

6. **Wire Dependencies**

   ```typescript
   // src/app/composition/composition.ts
   getNewFeatureUseCase(): GetNewFeatureUseCase {
     if (!this.newFeatureUseCase) {
       this.newFeatureUseCase = new GetNewFeatureUseCase(
         new HttpNewFeatureRepository()
       );
     }
     return this.newFeatureUseCase;
   }
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
const response = await httpFetch('/generateRecipe', {
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

## 🎮 Game Flow

The Recipe Quest game follows these steps:

1. **Start Game**: Begin the cooking challenge
2. **Select Ingredients**: Choose exactly 4 ingredients from available options
3. **Generate Recipe**: AI creates a custom recipe using selected ingredients (streaming)
4. **Create Image**: AI generates an image visualization of the dish
5. **Evaluate Dish**: AI chef evaluates and scores the creation
6. **View Results**: See score, feedback, and achievements
