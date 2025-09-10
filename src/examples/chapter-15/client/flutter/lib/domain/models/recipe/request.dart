import 'package:json_annotation/json_annotation.dart';

part 'request.g.dart';

/// Request for createRecipe flow
@JsonSerializable()
class RecipeRequest {
  const RecipeRequest({required this.ingredients});

  /// List of ingredients selected by the user
  final List<String> ingredients;

  /// Factory constructor for creating a new `RecipeRequest` instance from a map.
  factory RecipeRequest.fromJson(Map<String, dynamic> json) =>
      _$RecipeRequestFromJson(json);

  /// Converts this `RecipeRequest` instance to a map.
  Map<String, dynamic> toJson() => _$RecipeRequestToJson(this);

  @override
  String toString() => 'RecipeRequest(ingredients: $ingredients)';

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other.runtimeType != runtimeType) return false;
    if (other is! RecipeRequest) return false;
    return _listEquals(ingredients, other.ingredients);
  }

  @override
  int get hashCode => Object.hashAll(ingredients);
}

// Helper function for list equality
bool _listEquals<T>(List<T>? a, List<T>? b) {
  if (a == null) return b == null;
  if (b == null || a.length != b.length) return false;
  for (int index = 0; index < a.length; index += 1) {
    if (a[index] != b[index]) return false;
  }
  return true;
}
