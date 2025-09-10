import { RecipeQuestRepository } from '../src/domain/repositories';
import { HttpRecipeQuestRepository } from '../src/infrastructure/http';
import {
  GenerateRecipeUseCase,
  CreateImageUseCase,
  EvaluateDishUseCase,
} from '../src/usecases';

/**
 * Dependency Injection Composition Root
 * 
 * This module is responsible for wiring up all dependencies in the application.
 * It creates instances of infrastructure components and injects them into use cases.
 * Presentation layer components should only import from this module, not from infrastructure.
 */

// Singleton instances
let repository: RecipeQuestRepository | null = null;
let generateRecipeUseCase: GenerateRecipeUseCase | null = null;
let createImageUseCase: CreateImageUseCase | null = null;
let evaluateDishUseCase: EvaluateDishUseCase | null = null;

/**
 * Get or create the repository instance
 * In production, you might want to make this configurable for different environments
 */
function getRepository(): RecipeQuestRepository {
  if (!repository) {
    repository = new HttpRecipeQuestRepository();
  }
  return repository;
}

/**
 * Get or create the generate recipe use case instance
 */
export function getGenerateRecipeUseCase(): GenerateRecipeUseCase {
  if (!generateRecipeUseCase) {
    generateRecipeUseCase = new GenerateRecipeUseCase(getRepository());
  }
  return generateRecipeUseCase;
}

/**
 * Get or create the create image use case instance
 */
export function getCreateImageUseCase(): CreateImageUseCase {
  if (!createImageUseCase) {
    createImageUseCase = new CreateImageUseCase(getRepository());
  }
  return createImageUseCase;
}

/**
 * Get or create the evaluate dish use case instance
 */
export function getEvaluateDishUseCase(): EvaluateDishUseCase {
  if (!evaluateDishUseCase) {
    evaluateDishUseCase = new EvaluateDishUseCase(getRepository());
  }
  return evaluateDishUseCase;
}

/**
 * Reset all instances (useful for testing or switching environments)
 */
export function resetComposition(): void {
  repository = null;
  generateRecipeUseCase = null;
  createImageUseCase = null;
  evaluateDishUseCase = null;
}

/**
 * Configure the composition with a custom repository (useful for testing)
 * 
 * @param customRepository A custom repository implementation
 */
export function configureRepository(customRepository: RecipeQuestRepository): void {
  repository = customRepository;
  // Reset use cases to use the new repository
  generateRecipeUseCase = null;
  createImageUseCase = null;
  evaluateDishUseCase = null;
}
