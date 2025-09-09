import { Injectable } from '@angular/core';
import { RecipeQuestRepository } from '../../domain/repositories';
import { HttpRecipeQuestRepository } from '../../infrastructure/http';
import {
  GenerateRecipeUseCase,
  CreateImageUseCase,
  EvaluateDishUseCase,
} from '../../usecases';

/**
 * Dependency Injection Composition Root
 * 
 * This service is responsible for wiring up all dependencies in the application.
 * It creates instances of infrastructure components and injects them into use cases.
 * Presentation layer components should only import from this service, not from infrastructure.
 */
@Injectable({
  providedIn: 'root'
})
export class CompositionService {
  private repository: RecipeQuestRepository | null = null;
  private generateRecipeUseCase: GenerateRecipeUseCase | null = null;
  private createImageUseCase: CreateImageUseCase | null = null;
  private evaluateDishUseCase: EvaluateDishUseCase | null = null;

  /**
   * Get or create the repository instance
   * In production, you might want to make this configurable for different environments
   */
  private getRepository(): RecipeQuestRepository {
    if (!this.repository) {
      this.repository = new HttpRecipeQuestRepository();
    }
    return this.repository;
  }

  /**
   * Get or create the generate recipe use case instance
   */
  getGenerateRecipeUseCase(): GenerateRecipeUseCase {
    if (!this.generateRecipeUseCase) {
      this.generateRecipeUseCase = new GenerateRecipeUseCase(this.getRepository());
    }
    return this.generateRecipeUseCase;
  }

  /**
   * Get or create the create image use case instance
   */
  getCreateImageUseCase(): CreateImageUseCase {
    if (!this.createImageUseCase) {
      this.createImageUseCase = new CreateImageUseCase(this.getRepository());
    }
    return this.createImageUseCase;
  }

  /**
   * Get or create the evaluate dish use case instance
   */
  getEvaluateDishUseCase(): EvaluateDishUseCase {
    if (!this.evaluateDishUseCase) {
      this.evaluateDishUseCase = new EvaluateDishUseCase(this.getRepository());
    }
    return this.evaluateDishUseCase;
  }

  /**
   * Reset all instances (useful for testing or switching environments)
   */
  resetComposition(): void {
    this.repository = null;
    this.generateRecipeUseCase = null;
    this.createImageUseCase = null;
    this.evaluateDishUseCase = null;
  }

  /**
   * Configure the composition with a custom repository (useful for testing)
   * 
   * @param customRepository A custom repository implementation
   */
  configureRepository(customRepository: RecipeQuestRepository): void {
    this.repository = customRepository;
    // Reset use cases to use the new repository
    this.generateRecipeUseCase = null;
    this.createImageUseCase = null;
    this.evaluateDishUseCase = null;
  }
}
