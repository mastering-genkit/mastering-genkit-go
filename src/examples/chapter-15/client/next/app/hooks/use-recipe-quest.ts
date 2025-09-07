'use client';

import { useState, useCallback, useRef } from 'react';
import {
  RecipeRequestDomain,
  RecipeResponseDomain,
  ImageRequestDomain,
  ImageResponseDomain,
  EvaluateRequestDomain,
  EvaluateResponseDomain,
  DomainError,
} from '../../src/domain/models';
import { 
  getGenerateRecipeUseCase, 
  getCreateImageUseCase,
  getEvaluateDishUseCase 
} from '../composition';

/**
 * Custom hook for generating recipes (streaming)
 */
export function useGenerateRecipe() {
  const [isStreaming, setIsStreaming] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const abortControllerRef = useRef<AbortController | null>(null);

  const generateRecipe = useCallback(async function* (
    request: RecipeRequestDomain
  ): AsyncGenerator<RecipeResponseDomain, void, unknown> {
    setIsStreaming(true);
    setError(null);

    // Create abort controller for cancellation
    abortControllerRef.current = new AbortController();

    try {
      const useCase = getGenerateRecipeUseCase();
      
      for await (const response of useCase.execute(request)) {
        // Check if cancelled
        if (abortControllerRef.current?.signal.aborted) {
          break;
        }
        
        yield response;
      }
    } catch (err) {
      const errorMessage = err instanceof DomainError 
        ? err.message 
        : 'An unexpected error occurred';
      setError(errorMessage);
      console.error('Recipe generation error:', err);
    } finally {
      setIsStreaming(false);
      abortControllerRef.current = null;
    }
  }, []);

  const cancelGeneration = useCallback(() => {
    if (abortControllerRef.current) {
      abortControllerRef.current.abort();
      abortControllerRef.current = null;
      setIsStreaming(false);
    }
  }, []);

  return {
    generateRecipe,
    cancelGeneration,
    isStreaming,
    error,
  };
}

/**
 * Custom hook for creating images
 */
export function useCreateImage() {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const createImage = useCallback(async (
    request: ImageRequestDomain
  ): Promise<ImageResponseDomain | null> => {
    setIsLoading(true);
    setError(null);

    try {
      const useCase = getCreateImageUseCase();
      const response = await useCase.execute(request);
      return response;
    } catch (err) {
      const errorMessage = err instanceof DomainError 
        ? err.message 
        : 'An unexpected error occurred';
      setError(errorMessage);
      console.error('Image creation error:', err);
      return null;
    } finally {
      setIsLoading(false);
    }
  }, []);

  return {
    createImage,
    isLoading,
    error,
  };
}

/**
 * Custom hook for evaluating dishes
 */
export function useEvaluateDish() {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const evaluateDish = useCallback(async (
    request: EvaluateRequestDomain
  ): Promise<EvaluateResponseDomain | null> => {
    setIsLoading(true);
    setError(null);

    try {
      const useCase = getEvaluateDishUseCase();
      const response = await useCase.execute(request);
      return response;
    } catch (err) {
      const errorMessage = err instanceof DomainError 
        ? err.message 
        : 'An unexpected error occurred';
      setError(errorMessage);
      console.error('Dish evaluation error:', err);
      return null;
    } finally {
      setIsLoading(false);
    }
  }, []);

  return {
    evaluateDish,
    isLoading,
    error,
  };
}

/**
 * Combined hook for all cooking battle functionality
 */
export function useCookingBattle() {
  const recipe = useGenerateRecipe();
  const image = useCreateImage();
  const evaluate = useEvaluateDish();

  return {
    recipe,
    image, 
    evaluate,
  };
}
