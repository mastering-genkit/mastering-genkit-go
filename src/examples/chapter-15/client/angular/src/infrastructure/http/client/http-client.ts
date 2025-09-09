import { NetworkError } from '../../../domain/models';

/**
 * Common HTTP client utilities
 */

export interface RequestOptions extends RequestInit {
  // Simplified for local development only
}

/**
 * Make an HTTP request
 */
export async function httpFetch(
  url: string,
  options: RequestOptions = {}
): Promise<Response> {
  const {
    headers: providedHeaders,
    ...fetchOptions
  } = options;

  try {
    // Prepare headers with basic Content-Type
    const headers: HeadersInit = {
      'Content-Type': 'application/json',
      ...providedHeaders,
    };

    // Make request using global fetch
    const response = await globalThis.fetch(url, {
      ...fetchOptions,
      headers,
    });

    return response;
  } catch (error) {
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
