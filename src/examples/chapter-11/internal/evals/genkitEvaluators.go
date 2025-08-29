package evals

import (
	"context"
	"fmt"
	"log"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

// GenkitEvaluatorService handles built-in Genkit evaluator operations
type GenkitEvaluatorService struct {
	genkit *genkit.Genkit
}

// NewGenkitEvaluatorService creates a new Genkit evaluator service
func NewGenkitEvaluatorService(g *genkit.Genkit) *GenkitEvaluatorService {
	return &GenkitEvaluatorService{
		genkit: g,
	}
}

// GetDeepEqualEvaluator gets the built-in deep equal evaluator
func (ges *GenkitEvaluatorService) GetDeepEqualEvaluator() (ai.Evaluator, error) {
	genkitEvaluator := genkit.LookupEvaluator(ges.genkit, "genkitEval/deep_equal")
	if genkitEvaluator == nil {
		return nil, fmt.Errorf("deep_equal evaluator not found")
	}
	return genkitEvaluator, nil
}

// RunGenkitEvaluator runs a Genkit built-in evaluator against a dataset
func (ges *GenkitEvaluatorService) RunGenkitEvaluator(evaluator ai.Evaluator, ctx context.Context, dataset []*ai.Example) (*ai.EvaluatorResponse, error) {
	response, err := evaluator.Evaluate(ctx, &ai.EvaluatorRequest{
		Dataset: dataset,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate: %w", err)
	}

	// Log evaluation results
	for _, evaluation := range *response {
		for _, score := range evaluation.Evaluation {
			log.Printf("TestCaseId: %s, Score: %v, Status: %s",
				evaluation.TestCaseId, score.Score, score.Status)
		}
	}

	return response, nil
}

// GetSampleDataset returns a sample dataset for testing deep equal evaluator
func (ges *GenkitEvaluatorService) GetSampleDataset() []*ai.Example {
	return []*ai.Example{
		{
			TestCaseId: "test-case-1",
			Input:      "hola",
			Output:     "hola",
			Reference:  "hola",
		},
		{
			TestCaseId: "test-case-2",
			Input:      "What is the capital of France?",
			Output:     "Paris",
			Reference:  "Paris",
		},
		{
			TestCaseId: "test-case-3",
			Input:      "What is the capital of France?",
			Output:     "The capital of France is Paris.",
			Reference:  "Paris",
		},
	}
}
