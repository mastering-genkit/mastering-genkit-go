import { NetworkError, AuthError } from '../../../domain/models';
import { getIdToken } from '../../auth';
import { httpConfig } from '../config';

/**
 * Common HTTP client utilities
 */

export interface RequestOptions extends RequestInit {
  includeAuth?: boolean;
}

/**
 * Create headers with authentication
 */
export async function createAuthHeaders(additionalHeaders?: HeadersInit): Promise<Headers> {
  const token = await getIdToken();
  const headers = new Headers(additionalHeaders);
  headers.set('Authorization', `Bearer ${token}`);
  return headers;
}

/**
 * Make an authenticated HTTP request
 */
export async function authenticatedFetch(
  url: string,
  options: RequestOptions = {}
): Promise<Response> {
  const {
    includeAuth = true,
    headers: providedHeaders,
    ...fetchOptions
  } = options;

  try {
    // Prepare headers
    let headers = new Headers(providedHeaders);
    if (includeAuth) {
      headers = await createAuthHeaders(headers);
    }

    // Make request
    const response = await fetch(url, {
      ...fetchOptions,
      headers,
    });

    return response;
  } catch (error) {

    // Re-throw auth errors
    if (error instanceof AuthError) {
      throw error;
    }

    // Wrap other errors
    throw new NetworkError(
      error instanceof Error ? error.message : 'Unknown error occurred'
    );
  }
}

/**
 * Check response status and throw appropriate errors
 */
export async function checkResponseStatus(response: Response): Promise<void> {
  if (response.ok) {
    return;
  }

  // Handle auth errors
  if (response.status === 401 || response.status === 403) {
    throw new AuthError(`Authentication failed: ${response.statusText}`);
  }

  // Try to parse error response
  let errorMessage = response.statusText;
  try {
    const errorData = await response.json();
    if (errorData.error) {
      errorMessage = errorData.error;
    }
  } catch {
    // Ignore parse error, use status text
  }

  throw new NetworkError(`HTTP error: ${errorMessage}`, response.status);
}

/**
 * Retry logic with exponential backoff
 */
export async function retryWithBackoff<T>(
  operation: () => Promise<T>,
  maxAttempts: number = 3,
  shouldRetry: (error: unknown) => boolean = (error) => error instanceof NetworkError
): Promise<T> {
  let attempt = 0;

  while (attempt < maxAttempts) {
    try {
      return await operation();
    } catch (error) {
      attempt++;

      // Check if we should retry
      if (!shouldRetry(error) || attempt >= maxAttempts) {
        throw error;
      }

      // Calculate delay with exponential backoff
      const delay = Math.min(
        1000 * Math.pow(2, attempt - 1), // 1s, 2s, 4s...
        10000 // max 10 seconds
      );

      console.warn(`Retrying operation (attempt ${attempt}/${maxAttempts}) after ${delay}ms`);
      await new Promise(resolve => setTimeout(resolve, delay));
    }
  }

  throw new NetworkError(`Failed after ${maxAttempts} attempts`);
}
