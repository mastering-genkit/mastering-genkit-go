// DTO for cookingEvaluate flow response
export interface EvaluateResponseDTO {
  success: boolean;
  score?: number;
  feedback?: string;
  creativityScore?: number;
  techniqueScore?: number;
  appealScore?: number;
  title?: string;
  achievement?: string;
  error?: string;
}
