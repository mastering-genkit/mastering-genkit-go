import 'package:json_annotation/json_annotation.dart';

part 'game_state.g.dart';

/// Quest step enumeration for Recipe Quest flow
@JsonEnum()
enum GameStep {
  @JsonValue('ready')
  ready,
  @JsonValue('select_ingredients')
  selectIngredients,
  @JsonValue('recipe')
  recipe,
  @JsonValue('image')
  image,
  @JsonValue('evaluation')
  evaluation,
  @JsonValue('result')
  result,
}

/// Quest state interface for Recipe Quest game
@JsonSerializable()
class GameState {
  const GameState({
    required this.currentStep,
    required this.progress,
    required this.selectedIngredients,
    this.recipe,
    this.imageUrl,
    this.score,
    this.feedback,
    this.title,
    this.achievement,
    required this.isLoading,
    this.error,
    this.isGeneratingImage,
    this.imageGenerationProgress,
  });

  /// Current step in the quest flow
  final GameStep currentStep;

  /// Overall progress of the quest (0-100)
  final double progress;

  /// User-selected ingredients
  final List<String> selectedIngredients;

  /// Generated recipe content
  final String? recipe;

  /// URL of the generated dish image
  final String? imageUrl;

  /// Evaluation score for the dish
  final double? score;

  /// Evaluation feedback message
  final String? feedback;

  /// Title given to the dish evaluation
  final String? title;

  /// Achievement badge earned
  final String? achievement;

  /// Whether any operation is currently loading
  final bool isLoading;

  /// Current error message, if any
  final String? error;

  /// Whether image generation is in progress
  final bool? isGeneratingImage;

  /// Image generation progress (0-100)
  final double? imageGenerationProgress;

  /// Factory constructor for creating a new `GameState` instance from a map.
  factory GameState.fromJson(Map<String, dynamic> json) =>
      _$GameStateFromJson(json);

  /// Converts this `GameState` instance to a map.
  Map<String, dynamic> toJson() => _$GameStateToJson(this);

  /// Creates the initial game state
  factory GameState.initial() => const GameState(
    currentStep: GameStep.ready,
    progress: 0.0,
    selectedIngredients: [],
    isLoading: false,
  );

  /// Creates a copy of this state with some fields replaced by the non-null parameter values.
  GameState copyWith({
    GameStep? currentStep,
    double? progress,
    List<String>? selectedIngredients,
    String? recipe,
    String? imageUrl,
    double? score,
    String? feedback,
    String? title,
    String? achievement,
    bool? isLoading,
    String? error,
    bool? isGeneratingImage,
    double? imageGenerationProgress,
  }) {
    return GameState(
      currentStep: currentStep ?? this.currentStep,
      progress: progress ?? this.progress,
      selectedIngredients: selectedIngredients ?? this.selectedIngredients,
      recipe: recipe ?? this.recipe,
      imageUrl: imageUrl ?? this.imageUrl,
      score: score ?? this.score,
      feedback: feedback ?? this.feedback,
      title: title ?? this.title,
      achievement: achievement ?? this.achievement,
      isLoading: isLoading ?? this.isLoading,
      error: error ?? this.error,
      isGeneratingImage: isGeneratingImage ?? this.isGeneratingImage,
      imageGenerationProgress:
          imageGenerationProgress ?? this.imageGenerationProgress,
    );
  }

  /// Creates a copy with null values for specified fields
  GameState copyWithNull({
    bool recipe = false,
    bool imageUrl = false,
    bool score = false,
    bool feedback = false,
    bool title = false,
    bool achievement = false,
    bool error = false,
    bool isGeneratingImage = false,
    bool imageGenerationProgress = false,
  }) {
    return GameState(
      currentStep: currentStep,
      progress: progress,
      selectedIngredients: selectedIngredients,
      recipe: recipe ? null : this.recipe,
      imageUrl: imageUrl ? null : this.imageUrl,
      score: score ? null : this.score,
      feedback: feedback ? null : this.feedback,
      title: title ? null : this.title,
      achievement: achievement ? null : this.achievement,
      isLoading: isLoading,
      error: error ? null : this.error,
      isGeneratingImage: isGeneratingImage ? null : this.isGeneratingImage,
      imageGenerationProgress: imageGenerationProgress
          ? null
          : this.imageGenerationProgress,
    );
  }

  @override
  String toString() =>
      'GameState(currentStep: $currentStep, progress: $progress, '
      'selectedIngredients: $selectedIngredients, recipe: ${recipe?.length} chars, '
      'imageUrl: $imageUrl, score: $score, feedback: ${feedback?.length} chars, '
      'title: $title, achievement: $achievement, isLoading: $isLoading, '
      'error: $error, isGeneratingImage: $isGeneratingImage, '
      'imageGenerationProgress: $imageGenerationProgress)';

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other.runtimeType != runtimeType) return false;
    if (other is! GameState) return false;
    return currentStep == other.currentStep &&
        progress == other.progress &&
        _listEquals(selectedIngredients, other.selectedIngredients) &&
        recipe == other.recipe &&
        imageUrl == other.imageUrl &&
        score == other.score &&
        feedback == other.feedback &&
        title == other.title &&
        achievement == other.achievement &&
        isLoading == other.isLoading &&
        error == other.error &&
        isGeneratingImage == other.isGeneratingImage &&
        imageGenerationProgress == other.imageGenerationProgress;
  }

  @override
  int get hashCode => Object.hash(
    currentStep,
    progress,
    Object.hashAll(selectedIngredients),
    recipe,
    imageUrl,
    score,
    feedback,
    title,
    achievement,
    isLoading,
    error,
    isGeneratingImage,
    imageGenerationProgress,
  );
}

/// Quest actions for state management
sealed class GameAction {
  const GameAction();
}

class StartGame extends GameAction {
  const StartGame();
}

class AddIngredient extends GameAction {
  const AddIngredient(this.ingredient);
  final String ingredient;
}

class RemoveIngredient extends GameAction {
  const RemoveIngredient(this.ingredient);
  final String ingredient;
}

class StartRecipeGeneration extends GameAction {
  const StartRecipeGeneration();
}

class SetRecipe extends GameAction {
  const SetRecipe(this.recipe);
  final String recipe;
}

class StartImageGeneration extends GameAction {
  const StartImageGeneration();
}

class SetImageProgress extends GameAction {
  const SetImageProgress(this.progress);
  final double progress;
}

class SetImage extends GameAction {
  const SetImage(this.imageUrl);
  final String imageUrl;
}

class StartEvaluation extends GameAction {
  const StartEvaluation();
}

class SetEvaluation extends GameAction {
  const SetEvaluation({
    required this.score,
    required this.feedback,
    this.title,
    this.achievement,
  });
  final double score;
  final String feedback;
  final String? title;
  final String? achievement;
}

class SetLoading extends GameAction {
  const SetLoading(this.isLoading);
  final bool isLoading;
}

class SetError extends GameAction {
  const SetError(this.error);
  final String error;
}

class ResetGame extends GameAction {
  const ResetGame();
}

// Helper function for list equality
bool _listEquals<T>(List<T>? a, List<T>? b) {
  if (a == null) return b == null;
  if (b == null || a.length != b.length) return false;
  for (int index = 0; index < a.length; index += 1) {
    if (a[index] != b[index]) return false;
  }
  return true;
}
