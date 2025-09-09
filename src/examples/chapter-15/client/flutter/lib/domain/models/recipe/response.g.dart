// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'response.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

RecipeResponse _$RecipeResponseFromJson(Map<String, dynamic> json) =>
    RecipeResponse(
      type: $enumDecode(_$RecipeResponseTypeEnumMap, json['type']),
      content: json['content'] as String?,
      error: json['error'] as String?,
    );

Map<String, dynamic> _$RecipeResponseToJson(RecipeResponse instance) =>
    <String, dynamic>{
      'type': _$RecipeResponseTypeEnumMap[instance.type]!,
      'content': instance.content,
      'error': instance.error,
    };

const _$RecipeResponseTypeEnumMap = {
  RecipeResponseType.content: 'content',
  RecipeResponseType.done: 'done',
  RecipeResponseType.error: 'error',
};
