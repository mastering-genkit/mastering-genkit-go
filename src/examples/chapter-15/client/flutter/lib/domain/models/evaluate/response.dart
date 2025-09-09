import 'package:json_annotation/json_annotation.dart';

part 'response.g.dart';

/// Response for cookingEvaluate flow
@JsonSerializable()
class EvaluateResponse {
  const EvaluateResponse({
    required this.success,
    this.score,
    this.feedback,
    this.creativityScore,
    this.techniqueScore,
    this.appealScore,
    this.title,
    this.achievement,
    this.error,
  });

  /// Whether the evaluation was successful
  final bool success;

  /// Overall score of the dish
  final double? score;

  /// Feedback message for the dish
  final String? feedback;

  /// Creativity aspect score
  final double? creativityScore;

  /// Technique aspect score
  final double? techniqueScore;

  /// Appeal aspect score
  final double? appealScore;

  /// Title or name given to the evaluation
  final String? title;

  /// Achievement or badge earned
  final String? achievement;

  /// Error message (when success is false)
  final String? error;

  /// Factory constructor for creating a new `EvaluateResponse` instance from a map.
  factory EvaluateResponse.fromJson(Map<String, dynamic> json) =>
      _$EvaluateResponseFromJson(json);

  /// Converts this `EvaluateResponse` instance to a map.
  Map<String, dynamic> toJson() => _$EvaluateResponseToJson(this);

  /// Creates a successful evaluation response
  factory EvaluateResponse.success({
    required double score,
    required String feedback,
    double? creativityScore,
    double? techniqueScore,
    double? appealScore,
    String? title,
    String? achievement,
  }) => EvaluateResponse(
    success: true,
    score: score,
    feedback: feedback,
    creativityScore: creativityScore,
    techniqueScore: techniqueScore,
    appealScore: appealScore,
    title: title,
    achievement: achievement,
  );

  /// Creates an error response
  factory EvaluateResponse.error(String error) =>
      EvaluateResponse(success: false, error: error);

  @override
  String toString() =>
      'EvaluateResponse(success: $success, score: $score, feedback: $feedback, '
      'creativityScore: $creativityScore, techniqueScore: $techniqueScore, '
      'appealScore: $appealScore, title: $title, achievement: $achievement, error: $error)';

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other.runtimeType != runtimeType) return false;
    if (other is! EvaluateResponse) return false;
    return success == other.success &&
        score == other.score &&
        feedback == other.feedback &&
        creativityScore == other.creativityScore &&
        techniqueScore == other.techniqueScore &&
        appealScore == other.appealScore &&
        title == other.title &&
        achievement == other.achievement &&
        error == other.error;
  }

  @override
  int get hashCode => Object.hash(
    success,
    score,
    feedback,
    creativityScore,
    techniqueScore,
    appealScore,
    title,
    achievement,
    error,
  );
}
