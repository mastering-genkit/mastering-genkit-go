import 'package:flutter/material.dart';

/// Widget for displaying generated images with support for both Data URI and HTTP URLs
class GeneratedImageView extends StatelessWidget {
  const GeneratedImageView({
    super.key,
    required this.imageUrl,
    this.width,
    this.height,
    this.fit = BoxFit.cover,
  });

  final String imageUrl;
  final double? width;
  final double? height;
  final BoxFit fit;

  @override
  Widget build(BuildContext context) {
    return Container(
      width: width,
      height: height,
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(16),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withValues(alpha: 0.1),
            blurRadius: 10,
            spreadRadius: 1,
            offset: const Offset(0, 4),
          ),
        ],
      ),
      child: ClipRRect(
        borderRadius: BorderRadius.circular(16),
        child: _buildImage(),
      ),
    );
  }

  /// Build appropriate image widget based on URL type
  Widget _buildImage() {
    if (_isDataUri(imageUrl)) {
      // Handle Data URI (base64 encoded images)
      return _buildDataUriImage();
    } else {
      // Handle HTTP/HTTPS URLs
      return _buildNetworkImage();
    }
  }

  /// Build image from Data URI
  Widget _buildDataUriImage() {
    try {
      final uri = Uri.parse(imageUrl);
      if (uri.data != null) {
        final bytes = uri.data!.contentAsBytes();
        return Image.memory(
          bytes,
          fit: fit,
          errorBuilder: (context, error, stackTrace) {
            return _buildErrorWidget('Failed to decode image data');
          },
        );
      } else {
        return _buildErrorWidget('Invalid image data');
      }
    } catch (e) {
      return _buildErrorWidget('Failed to parse image URI: $e');
    }
  }

  /// Build image from network URL
  Widget _buildNetworkImage() {
    return Image.network(
      imageUrl,
      fit: fit,
      loadingBuilder: (context, child, loadingProgress) {
        if (loadingProgress == null) return child;

        return Container(
          width: width,
          height: height ?? 200,
          decoration: BoxDecoration(
            color: Colors.grey[200],
            borderRadius: BorderRadius.circular(16),
          ),
          child: Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                CircularProgressIndicator(
                  value: loadingProgress.expectedTotalBytes != null
                      ? loadingProgress.cumulativeBytesLoaded /
                            loadingProgress.expectedTotalBytes!
                      : null,
                  valueColor: const AlwaysStoppedAnimation<Color>(
                    Colors.orange,
                  ),
                ),
                const SizedBox(height: 12),
                Text(
                  'Loading image...',
                  style: TextStyle(fontSize: 14, color: Colors.grey[600]),
                ),
              ],
            ),
          ),
        );
      },
      errorBuilder: (context, error, stackTrace) {
        return _buildErrorWidget('Failed to load image from URL');
      },
    );
  }

  /// Build error widget when image fails to load
  Widget _buildErrorWidget(String message) {
    return Container(
      width: width,
      height: height ?? 200,
      decoration: BoxDecoration(
        color: Colors.red.withValues(alpha: 0.1),
        borderRadius: BorderRadius.circular(16),
        border: Border.all(color: Colors.red.withValues(alpha: 0.3)),
      ),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          const Icon(Icons.error_outline, color: Colors.red, size: 48),
          const SizedBox(height: 12),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 16.0),
            child: Text(
              message,
              style: const TextStyle(
                color: Colors.red,
                fontSize: 14,
                fontWeight: FontWeight.w500,
              ),
              textAlign: TextAlign.center,
              maxLines: 2,
              overflow: TextOverflow.ellipsis,
            ),
          ),
        ],
      ),
    );
  }

  /// Check if the URL is a Data URI
  bool _isDataUri(String url) {
    return url.startsWith('data:');
  }
}
