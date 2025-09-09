import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../domain/repositories/recipe_quest_repository.dart';
import '../repository/genkit_recipe_quest_repository.dart';

/// Provider for RecipeQuestRepository implementation using Dart client for Genkit
///
/// This provider creates a repository that uses defineRemoteAction internally
/// to communicate with Genkit flows deployed on the server.
final recipeQuestRepositoryProvider = Provider<RecipeQuestRepository>((ref) {
  return GenkitRecipeQuestRepository();
});
