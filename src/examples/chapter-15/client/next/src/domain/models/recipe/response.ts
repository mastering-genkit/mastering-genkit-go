// Response for createRecipe flow (streaming)
export interface RecipeResponseDomain {
  type: 'content' | 'done' | 'error';
  content?: string;
  error?: string;
}
