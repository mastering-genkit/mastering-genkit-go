package evals

import (
	"context"
	"fmt"
	"log"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

// NonBilledEvaluatorService handles evaluation operations
type NonBilledEvaluatorService struct {
	genkit *genkit.Genkit
}

// NewNonBilledEvaluatorService creates a new non-billed evaluator service
func NewNonBilledEvaluatorService(g *genkit.Genkit) *NonBilledEvaluatorService {
	return &NonBilledEvaluatorService{
		genkit: g,
	}
}

// NewResponseQualityEvaluator creates a response quality evaluator
func (es *NonBilledEvaluatorService) NewResponseQualityEvaluator() (ai.Evaluator, error) {
	return genkit.DefineEvaluator(es.genkit,
		"custom-non-billed", "response-quality",
		&ai.EvaluatorOptions{
			Definition:  "Evaluates the quality of AI responses based on length, relevance, and coherence.",
			IsBilled:    false,
			DisplayName: "Response Quality Evaluator",
		},
		func(ctx context.Context, req *ai.EvaluatorCallbackRequest) (*ai.EvaluatorCallbackResponse, error) {
			// Simple evaluation logic - check if response is meaningful
			score := 0.0
			reasoning := "Response is empty or invalid"

			// Extract input and output from the example
			input := req.Input
			inputText := ""
			outputText := ""

			// Convert input to string if it exists
			if input.Input != nil {
				if str, ok := input.Input.(string); ok {
					inputText = str
				}
			}

			// Convert output to string if it exists
			if input.Output != nil {
				if str, ok := input.Output.(string); ok {
					outputText = str
				}
			}

			// Basic evaluation criteria
			if len(outputText) > 10 {
				score += 0.3 // Has reasonable length
			}
			if outputText != inputText {
				score += 0.3 // Not just echoing input
			}
			if len(outputText) < 1000 {
				score += 0.2 // Not too verbose
			}
			if len(outputText) > 0 && (outputText[len(outputText)-1] == '.' || outputText[len(outputText)-1] == '!' || outputText[len(outputText)-1] == '?') {
				score += 0.2 // Ends with proper punctuation
			}

			reasoning = fmt.Sprintf("Response evaluated: length=%d, score=%.2f", len(outputText), score)

			return &ai.EvaluatorCallbackResponse{
				TestCaseId: input.TestCaseId,
				Evaluation: []ai.Score{
					{
						Id:     "response-quality",
						Score:  score,
						Status: "PASS",
						Details: map[string]any{
							"reasoning": reasoning,
							"length":    len(outputText),
						},
					},
				},
			}, nil
		})
}

// RunResponseQualityEvaluator runs an evaluator against a dataset
func (es *NonBilledEvaluatorService) RunResponseQualityEvaluator(evaluator ai.Evaluator, ctx context.Context, dataset []*ai.Example) (*ai.EvaluatorResponse, error) {
	response, err := evaluator.Evaluate(ctx, &ai.EvaluatorRequest{
		Dataset: dataset,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate: %w", err)
	}

	// Log evaluation results
	for _, evaluation := range *response {
		for _, score := range evaluation.Evaluation {
			log.Printf("TestCaseId: %s, Score: %v, Status: %s, Reasoning: %s",
				evaluation.TestCaseId, score.Score, score.Status, score.Details["reasoning"])
		}
	}

	return response, nil
}

// GetResponseQualityEvaluatorDataset returns a sample dataset for testing
func (es *NonBilledEvaluatorService) GetResponseQualityEvaluatorDataset() []*ai.Example {
	return []*ai.Example{
		{
			TestCaseId: "test-case-1",
			Input:      "What is the capital of France?",
			Output:     "The capital of France is Paris.",
		},
	}
}
