import {
  RecipeRequestDomain,
  RecipeResponseDomain,
  ImageRequestDomain,
  ImageResponseDomain,
  EvaluateRequestDomain,
  EvaluateResponseDomain,
} from './models';

/**
 * Repository interface for Recipe Quest operations with individual flow methods.
 * This is a domain-level interface with no infrastructure dependencies.
 * Implementations will handle the actual network communication.
 */
export interface RecipeQuestRepository {
  /**
   * Generate a recipe using ingredients (streaming).
   * 
   * @param request The recipe generation request with ingredients
   * @returns An async generator for streaming recipe content
   * @throws {NetworkError} If the request fails
   */
  generateRecipe(request: RecipeRequestDomain): AsyncGenerator<RecipeResponseDomain, void, unknown>;

  /**
   * Create an image for a dish.
   * 
   * @param request The image creation request with dish details
   * @returns A promise resolving to the image response
   * @throws {NetworkError} If the request fails
   */
  createImage(request: ImageRequestDomain): Promise<ImageResponseDomain>;

  /**
   * Evaluate a dish and provide scoring.
   * 
   * @param request The evaluation request with dish details
   * @returns A promise resolving to the evaluation response
   * @throws {NetworkError} If the request fails
   */
  evaluateDish(request: EvaluateRequestDomain): Promise<EvaluateResponseDomain>;
}
