import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../domain/entities/game_state.dart';
import '../../domain/models/recipe/request.dart';
import '../../domain/models/recipe/response.dart';
import '../../domain/models/image/request.dart';
import '../../domain/models/evaluate/request.dart';
import '../../usecases/providers.dart';

/// Available ingredients for Recipe Quest with emojis
class Ingredient {
  const Ingredient({required this.name, required this.emoji});

  final String name;
  final String emoji;
}

/// List of available ingredients for user selection (33 items matching Angular)
const availableIngredients = [
  Ingredient(name: 'chicken', emoji: '🐔'),
  Ingredient(name: 'beef', emoji: '🥩'),
  Ingredient(name: 'pork', emoji: '🐷'),
  Ingredient(name: 'salmon', emoji: '🐟'),
  Ingredient(name: 'shrimp', emoji: '🦐'),
  Ingredient(name: 'tofu', emoji: '🥡'),
  Ingredient(name: 'rice', emoji: '🍚'),
  Ingredient(name: 'noodles', emoji: '🍜'),
  Ingredient(name: 'pasta', emoji: '🍝'),
  Ingredient(name: 'potatoes', emoji: '🥔'),
  Ingredient(name: 'onions', emoji: '🧅'),
  Ingredient(name: 'garlic', emoji: '🧄'),
  Ingredient(name: 'ginger', emoji: '🫚'),
  Ingredient(name: 'carrots', emoji: '🥕'),
  Ingredient(name: 'peppers', emoji: '🌶️'),
  Ingredient(name: 'vegetables', emoji: '🥬'),
  Ingredient(name: 'mushrooms', emoji: '🍄'),
  Ingredient(name: 'tomatoes', emoji: '🍅'),
  Ingredient(name: 'lemon', emoji: '🍋'),
  Ingredient(name: 'herbs', emoji: '🌿'),
  Ingredient(name: 'sesame oil', emoji: '🫗'),
  Ingredient(name: 'soy sauce', emoji: '🥢'),
  Ingredient(name: 'miso', emoji: '🍲'),
  Ingredient(name: 'quinoa', emoji: '🌾'),
  Ingredient(name: 'avocado', emoji: '🥑'),
  Ingredient(name: 'lime', emoji: '🍈'),
];

/// Game State Notifier for Recipe Quest game flow
class GameStateNotifier extends StateNotifier<GameState> {
  GameStateNotifier(this._ref) : super(GameState.initial());

  final Ref _ref;

  /// Start the Recipe Quest game
  void startGame() {
    state = state
        .copyWith(
          currentStep: GameStep.selectIngredients,
          progress: 10.0,
          isLoading: false,
        )
        .copyWithNull(error: true);
  }

  /// Add an ingredient to selection (max 4)
  void addIngredient(String ingredient) {
    if (state.selectedIngredients.length >= 4) return;
    if (state.selectedIngredients.contains(ingredient)) return;

    final newIngredients = [...state.selectedIngredients, ingredient];
    state = state.copyWith(
      selectedIngredients: newIngredients,
      progress: (10 + (newIngredients.length * 2.5)).clamp(0.0, 20.0),
    );
  }

  /// Remove an ingredient from selection
  void removeIngredient(String ingredient) {
    final newIngredients = state.selectedIngredients
        .where((ing) => ing != ingredient)
        .toList();

    state = state.copyWith(
      selectedIngredients: newIngredients,
      progress: (10 + (newIngredients.length * 2.5)).clamp(0.0, 20.0),
    );
  }

  /// Toggle ingredient selection (add if not selected, remove if selected)
  void toggleIngredient(String ingredient) {
    if (state.selectedIngredients.contains(ingredient)) {
      removeIngredient(ingredient);
    } else {
      addIngredient(ingredient);
    }
  }

  /// Start recipe generation using selected ingredients
  Future<void> startRecipeGeneration() async {
    if (state.selectedIngredients.length != 4) return;

    // Update state to recipe generation
    state = state
        .copyWith(currentStep: GameStep.recipe, progress: 30.0, isLoading: true)
        .copyWithNull(error: true);

    try {
      // Get use case and execute
      final useCase = _ref.read(generateRecipeUseCaseProvider);
      final request = RecipeRequest(ingredients: state.selectedIngredients);

      String accumulatedRecipe = '';
      await for (final response in useCase.execute(request)) {
        if (response.type == RecipeResponseType.content &&
            response.content != null) {
          // Accumulate recipe content
          accumulatedRecipe += response.content!;
          state = state.copyWith(
            recipe: accumulatedRecipe,
            progress: (40 + (accumulatedRecipe.length * 0.01)).clamp(
              30.0,
              45.0,
            ),
          );
        } else if (response.type == RecipeResponseType.done) {
          // Recipe generation completed - move to image generation
          await _startImageGeneration();
          break;
        } else if (response.type == RecipeResponseType.error) {
          // Handle error
          _setError(response.error ?? 'Recipe generation failed');
          break;
        }
      }
    } catch (e) {
      _setError('Recipe generation failed: $e');
    }
  }

