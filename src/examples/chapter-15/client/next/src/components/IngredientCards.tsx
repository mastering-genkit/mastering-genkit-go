'use client';

interface IngredientCardsProps {
  ingredients: string[];
  isVisible: boolean;
}

// Ingredient emojis mapping
const INGREDIENT_EMOJIS: Record<string, string> = {
  // Proteins
  'chicken': '🐔',
  'beef': '🥩',
  'pork': '🐷',
  'salmon': '🐟',
  'shrimp': '🦐',
  'tofu': '🥡',
  'eggs': '🥚',
  'cheese': '🧀',
  
  // Grains & Starches
  'rice': '🍚',
  'pasta': '🍝',
  'noodles': '🍜',
  'potatoes': '🥔',
  'quinoa': '🌾',
  'bread': '🍞',
  
  // Vegetables
  'onions': '🧅',
  'garlic': '🧄',
  'carrots': '🥕',
  'peppers': '🌶️',
  'vegetables': '🥬',
  'avocado': '🥑',
  'tomatoes': '🍅',
  'mushrooms': '🍄',
  
  // Seasonings & Condiments
  'soy sauce': '🥢',
  'miso': '🍲',
  'herbs': '🌿',
  'lemon': '🍋',
  'lime': '🫐',
  'ginger': '🫚',
  'sesame oil': '🫗',
  'olive oil': '🫒',
  
  // Fallback
  'default': '🥘'
};

const getIngredientEmoji = (ingredient: string): string => {
  const key = ingredient.toLowerCase();
  return INGREDIENT_EMOJIS[key] || INGREDIENT_EMOJIS.default;
};

export default function IngredientCards({ ingredients, isVisible }: IngredientCardsProps) {
  if (!isVisible || ingredients.length === 0) {
    return null;
  }

  return (
    <div className="w-full max-w-4xl mx-auto mb-8">
      <div className="text-center mb-6">
        <h3 className="text-2xl font-bold text-gray-800 mb-2">
          🎲 Today&apos;s Challenge Ingredients
        </h3>
        <p className="text-gray-600">
          Your mission: Create something amazing with these ingredients!
        </p>
      </div>
      
      <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
        {ingredients.map((ingredient, index) => (
          <div
            key={ingredient}
            className={`
              bg-white rounded-xl shadow-lg p-6 text-center border-2 border-orange-100
              transform transition-all duration-500 hover:scale-105 hover:shadow-xl
            `}
            style={{
              animationDelay: `${index * 200}ms`
            }}
          >
            <div className="text-4xl mb-3">
              {getIngredientEmoji(ingredient)}
            </div>
            <h4 className="text-lg font-semibold text-gray-800 capitalize">
              {ingredient}
            </h4>
          </div>
        ))}
      </div>
      
      <div className="text-center mt-6">
        <p className="text-sm text-gray-500 italic">
          ✨ Ingredients selected! Watch as we create your custom recipe...
        </p>
      </div>
      
    </div>
  );
}
