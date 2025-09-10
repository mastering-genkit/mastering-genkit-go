package tools

import (
	"fmt"
	"log"
	"strings"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"

	toolstructs "mastering-genkit-go/example/chapter-15/internal/structs/tools"
)

// NewEstimateCookingDifficulty creates a tool that estimates cooking difficulty
func NewEstimateCookingDifficulty(genkitClient *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(
		genkitClient,
		"estimateCookingDifficulty",
		"Estimate cooking difficulty based on ingredients and cooking methods",
		func(ctx *ai.ToolContext, input toolstructs.EstimateDifficultyInput) (*toolstructs.DifficultyEstimate, error) {
			log.Printf("Estimating difficulty for %d ingredients with %d steps", len(input.Ingredients), input.CookingSteps)

			// Calculate difficulty score based on various factors
			score := calculateDifficultyScore(input)

			// Determine difficulty level
			level := getDifficultyLevel(score)

			// Estimate cooking time
			timeEstimate := estimateCookingTime(input)

			// Determine required skills
			skills := determineRequiredSkills(input)

			// Determine required equipment
			equipment := determineRequiredEquipment(input)

			// Generate reasoning
			reasoning := generateReasoning(input, score, level)

			// Generate helpful tips
			tips := generateTips(level, input)

			return &toolstructs.DifficultyEstimate{
				Level:             level,
				Score:             score,
				Reasoning:         reasoning,
				TimeEstimate:      timeEstimate,
				SkillsRequired:    skills,
				EquipmentRequired: equipment,
				Tips:              tips,
			}, nil
		},
	)
}

// calculateDifficultyScore computes difficulty score based on multiple factors
func calculateDifficultyScore(input toolstructs.EstimateDifficultyInput) int {
	score := 3 // Base score

	// Factor 1: Ingredient count
	ingredientCount := len(input.Ingredients)
	if ingredientCount <= 3 {
		score += 1 // Simple
	} else if ingredientCount >= 6 {
		score += 3 // Complex
	} else {
		score += 2 // Medium
	}

	// Factor 2: Cooking steps
	if input.CookingSteps <= 3 {
		score += 1
	} else if input.CookingSteps >= 8 {
		score += 3
	} else {
		score += 2
	}

	// Factor 3: Cooking methods complexity
	for _, method := range input.CookingMethods {
		switch strings.ToLower(method) {
		case "boil", "steam", "bake":
			score += 1 // Simple methods
		case "fry", "sauté", "roast":
			score += 2 // Medium methods
		case "grill", "braise", "confit", "sous-vide":
			score += 3 // Complex methods
		}
	}

	// Cap the score between 1-10
	if score < 1 {
		score = 1
	} else if score > 10 {
		score = 10
	}

	return score
}

// getDifficultyLevel converts score to difficulty level
func getDifficultyLevel(score int) string {
	switch {
	case score <= 4:
		return "Easy"
	case score <= 7:
		return "Medium"
	default:
		return "Hard"
	}
}

// estimateCookingTime estimates total cooking time
func estimateCookingTime(input toolstructs.EstimateDifficultyInput) int {
	baseTime := 10 // Base 10 minutes

	// Add time based on ingredients
	baseTime += len(input.Ingredients) * 3

	// Add time based on cooking steps
	baseTime += input.CookingSteps * 5

	// Add time based on methods
	for _, method := range input.CookingMethods {
		switch strings.ToLower(method) {
		case "boil":
			baseTime += 10
		case "fry", "sauté":
			baseTime += 8
		case "bake", "roast":
			baseTime += 25
		case "grill":
			baseTime += 15
		case "steam":
			baseTime += 12
		default:
			baseTime += 10
		}
	}

	return baseTime
}

// determineRequiredSkills determines cooking skills needed
func determineRequiredSkills(input toolstructs.EstimateDifficultyInput) []string {
	skills := []string{"basic knife skills"} // Always needed

	for _, method := range input.CookingMethods {
		switch strings.ToLower(method) {
		case "fry", "sauté":
			skills = append(skills, "heat control")
		case "grill":
			skills = append(skills, "grill management")
		case "bake":
			skills = append(skills, "oven timing")
		case "braise":
			skills = append(skills, "liquid cooking")
		}
	}

	return removeDuplicates(skills)
}

// determineRequiredEquipment determines cooking equipment needed
func determineRequiredEquipment(input toolstructs.EstimateDifficultyInput) []string {
	equipment := []string{"cutting board", "knife"} // Always needed

	for _, method := range input.CookingMethods {
		switch strings.ToLower(method) {
		case "boil", "steam":
			equipment = append(equipment, "pot")
		case "fry", "sauté":
			equipment = append(equipment, "frying pan")
		case "bake", "roast":
			equipment = append(equipment, "oven")
		case "grill":
			equipment = append(equipment, "grill")
		}
	}

	return removeDuplicates(equipment)
}

// generateReasoning creates explanation for the difficulty rating
func generateReasoning(input toolstructs.EstimateDifficultyInput, score int, level string) string {
	reasons := []string{}

	reasons = append(reasons, fmt.Sprintf("Using %d ingredients", len(input.Ingredients)))
	reasons = append(reasons, fmt.Sprintf("%d cooking steps required", input.CookingSteps))

	if len(input.CookingMethods) > 0 {
		reasons = append(reasons, fmt.Sprintf("Involves %s", strings.Join(input.CookingMethods, ", ")))
	}

	return fmt.Sprintf("%s difficulty (score %d/10): %s", level, score, strings.Join(reasons, ", "))
}

// generateTips provides helpful tips based on difficulty level
func generateTips(level string, input toolstructs.EstimateDifficultyInput) string {
	switch level {
	case "Easy":
		return "Perfect for beginners! Take your time and enjoy the process. Pre-read the recipe once before starting."
	case "Medium":
		return "Good challenge for developing skills. Prep all ingredients before cooking and watch your timing."
	case "Hard":
		return "Advanced recipe requiring focus. Plan each step carefully, prep everything in advance, and don't rush."
	default:
		return "Follow the recipe step by step and you'll do great!"
	}
}

// removeDuplicates removes duplicate strings from slice
func removeDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	result := []string{}

	for _, item := range slice {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}

	return result
}
