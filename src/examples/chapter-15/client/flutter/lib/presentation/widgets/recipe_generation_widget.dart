import 'package:flutter/material.dart';
import '../../domain/entities/game_state.dart';

/// Widget for displaying recipe generation progress and content
class RecipeGenerationWidget extends StatelessWidget {
  const RecipeGenerationWidget({super.key, required this.state});

  final GameState state;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        // Header
        _buildHeader(),
        const SizedBox(height: 32),

        // Recipe card
        _buildRecipeCard(),
      ],
    );
  }

  /// Build header with animation
  Widget _buildHeader() {
    return Column(
      children: [
        // Animated cooking emoji
        TweenAnimationBuilder<double>(
          duration: const Duration(seconds: 2),
          tween: Tween(begin: 0.8, end: 1.2),
          builder: (context, scale, child) {
            return Transform.scale(
              scale: scale,
              child: const Text('🍳', style: TextStyle(fontSize: 80)),
            );
          },
        ),
        const SizedBox(height: 16),

        const Text(
          'Creating Your Recipe',
          style: TextStyle(
            fontSize: 32,
            fontWeight: FontWeight.bold,
            color: Colors.black87,
          ),
          textAlign: TextAlign.center,
        ),
        const SizedBox(height: 8),

        Text(
          'AI chef is working on your custom recipe...',
          style: TextStyle(fontSize: 18, color: Colors.grey[600]),
          textAlign: TextAlign.center,
        ),
      ],
    );
  }

  /// Build recipe display card
  Widget _buildRecipeCard() {
    return Card(
      elevation: 8,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
      child: Padding(
        padding: const EdgeInsets.all(24.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Ingredients summary
            _buildIngredientsSection(),
            const SizedBox(height: 24),

            // Recipe content
            _buildRecipeContent(),
          ],
        ),
      ),
    );
  }

  /// Build ingredients summary section
  Widget _buildIngredientsSection() {
    return Container(
      padding: const EdgeInsets.all(16.0),
      decoration: BoxDecoration(
        color: Colors.orange.withValues(alpha: 0.1),
        borderRadius: BorderRadius.circular(12),
        border: Border.all(color: Colors.orange.withValues(alpha: 0.3)),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            children: [
              const Icon(Icons.inventory_2, color: Colors.orange, size: 20),
              const SizedBox(width: 8),
              const Text(
                'Your Selected Ingredients:',
                style: TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.bold,
                  color: Colors.black87,
                ),
              ),
            ],
          ),
          const SizedBox(height: 12),

          // Ingredients with emojis
          Wrap(
            spacing: 12,
            runSpacing: 8,
            children: state.selectedIngredients.map((ingredient) {
              return Container(
                padding: const EdgeInsets.symmetric(
                  horizontal: 12,
                  vertical: 6,
                ),
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.circular(20),
                  border: Border.all(
                    color: Colors.orange.withValues(alpha: 0.5),
                  ),
                ),
                child: Row(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    Text(
                      _getIngredientEmoji(ingredient),
                      style: const TextStyle(fontSize: 16),
                    ),
                    const SizedBox(width: 6),
                    Text(
                      ingredient,
                      style: const TextStyle(
                        fontSize: 14,
                        fontWeight: FontWeight.w500,
                        color: Colors.black87,
                      ),
                    ),
                  ],
                ),
              );
            }).toList(),
          ),
        ],
      ),
    );
  }

  /// Build recipe content area
  Widget _buildRecipeContent() {
    return Container(
      constraints: const BoxConstraints(minHeight: 200),
      padding: const EdgeInsets.all(20.0),
      decoration: BoxDecoration(
        color: Colors.grey[50],
        borderRadius: BorderRadius.circular(12),
        border: Border.all(color: Colors.grey[300]!),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            children: [
              const Icon(Icons.receipt_long, color: Colors.black54, size: 20),
              const SizedBox(width: 8),
              const Text(
                'Recipe:',
                style: TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.bold,
                  color: Colors.black87,
                ),
              ),
            ],
          ),
          const SizedBox(height: 16),

          if (state.recipe != null && state.recipe!.isNotEmpty)
            // Recipe content with typing animation
            _buildRecipeText()
          else if (state.isLoading)
            // Loading state
            _buildLoadingState()
          else
            // Empty state
            _buildEmptyState(),
        ],
      ),
    );
  }

  /// Build recipe text with streaming effect
  Widget _buildRecipeText() {
    return SelectableText(
      state.recipe!,
      style: const TextStyle(fontSize: 16, height: 1.6, color: Colors.black87),
    );
  }

  /// Build loading state
  Widget _buildLoadingState() {
    return Column(
      children: [
        Row(
          children: [
            SizedBox(
              width: 20,
              height: 20,
              child: CircularProgressIndicator(
                strokeWidth: 2.0,
                valueColor: AlwaysStoppedAnimation<Color>(Colors.grey[600]!),
              ),
            ),
            const SizedBox(width: 12),
            Expanded(
              child: Text(
                'Generating your personalized recipe...',
                style: TextStyle(
                  fontSize: 16,
                  color: Colors.grey[600],
                  fontStyle: FontStyle.italic,
                ),
                overflow: TextOverflow.ellipsis,
                maxLines: 2,
              ),
            ),
          ],
        ),
        const SizedBox(height: 16),

        // Animated dots
        _buildTypingAnimation(),
      ],
    );
  }

  /// Build empty state
  Widget _buildEmptyState() {
    return Text(
      'Recipe will appear here...',
      style: TextStyle(
        fontSize: 16,
        color: Colors.grey[500],
        fontStyle: FontStyle.italic,
      ),
    );
  }

  /// Build typing animation with dots
  Widget _buildTypingAnimation() {
    return TweenAnimationBuilder<int>(
      duration: const Duration(seconds: 2),
      tween: IntTween(begin: 0, end: 3),
      builder: (context, dotCount, child) {
        return Text(
          'Cooking up something delicious${'.' * dotCount}',
          style: TextStyle(
            fontSize: 14,
            color: Colors.grey[500],
            fontStyle: FontStyle.italic,
          ),
        );
      },
    );
  }

  /// Get emoji for ingredient name
  String _getIngredientEmoji(String ingredientName) {
    // This should match the availableIngredients list
    final ingredientMap = <String, String>{
      'chicken': '🐔',
      'beef': '🥩',
      'pork': '🐷',
      'salmon': '🐟',
      'shrimp': '🦐',
      'tofu': '🥡',
      'rice': '🍚',
      'noodles': '🍜',
      'pasta': '🍝',
      'potatoes': '🥔',
      'onions': '🧅',
      'garlic': '🧄',
      'ginger': '🫚',
      'carrots': '🥕',
      'peppers': '🌶️',
      'vegetables': '🥬',
      'mushrooms': '🍄',
      'tomatoes': '🍅',
      'lemon': '🍋',
      'herbs': '🌿',
      'sesame oil': '🫗',
      'soy sauce': '🥢',
      'miso': '🍲',
      'quinoa': '🌾',
      'avocado': '🥑',
      'lime': '🍈',
    };
    return ingredientMap[ingredientName] ?? '🍽️';
  }
}
