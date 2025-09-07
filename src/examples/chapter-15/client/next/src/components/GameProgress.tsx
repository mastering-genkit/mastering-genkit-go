'use client';

import { GameStep } from '../domain/models/game';

interface GameProgressProps {
  currentStep: GameStep;
  progress: number;
  isLoading: boolean;
}

// Step information with emojis and descriptions
const STEP_INFO = {
  [GameStep.Ready]: {
    emoji: '🎯',
    title: 'Ready to Start',
    description: 'Press the button to begin your cooking challenge!'
  },
  [GameStep.SelectIngredients]: {
    emoji: '🎲',
    title: 'Choose Your Ingredients',
    description: 'Select 4 ingredients for your recipe quest...'
  },
  [GameStep.Recipe]: {
    emoji: '🍳',
    title: 'Generating Recipe',
    description: 'AI chef is creating your custom recipe...'
  },
  [GameStep.Image]: {
    emoji: '📸',
    title: 'Creating Dish Image',
    description: 'Visualizing your delicious creation...'
  },
  [GameStep.Evaluation]: {
    emoji: '⭐',
    title: 'AI Judge Evaluation',
    description: 'Professional chef AI is rating your dish...'
  },
  [GameStep.Result]: {
    emoji: '🏆',
    title: 'Battle Complete',
    description: 'See how well you did in today\'s challenge!'
  }
};

export default function GameProgress({ currentStep, progress, isLoading }: GameProgressProps) {
  const stepInfo = STEP_INFO[currentStep];
  
  return (
    <div className="w-full max-w-2xl mx-auto mb-8">
      {/* Progress Bar */}
      <div className="mb-6">
        <div className="flex justify-between text-sm text-gray-600 mb-2">
          <span>Progress</span>
          <span>{Math.round(progress)}%</span>
        </div>
        <div className="w-full bg-gray-200 rounded-full h-3 overflow-hidden">
          <div 
            className="bg-gradient-to-r from-orange-400 to-red-500 h-full rounded-full transition-all duration-1000 ease-out"
            style={{ width: `${progress}%` }}
          />
        </div>
      </div>

      {/* Current Step Info */}
      <div className="text-center">
        <div className="text-6xl mb-4 animate-bounce">
          {stepInfo.emoji}
        </div>
        
        <h2 className="text-2xl font-bold text-gray-800 mb-2 flex items-center justify-center gap-3">
          {stepInfo.title}
          {isLoading && (
            <div className="w-6 h-6 border-2 border-orange-400 border-t-transparent rounded-full animate-spin" />
          )}
        </h2>
        
        <p className="text-gray-600 text-lg">
          {stepInfo.description}
        </p>
      </div>

      {/* Step Indicators */}
      <div className="flex justify-center items-center mt-8 gap-3">
        {Object.values(GameStep).map((step, index) => {
          const isActive = step === currentStep;
          const isCompleted = Object.values(GameStep).indexOf(currentStep) > index;
          
          return (
            <div key={step} className="flex items-center">
              <div
                className={`
                  w-4 h-4 rounded-full border-2 transition-all duration-300
                  ${isActive 
                    ? 'bg-orange-500 border-orange-500 scale-125' 
                    : isCompleted 
                      ? 'bg-green-500 border-green-500' 
                      : 'bg-gray-200 border-gray-300'
                  }
                `}
              />
              {index < Object.values(GameStep).length - 1 && (
                <div
                  className={`
                    w-8 h-0.5 mx-2 transition-colors duration-300
                    ${isCompleted ? 'bg-green-500' : 'bg-gray-300'}
                  `}
                />
              )}
            </div>
          );
        })}
      </div>
    </div>
  );
}
