import 'package:flutter/material.dart';
import '../../domain/entities/game_state.dart';
import 'generated_image_view.dart';

/// Widget for displaying dish evaluation in progress
class EvaluationWidget extends StatelessWidget {
  const EvaluationWidget({super.key, required this.state});

  final GameState state;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        // Header
        _buildHeader(),
        const SizedBox(height: 32),

        // Evaluation card
        _buildEvaluationCard(),
      ],
    );
  }

  /// Build header with animation
  Widget _buildHeader() {
    return Column(
      children: [
        // Animated balance scale emoji
        TweenAnimationBuilder<double>(
          duration: const Duration(seconds: 2),
          tween: Tween(begin: -0.1, end: 0.1),
          builder: (context, angle, child) {
            return Transform.rotate(
              angle: angle,
              child: const Text('⚖️', style: TextStyle(fontSize: 80)),
            );
          },
        ),
        const SizedBox(height: 16),

        const Text(
          'AI Judge Evaluation',
          style: TextStyle(
            fontSize: 32,
            fontWeight: FontWeight.bold,
            color: Colors.black87,
          ),
          textAlign: TextAlign.center,
        ),
        const SizedBox(height: 8),

        Text(
          'Professional chef AI is rating your dish...',
          style: TextStyle(fontSize: 18, color: Colors.grey[600]),
          textAlign: TextAlign.center,
        ),
      ],
    );
  }

  /// Build evaluation progress card
  Widget _buildEvaluationCard() {
    return Card(
      elevation: 8,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
      child: Padding(
        padding: const EdgeInsets.all(24.0),
        child: Column(
          children: [
            // Loading section
            _buildLoadingSection(),
            const SizedBox(height: 32),

            // Dish being evaluated
            _buildDishSection(),
          ],
        ),
      ),
    );
  }

  /// Build loading section with animated indicators
  Widget _buildLoadingSection() {
    return Container(
      padding: const EdgeInsets.all(24),
      decoration: BoxDecoration(
        color: Colors.purple.withValues(alpha: 0.05),
        borderRadius: BorderRadius.circular(16),
        border: Border.all(color: Colors.purple.withValues(alpha: 0.2)),
      ),
      child: Column(
        children: [
          // Main loading indicator
          SizedBox(
            width: 60,
            height: 60,
            child: CircularProgressIndicator(
              strokeWidth: 6.0,
              valueColor: AlwaysStoppedAnimation<Color>(Colors.purple[400]!),
            ),
          ),
          const SizedBox(height: 20),

          // Status text
          const Text(
            'Analyzing flavors, presentation, and creativity...',
            style: TextStyle(
              fontSize: 18,
              fontWeight: FontWeight.w600,
              color: Colors.black87,
            ),
            textAlign: TextAlign.center,
          ),
          const SizedBox(height: 20),

          // Evaluation criteria animation
          _buildEvaluationCriteria(),
        ],
      ),
    );
  }

  /// Build evaluation criteria with animated indicators
  Widget _buildEvaluationCriteria() {
    const criteria = [
      {
        'name': 'Flavor Balance',
        'icon': Icons.restaurant,
        'color': Colors.orange,
      },
      {
        'name': 'Presentation',
        'icon': Icons.photo_camera,
        'color': Colors.blue,
      },
      {'name': 'Creativity', 'icon': Icons.lightbulb, 'color': Colors.green},
      {
        'name': 'Technique',
        'icon': Icons.precision_manufacturing,
        'color': Colors.red,
      },
    ];

    return Column(
      children: criteria.map((criterion) {
        return Padding(
          padding: const EdgeInsets.symmetric(vertical: 6.0),
          child: Row(
            children: [
              // Animated indicator
              TweenAnimationBuilder<double>(
                duration: const Duration(seconds: 2),
                tween: Tween(begin: 0.0, end: 1.0),
                builder: (context, value, child) {
                  return SizedBox(
                    width: 20,
                    height: 20,
                    child: CircularProgressIndicator(
                      value: value,
                      strokeWidth: 3.0,
                      valueColor: AlwaysStoppedAnimation<Color>(
                        criterion['color'] as Color,
                      ),
                      backgroundColor: (criterion['color'] as Color).withValues(
                        alpha: 0.2,
                      ),
                    ),
                  );
                },
              ),
              const SizedBox(width: 12),

              // Icon
              Icon(
                criterion['icon'] as IconData,
                color: criterion['color'] as Color,
                size: 20,
              ),
              const SizedBox(width: 12),

              // Criterion name
              Text(
                criterion['name'] as String,
                style: const TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.w500,
                  color: Colors.black87,
                ),
              ),
              const Spacer(),

              // Status indicator
              _buildStatusIndicator(),
            ],
          ),
        );
      }).toList(),
    );
  }

  /// Build animated status indicator
  Widget _buildStatusIndicator() {
    return TweenAnimationBuilder<int>(
      duration: const Duration(seconds: 1),
      tween: IntTween(begin: 0, end: 3),
      builder: (context, dotCount, child) {
        return SizedBox(
          width: 30,
          child: Text(
            'analyzing${'.' * dotCount}',
            style: TextStyle(
              fontSize: 12,
              color: Colors.grey[500],
              fontStyle: FontStyle.italic,
            ),
          ),
        );
      },
    );
  }

  /// Build dish section showing recipe and image
  Widget _buildDishSection() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const Text(
          'Dish Under Evaluation',
          style: TextStyle(
            fontSize: 20,
            fontWeight: FontWeight.bold,
            color: Colors.black87,
          ),
        ),
        const SizedBox(height: 16),

        // Two-column layout for recipe and image
        LayoutBuilder(
          builder: (context, constraints) {
            if (constraints.maxWidth > 600) {
              // Desktop/tablet: side by side
              return Row(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Expanded(child: _buildRecipePreview()),
                  const SizedBox(width: 24),
                  Expanded(child: _buildImagePreview()),
                ],
              );
            } else {
              // Mobile: stacked
              return Column(
                children: [
                  _buildRecipePreview(),
                  const SizedBox(height: 24),
                  _buildImagePreview(),
                ],
              );
            }
          },
        ),
      ],
    );
  }

  /// Build recipe preview
  Widget _buildRecipePreview() {
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: Colors.grey[50],
        borderRadius: BorderRadius.circular(12),
        border: Border.all(color: Colors.grey[300]!),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Row(
            children: [
              Icon(Icons.receipt_long, color: Colors.black54, size: 20),
              SizedBox(width: 8),
              Text(
                'Your Recipe',
                style: TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.bold,
                  color: Colors.black87,
                ),
              ),
            ],
          ),
          const SizedBox(height: 12),

          Container(
            constraints: const BoxConstraints(maxHeight: 120),
            child: SingleChildScrollView(
              child: Text(
                state.recipe ?? 'No recipe available',
                style: const TextStyle(
                  fontSize: 14,
                  height: 1.4,
                  color: Colors.black87,
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }

  /// Build image preview
  Widget _buildImagePreview() {
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: Colors.grey[50],
        borderRadius: BorderRadius.circular(12),
        border: Border.all(color: Colors.grey[300]!),
      ),
      child: Column(
        children: [
          const Row(
            children: [
              Icon(Icons.image, color: Colors.black54, size: 20),
              SizedBox(width: 8),
              Text(
                'Your Creation',
                style: TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.bold,
                  color: Colors.black87,
                ),
              ),
            ],
          ),
          const SizedBox(height: 12),

          if (state.imageUrl != null && state.imageUrl!.isNotEmpty)
            GeneratedImageView(
              imageUrl: state.imageUrl!,
              width: 120,
              height: 120,
              fit: BoxFit.cover,
            )
          else
            Container(
              width: 120,
              height: 120,
              decoration: BoxDecoration(
                color: Colors.grey[200],
                borderRadius: BorderRadius.circular(8),
              ),
              child: const Icon(
                Icons.image_outlined,
                color: Colors.grey,
                size: 32,
              ),
            ),
        ],
      ),
    );
  }
}
