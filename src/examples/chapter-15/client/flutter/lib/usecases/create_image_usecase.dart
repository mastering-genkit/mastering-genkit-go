import '../domain/models/image/request.dart';
import '../domain/models/image/response.dart';
import '../domain/models/error/domain_error.dart';
import '../domain/repositories/recipe_quest_repository.dart';

/// Use case for creating dish images
class CreateImageUseCase {
  const CreateImageUseCase(this._repository);

  final RecipeQuestRepository _repository;

  /// Execute image creation
  ///
  /// Delegates to repository for image generation.
  /// Returns ImageResponse with either success result or error.
  Future<ImageResponse> execute(ImageRequest request) async {
    try {
      // Delegate to repository
      return await _repository.createImage(request);
    } on DomainError {
      // Re-throw domain errors as-is
      rethrow;
    } catch (error) {
      // Wrap unexpected errors in ImageResponse.error
      return ImageResponse.error('Image creation failed: ${error.toString()}');
    }
  }
}
