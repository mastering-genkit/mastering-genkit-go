import 'package:genkit/client.dart';
import '../../domain/models/recipe/response.dart';
import '../../domain/models/image/response.dart';
import '../../domain/models/evaluate/response.dart';

/// Configuration for Recipe Quest Genkit server
class RecipeQuestConfig {
  /// Base URL for local Genkit development server
  static const String baseUrl = 'http://127.0.0.1:9090';

  /// Individual flow endpoints
  static const String generateRecipeUrl = '$baseUrl/generateRecipe';
  static const String createImageUrl = '$baseUrl/createImage';
  static const String evaluateDishUrl = '$baseUrl/evaluateDish';
}

/// Remote actions for Recipe Quest Genkit flows using Dart client for Genkit
class RecipeQuestActions {
  /// Generate recipe action (streaming)
  ///
  /// Uses Genkit streaming flow to receive recipe content in real-time. The
  /// returned action emits an [ActionStream] when invoked, which is both a
  /// [Stream] of intermediate chunks and a [Future] via `result` for the final
  /// response payload.
  static RemoteAction<RecipeResponse, RecipeResponse> get generateRecipe {
    return defineRemoteAction<RecipeResponse, RecipeResponse>(
      url: RecipeQuestConfig.generateRecipeUrl,
      fromResponse: (res) => RecipeResponse.fromJson(res),
      fromStreamChunk: (chunk) => RecipeResponse.fromJson(chunk),
    );
  }

  /// Create image action (unary)
  ///
  /// Uses Genkit unary flow to generate an image for a dish
  static RemoteAction<ImageResponse, void> get createImage {
    return defineRemoteAction<ImageResponse, void>(
      url: RecipeQuestConfig.createImageUrl,
      fromResponse: (res) => ImageResponse.fromJson(res),
    );
  }

  /// Evaluate dish action (unary)
  ///
  /// Uses Genkit unary flow to evaluate and score a prepared dish
  static RemoteAction<EvaluateResponse, void> get evaluateDish {
    return defineRemoteAction<EvaluateResponse, void>(
      url: RecipeQuestConfig.evaluateDishUrl,
      fromResponse: (res) => EvaluateResponse.fromJson(res),
    );
  }
}
