'use client';

import { 
  GameProgress,
  IngredientCards,
  GameResult,
  RecipeDisplay,
  ImageDisplay
} from '../../src/components';
import { useRecipeQuestGame } from '../hooks/use-recipe-quest-game';
import { GameStep } from '../../src/domain/models/game';

export default function RecipeQuestPage() {
  const { state, availableIngredients, actions } = useRecipeQuestGame();

  return (
    <div className="min-h-screen bg-gradient-to-br from-orange-50 to-red-50">
      {/* Header */}
      <header className="text-center py-8 px-4">
        <h1 className="text-5xl font-bold text-gray-800 mb-4">
          🍳 Recipe Quest
        </h1>
        <p className="text-xl text-gray-600 max-w-2xl mx-auto">
          Test your culinary creativity! Get random ingredients and let AI help you create, visualize, and rate your dish.
        </p>
      </header>

      <div className="container mx-auto px-4 pb-8">
        {/* Game Progress */}
        <GameProgress 
          currentStep={state.currentStep}
          progress={state.progress}
          isLoading={state.isLoading}
        />

        {/* Error State */}
        {state.error && (
          <div className="max-w-2xl mx-auto mb-8 p-4 bg-red-100 border-2 border-red-200 rounded-xl">
            <div className="text-center">
              <div className="text-2xl mb-2">😅</div>
              <p className="text-red-800 font-medium">Oops! Something went wrong</p>
              <p className="text-red-600 text-sm mt-1">{state.error}</p>
              <button
                onClick={actions.resetGame}
                className="mt-3 px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 transition-colors"
              >
                Try Again
              </button>
            </div>
          </div>
        )}

        {/* Start Game Button */}
        {state.currentStep === GameStep.Ready && (
          <div className="text-center max-w-lg mx-auto">
            <div className="bg-white rounded-2xl shadow-xl p-8 border-4 border-orange-200">
              <div className="text-6xl mb-6">🎯</div>
              <h2 className="text-2xl font-bold text-gray-800 mb-4">
                Ready for the Challenge?
              </h2>
              <p className="text-gray-600 mb-8 leading-relaxed">
                Choose your 4 favorite ingredients and start your recipe quest! 
                The AI will help you create a recipe, visualize your dish, and give you a professional rating.
              </p>
              <button
                onClick={actions.startGame}
                disabled={state.isLoading}
                className="
                  bg-gradient-to-r from-orange-500 to-red-500 hover:from-orange-600 hover:to-red-600
                  disabled:from-gray-400 disabled:to-gray-500
                  text-white font-bold py-4 px-8 rounded-xl text-xl
                  transform transition-all duration-200 hover:scale-105 hover:shadow-lg
                  disabled:cursor-not-allowed disabled:transform-none
                "
              >
                {state.isLoading ? (
                  <span className="flex items-center gap-3">
                    <div className="w-6 h-6 border-2 border-white border-t-transparent rounded-full animate-spin" />
                    Starting Challenge...
                  </span>
                ) : (
                  '🚀 Start Recipe Quest'
                )}
              </button>
            </div>
          </div>
        )}

        {/* Ingredient Selection */}
        {state.currentStep === GameStep.SelectIngredients && (
          <div className="max-w-6xl mx-auto">
            <div className="bg-white rounded-2xl shadow-xl p-8 mb-8">
              <div className="text-center mb-8">
                <div className="text-4xl mb-3">🛒</div>
                <h2 className="text-2xl font-bold text-gray-800 mb-2">
                  Choose Your Ingredients
                </h2>
                <p className="text-gray-600">
                  Select exactly 4 ingredients to create your recipe ({state.selectedIngredients.length}/4 selected)
                </p>
              </div>

              {/* Available Ingredients Grid */}
              <div className="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-3 mb-8">
                {availableIngredients.map((ingredient) => {
                  const isSelected = state.selectedIngredients.includes(ingredient.name);
                  const canSelect = state.selectedIngredients.length < 4 || isSelected;
                  
                  return (
                    <button
                      key={ingredient.name}
                      onClick={() => 
                        isSelected 
                          ? actions.removeIngredient(ingredient.name)
                          : actions.addIngredient(ingredient.name)
                      }
                      disabled={!canSelect}
                      className={`
                        p-4 rounded-xl border-2 transition-all duration-200 text-center
                        ${isSelected 
                          ? 'bg-orange-100 border-orange-500 text-orange-800 shadow-md' 
                          : canSelect
                            ? 'bg-gray-50 border-gray-200 hover:border-orange-300 hover:bg-orange-50 text-gray-700'
                            : 'bg-gray-100 border-gray-200 text-gray-400 cursor-not-allowed'
                        }
                        ${canSelect ? 'hover:scale-105' : ''}
                      `}
                    >
                      <div className="text-2xl mb-1">{ingredient.emoji}</div>
                      <div className="text-sm font-medium capitalize">
                        {ingredient.name}
                      </div>
                    </button>
                  );
                })}
              </div>

              {/* Selected Ingredients Display */}
              {state.selectedIngredients.length > 0 && (
                <div className="bg-gradient-to-r from-orange-50 to-red-50 rounded-xl p-6 mb-6">
                  <h3 className="text-lg font-semibold text-gray-800 mb-3 text-center">
                    Selected Ingredients
                  </h3>
                  <div className="flex flex-wrap gap-2 justify-center">
                    {state.selectedIngredients.map((ingredient) => {
                      const ingredientData = availableIngredients.find(ing => ing.name === ingredient);
                      return (
                        <span
                          key={ingredient}
                          className="inline-flex items-center gap-2 bg-white px-3 py-2 rounded-lg border border-orange-200 text-sm font-medium"
                        >
                          <span>{ingredientData?.emoji}</span>
                          <span className="capitalize">{ingredient}</span>
                          <button
                            onClick={() => actions.removeIngredient(ingredient)}
                            className="text-red-500 hover:text-red-700 ml-1"
                          >
                            ×
                          </button>
                        </span>
                      );
                    })}
                  </div>
                </div>
              )}

              {/* Start Recipe Generation Button */}
              <div className="text-center">
                <button
                  onClick={actions.startRecipeGeneration}
                  disabled={state.selectedIngredients.length < 1 || state.isLoading}
                  className="
                    bg-gradient-to-r from-green-500 to-blue-500 hover:from-green-600 hover:to-blue-600
                    disabled:from-gray-400 disabled:to-gray-500
                    text-white font-bold py-4 px-8 rounded-xl text-lg
                    transform transition-all duration-200 hover:scale-105 hover:shadow-lg
                    disabled:cursor-not-allowed disabled:transform-none
                  "
                >
                  {state.isLoading ? (
                    <span className="flex items-center gap-3">
                      <div className="w-5 h-5 border-2 border-white border-t-transparent rounded-full animate-spin" />
                      Creating Recipe...
                    </span>
                  ) : (
                    `🍳 Create Recipe ${state.selectedIngredients.length > 0 ? `(${state.selectedIngredients.length} ingredients)` : ''}`
                  )}
                </button>
              </div>
            </div>
          </div>
        )}

        {/* Content Area - Shows different content based on game step */}
        <div className="space-y-8">
          {/* Ingredients Display */}
          <IngredientCards
            ingredients={state.selectedIngredients}
            isVisible={state.currentStep === GameStep.SelectIngredients || 
                      (state.currentStep !== GameStep.Ready && state.currentStep !== GameStep.Result)}
          />

          {/* Recipe Display */}
          <RecipeDisplay
            recipe={state.recipe || ''}
            isVisible={state.currentStep === GameStep.Recipe || 
                      (!!state.recipe && state.currentStep !== GameStep.Ready && state.currentStep !== GameStep.Result)}
          />

          {/* Image Display */}
          <ImageDisplay
            imageUrl={state.imageUrl || ''}
            dishName={state.recipe?.split('\n')[0]?.replace(/Recipe name:\s*/i, '') || 'Your Dish'}
            isVisible={state.currentStep === GameStep.Image || 
                      (!!state.imageUrl && state.currentStep !== GameStep.Ready && state.currentStep !== GameStep.Result)}
          />

          {/* Game Result */}
          {state.currentStep === GameStep.Result && state.score !== undefined && (
            <GameResult
              score={state.score}
              feedback={state.feedback || 'Great job on your Recipe Quest!'}
              ingredients={state.selectedIngredients}
              onPlayAgain={actions.resetGame}
            />
          )}
        </div>

        {/* Footer */}
        <footer className="text-center mt-16 text-gray-500">
          <p className="text-sm">
            Powered by Genkit + Next.js 15 | Embark on your Recipe Quest!
          </p>
        </footer>
      </div>
    </div>
  );
}