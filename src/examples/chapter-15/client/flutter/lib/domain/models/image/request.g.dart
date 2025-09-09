// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'request.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

ImageRequest _$ImageRequestFromJson(Map<String, dynamic> json) => ImageRequest(
  dishName: json['dishName'] as String,
  description: json['description'] as String?,
);

Map<String, dynamic> _$ImageRequestToJson(ImageRequest instance) =>
    <String, dynamic>{
      'dishName': instance.dishName,
      'description': instance.description,
    };
