'use client';

interface GameResultProps {
  score: number;
  feedback: string;
  ingredients: string[];
  onPlayAgain: () => void;
}

// Score tier thresholds and styling
const getScoreTier = (score: number) => {
  if (score >= 90) return {
    tier: 'Legendary Chef',
    emoji: '👨‍🍳',
    color: 'text-yellow-500',
    bgColor: 'bg-gradient-to-r from-yellow-100 to-orange-100',
    borderColor: 'border-yellow-300'
  };
  if (score >= 80) return {
    tier: 'Master Cook',
    emoji: '⭐',
    color: 'text-purple-500',
    bgColor: 'bg-gradient-to-r from-purple-100 to-pink-100',
    borderColor: 'border-purple-300'
  };
  if (score >= 70) return {
    tier: 'Skilled Baker',
    emoji: '🌟',
    color: 'text-blue-500',
    bgColor: 'bg-gradient-to-r from-blue-100 to-cyan-100',
    borderColor: 'border-blue-300'
  };
  if (score >= 60) return {
    tier: 'Home Chef',
    emoji: '👍',
    color: 'text-green-500',
    bgColor: 'bg-gradient-to-r from-green-100 to-teal-100',
    borderColor: 'border-green-300'
  };
  return {
    tier: 'Beginner Cook',
    emoji: '💪',
    color: 'text-gray-500',
    bgColor: 'bg-gradient-to-r from-gray-100 to-slate-100',
    borderColor: 'border-gray-300'
  };
};

const getEncouragementMessage = (score: number): string => {
  if (score >= 90) return "Absolutely incredible! You're a culinary genius! 🔥";
  if (score >= 80) return "Outstanding performance! Your skills are impressive! 🎉";
  if (score >= 70) return "Well done! You've got some serious cooking talent! 👏";
  if (score >= 60) return "Great job! Keep practicing and you'll be amazing! 💫";
  return "Good effort! Every great chef started somewhere! 🚀";
};

export default function GameResult({ score, feedback, ingredients, onPlayAgain }: GameResultProps) {
  const scoreTier = getScoreTier(score);
  const encouragement = getEncouragementMessage(score);

  return (
    <div className="w-full max-w-4xl mx-auto">
      {/* Main Result Card */}
      <div className={`
        ${scoreTier.bgColor} rounded-2xl shadow-xl p-8 border-4 ${scoreTier.borderColor}
        transform
      `}>
        {/* Score Circle */}
        <div className="text-center mb-6">
          <div className="inline-flex items-center justify-center w-32 h-32 rounded-full bg-white shadow-lg mb-4">
            <div className="text-center">
              <div className="text-4xl font-bold text-gray-800">{score}</div>
              <div className="text-sm text-gray-600">/ 100</div>
            </div>
          </div>
          
          <div className="text-3xl mb-2">{scoreTier.emoji}</div>
          <h2 className={`text-2xl font-bold ${scoreTier.color} mb-1`}>
            {scoreTier.tier}
          </h2>
          <p className="text-lg text-gray-700 font-medium">
            {encouragement}
          </p>
        </div>

        {/* Challenge Summary */}
        <div className="bg-white/70 rounded-xl p-6 mb-6">
          <h3 className="text-lg font-semibold text-gray-800 mb-3 text-center">
            🎯 Challenge Summary
          </h3>
          <div className="flex flex-wrap justify-center gap-2 mb-4">
            {ingredients.map((ingredient) => (
              <span
                key={ingredient}
                className="px-3 py-1 bg-orange-100 text-orange-800 rounded-full text-sm font-medium capitalize"
              >
                {ingredient}
              </span>
            ))}
          </div>
        </div>

        {/* AI Feedback */}
        <div className="bg-white/70 rounded-xl p-6 mb-8">
          <h3 className="text-lg font-semibold text-gray-800 mb-3 text-center">
            👨‍⚖️ AI Judge Feedback
          </h3>
          <p className="text-gray-700 leading-relaxed text-center italic">
            &ldquo;{feedback}&rdquo;
          </p>
        </div>

        {/* Action Buttons */}
        <div className="text-center space-y-4">
          <button
            onClick={onPlayAgain}
            className="
              bg-gradient-to-r from-orange-500 to-red-500 hover:from-orange-600 hover:to-red-600
              text-white font-bold py-4 px-8 rounded-xl
              transform transition-all duration-200 hover:scale-105 hover:shadow-lg
              text-lg
            "
          >
            🎲 Try Another Challenge
          </button>
          
          <div className="text-sm text-gray-600">
            <p>Challenge yourself again with new ingredients!</p>
          </div>
        </div>
      </div>

      {/* Achievement Badges (future feature) */}
      <div className="mt-6 text-center">
        <div className="text-sm text-gray-500 italic">
          🏆 Collect achievements as you play more challenges!
        </div>
      </div>

    </div>
  );
}
