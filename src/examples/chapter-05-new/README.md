# Chapter 5: Structured Data - Example Code

This directory contains example code for Chapter 5, demonstrating structured input/output patterns with Genkit Go. The examples focus on `GenerateData` for type-safe structured data handling with multimodal support.

## Overview

This chapter moves beyond simple text generation to demonstrate how Genkit Go handles complex structured data across different modalities (text, image, audio, and image generation). All examples use `GenerateData` for deterministic, type-safe data extraction and generation.

## Project Structure

```text
chapter-05-new/
├── main.go                  # Application entry point
├── internal/
│   ├── flows/              # Flow definitions
│   │   ├── 1-simple.go     # Flow 1: Simple structured output
│   │   ├── 2-nested.go     # Flow 2: Nested structures
│   │   ├── 3-image.go      # Flow 3: Image analysis
│   │   ├── 4-audio.go      # Flow 4: Audio transcription
│   │   └── 5-generation.go # Flow 5: Image generation
│   └── schemas/            # Data structures
│       ├── 1-simple/       # Review analysis schemas
│       │   ├── input.go    # ReviewInput
│       │   └── output.go   # ReviewAnalysis
│       ├── 2-nested/       # Recipe generation schemas
│       │   ├── input.go    # RecipeRequest
│       │   └── output.go   # Recipe, Ingredient, CookingStep
│       ├── 3-image/        # Product analysis schemas
│       │   ├── input.go    # ImageAnalysisRequest
│       │   └── output.go   # ProductInfo
│       ├── 4-audio/        # Meeting minutes schemas
│       │   ├── input.go    # AudioRequest
│       │   └── output.go   # MeetingMinutes, SpeakerSummary, ActionItem
│       └── 5-generation/   # Character generation schemas
│           ├── input.go    # CharacterRequest
│           └── output.go   # CharacterResult, Character
├── go.mod
└── go.sum
```

## Five Independent Flows

### Flow 1: simpleStructuredFlow

Purpose: Basic structured output pattern demonstration

Input Structure:

```go
type ReviewInput struct {
    ReviewText string `json:"review_text"`
}
```

Output Structure:

```go
type ReviewAnalysis struct {
    Sentiment string   `json:"sentiment"` // "positive", "negative", "neutral"
    Score     int      `json:"score"`     // 1-5
    Keywords  []string `json:"keywords"`  // Key topics/features mentioned
    Summary   string   `json:"summary"`   // Brief summary
}
```

Example Usage:

```json
// Input
{
  "review_text": "This product exceeded my expectations! The quality is outstanding and shipping was fast. Highly recommend!"
}

// Expected Output
{
  "sentiment": "positive",
  "score": 5,
  "keywords": ["quality", "shipping", "recommend"],
  "summary": "Very satisfied customer praising product quality and fast shipping"
}
```

### Flow 2: nestedStructuredFlow

Purpose: Complex nested structures, arrays, and maps

Input Structure:

```go
type RecipeRequest struct {
    DishName            string   `json:"dish_name"`
    Servings           int      `json:"servings"`
    DietaryRestrictions []string `json:"dietary_restrictions"`
}
```

Output Structure:

```go
type Recipe struct {
    Name        string        `json:"name"`
    Ingredients []Ingredient  `json:"ingredients"`
    Steps       []CookingStep `json:"steps"`
    Nutrition   Nutrition     `json:"nutrition"`
    PrepTime    int           `json:"prep_time"`  // minutes
    CookTime    int           `json:"cook_time"`  // minutes
    Difficulty  string        `json:"difficulty"` // easy, medium, hard
}

type Nutrition struct {
    Calories float64 `json:"calories"`
    Protein  float64 `json:"protein"`  // grams
    Carbs    float64 `json:"carbs"`    // grams
    Fat      float64 `json:"fat"`      // grams
    Fiber    float64 `json:"fiber"`    // grams
}

type Ingredient struct {
    Name     string  `json:"name"`
    Amount   float64 `json:"amount"`
    Unit     string  `json:"unit"`
    Optional bool    `json:"optional"`
}

type CookingStep struct {
    Number      int    `json:"number"`
    Instruction string `json:"instruction"`
    Duration    int    `json:"duration"` // minutes
    Tips        string `json:"tips,omitempty"`
}
```

Example Usage:

