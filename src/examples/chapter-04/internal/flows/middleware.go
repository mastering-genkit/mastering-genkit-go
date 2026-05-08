package flows

import (
	"context"
	"log"

	"github.com/firebase/genkit/go/ai"
)

// CostTracker logs the dollar cost of every model API call by combining
// per-million-token rates with the token usage that the model provider
// reports on each response. It implements the Middleware V2 interface
// (Name + New) so it can also be registered with ai.DefineMiddleware to
// surface in the Dev UI if desired.
type CostTracker struct {
	InputUSDPer1M  float64 `json:"inputUsdPer1M,omitempty"`
	OutputUSDPer1M float64 `json:"outputUsdPer1M,omitempty"`
}

func (CostTracker) Name() string { return "cost-tracker" }

func (c CostTracker) New(ctx context.Context) (*ai.Hooks, error) {
	return &ai.Hooks{
		WrapModel: func(ctx context.Context, params *ai.ModelParams, next ai.ModelNext) (*ai.ModelResponse, error) {
			resp, err := next(ctx, params)
			if err != nil {
				return nil, err
			}
			if resp.Usage != nil {
				inputCost := float64(resp.Usage.InputTokens) / 1_000_000 * c.InputUSDPer1M
				outputCost := float64(resp.Usage.OutputTokens) / 1_000_000 * c.OutputUSDPer1M
				log.Printf("Cost: $%.6f (in: $%.6f, out: $%.6f, latency: %.0fms)",
					inputCost+outputCost, inputCost, outputCost, resp.LatencyMs)
			}
			return resp, nil
		},
	}, nil
}
