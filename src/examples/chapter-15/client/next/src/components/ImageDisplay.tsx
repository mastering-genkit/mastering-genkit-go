'use client';

import { useState } from 'react';

interface ImageDisplayProps {
  imageUrl: string;
  dishName: string;
  isVisible: boolean;
}

export default function ImageDisplay({ imageUrl, dishName, isVisible }: ImageDisplayProps) {
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
    <div className="w-full max-w-4xl mx-auto mb-8">
      <div className="bg-white rounded-2xl shadow-xl p-8 border-l-4 border-purple-500">
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
        <div className="bg-gray-50 rounded-xl p-6 text-center">
          {!imageError ? (
            <div className="relative inline-block">
              {!imageLoaded && (
                <div className="w-full h-64 bg-gray-200 rounded-lg flex items-center justify-center">
                  <div className="text-center">
                    <div className="w-8 h-8 border-2 border-purple-400 border-t-transparent rounded-full animate-spin mx-auto mb-2" />
                    <p className="text-gray-500">Loading your dish image...</p>
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
