import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../domain/entities/game_state.dart';
import '../providers/game_state_notifier.dart';
import '../widgets/progress_bar_widget.dart';
import '../widgets/ready_widget.dart';
import '../widgets/select_ingredients_widget.dart';
import '../widgets/recipe_generation_widget.dart';
import '../widgets/image_generation_widget.dart';
import '../widgets/evaluation_widget.dart';
import '../widgets/result_widget.dart';
import '../widgets/error_widget.dart';

/// Main Recipe Quest page that displays different widgets based on game state
class RecipeQuestPage extends ConsumerWidget {
  const RecipeQuestPage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final gameState = ref.watch(gameStateProvider);
    final gameNotifier = ref.read(gameStateProvider.notifier);

    return Scaffold(
      body: Container(
        decoration: const BoxDecoration(
          gradient: LinearGradient(
            begin: Alignment.topLeft,
            end: Alignment.bottomRight,
            colors: [
              Color(0xFFFEF3E2), // Light orange
              Color(0xFFFDE8E8), // Light pink
            ],
          ),
        ),
        child: SafeArea(
          child: Column(
            children: [
              // Fixed Progress Bar (only visible during processing)
              if (_shouldShowProgressBar(gameState.currentStep))
                ProgressBarWidget(state: gameState),

              // Main Content
              Expanded(
                child: SingleChildScrollView(
                  padding: EdgeInsets.only(
                    left: 16.0,
                    right: 16.0,
                    top: _shouldShowProgressBar(gameState.currentStep)
                        ? 16.0
                        : 32.0,
                    bottom: 16.0,
                  ),
                  child: _buildCurrentStepWidget(gameState, gameNotifier),
                ),
              ),

              // Error Display (if present)
              if (gameState.error != null)
                RecipeQuestErrorWidget(
                  error: gameState.error!,
                  onReset: gameNotifier.resetGame,
                ),
            ],
          ),
        ),
      ),
    );
  }

  /// Determine if progress bar should be visible
  bool _shouldShowProgressBar(GameStep currentStep) {
    return currentStep != GameStep.ready &&
        currentStep != GameStep.selectIngredients &&
        currentStep != GameStep.result;
  }

  /// Build the appropriate widget for current game step
  Widget _buildCurrentStepWidget(GameState state, GameStateNotifier notifier) {
    switch (state.currentStep) {
      case GameStep.ready:
        return ReadyWidget(onStartGame: notifier.startGame);

      case GameStep.selectIngredients:
        return SelectIngredientsWidget(
          state: state,
          onToggleIngredient: notifier.toggleIngredient,
          onStartRecipeGeneration: notifier.startRecipeGeneration,
        );

      case GameStep.recipe:
        return RecipeGenerationWidget(state: state);

      case GameStep.image:
        return ImageGenerationWidget(state: state);

      case GameStep.evaluation:
        return EvaluationWidget(state: state);

      case GameStep.result:
        return ResultWidget(state: state, onPlayAgain: notifier.resetGame);
    }
  }
}
