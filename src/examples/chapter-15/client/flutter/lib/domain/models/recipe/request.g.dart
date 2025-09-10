// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'request.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

RecipeRequest _$RecipeRequestFromJson(Map<String, dynamic> json) =>
    RecipeRequest(
      ingredients: (json['ingredients'] as List<dynamic>)
          .map((e) => e as String)
          .toList(),
    );

Map<String, dynamic> _$RecipeRequestToJson(RecipeRequest instance) =>
    <String, dynamic>{'ingredients': instance.ingredients};
