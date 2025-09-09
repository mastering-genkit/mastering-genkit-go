import 'package:json_annotation/json_annotation.dart';

part 'request.g.dart';

/// Request for cookingEvaluate flow
@JsonSerializable()
class EvaluateRequest {
  const EvaluateRequest({required this.dishName, required this.description});

  /// Name of the dish to evaluate
  final String dishName;

  /// Description of the prepared dish
  final String description;

  /// Factory constructor for creating a new `EvaluateRequest` instance from a map.
  factory EvaluateRequest.fromJson(Map<String, dynamic> json) =>
      _$EvaluateRequestFromJson(json);

  /// Converts this `EvaluateRequest` instance to a map.
  Map<String, dynamic> toJson() => _$EvaluateRequestToJson(this);

  @override
  String toString() =>
      'EvaluateRequest(dishName: $dishName, description: $description)';

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other.runtimeType != runtimeType) return false;
    if (other is! EvaluateRequest) return false;
    return dishName == other.dishName && description == other.description;
  }

  @override
  int get hashCode => Object.hash(dishName, description);
}
