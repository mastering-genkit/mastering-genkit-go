import 'package:flutter/material.dart';

/// Widget for displaying errors in Recipe Quest
class RecipeQuestErrorWidget extends StatelessWidget {
  const RecipeQuestErrorWidget({
    super.key,
    required this.error,
    required this.onReset,
  });

  final String error;
  final VoidCallback onReset;

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.all(16.0),
      child: Card(
        elevation: 8,
        color: Colors.red[50],
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(16),
          side: BorderSide(color: Colors.red[200]!, width: 2),
        ),
        child: Padding(
          padding: const EdgeInsets.all(20.0),
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              // Error icon and title
              Row(
                children: [
                  Container(
                    padding: const EdgeInsets.all(8),
                    decoration: BoxDecoration(
                      color: Colors.red[100],
                      shape: BoxShape.circle,
                    ),
                    child: Icon(
                      Icons.error_outline,
                      color: Colors.red[700],
                      size: 32,
                    ),
                  ),
                  const SizedBox(width: 16),

                  Expanded(
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text(
                          'Oops! Something went wrong',
                          style: TextStyle(
                            fontSize: 18,
                            fontWeight: FontWeight.bold,
                            color: Colors.red[800],
                          ),
                        ),
                        const SizedBox(height: 4),

                        Text(
                          'Don\'t worry, these things happen!',
                          style: TextStyle(
                            fontSize: 14,
                            color: Colors.red[600],
                          ),
                        ),
                      ],
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 20),

              // Error message
              Container(
                width: double.infinity,
                padding: const EdgeInsets.all(16),
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.circular(12),
                  border: Border.all(color: Colors.red[300]!),
                ),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      'Error Details:',
                      style: TextStyle(
                        fontSize: 14,
                        fontWeight: FontWeight.w600,
                        color: Colors.red[700],
                      ),
                    ),
                    const SizedBox(height: 8),

                    SelectableText(
                      error,
                      style: TextStyle(
                        fontSize: 14,
                        color: Colors.red[800],
                        height: 1.4,
                      ),
                    ),
                  ],
                ),
              ),
              const SizedBox(height: 20),

              // Helpful suggestions
              _buildSuggestions(),
              const SizedBox(height: 20),

              // Reset button
              SizedBox(
                width: double.infinity,
                child: ElevatedButton.icon(
                  onPressed: onReset,
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Colors.red[600],
                    foregroundColor: Colors.white,
                    padding: const EdgeInsets.symmetric(vertical: 16),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(12),
                    ),
                    elevation: 4,
                  ),
                  icon: const Icon(Icons.refresh, size: 20),
                  label: const Text(
                    'Start Over',
                    style: TextStyle(fontSize: 16, fontWeight: FontWeight.w600),
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  /// Build helpful suggestions section
  Widget _buildSuggestions() {
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: Colors.blue[50],
        borderRadius: BorderRadius.circular(12),
        border: Border.all(color: Colors.blue[200]!),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            children: [
              Icon(Icons.lightbulb_outline, color: Colors.blue[700], size: 20),
              const SizedBox(width: 8),
              Text(
                'Troubleshooting Tips:',
                style: TextStyle(
                  fontSize: 14,
                  fontWeight: FontWeight.w600,
                  color: Colors.blue[700],
                ),
              ),
            ],
          ),
          const SizedBox(height: 12),

          ..._getSuggestions().map((suggestion) {
            return Padding(
              padding: const EdgeInsets.symmetric(vertical: 4.0),
              child: Row(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Icon(
                    Icons.check_circle_outline,
                    color: Colors.blue[600],
                    size: 16,
                  ),
                  const SizedBox(width: 8),

                  Expanded(
                    child: Text(
                      suggestion,
                      style: TextStyle(
                        fontSize: 13,
                        color: Colors.blue[800],
                        height: 1.3,
                      ),
                    ),
                  ),
                ],
              ),
            );
          }),
        ],
      ),
    );
  }

  /// Get contextual suggestions based on error message
  List<String> _getSuggestions() {
    final errorLower = error.toLowerCase();

    if (errorLower.contains('network') || errorLower.contains('connection')) {
      return [
        'Check your internet connection',
        'Make sure the Recipe Quest server is running',
        'Verify the server URL (http://127.0.0.1:9090)',
        'Try again in a few moments',
      ];
    }

    if (errorLower.contains('recipe') || errorLower.contains('generation')) {
      return [
        'Try selecting different ingredients',
        'Check if the AI service is available',
        'Make sure you selected exactly 4 ingredients',
        'The server might be busy - try again',
      ];
    }

    if (errorLower.contains('image') || errorLower.contains('visualization')) {
      return [
        'Image generation can take some time',
        'Check if the image AI service is working',
        'Try with a different recipe combination',
        'Server might need a moment to process',
      ];
    }

    if (errorLower.contains('evaluation') || errorLower.contains('rating')) {
      return [
        'The evaluation AI might be busy',
        'Try with a complete recipe and image',
        'Check server connectivity',
        'Some evaluations take longer than others',
      ];
    }

    if (errorLower.contains('validation') || errorLower.contains('invalid')) {
      return [
        'Check that you selected valid ingredients',
        'Make sure your input meets the requirements',
        'Try selecting different combinations',
        'Some special characters might not be supported',
      ];
    }

    // Generic suggestions
    return [
      'Make sure the Recipe Quest server is running',
      'Check your internet connection',
      'Try refreshing the page',
      'Contact support if the problem persists',
    ];
  }
}
