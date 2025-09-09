/// Base class for all domain errors in Recipe Quest
abstract class DomainError extends Error {
  DomainError(this.message);

  final String message;

  @override
  String toString() => '$runtimeType: $message';
}

/// Error related to network operations
class NetworkError extends DomainError {
  NetworkError(super.message);
}

/// Error related to recipe generation
class RecipeGenerationError extends DomainError {
  RecipeGenerationError(super.message);
}

/// Error related to image generation
class ImageGenerationError extends DomainError {
  ImageGenerationError(super.message);
}

/// Error related to dish evaluation
class EvaluationError extends DomainError {
  EvaluationError(super.message);
}

/// Error related to invalid input or validation
class ValidationError extends DomainError {
  ValidationError(super.message);
}

/// Error related to parsing or serialization
class ParseError extends DomainError {
  ParseError(super.message);
}
