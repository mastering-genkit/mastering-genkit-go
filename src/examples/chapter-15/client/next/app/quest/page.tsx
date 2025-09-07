'use client';

import { 
  GameProgress,
  GameResult,
  RecipeDisplay,
  ImageDisplay
} from '../../src/components';
import { useRecipeQuestGame } from '../hooks/use-recipe-quest-game';
import { GameStep } from '../../src/domain/models/game';

export default function RecipeQuestPage() {
  const { state, availableIngredients, actions } = useRecipeQuestGame();

  // Step info for fixed progress bar
  const stepInfo = {
    [GameStep.Ready]: { emoji: '🎯', title: 'Ready to Start' },
    [GameStep.SelectIngredients]: { emoji: '🎲', title: 'Choose Ingredients' },
    [GameStep.Recipe]: { emoji: '🍳', title: 'Creating Recipe' },
    [GameStep.Image]: { emoji: '📸', title: 'Generating Image' },
    [GameStep.Evaluation]: { emoji: '⚖️', title: 'Evaluating Dish' },
    [GameStep.Result]: { emoji: '🏆', title: 'Quest Complete' },
  }[state.currentStep] || { emoji: '🎯', title: 'Recipe Quest' };

  return (
    <div className="min-h-screen bg-gradient-to-br from-orange-50 to-red-50">
      {/* Fixed Progress Bar (only visible during processing) */}
      {(state.currentStep !== GameStep.Ready && state.currentStep !== GameStep.SelectIngredients) && (
        <div className="fixed top-0 left-0 right-0 bg-white/95 backdrop-blur-sm shadow-md z-50 border-b-2 border-orange-200">
          <div className="max-w-4xl mx-auto px-4 py-3">
            <div className="flex items-center justify-between">
              <div className="flex items-center gap-3">
                <div className="text-2xl">{stepInfo.emoji}</div>
                <div>
                  <div className="font-bold text-gray-800">{stepInfo.title}</div>
                  <div className="text-sm text-gray-600">
                    {state.selectedIngredients.map(ing => {
                      const data = availableIngredients.find(a => a.name === ing);
                      return data?.emoji;
                    }).join(' ')} • {state.selectedIngredients.join(', ')}
                  </div>
                </div>
              </div>
              <div className="flex items-center gap-3">
                {state.isLoading && (
                  <div className="w-5 h-5 border-2 border-orange-400 border-t-transparent rounded-full animate-spin" />
                )}
                <div className="text-right">
                  <div className="font-bold text-lg text-orange-600">{Math.round(state.progress)}%</div>
                  <div className="w-24 bg-gray-200 rounded-full h-2">
                    <div
                      className={`
                        h-full rounded-full transition-all duration-500
                        ${state.isLoading 
                          ? 'bg-gradient-to-r from-orange-400 via-red-500 to-orange-400 animate-pulse'
                          : 'bg-gradient-to-r from-orange-400 to-red-500'
                        }
                      `}
                      style={{ width: `${state.progress}%` }}
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      )}
      
      {/* Header */}
      <header className="text-center py-8 px-4">
        <h1 className="text-5xl font-bold text-gray-800 mb-4">
          🍳 Recipe Quest
        </h1>
        <p className="text-xl text-gray-600 max-w-2xl mx-auto">
          Test your culinary creativity! Get random ingredients and let AI help you create, visualize, and rate your dish.
        </p>
      </header>

      <div className={`container mx-auto px-4 pb-8 ${
        (state.currentStep !== GameStep.Ready && state.currentStep !== GameStep.SelectIngredients) ? 'pt-20' : ''
      }`}>
        {/* Game Progress */}
        <GameProgress 
          currentStep={state.currentStep}
          progress={state.progress}
          isLoading={state.isLoading}
        />

        {/* Processing Status Card */}
        {state.isLoading && state.currentStep !== GameStep.SelectIngredients && (
          <div className="max-w-2xl mx-auto mb-8 p-6 bg-blue-50 border-2 border-blue-200 rounded-xl shadow-lg">
            <div className="text-center">
              <div className="flex items-center justify-center gap-3 mb-4">
                <div className="w-8 h-8 border-3 border-blue-400 border-t-transparent rounded-full animate-spin" />
                <div className="text-xl font-bold text-blue-800">Processing...</div>
              </div>
              
              <div className="space-y-2">
                {state.currentStep === GameStep.Recipe && (
                  <>
                    <p className="text-blue-700 font-medium">🤖 AI Chef is analyzing your ingredients</p>
                    <p className="text-sm text-blue-600">✨ Creating a custom recipe with: {state.selectedIngredients.join(', ')}</p>
                    <p className="text-xs text-blue-500 mt-3">📡 Receiving live updates from the kitchen...</p>
                  </>
                )}
                
                {state.currentStep === GameStep.Image && (
                  <>
                    <p className="text-blue-700 font-medium">🎨 Visualizing your dish</p>
                    <p className="text-sm text-blue-600">📸 Creating an appetizing image of your recipe</p>
                    <p className="text-xs text-blue-500 mt-3">🖼️ This may take a moment...</p>
                  </>
                )}
                
                {state.currentStep === GameStep.Evaluation && (
                  <>
                    <p className="text-blue-700 font-medium">⚖️ Professional chef evaluation</p>
                    <p className="text-sm text-blue-600">👨‍🍳 Scoring your creativity, technique, and flavor</p>
                    <p className="text-xs text-blue-500 mt-3">🏆 Preparing your final results...</p>
                  </>
                )}
              </div>
            </div>
          </div>
        )}

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
                        p-4 rounded-xl border-3 transition-all duration-200 text-center min-h-[80px]
                        ${isSelected 
                          ? 'bg-gradient-to-br from-orange-400 to-red-500 border-orange-600 text-white shadow-lg transform scale-105 font-bold' 
                          : canSelect
                            ? 'bg-white border-gray-300 hover:border-orange-400 hover:bg-orange-50 text-gray-800 shadow hover:shadow-md'
                            : 'bg-gray-50 border-gray-200 text-gray-400 cursor-not-allowed'
                        }
                        ${canSelect && !isSelected ? 'hover:scale-102' : ''}
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
                <div className="bg-gradient-to-r from-orange-100 to-red-100 rounded-2xl p-6 mb-6 border-2 border-orange-300 shadow-lg">
                  <h3 className="text-xl font-bold text-gray-800 mb-4 text-center flex items-center justify-center gap-2">
                    <span>🍽️</span>
                    Selected Ingredients
                    <span className="bg-orange-500 text-white text-sm px-2 py-1 rounded-full">
                      {state.selectedIngredients.length}/4
                    </span>
                  </h3>
                  <div className="grid grid-cols-2 md:grid-cols-4 gap-3">
                    {state.selectedIngredients.map((ingredient) => {
                      const ingredientData = availableIngredients.find(ing => ing.name === ingredient);
                      return (
                        <div
                          key={ingredient}
                          className="bg-white rounded-xl p-4 border-2 border-orange-300 shadow-md hover:shadow-lg transition-all duration-200"
                        >
                          <div className="text-center">
                            <div className="text-3xl mb-2">{ingredientData?.emoji}</div>
                            <div className="text-sm font-semibold text-gray-800 capitalize mb-2">
                              {ingredient}
                            </div>
                            <button
                              onClick={() => actions.removeIngredient(ingredient)}
                              className="bg-red-500 hover:bg-red-600 text-white text-xs px-3 py-1 rounded-full transition-colors duration-200"
                            >
                              Remove
                            </button>
                          </div>
                        </div>
                      );
                    })}
                  </div>
                </div>
              )}

              {/* Start Recipe Generation Button */}
              <div className="text-center">
                {state.selectedIngredients.length === 4 && (
                  <div className="mb-4 p-3 bg-green-50 border border-green-300 rounded-lg">
                    <div className="text-green-800 font-semibold text-sm">
                      ✨ Perfect! You&apos;ve selected 4 ingredients. Ready to create your masterpiece?
                    </div>
                  </div>
                )}
                
                <button
                  onClick={actions.startRecipeGeneration}
                  disabled={state.selectedIngredients.length < 1 || state.isLoading}
                  className="
                    bg-gradient-to-r from-green-500 to-blue-500 hover:from-green-600 hover:to-blue-600
                    disabled:from-gray-400 disabled:to-gray-500
                    text-white font-bold py-4 px-8 rounded-xl text-lg
                    transform transition-all duration-200 hover:scale-105 hover:shadow-lg
                    disabled:cursor-not-allowed disabled:transform-none
                    min-w-[280px]
                  "
                >
                  {state.isLoading ? (
                    <span className="flex items-center justify-center gap-3">
                      <div className="w-5 h-5 border-2 border-white border-t-transparent rounded-full animate-spin" />
                      <div className="flex flex-col items-start text-sm">
                        <span>Creating Recipe...</span>
                        <span className="text-xs opacity-80">🤖 AI Chef is working...</span>
                      </div>
                    </span>
                  ) : (
                    <span className="flex items-center justify-center gap-2">
                      <span>🍳</span>
                      <span>Create Recipe</span>
                      {state.selectedIngredients.length > 0 && (
                        <span className="bg-white/20 px-2 py-1 rounded-full text-sm">
                          {state.selectedIngredients.length} ingredients
                        </span>
                      )}
                    </span>
                  )}
                </button>
                
                {state.selectedIngredients.length > 0 && state.selectedIngredients.length < 4 && (
                  <p className="mt-3 text-sm text-gray-600">
                    💡 Tip: Select {4 - state.selectedIngredients.length} more ingredient{4 - state.selectedIngredients.length > 1 ? 's' : ''} for the best recipe results!
                  </p>
                )}
              </div>
            </div>
          </div>
        )}

        {/* Content Area - Shows different content based on game step */}
        <div className="space-y-8">

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