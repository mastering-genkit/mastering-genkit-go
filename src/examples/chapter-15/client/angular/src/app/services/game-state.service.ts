import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { GameStep, GameState, GameAction } from '../../domain/models/game';

// Available ingredients for user selection
const AVAILABLE_INGREDIENTS = [
  { name: 'chicken', emoji: '🐔' },
  { name: 'beef', emoji: '🥩' },
  { name: 'pork', emoji: '🐷' },
  { name: 'salmon', emoji: '🐟' },
  { name: 'shrimp', emoji: '🦐' },
  { name: 'tofu', emoji: '🥡' },
  { name: 'rice', emoji: '🍚' },
  { name: 'noodles', emoji: '🍜' },
  { name: 'pasta', emoji: '🍝' },
  { name: 'potatoes', emoji: '🥔' },
  { name: 'onions', emoji: '🧅' },
  { name: 'garlic', emoji: '🧄' },
  { name: 'ginger', emoji: '🫚' },
  { name: 'carrots', emoji: '🥕' },
  { name: 'peppers', emoji: '🌶️' },
  { name: 'vegetables', emoji: '🥬' },
  { name: 'mushrooms', emoji: '🍄' },
  { name: 'tomatoes', emoji: '🍅' },
  { name: 'lemon', emoji: '🍋' },
  { name: 'herbs', emoji: '🌿' },
  { name: 'sesame oil', emoji: '🫗' },
  { name: 'soy sauce', emoji: '🥢' },
  { name: 'miso', emoji: '🍲' },
  { name: 'quinoa', emoji: '🌾' },
  { name: 'avocado', emoji: '🥑' },
  { name: 'lime', emoji: '🍈' },
];

// Initial game state
const initialState: GameState = {
  currentStep: GameStep.Ready,
  progress: 0,
  selectedIngredients: [],
  isLoading: false,
};

/**
 * Service for managing game state
 */
@Injectable({
  providedIn: 'root'
})
export class GameStateService {
  private stateSubject = new BehaviorSubject<GameState>(initialState);
  
  public state$ = this.stateSubject.asObservable();
  public availableIngredients = AVAILABLE_INGREDIENTS;

  get currentState(): GameState {
    return this.stateSubject.value;
  }

  // Game state reducer
  private gameReducer(state: GameState, action: GameAction): GameState {
    switch (action.type) {
      case 'START_GAME':
        return {
          ...state,
          currentStep: GameStep.SelectIngredients,
          progress: 10,
          isLoading: false,
          error: undefined,
        };
        
      case 'ADD_INGREDIENT':
        const newIngredients = [...state.selectedIngredients, action.payload];
        return {
          ...state,
          selectedIngredients: newIngredients,
          progress: Math.min(10 + (newIngredients.length * 2.5), 20), // Progress as ingredients are added
        };
        
      case 'REMOVE_INGREDIENT':
        const filteredIngredients = state.selectedIngredients.filter(ing => ing !== action.payload);
        return {
          ...state,
          selectedIngredients: filteredIngredients,
          progress: Math.min(10 + (filteredIngredients.length * 2.5), 20),
        };
        
      case 'START_RECIPE_GENERATION':
        return {
          ...state,
          currentStep: GameStep.Recipe,
          progress: 30,
          isLoading: true,
        };
        
      case 'SET_RECIPE':
        return {
          ...state,
          currentStep: GameStep.Recipe,
          progress: 40,
          recipe: action.payload,
          isLoading: false,
        };
        
      case 'START_IMAGE_GENERATION':
        return {
          ...state,
          currentStep: GameStep.Image,
          progress: 50,
          isLoading: true,
          isGeneratingImage: true,
          imageGenerationProgress: 0,
        };
        
      case 'SET_IMAGE_PROGRESS':
        return {
          ...state,
          imageGenerationProgress: action.payload,
          progress: 50 + (action.payload * 0.1), // Image progress contributes 10% to total (50-60%)
        };
        
      case 'SET_IMAGE':
        return {
          ...state,
          currentStep: GameStep.Image,
          progress: 60,
          imageUrl: action.payload,
          isLoading: false,
          isGeneratingImage: false,
          imageGenerationProgress: 100,
        };
        
      case 'START_EVALUATION':
        return {
          ...state,
          currentStep: GameStep.Evaluation,
          progress: 80,
          isLoading: false,
        };
        
      case 'SET_EVALUATION':
        return {
          ...state,
          currentStep: GameStep.Result,
          progress: 100,
          score: action.payload.score,
          feedback: action.payload.feedback,
          title: action.payload.title,
          achievement: action.payload.achievement,
          isLoading: false,
        };
        
      case 'SET_LOADING':
        return {
          ...state,
          isLoading: action.payload,
        };
        
      case 'SET_ERROR':
        return {
          ...state,
          error: action.payload,
          isLoading: false,
        };
        
      case 'RESET_GAME':
        return initialState;
        
      default:
        return state;
    }
  }

  // Action dispatch methods
  startGame(): void {
    this.dispatch({ type: 'START_GAME' });
  }

  addIngredient(ingredient: string): void {
    const state = this.currentState;
    if (state.selectedIngredients.length < 4 && !state.selectedIngredients.includes(ingredient)) {
      this.dispatch({ type: 'ADD_INGREDIENT', payload: ingredient });
    }
  }

  removeIngredient(ingredient: string): void {
    this.dispatch({ type: 'REMOVE_INGREDIENT', payload: ingredient });
  }

  startRecipeGeneration(): void {
    this.dispatch({ type: 'START_RECIPE_GENERATION' });
  }

  setRecipe(recipe: string): void {
    this.dispatch({ type: 'SET_RECIPE', payload: recipe });
  }

  startImageGeneration(): void {
    this.dispatch({ type: 'START_IMAGE_GENERATION' });
  }

  setImageProgress(progress: number): void {
    this.dispatch({ type: 'SET_IMAGE_PROGRESS', payload: progress });
  }

  setImage(imageUrl: string): void {
    this.dispatch({ type: 'SET_IMAGE', payload: imageUrl });
  }

  startEvaluation(): void {
    this.dispatch({ type: 'START_EVALUATION' });
  }

  setEvaluation(evaluation: { score: number; feedback: string; title?: string; achievement?: string }): void {
    this.dispatch({ type: 'SET_EVALUATION', payload: evaluation });
  }

  setLoading(isLoading: boolean): void {
    this.dispatch({ type: 'SET_LOADING', payload: isLoading });
  }

  setError(error: string): void {
    this.dispatch({ type: 'SET_ERROR', payload: error });
  }

  resetGame(): void {
    this.dispatch({ type: 'RESET_GAME' });
  }

  private dispatch(action: GameAction): void {
    const currentState = this.currentState;
    const newState = this.gameReducer(currentState, action);
    this.stateSubject.next(newState);
  }
}
