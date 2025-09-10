/**
 * HTTP client configuration
 */
export const httpConfig = {
  /**
   * Base URL for API endpoints
   * Defaults to local development server if not specified
   */
  baseUrl: (globalThis as any)?.env?.['ANGULAR_API_URL'] || 'http://127.0.0.1:9090',
  
  /**
   * API endpoints for individual flows
   */
  endpoints: {
    generateRecipe: '/generateRecipe',
    createImage: '/createImage',
    evaluateDish: '/evaluateDish',
  },
} as const;

/**
 * Get full URL for an endpoint
 */
export function getEndpointUrl(endpoint: keyof typeof httpConfig.endpoints): string {
  return `${httpConfig.baseUrl}${httpConfig.endpoints[endpoint]}`;
}
