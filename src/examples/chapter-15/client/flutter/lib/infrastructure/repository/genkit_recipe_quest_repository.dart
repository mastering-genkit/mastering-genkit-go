import '../../domain/models/recipe/request.dart';
import '../../domain/models/recipe/response.dart';
import '../../domain/models/image/request.dart';
import '../../domain/models/image/response.dart';
import '../../domain/models/evaluate/request.dart';
import '../../domain/models/evaluate/response.dart';
import '../../domain/repositories/recipe_quest_repository.dart';
import '../actions/recipe_quest_actions.dart';

/// Repository implementation using Dart client for Genkit
///
/// This implementation uses defineRemoteAction to call Genkit flows
/// deployed on the server with proper type safety and streaming support.
class GenkitRecipeQuestRepository implements RecipeQuestRepository {
  @override
  Stream<RecipeResponse> generateRecipe(RecipeRequest request) {
    try {
      // Get the streaming remote action
      final action = RecipeQuestActions.generateRecipe;

      // Execute streaming flow
      final (:stream, :response) = action.stream(input: request);

      // Return the stream of recipe responses
      // Note: The final response from :response can be ignored for streaming use case
      // or could be used for completion notification if needed
      return stream.handleError((error) {
        // Convert any Genkit errors to RecipeResponse.error
        return RecipeResponse.error('Failed to generate recipe: $error');
      });
    } catch (e) {
      // Return error stream if action creation fails
      return Stream.fromIterable([
        RecipeResponse.error('Failed to initialize recipe generation: $e'),
      ]);
    }
  }

  @override
  Future<ImageResponse> createImage(ImageRequest request) async {
    try {
      // Get the unary remote action
      final action = RecipeQuestActions.createImage;

      // Execute unary flow
      final response = await action(input: request);

      return response;
    } catch (e) {
      // Convert any errors to ImageResponse.error
      return ImageResponse.error('Failed to create image: $e');
    }
  }

  @override
  Future<EvaluateResponse> evaluateDish(EvaluateRequest request) async {
    try {
      // Get the unary remote action
      final action = RecipeQuestActions.evaluateDish;

      // Execute unary flow
      final response = await action(input: request);

      return response;
    } catch (e) {
      // Convert any errors to EvaluateResponse.error
      return EvaluateResponse.error('Failed to evaluate dish: $e');
    }
  }
}
