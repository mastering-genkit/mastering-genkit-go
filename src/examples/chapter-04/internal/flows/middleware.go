package flows

import (
	"context"
	"log"
	"time"

	"github.com/firebase/genkit/go/ai"
)

// LoggingMiddleware logs execution time for all generation requests
var LoggingMiddleware = func(fn ai.ModelFunc) ai.ModelFunc {
	return func(ctx context.Context, req *ai.ModelRequest, cb ai.ModelStreamCallback) (*ai.ModelResponse, error) {
		start := time.Now()
		resp, err := fn(ctx, req, cb)
		if err != nil {
			log.Printf("Error after %v", time.Since(start))
			return nil, err
		}
		log.Printf("Success in %v", time.Since(start))
		return resp, nil
	}
}