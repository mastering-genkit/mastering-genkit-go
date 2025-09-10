'use client';

import { useState } from 'react';

interface ImageDisplayProps {
  imageUrl: string;
  dishName: string;
  isVisible: boolean;
  isGenerating?: boolean;
  generationProgress?: number;
}

export default function ImageDisplay({ imageUrl, dishName, isVisible, isGenerating, generationProgress }: ImageDisplayProps) {
  const [imageLoaded, setImageLoaded] = useState(false);
  const [imageError, setImageError] = useState(false);

  if (!isVisible || !imageUrl) {
    return null;
  }

  const handleImageLoad = () => {
    setImageLoaded(true);
    setImageError(false);
  };

  const handleImageError = () => {
    setImageError(true);
    setImageLoaded(false);
  };

  return (
    <div className="w-full max-w-4xl mx-auto mb-8" data-image-section>
      <div className="bg-white rounded-2xl shadow-xl p-8 border-l-4 border-purple-500 relative overflow-hidden">
        {/* Header */}
        <div className="text-center mb-6">
          <div className="text-4xl mb-3">🎨</div>
          <h3 className="text-2xl font-bold text-gray-800 mb-2">
            Your Dish Visualization
          </h3>
          <h4 className="text-lg font-medium text-purple-600">
            {dishName}
          </h4>
        </div>

        {/* Image Container */}
        <div className="bg-gray-50 rounded-xl p-6 text-center relative">
          {/* Generation Progress Bar */}
          {isGenerating && generationProgress !== undefined && (
            <div className="absolute top-0 left-0 right-0 bg-white/95 backdrop-blur-sm p-4 border-b border-purple-200">
              <div className="max-w-md mx-auto">
                <div className="flex items-center justify-between mb-2">
                  <span className="text-sm font-medium text-gray-700">Generating Image...</span>
                  <span className="text-sm font-bold text-purple-600">{Math.round(generationProgress)}%</span>
                </div>
                <div className="w-full bg-gray-200 rounded-full h-2 overflow-hidden">
                  <div
                    className="h-full bg-gradient-to-r from-purple-400 to-pink-500 rounded-full transition-all duration-300 ease-out"
                    style={{ width: `${generationProgress}%` }}
                  >
                    <div className="h-full bg-white/30 animate-pulse" />
                  </div>
                </div>
                <div className="mt-2 text-xs text-gray-600 text-center">
                  🎨 AI artist is creating your dish visualization...
                </div>
              </div>
            </div>
          )}
          
          {!imageError ? (
            <div className="relative inline-block" style={{ marginTop: isGenerating ? '80px' : '0' }}>
              {!imageLoaded && !isGenerating && (
                <div className="w-full h-64 bg-gray-200 rounded-lg flex items-center justify-center">
                  <div className="text-center">
                    <div className="w-8 h-8 border-2 border-purple-400 border-t-transparent rounded-full animate-spin mx-auto mb-2" />
                    <p className="text-gray-500">Loading your dish image...</p>
                  </div>
                </div>
              )}
              
              {!imageLoaded && isGenerating && (
                <div className="w-full h-64 bg-gradient-to-br from-purple-50 to-pink-50 rounded-lg flex items-center justify-center border-2 border-dashed border-purple-300 animate-pulse">
                  <div className="text-center">
                    <div className="text-6xl mb-4 animate-bounce">🎨</div>
                    <p className="text-purple-700 font-medium">Creating your masterpiece...</p>
                    <p className="text-sm text-purple-600 mt-2">The AI artist is at work</p>
                  </div>
                </div>
              )}
              
              <img
                src={imageUrl}
                alt={dishName}
                className={`
                  max-w-full h-auto rounded-lg shadow-lg transition-opacity duration-500
                  ${imageLoaded ? 'opacity-100' : 'opacity-0 absolute inset-0'}
                `}
                onLoad={handleImageLoad}
                onError={handleImageError}
                style={{ maxHeight: '400px' }}
              />
              
              {imageLoaded && (
                <div className="absolute inset-0 bg-gradient-to-t from-black/20 to-transparent rounded-lg pointer-events-none" />
              )}
            </div>
          ) : (
            <div className="w-full h-64 bg-gray-100 rounded-lg flex items-center justify-center border-2 border-dashed border-gray-300">
              <div className="text-center">
                <div className="text-6xl mb-4">🖼️</div>
                <p className="text-gray-600 font-medium">Image Preview</p>
                <p className="text-sm text-gray-500 mt-2">
                  Imagine your beautiful {dishName} here!
                </p>
              </div>
            </div>
          )}
        </div>

        {/* Status Message */}
        <div className="text-center mt-6">
          {imageLoaded && !imageError ? (
            <p className="text-sm text-gray-500 italic">
              📸 Perfect! Now let&apos;s see what the AI judge thinks...
            </p>
          ) : imageError ? (
            <p className="text-sm text-gray-500 italic">
              ✨ Your dish looks amazing! Ready for evaluation...
            </p>
          ) : (
            <p className="text-sm text-gray-500 italic">
              🎨 Creating your visual masterpiece...
            </p>
          )}
        </div>
      </div>

    </div>
  );
}
