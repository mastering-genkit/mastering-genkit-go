import 'package:flutter/material.dart';
import '../../domain/entities/game_state.dart';
import 'generated_image_view.dart';

/// Widget for displaying quest completion results
class ResultWidget extends StatelessWidget {
  const ResultWidget({
    super.key,
    required this.state,
    required this.onPlayAgain,
  });

  final GameState state;
  final VoidCallback onPlayAgain;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        // Header
        _buildHeader(),
        const SizedBox(height: 32),

        // Results card
        _buildResultsCard(),
        const SizedBox(height: 32),

        // Play Again button
        _buildPlayAgainButton(),
      ],
    );
  }

  /// Build header with trophy and title
  Widget _buildHeader() {
    return Column(
      children: [
        // Animated trophy emoji
        TweenAnimationBuilder<double>(
          duration: const Duration(seconds: 2),
          tween: Tween(begin: 0.8, end: 1.2),
          builder: (context, scale, child) {
            return Transform.scale(
              scale: scale,
              child: const Text('🏆', style: TextStyle(fontSize: 96)),
            );
          },
        ),
        const SizedBox(height: 16),

        const Text(
          'Quest Complete!',
          style: TextStyle(
            fontSize: 36,
            fontWeight: FontWeight.bold,
            color: Colors.black87,
          ),
          textAlign: TextAlign.center,
        ),
        const SizedBox(height: 12),

        // Title and achievement from evaluation
        if (state.title != null) ...[
          Text(
            state.title!,
            style: const TextStyle(
              fontSize: 24,
              fontWeight: FontWeight.w600,
              color: Colors.orange,
            ),
            textAlign: TextAlign.center,
          ),
          const SizedBox(height: 8),
        ],

        if (state.achievement != null) ...[
          Container(
            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
            decoration: BoxDecoration(
              color: Colors.purple.withValues(alpha: 0.1),
              borderRadius: BorderRadius.circular(20),
              border: Border.all(color: Colors.purple.withValues(alpha: 0.3)),
            ),
            child: Row(
              mainAxisSize: MainAxisSize.min,
              children: [
                const Text('🎖️', style: TextStyle(fontSize: 20)),
                const SizedBox(width: 8),
                Flexible(
                  child: Text(
                    state.achievement!,
                    style: const TextStyle(
                      fontSize: 18,
                      fontWeight: FontWeight.w600,
                      color: Colors.purple,
                    ),
                    maxLines: 2,
                    overflow: TextOverflow.ellipsis,
                    textAlign: TextAlign.center,
                  ),
                ),
              ],
            ),
          ),
        ],
      ],
    );
  }

  /// Build results display card
  Widget _buildResultsCard() {
    return Card(
      elevation: 8,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
      child: Padding(
        padding: const EdgeInsets.all(24.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Recipe section
            _buildRecipeSection(),

            if (state.imageUrl != null) ...[
              const SizedBox(height: 32),
              _buildImageSection(),
            ],

            const SizedBox(height: 32),
            _buildScoreAndFeedbackSection(),
          ],
        ),
      ),
    );
  }

  /// Build recipe section
  Widget _buildRecipeSection() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const Row(
          children: [
            Icon(Icons.receipt_long, color: Colors.orange, size: 24),
            SizedBox(width: 12),
            Text(
              'Your Recipe',
              style: TextStyle(
                fontSize: 24,
                fontWeight: FontWeight.bold,
                color: Colors.black87,
              ),
            ),
          ],
        ),
        const SizedBox(height: 16),

        Container(
          width: double.infinity,
          padding: const EdgeInsets.all(20),
          decoration: BoxDecoration(
            color: Colors.grey[50],
            borderRadius: BorderRadius.circular(12),
            border: Border.all(color: Colors.grey[300]!),
          ),
          child: SelectableText(
            state.recipe ?? 'No recipe available',
            style: const TextStyle(
              fontSize: 16,
              height: 1.6,
              color: Colors.black87,
            ),
          ),
        ),
      ],
    );
  }

  /// Build image section
  Widget _buildImageSection() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const Row(
          children: [
            Icon(Icons.image, color: Colors.blue, size: 24),
            SizedBox(width: 12),
            Text(
              'Your Creation',
              style: TextStyle(
                fontSize: 24,
                fontWeight: FontWeight.bold,
                color: Colors.black87,
              ),
            ),
          ],
        ),
        const SizedBox(height: 16),

        Center(
          child: GeneratedImageView(
            imageUrl: state.imageUrl!,
            width: 400,
            height: 400,
            fit: BoxFit.cover,
          ),
        ),
      ],
    );
  }

  /// Build score and feedback section
  Widget _buildScoreAndFeedbackSection() {
    return Column(
      children: [
        // Score display
        if (state.score != null) ...[
          Container(
            padding: const EdgeInsets.all(24),
            decoration: BoxDecoration(
              gradient: LinearGradient(
                colors: [
                  Colors.orange.withValues(alpha: 0.1),
                  Colors.red.withValues(alpha: 0.1),
                ],
                begin: Alignment.topLeft,
                end: Alignment.bottomRight,
              ),
              borderRadius: BorderRadius.circular(16),
              border: Border.all(color: Colors.orange.withValues(alpha: 0.3)),
            ),
            child: Column(
              children: [
                const Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Icon(Icons.star, color: Colors.orange, size: 32),
                    SizedBox(width: 8),
                    Text(
                      'Your Score',
                      style: TextStyle(
                        fontSize: 24,
                        fontWeight: FontWeight.bold,
                        color: Colors.black87,
                      ),
                    ),
                  ],
                ),
                const SizedBox(height: 16),

                // Animated score display
                TweenAnimationBuilder<double>(
                  duration: const Duration(seconds: 2),
                  tween: Tween(begin: 0.0, end: state.score!),
                  builder: (context, animatedScore, child) {
                    return Text(
                      '${animatedScore.round()}/100',
                      style: const TextStyle(
                        fontSize: 48,
                        fontWeight: FontWeight.bold,
                        color: Colors.orange,
                      ),
                    );
                  },
                ),
                const SizedBox(height: 8),

                // Score interpretation
                _buildScoreInterpretation(state.score!),
              ],
            ),
          ),
          const SizedBox(height: 24),
        ],

        // Feedback display
        if (state.feedback != null) ...[
          Container(
            width: double.infinity,
            padding: const EdgeInsets.all(20),
            decoration: BoxDecoration(
              color: Colors.blue.withValues(alpha: 0.05),
              borderRadius: BorderRadius.circular(16),
              border: Border.all(color: Colors.blue.withValues(alpha: 0.2)),
            ),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const Row(
                  children: [
                    Icon(Icons.comment, color: Colors.blue, size: 24),
                    SizedBox(width: 12),
                    Text(
                      "Chef's Feedback:",
                      style: TextStyle(
                        fontSize: 20,
                        fontWeight: FontWeight.bold,
                        color: Colors.black87,
                      ),
                    ),
                  ],
                ),
                const SizedBox(height: 16),

                Text(
                  state.feedback!,
                  style: const TextStyle(
                    fontSize: 16,
                    height: 1.6,
                    color: Colors.black87,
                  ),
                ),
              ],
            ),
          ),
        ],
      ],
    );
  }

  /// Build score interpretation widget
  Widget _buildScoreInterpretation(double score) {
    String interpretation;
    Color color;
    IconData icon;

    if (score >= 90) {
      interpretation = 'Outstanding! Master Chef Level!';
      color = Colors.purple;
      icon = Icons.emoji_events;
    } else if (score >= 80) {
      interpretation = 'Excellent! Professional Quality!';
      color = Colors.green;
      icon = Icons.thumb_up;
    } else if (score >= 70) {
      interpretation = 'Great Job! Very Tasty!';
      color = Colors.blue;
      icon = Icons.favorite;
    } else if (score >= 60) {
      interpretation = 'Good Effort! Keep Cooking!';
      color = Colors.orange;
      icon = Icons.local_fire_department;
    } else {
      interpretation = 'Keep Practicing! You\'ll Get There!';
      color = Colors.grey;
      icon = Icons.trending_up;
    }

    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Icon(icon, color: color, size: 20),
        const SizedBox(width: 8),
        Text(
          interpretation,
          style: TextStyle(
            fontSize: 16,
            fontWeight: FontWeight.w600,
            color: color,
          ),
        ),
      ],
    );
  }

  /// Build play again button
  Widget _buildPlayAgainButton() {
    return ElevatedButton.icon(
      onPressed: onPlayAgain,
      style: ElevatedButton.styleFrom(
        backgroundColor: Colors.orange,
        foregroundColor: Colors.white,
        padding: const EdgeInsets.symmetric(horizontal: 48, vertical: 20),
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
        elevation: 8,
        shadowColor: Colors.orange.withValues(alpha: 0.5),
      ),
      icon: const Icon(Icons.refresh, size: 24),
      label: const Text(
        'Play Again',
        style: TextStyle(fontSize: 20, fontWeight: FontWeight.w600),
      ),
    );
  }
}
