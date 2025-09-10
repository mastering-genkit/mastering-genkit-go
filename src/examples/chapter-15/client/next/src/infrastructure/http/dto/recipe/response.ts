// DTO for createRecipe flow response (streaming)
export interface RecipeResponseDTO {
  type: string; // "content", "done", "error"
  content?: string;
  error?: string;
}
