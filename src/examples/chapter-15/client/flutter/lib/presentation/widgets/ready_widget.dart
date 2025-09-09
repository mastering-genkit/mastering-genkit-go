import 'package:flutter/material.dart';

/// Widget displayed in the ready state - game start screen
class ReadyWidget extends StatelessWidget {
  const ReadyWidget({super.key, required this.onStartGame});

  final VoidCallback onStartGame;

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Padding(
        padding: const EdgeInsets.all(24.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            // Animated target emoji
            TweenAnimationBuilder<double>(
              duration: const Duration(seconds: 2),
              tween: Tween(begin: 0.8, end: 1.2),
              builder: (context, scale, child) {
                return Transform.scale(
                  scale: scale,
                  child: const Text('🎯', style: TextStyle(fontSize: 96)),
                );
              },
            ),
            const SizedBox(height: 32),

            // Game title
            const Text(
              'Recipe Quest',
              style: TextStyle(
                fontSize: 48,
                fontWeight: FontWeight.bold,
                color: Colors.black87,
              ),
              textAlign: TextAlign.center,
            ),
            const SizedBox(height: 16),

            // Subtitle
            Text(
              'Challenge yourself to create amazing dishes!',
              style: TextStyle(fontSize: 20, color: Colors.grey[600]),
              textAlign: TextAlign.center,
            ),
            const SizedBox(height: 48),

            // Start Quest button
            ElevatedButton.icon(
              onPressed: onStartGame,
              style: ElevatedButton.styleFrom(
                backgroundColor: Colors.orange,
                foregroundColor: Colors.white,
                padding: const EdgeInsets.symmetric(
                  horizontal: 48,
                  vertical: 24,
                ),
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(16),
                ),
                elevation: 8,
                shadowColor: Colors.orange.withValues(alpha: 0.5),
              ),
              icon: const Icon(Icons.play_arrow, size: 28),
              label: const Text(
                'Start Quest',
                style: TextStyle(fontSize: 24, fontWeight: FontWeight.w600),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
