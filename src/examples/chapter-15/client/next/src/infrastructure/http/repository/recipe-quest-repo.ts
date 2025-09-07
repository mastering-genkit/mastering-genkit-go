import {
  RecipeRequestDomain,
  RecipeResponseDomain,
  ImageRequestDomain,
  ImageResponseDomain,
  EvaluateRequestDomain,
  EvaluateResponseDomain,
} from '../../../domain/models';
import { RecipeQuestRepository } from '../../../domain/repositories';
import { getEndpointUrl } from '../config';
import {
  mapRecipeRequestToDTO,
  mapRecipeResponseToDomain,
  mapImageRequestToDTO,
  mapImageResponseToDomain,
  mapEvaluateRequestToDTO,
  mapEvaluateResponseToDomain,
} from '../mappers';
import {
  authenticatedFetch,
  checkResponseStatus,
} from '../client/http-client';

/**
 * HTTP implementation of RecipeQuestRepository with individual flow methods
 */
export class HttpRecipeQuestRepository implements RecipeQuestRepository {
  /**
   * Generate a recipe using ingredients (streaming)
   */
  async *generateRecipe(request: RecipeRequestDomain): AsyncGenerator<RecipeResponseDomain, void, unknown> {
    const url = getEndpointUrl('generateRecipe');
    const dto = mapRecipeRequestToDTO(request);

    // Make authenticated request with streaming headers
    const response = await authenticatedFetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'text/event-stream',
      },
      body: JSON.stringify({ data: dto }), // Wrap in 'data' key for Genkit Flow
    });

    // Check response status
    await checkResponseStatus(response);

    // Handle streaming response
    if (!response.body) {
      throw new Error('No response body for streaming request');
    }

    const reader = response.body.getReader();
    const decoder = new TextDecoder();
    let buffer = '';

    try {
      while (true) {
        const { done, value } = await reader.read();
        
        if (done) break;
        
        buffer += decoder.decode(value, { stream: true });
        const lines = buffer.split('\n');
        
        // Keep the last incomplete line in the buffer
        buffer = lines.pop() || '';
        
        for (const line of lines) {
          if (line.trim() === '') continue;
          
          // Parse SSE format: "data: {...}"
          if (line.startsWith('data: ')) {
            try {
              const jsonData = line.slice(6); // Remove "data: " prefix
              const parsed = JSON.parse(jsonData);
              // Genkit wraps streaming content in a "message" field
              if (parsed.message) {
                yield mapRecipeResponseToDomain(parsed.message);
              } else if (parsed.result) {
                // Final result - ignore for streaming
                continue;
              }
            } catch (e) {
              console.warn('Failed to parse streaming chunk:', line);
            }
          }
        }
      }
    } finally {
      reader.releaseLock();
    }
  }

  /**
   * Create an image for a dish
   */
  async createImage(request: ImageRequestDomain): Promise<ImageResponseDomain> {
    const url = getEndpointUrl('createImage');
    const dto = mapImageRequestToDTO(request);

    // Make authenticated request
    const response = await authenticatedFetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
      body: JSON.stringify({ data: dto }), // Wrap in 'data' key for Genkit Flow
    });

    // Check response status
    await checkResponseStatus(response);

    // Parse and map response
    const responseData = await response.json();
    // Extract the result from Genkit Flow response wrapper
    const responseDto = responseData.result || responseData;
    return mapImageResponseToDomain(responseDto);
  }

  /**
   * Evaluate a dish and provide scoring
   */
  async evaluateDish(request: EvaluateRequestDomain): Promise<EvaluateResponseDomain> {
    const url = getEndpointUrl('evaluateDish');
    const dto = mapEvaluateRequestToDTO(request);

    // Make authenticated request
    const response = await authenticatedFetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
      body: JSON.stringify({ data: dto }), // Wrap in 'data' key for Genkit Flow
    });

    // Check response status
    await checkResponseStatus(response);

    // Parse and map response
    const responseData = await response.json();
    // Extract the result from Genkit Flow response wrapper
    const responseDto = responseData.result || responseData;
    return mapEvaluateResponseToDomain(responseDto);
  }
}