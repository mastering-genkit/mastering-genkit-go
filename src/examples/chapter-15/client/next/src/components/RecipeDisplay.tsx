'use client';

interface RecipeDisplayProps {
  recipe: string;
  isVisible: boolean;
}

export default function RecipeDisplay({ recipe, isVisible }: RecipeDisplayProps) {
  if (!isVisible || !recipe) {
    return null;
  }

  // Extract recipe name from the first line
  const lines = recipe.split('\n').filter(line => line.trim());
  const recipeName = lines[0]?.replace(/Recipe name:\s*/i, '') || 'Your Custom Recipe';
  const recipeContent = lines.slice(1).join('\n');

  return (
    <div className="w-full max-w-4xl mx-auto mb-8" data-recipe-section>
      <div className="bg-white rounded-2xl shadow-xl p-8 border-l-4 border-orange-500" data-recipe-display>
        {/* Header */}
        <div className="text-center mb-6">
          <div className="text-4xl mb-3">📋</div>
          <h3 className="text-2xl font-bold text-gray-800 mb-2">
            Your Custom Recipe
          </h3>
          <h4 className="text-xl font-semibold text-orange-600 mb-4">
            {recipeName}
          </h4>
        </div>

        {/* Recipe Content */}
        <div className="bg-gray-50 rounded-xl p-6">
          <pre className="whitespace-pre-wrap text-gray-700 leading-relaxed font-sans text-sm">
            {recipeContent}
          </pre>
        </div>

        {/* Status Message */}
        <div className="text-center mt-6">
          <p className="text-sm text-gray-500 italic">
            ✨ Recipe complete! Now creating a visual of your dish...
          </p>
        </div>
      </div>

    </div>
  );
}
