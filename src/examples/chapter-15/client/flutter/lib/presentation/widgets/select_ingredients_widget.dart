import 'package:flutter/material.dart';
import '../../domain/entities/game_state.dart';
import '../providers/game_state_notifier.dart';

/// Widget for selecting ingredients (step 2 of the quest)
class SelectIngredientsWidget extends StatelessWidget {
  const SelectIngredientsWidget({
    super.key,
    required this.state,
    required this.onToggleIngredient,
    required this.onStartRecipeGeneration,
  });

  final GameState state;
  final Function(String) onToggleIngredient;
  final VoidCallback onStartRecipeGeneration;

  static const int maxIngredients = 4;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        // Header section
        _buildHeader(),
        const SizedBox(height: 32),

        // Ingredients grid
        _buildIngredientsGrid(),
        const SizedBox(height: 32),

        // Generate Recipe button
        _buildGenerateButton(),
      ],
    );
  }

  /// Build header with title and selection counter
  Widget _buildHeader() {
    return Column(
      children: [
        // Dice emoji with animation
        TweenAnimationBuilder<double>(
          duration: const Duration(seconds: 1),
          tween: Tween(begin: 0.0, end: 1.0),
          builder: (context, value, child) {
            return Transform.rotate(
              angle: value * 0.5,
              child: const Text('🎲', style: TextStyle(fontSize: 80)),
            );
          },
        ),
        const SizedBox(height: 16),

        const Text(
          'Choose Your Ingredients',
          style: TextStyle(
            fontSize: 32,
            fontWeight: FontWeight.bold,
            color: Colors.black87,
          ),
          textAlign: TextAlign.center,
        ),
        const SizedBox(height: 8),

        Text(
          'Select exactly $maxIngredients ingredients for your recipe quest',
          style: TextStyle(fontSize: 18, color: Colors.grey[600]),
          textAlign: TextAlign.center,
        ),
        const SizedBox(height: 12),

        // Selection counter
        Container(
          padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
          decoration: BoxDecoration(
            color: state.selectedIngredients.length == maxIngredients
                ? Colors.green.withValues(alpha: 0.1)
                : Colors.orange.withValues(alpha: 0.1),
            borderRadius: BorderRadius.circular(20),
            border: Border.all(
              color: state.selectedIngredients.length == maxIngredients
                  ? Colors.green
                  : Colors.orange,
              width: 2,
            ),
          ),
          child: Text(
            'Selected: ${state.selectedIngredients.length}/$maxIngredients',
            style: TextStyle(
              fontSize: 16,
              fontWeight: FontWeight.w600,
              color: state.selectedIngredients.length == maxIngredients
                  ? Colors.green[700]
                  : Colors.orange[700],
            ),
          ),
        ),
      ],
    );
  }

  /// Build ingredients grid
  Widget _buildIngredientsGrid() {
    return LayoutBuilder(
      builder: (context, constraints) {
        // Responsive grid columns based on screen width
        int crossAxisCount = 2;
        if (constraints.maxWidth > 600) crossAxisCount = 4;
        if (constraints.maxWidth > 900) crossAxisCount = 6;

        return GridView.builder(
          shrinkWrap: true,
          physics: const NeverScrollableScrollPhysics(),
          gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
            crossAxisCount: crossAxisCount,
            childAspectRatio: 1.0,
            crossAxisSpacing: 12,
            mainAxisSpacing: 12,
          ),
          itemCount: availableIngredients.length,
          itemBuilder: (context, index) {
            final ingredient = availableIngredients[index];
            return _buildIngredientCard(ingredient);
          },
        );
      },
    );
  }

  /// Build individual ingredient card
  Widget _buildIngredientCard(Ingredient ingredient) {
    final isSelected = state.selectedIngredients.contains(ingredient.name);
    final canSelect =
        state.selectedIngredients.length < maxIngredients || isSelected;

    return AnimatedContainer(
      duration: const Duration(milliseconds: 200),
      curve: Curves.easeInOut,
      transform: Matrix4.identity()..scale(isSelected ? 1.05 : 1.0),
      child: Material(
        elevation: isSelected ? 8 : 2,
        borderRadius: BorderRadius.circular(16),
        color: isSelected ? Colors.orange.withValues(alpha: 0.1) : Colors.white,
        child: InkWell(
          onTap: canSelect ? () => onToggleIngredient(ingredient.name) : null,
          borderRadius: BorderRadius.circular(16),
          child: Container(
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(16),
              border: Border.all(
                color: isSelected
                    ? Colors.orange
                    : canSelect
                    ? Colors.grey[300]!
                    : Colors.grey[200]!,
                width: isSelected ? 3 : 1,
              ),
            ),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                // Ingredient emoji
                Text(
                  ingredient.emoji,
                  style: TextStyle(
                    fontSize: 40,
                    color: canSelect ? null : Colors.grey[400],
                  ),
                ),
                const SizedBox(height: 8),

                // Ingredient name
                Text(
                  ingredient.name,
                  style: TextStyle(
                    fontSize: 14,
                    fontWeight: FontWeight.w600,
                    color: canSelect ? Colors.black87 : Colors.grey[400],
                  ),
                  textAlign: TextAlign.center,
                  maxLines: 2,
                  overflow: TextOverflow.ellipsis,
                ),

                // Selection indicator
                if (isSelected)
                  Container(
                    margin: const EdgeInsets.only(top: 4),
                    width: 24,
                    height: 24,
                    decoration: const BoxDecoration(
                      shape: BoxShape.circle,
                      color: Colors.orange,
                    ),
                    child: const Icon(
                      Icons.check,
                      color: Colors.white,
                      size: 16,
                    ),
                  ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  /// Build generate recipe button
  Widget _buildGenerateButton() {
    final canGenerate = state.selectedIngredients.length == maxIngredients;

    return AnimatedContainer(
      duration: const Duration(milliseconds: 300),
      child: ElevatedButton.icon(
        onPressed: canGenerate ? onStartRecipeGeneration : null,
        style: ElevatedButton.styleFrom(
          backgroundColor: canGenerate ? Colors.green : Colors.grey[300],
          foregroundColor: canGenerate ? Colors.white : Colors.grey[500],
          padding: const EdgeInsets.symmetric(horizontal: 48, vertical: 20),
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(16),
          ),
          elevation: canGenerate ? 8 : 2,
          shadowColor: canGenerate
              ? Colors.green.withValues(alpha: 0.5)
              : Colors.transparent,
        ),
        icon: Icon(canGenerate ? Icons.restaurant_menu : Icons.lock, size: 24),
        label: Text(
          'Create Recipe (${state.selectedIngredients.length}/$maxIngredients)',
          style: const TextStyle(fontSize: 20, fontWeight: FontWeight.w600),
        ),
      ),
    );
  }
}
