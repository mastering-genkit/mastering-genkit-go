import {
  RecipeRequestDomain,
  RecipeResponseDomain,
  DomainError,
} from '../domain/models';
import { RecipeQuestRepository } from '../domain/repositories';

/**
 * Use case for generating recipes using ingredients (streaming)
 */
export class GenerateRecipeUseCase {
  constructor(private readonly repository: RecipeQuestRepository) {}

  /**
   * Execute recipe generation (streaming)
   */
  async *execute(request: RecipeRequestDomain): AsyncGenerator<RecipeResponseDomain, void, unknown> {
    // Validate request
    this.validateRequest(request);

    try {
      // Delegate to repository for streaming execution
      for await (const response of this.repository.generateRecipe(request)) {
        yield response;
      }
    } catch (error) {
      // Re-throw domain errors as-is
      if (error instanceof DomainError) {
        throw error;
      }
      // Wrap unexpected errors
      throw new DomainError(
        error instanceof Error ? error.message : 'Unknown error occurred',
        'USECASE_ERROR'
      );
    }
  }

  private validateRequest(request: RecipeRequestDomain): void {
    if (!request.ingredients || !Array.isArray(request.ingredients)) {
      throw new DomainError(
        'Ingredients array is required',
        'VALIDATION_ERROR'
      );
    }

    if (request.ingredients.length === 0) {
      throw new DomainError(
        'At least one ingredient is required',
        'VALIDATION_ERROR'
      );
    }

    // Validate each ingredient is a non-empty string
    request.ingredients.forEach((ingredient, index) => {
      if (typeof ingredient !== 'string' || ingredient.trim().length === 0) {
        throw new DomainError(
          `Invalid ingredient at index ${index}: must be a non-empty string`,
          'VALIDATION_ERROR'
        );
      }
    });
  }
}
