import {
  EvaluateRequestDomain,
  EvaluateResponseDomain,
  DomainError,
} from '../domain/models';
import { RecipeQuestRepository } from '../domain/repositories';

/**
 * Use case for evaluating dishes
 */
export class EvaluateDishUseCase {
  constructor(private readonly repository: RecipeQuestRepository) {}

  /**
   * Execute dish evaluation
   */
  async execute(request: EvaluateRequestDomain): Promise<EvaluateResponseDomain> {
    // Validate request
    this.validateRequest(request);

    try {
      // Delegate to repository
      return await this.repository.evaluateDish(request);
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

  private validateRequest(request: EvaluateRequestDomain): void {
    if (!request.dishName || typeof request.dishName !== 'string') {
      throw new DomainError(
        'Dish name is required',
        'VALIDATION_ERROR'
      );
    }

    if (request.dishName.trim().length === 0) {
      throw new DomainError(
        'Dish name cannot be empty',
        'VALIDATION_ERROR'
      );
    }

    if (!request.description || typeof request.description !== 'string') {
      throw new DomainError(
        'Recipe description is required',
        'VALIDATION_ERROR'
      );
    }

    if (request.description.trim().length === 0) {
      throw new DomainError(
        'Recipe description cannot be empty',
        'VALIDATION_ERROR'
      );
    }
  }
}
