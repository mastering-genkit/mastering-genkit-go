import 'package:flutter/material.dart';
import '../../domain/entities/game_state.dart';
import 'generated_image_view.dart';

/// Widget for displaying image generation progress and result
class ImageGenerationWidget extends StatelessWidget {
  const ImageGenerationWidget({super.key, required this.state});

  final GameState state;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        // Header
        _buildHeader(),
        const SizedBox(height: 32),

        // Image generation card
        _buildImageCard(),
      ],
    );
  }

  /// Build header with animation
  Widget _buildHeader() {
    return Column(
      children: [
        // Animated camera emoji
        TweenAnimationBuilder<double>(
          duration: const Duration(seconds: 2),
          tween: Tween(begin: 0.8, end: 1.2),
          builder: (context, scale, child) {
            return Transform.scale(
              scale: scale,
              child: const Text('📸', style: TextStyle(fontSize: 80)),
            );
          },
        ),
        const SizedBox(height: 16),

        const Text(
          'Creating Dish Image',
          style: TextStyle(
            fontSize: 32,
            fontWeight: FontWeight.bold,
            color: Colors.black87,
          ),
          textAlign: TextAlign.center,
        ),
        const SizedBox(height: 8),

        Text(
          'Visualizing your delicious creation...',
          style: TextStyle(fontSize: 18, color: Colors.grey[600]),
          textAlign: TextAlign.center,
        ),
      ],
    );
  }

  /// Build image generation card
  Widget _buildImageCard() {
    return Card(
      elevation: 8,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
      child: Padding(
        padding: const EdgeInsets.all(24.0),
        child: Column(
          children: [
            if (state.imageUrl != null && state.imageUrl!.isNotEmpty)
              // Generated image
              _buildGeneratedImage()
            else if (state.isGeneratingImage == true)
              // Image generation in progress
              _buildImageGenerationProgress()
            else
              // Waiting state
              _buildWaitingState(),
          ],
        ),
      ),
    );
  }

  /// Build generated image display
  Widget _buildGeneratedImage() {
    return Column(
      children: [
        // Success message
        Container(
          padding: const EdgeInsets.all(16),
          decoration: BoxDecoration(
            color: Colors.green.withValues(alpha: 0.1),
            borderRadius: BorderRadius.circular(12),
            border: Border.all(color: Colors.green.withValues(alpha: 0.3)),
          ),
          child: const Row(
            children: [
              Icon(Icons.check_circle, color: Colors.green, size: 24),
              SizedBox(width: 12),
              Text(
                'Image generated successfully!',
                style: TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.w600,
                  color: Colors.green,
                ),
              ),
            ],
          ),
        ),
        const SizedBox(height: 24),

        // Image display using GeneratedImageView
        GeneratedImageView(
          imageUrl: state.imageUrl!,
          width: 400,
          height: 400,
          fit: BoxFit.cover,
        ),
      ],
    );
  }

  /// Build image generation progress
  Widget _buildImageGenerationProgress() {
    final progress = state.imageGenerationProgress ?? 0.0;

    return Column(
      children: [
        // Progress indicator section
        Container(
          padding: const EdgeInsets.all(24),
          decoration: BoxDecoration(
            color: Colors.blue.withValues(alpha: 0.05),
            borderRadius: BorderRadius.circular(16),
            border: Border.all(color: Colors.blue.withValues(alpha: 0.2)),
          ),
          child: Column(
            children: [
              // Spinning loading indicator
              SizedBox(
                width: 64,
                height: 64,
                child: CircularProgressIndicator(
                  strokeWidth: 6.0,
                  valueColor: AlwaysStoppedAnimation<Color>(Colors.blue[400]!),
                ),
              ),
              const SizedBox(height: 20),

              // Status text
              const Text(
                'Creating your dish visualization...',
                style: TextStyle(
                  fontSize: 18,
                  fontWeight: FontWeight.w600,
                  color: Colors.black87,
                ),
                textAlign: TextAlign.center,
              ),
              const SizedBox(height: 8),

              // Progress percentage
              if (progress > 0)
                Text(
                  '${progress.round()}% complete',
                  style: TextStyle(fontSize: 14, color: Colors.grey[600]),
                ),
              const SizedBox(height: 20),

              // Progress bar
              if (progress > 0)
                Column(
                  children: [
                    LinearProgressIndicator(
                      value: progress / 100,
                      backgroundColor: Colors.grey[300],
                      valueColor: AlwaysStoppedAnimation<Color>(
                        Colors.blue[400]!,
                      ),
                      minHeight: 8,
                    ),
                    const SizedBox(height: 16),
                  ],
                ),

              // Motivational text
              _buildMotivationalText(),
            ],
          ),
        ),
      ],
    );
  }

  /// Build waiting state
  Widget _buildWaitingState() {
    return Container(
      padding: const EdgeInsets.all(32),
      child: Column(
        children: [
          Icon(Icons.image_outlined, size: 80, color: Colors.grey[400]),
          const SizedBox(height: 16),
          Text(
            'Ready to generate image...',
            style: TextStyle(fontSize: 18, color: Colors.grey[600]),
            textAlign: TextAlign.center,
          ),
        ],
      ),
    );
  }

  /// Build motivational text that changes based on progress
  Widget _buildMotivationalText() {
    final progress = state.imageGenerationProgress ?? 0.0;
    String message;

    if (progress < 25) {
      message = 'Analyzing your recipe ingredients...';
    } else if (progress < 50) {
      message = 'Composing the perfect dish presentation...';
    } else if (progress < 75) {
      message = 'Adding artistic touches and details...';
    } else if (progress < 95) {
      message = 'Finalizing your delicious creation...';
    } else {
      message = 'Almost ready! Just a few more seconds...';
    }

    return AnimatedSwitcher(
      duration: const Duration(milliseconds: 500),
      child: Text(
        message,
        key: ValueKey(message),
        style: TextStyle(
          fontSize: 14,
          color: Colors.grey[600],
          fontStyle: FontStyle.italic,
        ),
        textAlign: TextAlign.center,
      ),
    );
  }
}
