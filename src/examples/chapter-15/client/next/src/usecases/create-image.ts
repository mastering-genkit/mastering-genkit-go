import {
  ImageRequestDomain,
  ImageResponseDomain,
  DomainError,
} from '../domain/models';
import { RecipeQuestRepository } from '../domain/repositories';

/**
 * Use case for creating dish images
 */
export class CreateImageUseCase {
  constructor(private readonly repository: RecipeQuestRepository) {}

  /**
   * Execute image creation
   */
  async execute(request: ImageRequestDomain): Promise<ImageResponseDomain> {
    // Validate request
    this.validateRequest(request);

    try {
      // Delegate to repository
      return await this.repository.createImage(request);
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

  private validateRequest(request: ImageRequestDomain): void {
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

    // Description is optional but if provided, must be a string
    if (request.description !== undefined && typeof request.description !== 'string') {
      throw new DomainError(
        'Description must be a string if provided',
        'VALIDATION_ERROR'
      );
    }
  }
}
