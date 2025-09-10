'use client';

import { useEffect, useRef } from 'react';

interface RecipeDisplayProps {
  recipe: string;
  isVisible: boolean;
  isStreaming?: boolean;
}

export default function RecipeDisplay({ recipe, isVisible, isStreaming }: RecipeDisplayProps) {
  const contentRef = useRef<HTMLDivElement>(null);
  const lastScrollRef = useRef<number>(0);
  
  // Auto-scroll effect during streaming
  useEffect(() => {
    if (isStreaming && contentRef.current) {
      const element = contentRef.current;
      const currentTime = Date.now();
      
      // Throttle scrolling to every 100ms for smooth experience
      if (currentTime - lastScrollRef.current > 100) {
        element.scrollIntoView({ 
          behavior: 'smooth', 
          block: 'end',
          inline: 'nearest'
        });
        lastScrollRef.current = currentTime;
      }
    }
  }, [recipe, isStreaming]);
  if (!isVisible || !recipe) {
    return null;
  }

  // Extract recipe name from the first line
  const lines = recipe.split('\n').filter(line => line.trim());
  const recipeName = lines[0]?.replace(/Recipe name:\s*/i, '') || 'Your Custom Recipe';
  const recipeContent = lines.slice(1).join('\n');

  return (
    <div className="w-full max-w-4xl mx-auto mb-8" data-recipe-section>
      <div className="bg-white rounded-2xl shadow-xl p-8 border-l-4 border-orange-500 relative" data-recipe-display>
        {/* Streaming indicator */}
        {isStreaming && (
          <div className="absolute top-4 right-4 flex items-center gap-2 bg-orange-100 px-3 py-1 rounded-full">
            <div className="w-2 h-2 bg-orange-500 rounded-full animate-pulse" />
            <span className="text-xs font-medium text-orange-700">Streaming...</span>
          </div>
        )}
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
        <div className="bg-gray-50 rounded-xl p-6 max-h-[500px] overflow-y-auto custom-scrollbar">
          <pre className="whitespace-pre-wrap text-gray-700 leading-relaxed font-sans text-sm">
            {recipeContent}
          </pre>
          {/* Scroll anchor */}
          <div ref={contentRef} className="h-0" />
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
