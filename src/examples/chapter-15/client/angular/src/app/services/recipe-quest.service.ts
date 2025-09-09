import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, Subject, EMPTY } from 'rxjs';
import { catchError, finalize } from 'rxjs/operators';
import {
  RecipeRequestDomain,
  RecipeResponseDomain,
  ImageRequestDomain,
  ImageResponseDomain,
  EvaluateRequestDomain,
  EvaluateResponseDomain,
  DomainError,
} from '../../domain/models';
import { CompositionService } from '../composition';

/**
 * Service for generating recipes (streaming)
 */
@Injectable({
  providedIn: 'root'
})
export class GenerateRecipeService {
  private isStreamingSubject = new BehaviorSubject<boolean>(false);
  private errorSubject = new BehaviorSubject<string | null>(null);
  private recipeStreamSubject = new Subject<RecipeResponseDomain>();
  private abortController: AbortController | null = null;

  public isStreaming$ = this.isStreamingSubject.asObservable();
  public error$ = this.errorSubject.asObservable();
  public recipeStream$ = this.recipeStreamSubject.asObservable();

  constructor(private compositionService: CompositionService) {}

  async generateRecipe(request: RecipeRequestDomain): Promise<void> {
    this.isStreamingSubject.next(true);
    this.errorSubject.next(null);

    // Create abort controller for cancellation
    this.abortController = new AbortController();

    try {
      const useCase = this.compositionService.getGenerateRecipeUseCase();
      
      for await (const response of useCase.execute(request)) {
        // Check if cancelled
        if (this.abortController?.signal.aborted) {
          break;
        }
        
        this.recipeStreamSubject.next(response);
      }
    } catch (err) {
      const errorMessage = err instanceof DomainError 
        ? err.message 
        : 'An unexpected error occurred';
      this.errorSubject.next(errorMessage);
      console.error('Recipe generation error:', err);
    } finally {
      this.isStreamingSubject.next(false);
      this.abortController = null;
    }
  }

  cancelGeneration(): void {
    if (this.abortController) {
      this.abortController.abort();
      this.abortController = null;
      this.isStreamingSubject.next(false);
    }
  }
}

/**
 * Service for creating images
 */
@Injectable({
  providedIn: 'root'
})
export class CreateImageService {
  private isLoadingSubject = new BehaviorSubject<boolean>(false);
  private errorSubject = new BehaviorSubject<string | null>(null);

  public isLoading$ = this.isLoadingSubject.asObservable();
  public error$ = this.errorSubject.asObservable();

  constructor(private compositionService: CompositionService) {}

  async createImage(request: ImageRequestDomain): Promise<ImageResponseDomain | null> {
    this.isLoadingSubject.next(true);
    this.errorSubject.next(null);

    try {
      const useCase = this.compositionService.getCreateImageUseCase();
      const response = await useCase.execute(request);
      return response;
    } catch (err) {
      const errorMessage = err instanceof DomainError 
        ? err.message 
        : 'An unexpected error occurred';
      this.errorSubject.next(errorMessage);
      console.error('Image creation error:', err);
      return null;
    } finally {
      this.isLoadingSubject.next(false);
    }
  }
}

/**
 * Service for evaluating dishes
 */
@Injectable({
  providedIn: 'root'
})
export class EvaluateDishService {
  private isLoadingSubject = new BehaviorSubject<boolean>(false);
  private errorSubject = new BehaviorSubject<string | null>(null);

  public isLoading$ = this.isLoadingSubject.asObservable();
  public error$ = this.errorSubject.asObservable();

  constructor(private compositionService: CompositionService) {}

  async evaluateDish(request: EvaluateRequestDomain): Promise<EvaluateResponseDomain | null> {
    this.isLoadingSubject.next(true);
    this.errorSubject.next(null);

    try {
      const useCase = this.compositionService.getEvaluateDishUseCase();
      const response = await useCase.execute(request);
      return response;
    } catch (err) {
      const errorMessage = err instanceof DomainError 
        ? err.message 
        : 'An unexpected error occurred';
      this.errorSubject.next(errorMessage);
      console.error('Dish evaluation error:', err);
      return null;
    } finally {
      this.isLoadingSubject.next(false);
    }
  }
}
