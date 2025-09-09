import 'package:json_annotation/json_annotation.dart';

part 'response.g.dart';

/// Response for createRecipe flow (streaming)
@JsonSerializable()
class RecipeResponse {
  const RecipeResponse({required this.type, this.content, this.error});

  /// Type of the streaming response
  final RecipeResponseType type;

  /// Content of the recipe (when type is 'content')
  final String? content;

  /// Error message (when type is 'error')
  final String? error;

  /// Factory constructor for creating a new `RecipeResponse` instance from a map.
  factory RecipeResponse.fromJson(Map<String, dynamic> json) =>
      _$RecipeResponseFromJson(json);

  /// Converts this `RecipeResponse` instance to a map.
  Map<String, dynamic> toJson() => _$RecipeResponseToJson(this);

  /// Creates a content chunk response
  factory RecipeResponse.content(String content) =>
      RecipeResponse(type: RecipeResponseType.content, content: content);

  /// Creates a done response (end of stream)
  factory RecipeResponse.done() =>
      const RecipeResponse(type: RecipeResponseType.done);

  /// Creates an error response
  factory RecipeResponse.error(String error) =>
      RecipeResponse(type: RecipeResponseType.error, error: error);

  @override
  String toString() =>
      'RecipeResponse(type: $type, content: $content, error: $error)';

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other.runtimeType != runtimeType) return false;
    if (other is! RecipeResponse) return false;
    return type == other.type &&
        content == other.content &&
        error == other.error;
  }

  @override
  int get hashCode => Object.hash(type, content, error);
}

/// Types of streaming response for recipe generation
@JsonEnum()
enum RecipeResponseType {
  @JsonValue('content')
  content,
  @JsonValue('done')
  done,
  @JsonValue('error')
  error,
}
