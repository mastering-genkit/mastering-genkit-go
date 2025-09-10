import '../models/recipe/request.dart';
import '../models/recipe/response.dart';
import '../models/image/request.dart';
import '../models/image/response.dart';
import '../models/evaluate/request.dart';
import '../models/evaluate/response.dart';
import '../models/error/domain_error.dart';

/// Repository interface for Recipe Quest operations with individual flow methods.
/// This is a domain-level interface with no infrastructure dependencies.
/// Implementations will handle the actual network communication.
abstract class RecipeQuestRepository {
  /// Generate a recipe using ingredients (streaming).
  ///
  /// @param request The recipe generation request with ingredients
  /// @returns A stream for streaming recipe content
  /// @throws [NetworkError] If the request fails
  Stream<RecipeResponse> generateRecipe(RecipeRequest request);

  /// Create an image for a dish.
  ///
  /// @param request The image creation request with dish details
  /// @returns A future resolving to the image response
  /// @throws [NetworkError] If the request fails
  Future<ImageResponse> createImage(ImageRequest request);

  /// Evaluate a dish and provide scoring.
  ///
  /// @param request The evaluation request with dish details
  /// @returns A future resolving to the evaluation response
  /// @throws [NetworkError] If the request fails
  Future<EvaluateResponse> evaluateDish(EvaluateRequest request);
}
