package flows

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/localvec"
)

// RetrievalRequest represents the input for the retrieval flow
type RetrievalRequest struct {
	Query string `json:"query"`
	K     int    `json:"k,omitempty"` // number of results to return
}

// NewRetrievalFlow creates a flow that searches the indexed documents using the query.
func NewRetrievalFlow(g *genkit.Genkit, tools []ai.ToolRef, retriever ai.Retriever) *core.Flow[RetrievalRequest, string, struct{}] {
	return genkit.DefineFlow(g, "retrievalFlow", func(ctx context.Context, req RetrievalRequest) (string, error) {
		// Default K to 5 if not provided
		k := req.K
		if k == 0 {
			k = 5
		}

		// Create a document from the query text
		queryDoc := ai.DocumentFromText(req.Query, nil)

		// Retrieve similar documents
		retrieverOptions := &localvec.RetrieverOptions{
			K: k,
		}

		retrieverReq := &ai.RetrieverRequest{
			Query:   queryDoc,
			Options: retrieverOptions,
		}

		retrieverResp, err := retriever.Retrieve(ctx, retrieverReq)
		if err != nil {
			return "", fmt.Errorf("failed to retrieve documents: %w", err)
		}

		// Use the retrieved documents with Generate to provide expert Arduino assistance
		prompt := fmt.Sprintf(`You are an Arduino expert and analyst with deep knowledge of the Arduino ecosystem, open source hardware, and the Arduino community. 
You have access to the annual Arduino Open Source Report and can provide insights based on this comprehensive documentation.

Based on the provided Arduino Open Source Report documentation, please answer the following question: %s

Please provide a comprehensive answer that includes:
- Specific data and insights from the Arduino Open Source Report
- Trends and developments in the Arduino ecosystem
- Community statistics and growth metrics if mentioned
- Key findings and recommendations from the report
- Relevant comparisons or benchmarks discussed in the report

If the question is not directly answered in the report, please provide a reasoned analysis based on the available data. Do not make assumptions beyond the provided documentation.

Question: %s`, req.Query, req.Query)

		// Use the Genkit Generate function with the retrieved documents as context
		generateResp, err := genkit.Generate(ctx, g,
			ai.WithPrompt(prompt),
			ai.WithDocs(retrieverResp.Documents...))
		if err != nil {
			return "", fmt.Errorf("failed to generate response: %w", err)
		}

		return generateResp.Text(), nil
	})
}