  /// Start image generation for the recipe
  Future<void> _startImageGeneration() async {
    state = state.copyWith(
      currentStep: GameStep.image,
      progress: 50.0,
      isLoading: true,
      isGeneratingImage: true,
      imageGenerationProgress: 0.0,
    );

    // Simulate image generation progress (since Genkit doesn't provide real progress)
    _simulateImageProgress();

    try {
      // Extract dish name from recipe for image generation
      final dishName = _extractDishName(state.recipe ?? 'Recipe');
      final useCase = _ref.read(createImageUseCaseProvider);
      final request = ImageRequest(
        dishName: dishName,
        description: _createImageDescription(state.recipe ?? ''),
      );

      final response = await useCase.execute(request);

      if (response.success && response.imageUrl != null) {
        state = state.copyWith(
          imageUrl: response.imageUrl,
          progress: 60.0,
          isLoading: false,
          isGeneratingImage: false,
          imageGenerationProgress: 100.0,
        );
        // Move to evaluation
        await _startEvaluation();
      } else {
        _setError(response.error ?? 'Image generation failed');
      }
    } catch (e) {
      _setError('Image generation failed: $e');
    }
  }

  /// Simulate image generation progress
  void _simulateImageProgress() {
    // Simulate progress updates every 500ms
    const duration = Duration(milliseconds: 500);
    double progress = 0.0;

    void updateImageProgress(double prog) {
      if (mounted &&
          state.currentStep == GameStep.image &&
          state.isGeneratingImage == true) {
        progress = (prog + 15.0).clamp(0.0, 95.0);
        state = state.copyWith(
          imageGenerationProgress: progress,
          progress:
              50.0 +
              (progress * 0.1), // Image contributes 10% to total progress
        );

        if (progress < 95.0) {
          Future.delayed(duration, () => updateImageProgress(progress));
        }
      }
    }

    Future.delayed(duration, () {
      updateImageProgress(progress);
    });
  }

  /// Start dish evaluation
  Future<void> _startEvaluation() async {
    state = state.copyWith(
      currentStep: GameStep.evaluation,
      progress: 80.0,
      isLoading: true,
    );

    try {
      final dishName = _extractDishName(state.recipe ?? 'Recipe');
      final useCase = _ref.read(evaluateDishUseCaseProvider);
      final request = EvaluateRequest(
        dishName: dishName,
        description: state.recipe ?? 'A delicious recipe',
      );

      final response = await useCase.execute(request);

      if (response.success) {
        state = state.copyWith(
          currentStep: GameStep.result,
          progress: 100.0,
          score: response.score,
          feedback: response.feedback,
          title: response.title,
          achievement: response.achievement,
          isLoading: false,
        );
      } else {
        _setError(response.error ?? 'Evaluation failed');
      }
    } catch (e) {
      _setError('Evaluation failed: $e');
    }
  }

  /// Extract dish name from recipe text (simple heuristic)
  String _extractDishName(String recipe) {
    final lines = recipe.split('\n');
    for (final line in lines) {
      final trimmed = line.trim();
      if (trimmed.isNotEmpty &&
          !trimmed.startsWith('Ingredients:') &&
          !trimmed.startsWith('-')) {
        // Return first non-empty line as dish name
        return trimmed.replaceAll(RegExp(r'[^\w\s]'), '').trim();
      }
    }
    return 'Delicious Recipe';
  }

  /// Create concise description for image generation from recipe
  String _createImageDescription(String recipe) {
    if (recipe.isEmpty) return 'A delicious dish with selected ingredients';

    // Extract key information from recipe for image generation
    final ingredients = state.selectedIngredients.join(', ');
    final dishName = _extractDishName(recipe);

    // Create concise description under 200 characters
    String description = '$dishName made with $ingredients';

    // Add cooking method if found in recipe (keep it brief)
    final lowerRecipe = recipe.toLowerCase();
    if (lowerRecipe.contains('fried') || lowerRecipe.contains('fry')) {
      description += ', fried to perfection';
    } else if (lowerRecipe.contains('grilled') ||
        lowerRecipe.contains('grill')) {
      description += ', grilled beautifully';
    } else if (lowerRecipe.contains('baked') || lowerRecipe.contains('bake')) {
      description += ', baked golden';
    } else if (lowerRecipe.contains('steamed') ||
        lowerRecipe.contains('steam')) {
      description += ', steamed delicately';
    } else {
      description += ', cooked with care';
    }

    // Ensure it stays under 300 characters for image generation
    if (description.length > 300) {
      description = '${description.substring(0, 297)}...';
    }

    return description;
  }

  /// Set loading state
  void setLoading(bool isLoading) {
    state = state.copyWith(isLoading: isLoading);
  }

  /// Set error state
  void _setError(String error) {
    state = state.copyWith(
      error: error,
      isLoading: false,
      isGeneratingImage: false,
    );
  }

  /// Reset game to initial state
  void resetGame() {
    state = GameState.initial();
  }
}

/// Provider for GameStateNotifier
final gameStateProvider = StateNotifierProvider<GameStateNotifier, GameState>((
  ref,
) {
  return GameStateNotifier(ref);
});
