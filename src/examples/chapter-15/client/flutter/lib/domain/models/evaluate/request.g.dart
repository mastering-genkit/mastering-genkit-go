// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'request.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

EvaluateRequest _$EvaluateRequestFromJson(Map<String, dynamic> json) =>
    EvaluateRequest(
      dishName: json['dishName'] as String,
      description: json['description'] as String,
    );

Map<String, dynamic> _$EvaluateRequestToJson(EvaluateRequest instance) =>
    <String, dynamic>{
      'dishName': instance.dishName,
      'description': instance.description,
    };
