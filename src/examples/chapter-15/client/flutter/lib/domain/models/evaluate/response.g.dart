// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'response.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

EvaluateResponse _$EvaluateResponseFromJson(Map<String, dynamic> json) =>
    EvaluateResponse(
      success: json['success'] as bool,
      score: (json['score'] as num?)?.toDouble(),
      feedback: json['feedback'] as String?,
      creativityScore: (json['creativityScore'] as num?)?.toDouble(),
      techniqueScore: (json['techniqueScore'] as num?)?.toDouble(),
      appealScore: (json['appealScore'] as num?)?.toDouble(),
      title: json['title'] as String?,
      achievement: json['achievement'] as String?,
      error: json['error'] as String?,
    );

Map<String, dynamic> _$EvaluateResponseToJson(EvaluateResponse instance) =>
    <String, dynamic>{
      'success': instance.success,
      'score': instance.score,
      'feedback': instance.feedback,
      'creativityScore': instance.creativityScore,
      'techniqueScore': instance.techniqueScore,
      'appealScore': instance.appealScore,
      'title': instance.title,
      'achievement': instance.achievement,
      'error': instance.error,
    };
