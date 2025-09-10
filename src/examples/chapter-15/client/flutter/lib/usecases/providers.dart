import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../infrastructure/di/service_locator.dart';
import 'generate_recipe_usecase.dart';
import 'create_image_usecase.dart';
import 'evaluate_dish_usecase.dart';

/// Provider for GenerateRecipeUseCase with repository injection
final generateRecipeUseCaseProvider = Provider<GenerateRecipeUseCase>((ref) {
  final repository = ref.watch(recipeQuestRepositoryProvider);
  return GenerateRecipeUseCase(repository);
});

/// Provider for CreateImageUseCase with repository injection
final createImageUseCaseProvider = Provider<CreateImageUseCase>((ref) {
  final repository = ref.watch(recipeQuestRepositoryProvider);
  return CreateImageUseCase(repository);
});

/// Provider for EvaluateDishUseCase with repository injection
final evaluateDishUseCaseProvider = Provider<EvaluateDishUseCase>((ref) {
  final repository = ref.watch(recipeQuestRepositoryProvider);
  return EvaluateDishUseCase(repository);
});
