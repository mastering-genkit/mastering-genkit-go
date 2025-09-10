import 'package:json_annotation/json_annotation.dart';

part 'response.g.dart';

/// Response for createImage flow
@JsonSerializable()
class ImageResponse {
  const ImageResponse({
    required this.success,
    this.imageUrl,
    this.dishName,
    this.error,
  });

  /// Whether the image generation was successful
  final bool success;

  /// URL of the generated image (when success is true)
  final String? imageUrl;

  /// Name of the dish the image was generated for
  final String? dishName;

  /// Error message (when success is false)
  final String? error;

  /// Factory constructor for creating a new `ImageResponse` instance from a map.
  factory ImageResponse.fromJson(Map<String, dynamic> json) =>
      _$ImageResponseFromJson(json);

  /// Converts this `ImageResponse` instance to a map.
  Map<String, dynamic> toJson() => _$ImageResponseToJson(this);

  /// Creates a successful response
  factory ImageResponse.success({required String imageUrl, String? dishName}) =>
      ImageResponse(success: true, imageUrl: imageUrl, dishName: dishName);

  /// Creates an error response
  factory ImageResponse.error(String error) =>
      ImageResponse(success: false, error: error);

  @override
  String toString() =>
      'ImageResponse(success: $success, imageUrl: $imageUrl, dishName: $dishName, error: $error)';

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other.runtimeType != runtimeType) return false;
    if (other is! ImageResponse) return false;
    return success == other.success &&
        imageUrl == other.imageUrl &&
        dishName == other.dishName &&
        error == other.error;
  }

  @override
  int get hashCode => Object.hash(success, imageUrl, dishName, error);
}
