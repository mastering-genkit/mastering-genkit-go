# Recipe Quest - Flutter Client

A Flutter client application for the Recipe Quest cooking battle game, communicating with a Genkit Go server using Clean Architecture principles for iOS and Android platforms.

## 🎯 Features

- **Interactive Cooking Game**: Select ingredients, generate recipes, create dish images, and get AI evaluations
- **Real-time Streaming**: Watch recipes being generated in real-time using Dart client for Genkit
- **Clean Architecture**: Built with Clean Architecture principles for maintainability
- **Tech Stack**: Flutter 3.8+, Material 3, Riverpod, and JSON serialization
- **AI-Powered**: Integration with Genkit Go flows for recipe generation, image creation, and evaluation
- **Cross-Platform**: Native iOS and Android apps with consistent UX

## 📱 Supported Platforms

- ✅ **iOS** (iPhone/iPad Simulator)
- ✅ **Android** (Phone/Tablet Emulator)

## 🏗️ Architecture

### Clean Architecture (4 Layers)

```text
lib/
├── domain/              # Domain Layer (Core Business Logic)
├── usecases/           # Use Case Layer (Application Logic)
├── infrastructure/     # Infrastructure Layer (External Systems)
└── presentation/      # Presentation Layer (UI)
```

### Key Components

- **Dart client for Genkit**: Type-safe remote action definitions
- **Riverpod**: State management with dependency injection
- **JSON Serialization**: Auto-generated with `build_runner`
- **Material 3**: Google's latest design system

## 🚀 Quick Start

### Prerequisites

- **Flutter SDK** 3.8.1+
- **Genkit Go Server** running on `http://127.0.0.1:9090`

### Installation & Run

```bash
# Install dependencies
flutter pub get

# Generate schema code
flutter packages pub run build_runner build

# Launch app (auto-select device)
flutter run

# Or specific devices:
flutter run -d "iPhone 16 Pro Max"      # iOS
flutter run -d "sdk gphone64 arm64"     # Android
```

## 🎮 Game Flow

1. **Ready** → **Select Ingredients** → **Recipe** → **Image** → **Evaluation** → **Result**
2. Choose 4 ingredients from 33 options
3. AI streams recipe generation
4. AI creates dish visualization
5. AI evaluates and scores your creation

## 🔌 Genkit Integration

### Remote Actions using Dart client for Genkit

```dart
// Streaming recipe generation
final generateRecipe = defineRemoteAction<RecipeResponse, RecipeResponse>(
  url: 'http://127.0.0.1:9090/generateRecipe',
  fromResponse: (json) => RecipeResponse.fromJson(json),
  fromStreamChunk: (json) => RecipeResponse.fromJson(json),
);

// Execute streaming
final (:stream, :response) = generateRecipe.stream(
  input: RecipeRequest(ingredients: ['tomato', 'pasta', 'cheese', 'herbs'])
);

await for (final chunk in stream) {
  print('Recipe: ${chunk.content}');
}
```

## 🛠️ Development

### Hot Reload Commands

- **`r`** - Hot reload 🔥
- **`R`** - Hot restart
- **`q`** - Quit
- **`h`** - Help
