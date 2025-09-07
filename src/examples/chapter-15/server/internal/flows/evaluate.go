package flows

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"

	"mastering-genkit-go/example/chapter-15/internal/structs/client"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NewCookingEvaluateFlow creates a recipe evaluation flow
func NewCookingEvaluateFlow(g *genkit.Genkit) *core.Flow[client.EvaluateRequest, client.EvaluateResponse, struct{}] {
	return genkit.DefineFlow(
		g,
		"cookingEvaluate",
		func(ctx context.Context, input client.EvaluateRequest) (client.EvaluateResponse, error) {
			log.Printf("cookingEvaluate flow called for dish: %s", input.DishName)

			systemPrompt := `You are a professional food critic and cooking competition judge.
Evaluate dishes based on recipe creativity, technique, and ingredient utilization.
Provide constructive feedback and fair scoring (1-100 scale).`

			prompt := fmt.Sprintf(`Evaluate this recipe for a cooking battle:
Dish Name: %s
Recipe Description: %s

Please provide:
1. Recipe creativity score (1-10)
2. Ingredient utilization score (1-10) 
3. Cooking technique score (1-10)
4. Overall appeal score (1-10)
5. Detailed feedback and suggestions for improvement

Rate based on the recipe description and cooking approach.`,
				input.DishName, input.Description)

			resp, err := genkit.Generate(ctx, g,
				ai.WithSystem(systemPrompt),
				ai.WithPrompt(prompt),
			)
			if err != nil {
				log.Printf("Failed to evaluate dish: %v", err)
				return client.EvaluateResponse{
					Success: false,
					Error:   fmt.Sprintf("Failed to evaluate dish: %v", err),
				}, err
			}

			// Parse scores from AI response
			evaluation := resp.Text()
			creativityScore := extractScore(evaluation, "creativity")
			techniqueScore := extractScore(evaluation, "technique")
			appealScore := extractScore(evaluation, "appeal")

			// Calculate overall score (average of individual scores * 10)
			overallScore := (creativityScore + techniqueScore + appealScore) * 10 / 3

			// Award chef title and achievement based on score
			title, achievement := awardChefTitle(overallScore, creativityScore, techniqueScore, appealScore)

			return client.EvaluateResponse{
				Success:         true,
				Score:           overallScore,
				Feedback:        evaluation,
				CreativityScore: creativityScore,
				TechniqueScore:  techniqueScore,
				AppealScore:     appealScore,
				Title:           title,
				Achievement:     achievement,
				Error:           "",
			}, nil
		},
	)
}

// awardChefTitle determines chef title and achievement based on performance
func awardChefTitle(overallScore, creativity, technique, appeal int) (string, string) {
	// Special achievements based on individual scores
	var specialAchievements []string

	if creativity >= 9 {
		specialAchievements = append(specialAchievements, "Innovation Master")
	}
	if technique >= 9 {
		specialAchievements = append(specialAchievements, "Technique Virtuoso")
	}
	if appeal >= 9 {
		specialAchievements = append(specialAchievements, "Flavor Magician")
	}
	if creativity >= 8 && technique >= 8 && appeal >= 8 {
		specialAchievements = append(specialAchievements, "Triple Crown Winner")
	}

	// Main title based on overall score
	var title string
	switch {
	case overallScore >= 90:
		title = "🏆 Legendary Quest Master"
	case overallScore >= 85:
		title = "⭐ Elite Recipe Explorer"
	case overallScore >= 80:
		title = "🌟 Master Quest Chef"
	case overallScore >= 75:
		title = "👨‍🍳 Skilled Adventurer"
	case overallScore >= 70:
		title = "🎯 Promising Quester"
	case overallScore >= 65:
		title = "🥄 Kitchen Explorer"
	case overallScore >= 60:
		title = "📚 Recipe Student"
	default:
		title = "🌱 Culinary Beginner"
	}

	// Create achievement description
	achievement := "Completed Recipe Quest"
	if len(specialAchievements) > 0 {
		achievement = fmt.Sprintf("🏅 %s - %s", achievement, specialAchievements[0])
	}

	return title, achievement
}

// extractScore tries to extract numerical scores from AI evaluation text
func extractScore(text string, category string) int {
	// Try to find patterns like "creativity score: 8" or "creativity: 8/10"
	patterns := []string{
		fmt.Sprintf(`(?i)%s[^:]*:\s*(\d+)`, category),
		fmt.Sprintf(`(?i)\d+\.\s*%s[^:]*:\s*(\d+)`, category),
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(text)
		if len(matches) > 1 {
			if score, err := strconv.Atoi(matches[1]); err == nil && score >= 1 && score <= 10 {
				return score
			}
		}
	}

	// Fallback to random score in reasonable range
	return rand.Intn(4) + 6 // 6-9
}