```json
// Input
{
  "dish_name": "Vegetable Stir Fry",
  "servings": 4,
  "dietary_restrictions": ["vegetarian", "gluten-free"]
}

// Expected Output (abbreviated)
{
  "name": "Quick Vegetable Stir Fry",
  "ingredients": [
    {"name": "broccoli", "amount": 2, "unit": "cups", "optional": false},
    {"name": "soy sauce", "amount": 2, "unit": "tbsp", "optional": false}
  ],
  "steps": [...],
  "nutrition": {"calories": 250, "protein": 8, "carbs": 35, "fat": 10, "fiber": 6},
  "prep_time": 10,
  "cook_time": 15,
  "difficulty": "easy"
}
```

### Flow 3: imageAnalysisFlow

Purpose: Image input with structured output

Input Structure:

```go
type ImageAnalysisRequest struct {
    ImageURL string `json:"image_url"` // URL of the image to analyze
    Language string `json:"language"` // Output language preference
}
```

Output Structure:

```go
type ProductInfo struct {
    Category       string   `json:"category"`
    Features       []string `json:"features"`
    Colors         []string `json:"colors"`
    EstimatedPrice string   `json:"estimated_price"` // price range
    TargetAudience string   `json:"target_audience"`
    Description    string   `json:"description"`
}
```

Example Usage:

```json
// Input
{
  "image_url": "https://example.com/product-image.jpg",
  "language": "en"
}

// Expected Output
{
  "category": "Electronics - Headphones",
  "features": ["wireless", "noise-canceling", "over-ear"],
  "colors": ["black", "silver"],
  "estimated_price": "$150-$250",
  "target_audience": "Music enthusiasts and remote workers",
  "description": "Premium wireless headphones with active noise cancellation"
}
```

### Flow 4: audioTranscriptionFlow

Purpose: Audio input with structured output

Input Structure:

```go
type AudioRequest struct {
    AudioURL string `json:"audio_url"` // URL of the audio file
    Topic    string `json:"topic"`     // Meeting topic/context
    Language string `json:"language"`  // Audio language
}
```

Output Structure:

```go
type MeetingMinutes struct {
    Transcription     string            `json:"transcription"`
    SpeakerSummaries []SpeakerSummary  `json:"speaker_summaries"`
    Decisions        []string          `json:"decisions"`
    ActionItems      []ActionItem      `json:"action_items"`
    Keywords         []string          `json:"keywords"`
}

type SpeakerSummary struct {
    Speaker string   `json:"speaker"`
    Points  []string `json:"points"`
}

type ActionItem struct {
    Task       string `json:"task"`
    Assignee   string `json:"assignee"`
    DueDate    string `json:"due_date"`
    Priority   string `json:"priority"` // high, medium, low
}
```

Example Usage:

```json
// Input
{
  "audio_url": "https://example.com/meeting-recording.mp3",
  "topic": "Q4 Product Planning",
  "language": "en"
}

// Expected Output
{
  "transcription": "Full meeting transcript...",
  "speaker_summaries": [
    {
      "speaker": "John",
      "points": ["Proposed new feature X", "Suggested timeline of 3 weeks"]
    }
  ],
  "decisions": ["Launch feature X in December", "Allocate 2 engineers"],
  "action_items": [
    {
      "task": "Create feature specification",
      "assignee": "Sarah",
      "due_date": "2024-11-15",
      "priority": "high"
    }
  ],
  "keywords": ["feature X", "December launch", "engineering resources"]
}
```

### Flow 5: characterGenerationFlow

Purpose: Image generation with structured metadata

Input Structure:

```go
type CharacterRequest struct {
    Description string `json:"description"` // "blue-haired wizard"
    Style       string `json:"style"`       // "anime", "realistic", "cartoon"
}
```

Output Structure:

```go
type CharacterResult struct {
    ImageData string    `json:"image_data"` // base64 encoded image (data URI format)
    Character Character `json:"character"`  // Character details
}

type Character struct {
    Name        string   `json:"name"`        // AI-generated name
    Description string   `json:"description"` // Visual description
    Traits      []string `json:"traits"`      // Character traits
    Story       string   `json:"story"`       // Brief backstory
}
```

Example Usage:

```json
// Input
{
  "description": "blue-haired wizard with starry robe",
  "style": "anime"
}

// Expected Output
{
  "image_data": "data:image/png;base64,iVBORw0KGgoAAAANS...",
  "character": {
    "name": "Luna Stellaris",
    "description": "A young wizard with flowing blue hair and a mystical starry robe",
    "traits": ["mysterious", "wise", "gentle"],
    "story": "A young wizard from the Celestial Academy who masters star magic..."
  }
}
```
