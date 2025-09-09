import 'package:flutter/material.dart';
import '../../domain/entities/game_state.dart';
import '../providers/game_state_notifier.dart';

/// Fixed progress bar widget shown during processing steps
class ProgressBarWidget extends StatelessWidget {
  const ProgressBarWidget({super.key, required this.state});

  final GameState state;

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        color: Colors.white.withValues(alpha: 0.95),
        border: const Border(
          bottom: BorderSide(color: Colors.orange, width: 4.0),
        ),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withValues(alpha: 0.1),
            blurRadius: 8.0,
            offset: const Offset(0, 2),
          ),
        ],
      ),
      child: SafeArea(
        bottom: false,
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            children: [
              // Main Progress Section
              _buildMainProgressSection(),
              const SizedBox(height: 12),

              // Progress Bar
              LinearProgressIndicator(
                value: state.progress / 100,
                backgroundColor: Colors.grey[300],
                valueColor: const AlwaysStoppedAnimation<Color>(Colors.orange),
                minHeight: 6,
              ),
              const SizedBox(height: 12),

              // Step Indicators
              _buildStepIndicators(),
            ],
          ),
        ),
      ),
    );
  }

  /// Build main progress section with proper overflow handling
  Widget _buildMainProgressSection() {
    return Row(
      children: [
        // Left side: Icon and info
        Expanded(
          flex: 3,
          child: Row(
            children: [
              _buildStepIcon(),
              const SizedBox(width: 12),
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    Text(
                      _getStepTitle(),
                      style: const TextStyle(
                        fontSize: 16,
                        fontWeight: FontWeight.bold,
                        color: Colors.black87,
                      ),
                      maxLines: 1,
                      overflow: TextOverflow.ellipsis,
                    ),
                    const SizedBox(height: 2),
                    _buildIngredientsInfo(),
                  ],
                ),
              ),
            ],
          ),
        ),

        // Right side: Loading and progress
        _buildProgressInfo(),
      ],
    );
  }

  /// Build step icon with animation
  Widget _buildStepIcon() {
    String emoji;
    switch (state.currentStep) {
      case GameStep.recipe:
        emoji = '🍳';
        break;
      case GameStep.image:
        emoji = '📸';
        break;
      case GameStep.evaluation:
        emoji = '⚖️';
        break;
      default:
        emoji = '🎯';
    }

    return TweenAnimationBuilder<double>(
      duration: const Duration(milliseconds: 1000),
      tween: Tween(begin: 0.8, end: 1.2),
      builder: (context, scale, child) {
        return Transform.scale(
          scale: scale,
          child: Text(emoji, style: const TextStyle(fontSize: 28)),
        );
      },
    );
  }

  /// Get current step title
  String _getStepTitle() {
    switch (state.currentStep) {
      case GameStep.recipe:
        return 'Creating Recipe';
      case GameStep.image:
        return 'Generating Image';
      case GameStep.evaluation:
        return 'Evaluating Dish';
      default:
        return 'Processing';
    }
  }

  /// Build ingredients info with safe overflow handling
  Widget _buildIngredientsInfo() {
    // Count emojis and text length to estimate if it fits
    final emojiCount = state.selectedIngredients.length;
    final textContent = state.selectedIngredients.join(', ');

    return LayoutBuilder(
      builder: (context, constraints) {
        // Use a more compact display if we have limited space
        if (constraints.maxWidth < 200 || textContent.length > 30) {
          return Text(
            '$emojiCount ingredients selected',
            style: const TextStyle(fontSize: 12, color: Colors.black54),
            maxLines: 1,
            overflow: TextOverflow.ellipsis,
          );
        }

        // If we have space, show emojis + truncated text
        return Row(
          children: [
            // Show max 4 emojis to avoid overflow
            for (int i = 0; i < emojiCount && i < 4; i++)
              Padding(
                padding: const EdgeInsets.only(right: 2.0),
                child: Text(
                  _getIngredientEmoji(state.selectedIngredients[i]),
                  style: const TextStyle(fontSize: 14),
                ),
              ),
            if (emojiCount > 4)
              Text(
                '+${emojiCount - 4}',
                style: const TextStyle(fontSize: 12, color: Colors.black54),
              ),
            const SizedBox(width: 6),
            Expanded(
              child: Text(
                '• $textContent',
                style: const TextStyle(fontSize: 12, color: Colors.black54),
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
              ),
            ),
          ],
        );
      },
    );
  }

  /// Build progress info section
  Widget _buildProgressInfo() {
    return Row(
      mainAxisSize: MainAxisSize.min,
      children: [
        if (state.isLoading) ...[
          const SizedBox(
            width: 16,
            height: 16,
            child: CircularProgressIndicator(
              strokeWidth: 2.0,
              valueColor: AlwaysStoppedAnimation<Color>(Colors.orange),
            ),
          ),
          const SizedBox(width: 8),
          const Text(
            'Processing...',
            style: TextStyle(fontSize: 12, color: Colors.black54),
          ),
          const SizedBox(width: 12),
        ],
        Column(
          crossAxisAlignment: CrossAxisAlignment.end,
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(
              '${state.progress.round()}%',
              style: const TextStyle(
                fontSize: 18,
                fontWeight: FontWeight.bold,
                color: Colors.orange,
              ),
            ),
          ],
        ),
      ],
    );
  }

  /// Get emoji for ingredient
  String _getIngredientEmoji(String ingredientName) {
    final ingredient = availableIngredients.firstWhere(
      (ing) => ing.name == ingredientName,
      orElse: () => const Ingredient(name: '', emoji: '🍽️'),
    );
    return ingredient.emoji;
  }

  /// Build step indicators (circular progress indicators)
  Widget _buildStepIndicators() {
    final steps = [
      GameStep.recipe,
      GameStep.image,
      GameStep.evaluation,
      GameStep.result,
    ];

    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        for (int i = 0; i < steps.length; i++) ...[
          _buildStepDot(steps[i]),
          if (i < steps.length - 1) _buildStepConnector(steps[i]),
        ],
      ],
    );
  }

  /// Build individual step dot indicator
  Widget _buildStepDot(GameStep step) {
    final isCompleted = _getStepIndex(state.currentStep) > _getStepIndex(step);
    final isCurrent = state.currentStep == step;

    return AnimatedContainer(
      duration: const Duration(milliseconds: 300),
      width: isCurrent ? 16 : 12,
      height: isCurrent ? 16 : 12,
      decoration: BoxDecoration(
        shape: BoxShape.circle,
        color: isCompleted
            ? Colors.green
            : isCurrent
            ? Colors.orange
            : Colors.grey[400],
        border: Border.all(
          color: isCompleted
              ? Colors.green
              : isCurrent
              ? Colors.orange
              : Colors.grey[400]!,
          width: 2,
        ),
      ),
      child: isCurrent
          ? TweenAnimationBuilder<double>(
              duration: const Duration(seconds: 1),
              tween: Tween(begin: 0.5, end: 1.0),
              builder: (context, opacity, child) {
                return Container(
                  decoration: BoxDecoration(
                    shape: BoxShape.circle,
                    color: Colors.orange.withValues(alpha: opacity),
                  ),
                );
              },
            )
          : null,
    );
  }

  /// Build connector line between step dots
  Widget _buildStepConnector(GameStep step) {
    final isCompleted = _getStepIndex(state.currentStep) > _getStepIndex(step);

    return Container(
      width: 32,
      height: 2,
      margin: const EdgeInsets.symmetric(horizontal: 4),
      decoration: BoxDecoration(
        color: isCompleted ? Colors.green : Colors.grey[400],
        borderRadius: BorderRadius.circular(1),
      ),
    );
  }

  /// Get step index for comparison
  int _getStepIndex(GameStep step) {
    switch (step) {
      case GameStep.ready:
        return 0;
      case GameStep.selectIngredients:
        return 1;
      case GameStep.recipe:
        return 2;
      case GameStep.image:
        return 3;
      case GameStep.evaluation:
        return 4;
      case GameStep.result:
        return 5;
    }
  }
}
