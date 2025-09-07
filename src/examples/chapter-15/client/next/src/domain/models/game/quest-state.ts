/**
 * Quest step enumeration for Recipe Quest flow
 */
export enum GameStep {
  Ready = 'ready',
  SelectIngredients = 'select_ingredients',
  Recipe = 'recipe',
  Image = 'image',
  Evaluation = 'evaluation',
  Result = 'result'
}

/**
 * Quest state interface for Recipe Quest game
 */
export interface GameState {
  currentStep: GameStep;
  progress: number; // 0-100
  selectedIngredients: string[]; // User-selected ingredients
  recipe?: string;
  imageUrl?: string;
  score?: number;
  feedback?: string;
  title?: string;
  achievement?: string;
  isLoading: boolean;
  error?: string;
}

/**
 * Quest actions for state management
 */
export type GameAction = 
  | { type: 'START_GAME' }
  | { type: 'ADD_INGREDIENT'; payload: string }
  | { type: 'REMOVE_INGREDIENT'; payload: string }
  | { type: 'START_RECIPE_GENERATION' }
  | { type: 'SET_RECIPE'; payload: string }
  | { type: 'SET_IMAGE'; payload: string }
  | { type: 'START_EVALUATION' }
  | { type: 'SET_EVALUATION'; payload: { score: number; feedback: string; title?: string; achievement?: string } }
  | { type: 'SET_LOADING'; payload: boolean }
  | { type: 'SET_ERROR'; payload: string }
  | { type: 'RESET_GAME' };
