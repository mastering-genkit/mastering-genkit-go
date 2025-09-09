import 'package:json_annotation/json_annotation.dart';

part 'request.g.dart';

/// Request for createImage flow
@JsonSerializable()
class ImageRequest {
  const ImageRequest({required this.dishName, this.description});

  /// Name of the dish to generate image for
  final String dishName;

  /// Optional description of the dish
  final String? description;

  /// Factory constructor for creating a new `ImageRequest` instance from a map.
  factory ImageRequest.fromJson(Map<String, dynamic> json) =>
      _$ImageRequestFromJson(json);

  /// Converts this `ImageRequest` instance to a map.
  Map<String, dynamic> toJson() => _$ImageRequestToJson(this);

  @override
  String toString() =>
      'ImageRequest(dishName: $dishName, description: $description)';

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other.runtimeType != runtimeType) return false;
    if (other is! ImageRequest) return false;
    return dishName == other.dishName && description == other.description;
  }

  @override
  int get hashCode => Object.hash(dishName, description);
}
