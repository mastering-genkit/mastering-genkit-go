// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'response.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

ImageResponse _$ImageResponseFromJson(Map<String, dynamic> json) =>
    ImageResponse(
      success: json['success'] as bool,
      imageUrl: json['imageUrl'] as String?,
      dishName: json['dishName'] as String?,
      error: json['error'] as String?,
    );

Map<String, dynamic> _$ImageResponseToJson(ImageResponse instance) =>
    <String, dynamic>{
      'success': instance.success,
      'imageUrl': instance.imageUrl,
      'dishName': instance.dishName,
      'error': instance.error,
    };
