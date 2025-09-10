import '../domain/models/recipe/request.dart';
import '../domain/models/recipe/response.dart';
import '../domain/models/error/domain_error.dart';
import '../domain/repositories/recipe_quest_repository.dart';

/// Use case for generating recipes using ingredients (streaming)
class GenerateRecipeUseCase {
  const GenerateRecipeUseCase(this._repository);

  final RecipeQuestRepository _repository;

  /// Execute recipe generation (streaming)
  ///
  /// Delegates to repository for streaming execution.
  /// Yields RecipeResponse chunks as they arrive from the server.
  Stream<RecipeResponse> execute(RecipeRequest request) async* {
    try {
      // Delegate to repository for streaming execution
      await for (final response in _repository.generateRecipe(request)) {
        yield response;
      }
    } on DomainError {
      // Re-throw domain errors as-is
      rethrow;
    } catch (error) {
      // Wrap unexpected errors in RecipeResponse.error
      yield RecipeResponse.error(
        'Recipe generation failed: ${error.toString()}',
      );
    }
  }
}
