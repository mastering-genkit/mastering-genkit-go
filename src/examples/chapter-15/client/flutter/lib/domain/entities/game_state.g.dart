// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'game_state.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

GameState _$GameStateFromJson(Map<String, dynamic> json) => GameState(
  currentStep: $enumDecode(_$GameStepEnumMap, json['currentStep']),
  progress: (json['progress'] as num).toDouble(),
  selectedIngredients: (json['selectedIngredients'] as List<dynamic>)
      .map((e) => e as String)
      .toList(),
  recipe: json['recipe'] as String?,
  imageUrl: json['imageUrl'] as String?,
  score: (json['score'] as num?)?.toDouble(),
  feedback: json['feedback'] as String?,
  title: json['title'] as String?,
  achievement: json['achievement'] as String?,
  isLoading: json['isLoading'] as bool,
  error: json['error'] as String?,
  isGeneratingImage: json['isGeneratingImage'] as bool?,
  imageGenerationProgress: (json['imageGenerationProgress'] as num?)
      ?.toDouble(),
);

Map<String, dynamic> _$GameStateToJson(GameState instance) => <String, dynamic>{
  'currentStep': _$GameStepEnumMap[instance.currentStep]!,
  'progress': instance.progress,
  'selectedIngredients': instance.selectedIngredients,
  'recipe': instance.recipe,
  'imageUrl': instance.imageUrl,
  'score': instance.score,
  'feedback': instance.feedback,
  'title': instance.title,
  'achievement': instance.achievement,
  'isLoading': instance.isLoading,
  'error': instance.error,
  'isGeneratingImage': instance.isGeneratingImage,
  'imageGenerationProgress': instance.imageGenerationProgress,
};

const _$GameStepEnumMap = {
  GameStep.ready: 'ready',
  GameStep.selectIngredients: 'select_ingredients',
  GameStep.recipe: 'recipe',
  GameStep.image: 'image',
  GameStep.evaluation: 'evaluation',
  GameStep.result: 'result',
};
