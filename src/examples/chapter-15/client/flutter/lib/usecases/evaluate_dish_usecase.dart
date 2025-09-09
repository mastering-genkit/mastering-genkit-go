import '../domain/models/evaluate/request.dart';
import '../domain/models/evaluate/response.dart';
import '../domain/models/error/domain_error.dart';
import '../domain/repositories/recipe_quest_repository.dart';

/// Use case for evaluating dishes
class EvaluateDishUseCase {
  const EvaluateDishUseCase(this._repository);

  final RecipeQuestRepository _repository;

  /// Execute dish evaluation
  ///
  /// Delegates to repository for dish evaluation.
  /// Returns EvaluateResponse with either success result or error.
  Future<EvaluateResponse> execute(EvaluateRequest request) async {
    try {
      // Delegate to repository
      return await _repository.evaluateDish(request);
    } on DomainError {
      // Re-throw domain errors as-is
      rethrow;
    } catch (error) {
      // Wrap unexpected errors in EvaluateResponse.error
      return EvaluateResponse.error(
        'Dish evaluation failed: ${error.toString()}',
      );
    }
  }
}
