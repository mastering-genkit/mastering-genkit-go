'use client';

import { useReducer, useCallback } from 'react';
import { GameStep, GameState, GameAction } from '../../src/domain/models/game';
import { 
  useGenerateRecipe,
  useCreateImage,
  useEvaluateDish 
} from './use-recipe-quest';

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

// Game state reducer
function gameReducer(state: GameState, action: GameAction): GameState {
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
      
    case 'SET_IMAGE':
      return {
        ...state,
        currentStep: GameStep.Image,
        progress: 60,
        imageUrl: action.payload,
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
      
    case 'START_EVALUATION':
      return {
        ...state,
        currentStep: GameStep.Evaluation,
        progress: 80,
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

/**
 * Simplified custom hook for managing Recipe Quest game state and flow
 */
export function useRecipeQuestGame() {
  const [state, dispatch] = useReducer(gameReducer, initialState);
  const generateRecipeHook = useGenerateRecipe();
  const createImageHook = useCreateImage();
  const evaluateHook = useEvaluateDish();

  // Evaluate dish function (最後に実行されるので最初に定義)
  const startDishEvaluation = useCallback(async (dishName: string, recipe: string) => {
    try {
      dispatch({ type: 'START_EVALUATION' });
      
      const response = await evaluateHook.evaluateDish({
        dishName,
        description: recipe,
      });

      if (response && response.success) {
        const score = response.score || Math.floor(Math.random() * 40) + 60;
        const feedback = response.feedback || 'Great dish! Well executed with creative use of ingredients.';
        const title = response.title || 'Chef';
        const achievement = response.achievement || 'Recipe Quest Complete';
        
        dispatch({ 
          type: 'SET_EVALUATION', 
          payload: { score, feedback, title, achievement } 
        });
      } else {
        dispatch({ type: 'SET_ERROR', payload: 'Failed to evaluate dish' });
      }
    } catch (error) {
      dispatch({ type: 'SET_ERROR', payload: 'Error evaluating dish' });
    }
  }, [evaluateHook]);

  // Generate image function (startDishEvaluationを呼ぶので2番目)
  const startImageGeneration = useCallback(async (recipe: string) => {
    try {
      dispatch({ type: 'SET_LOADING', payload: true });
      const dishName = recipe.split('\n')[0]?.replace(/Recipe name:\s*/i, '') || 'Delicious Dish';
      
      const response = await createImageHook.createImage({
        dishName,
        description: `A beautiful dish made with: ${state.selectedIngredients.join(', ')}`,
      });

      if (response && response.success && response.imageUrl) {
        dispatch({ type: 'SET_IMAGE', payload: response.imageUrl });
        
        // Auto advance to evaluation
        setTimeout(() => {
          startDishEvaluation(dishName, recipe);
        }, 1000);
      } else {
        dispatch({ type: 'SET_ERROR', payload: 'Failed to generate image' });
      }
    } catch (error) {
      dispatch({ type: 'SET_ERROR', payload: 'Error generating image' });
    }
  }, [state.selectedIngredients, createImageHook, startDishEvaluation]);

  // Start recipe generation with selected ingredients (startImageGenerationを呼ぶので3番目)
  const startRecipeGeneration = useCallback(async () => {
    if (state.selectedIngredients.length === 0) {
      dispatch({ type: 'SET_ERROR', payload: 'Please select at least one ingredient' });
      return;
    }

    try {
      dispatch({ type: 'START_RECIPE_GENERATION' });
      let fullContent = '';
      
      for await (const response of generateRecipeHook.generateRecipe({
        ingredients: state.selectedIngredients,
      })) {
        if (response.type === 'content' && response.content) {
          fullContent += response.content;
          dispatch({ type: 'SET_RECIPE', payload: fullContent });
        } else if (response.type === 'done') {
          // Auto advance to image generation
          setTimeout(() => {
            startImageGeneration(fullContent);
          }, 1000);
          break;
        } else if (response.type === 'error') {
          dispatch({ type: 'SET_ERROR', payload: response.error || 'Failed to generate recipe' });
          break;
        }
      }
    } catch (error) {
      dispatch({ type: 'SET_ERROR', payload: 'Error generating recipe' });
    }
  }, [state.selectedIngredients, generateRecipeHook, startImageGeneration]);

  // Start game - enter ingredient selection mode (最後に定義)
  const startGame = useCallback(() => {
    dispatch({ type: 'START_GAME' });
  }, []);

  // Add ingredient to selection
  const addIngredient = useCallback((ingredient: string) => {
    if (state.selectedIngredients.length < 4 && !state.selectedIngredients.includes(ingredient)) {
      dispatch({ type: 'ADD_INGREDIENT', payload: ingredient });
    }
  }, [state.selectedIngredients]);

  // Remove ingredient from selection
  const removeIngredient = useCallback((ingredient: string) => {
    dispatch({ type: 'REMOVE_INGREDIENT', payload: ingredient });
  }, []);

  // Reset game to start over
  const resetGame = useCallback(() => {
    dispatch({ type: 'RESET_GAME' });
  }, []);

  return {
    state,
    availableIngredients: AVAILABLE_INGREDIENTS,
    actions: {
      startGame,
      addIngredient,
      removeIngredient,
      startRecipeGeneration,
      resetGame,
    },
  };
}