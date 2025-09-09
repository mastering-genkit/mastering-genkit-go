// Response for cookingEvaluate flow
export interface EvaluateResponseDomain {
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
