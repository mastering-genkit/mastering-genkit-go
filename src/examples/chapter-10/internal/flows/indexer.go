package flows

import (
	"context"
	"fmt"
	"mastering-genkit-go/example/chapter-10/internal/rag"
	"path/filepath"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/localvec"
)

// IndexerRequest represents the input for the indexer flow
type IndexerRequest struct {
	PDFPath string `json:"pdfPath,omitempty"`
}

// NewIndexerFlow creates a flow that reads PDF documents, creates embeddings, and stores them in localvec.
func NewIndexerFlow(g *genkit.Genkit, tools []ai.ToolRef, docStore *localvec.DocStore) *core.Flow[IndexerRequest, string, struct{}] {
	return genkit.DefineFlow(g, "indexerFlow", func(ctx context.Context, req IndexerRequest) (string, error) {
		// Default PDF path if not provided
		pdfPath := req.PDFPath
		if pdfPath == "" {
			// Use absolute path to the Arduino report
			pdfPath = "internal/docs/arduino_report.pdf"
		}

		// Make path absolute
		absPath, err := filepath.Abs(pdfPath)
		if err != nil {
			return "", fmt.Errorf("failed to get absolute path: %w", err)
		}

		// Parse PDF into chunks
		chunks, err := rag.ParsePDFToChunks(absPath, 1000) // 1000 character chunks
		if err != nil {
			return "", fmt.Errorf("failed to parse PDF: %w", err)
		}

		// Index the documents
		err = localvec.Index(ctx, chunks, docStore)
		if err != nil {
			return "", fmt.Errorf("failed to index documents: %w", err)
		}

		return fmt.Sprintf("Successfully indexed %d chunks from %s", len(chunks), pdfPath), nil
	})
}
