package rag

import (
	"fmt"
	"strings"

	"github.com/firebase/genkit/go/ai"
	"github.com/ledongthuc/pdf"
)

// ParsePDFToChunks reads a PDF file and returns it as chunks of documents
func ParsePDFToChunks(filePath string, maxChunkSize int) ([]*ai.Document, error) {
	// Open the PDF file
	f, r, err := pdf.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open PDF: %w", err)
	}
	defer f.Close()

	var allText strings.Builder

	// Extract text from all pages
	for pageIndex := 1; pageIndex <= r.NumPage(); pageIndex++ {
		page := r.Page(pageIndex)
		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to extract text from page %d: %w", pageIndex, err)
		}

		allText.WriteString(text)
		allText.WriteString("\n\n")
	}

	// Split the text into chunks
	fullText := allText.String()
	chunks := chunkText(fullText, maxChunkSize)

	// Convert chunks to AI documents
	var documents []*ai.Document
	for i, chunk := range chunks {
		doc := ai.DocumentFromText(chunk, map[string]any{
			"source":     filePath,
			"chunk_id":   i,
			"chunk_size": len(chunk),
		})
		documents = append(documents, doc)
	}

	return documents, nil
}

// chunkText splits text into chunks of approximately maxChunkSize characters
// while trying to preserve sentence boundaries
func chunkText(text string, maxChunkSize int) []string {
	if len(text) <= maxChunkSize {
		return []string{text}
	}

	var chunks []string
	sentences := strings.Split(text, ". ")

	var currentChunk strings.Builder

	for _, sentence := range sentences {
		// Add the period back except for the last sentence
		if !strings.HasSuffix(sentence, ".") && !strings.HasSuffix(sentence, "!") && !strings.HasSuffix(sentence, "?") {
			sentence += "."
		}

		// Check if adding this sentence would exceed the chunk size
		if currentChunk.Len()+len(sentence)+1 > maxChunkSize && currentChunk.Len() > 0 {
			chunks = append(chunks, strings.TrimSpace(currentChunk.String()))
			currentChunk.Reset()
		}

		if currentChunk.Len() > 0 {
			currentChunk.WriteString(" ")
		}
		currentChunk.WriteString(sentence)
	}

	// Add the last chunk if it has content
	if currentChunk.Len() > 0 {
		chunks = append(chunks, strings.TrimSpace(currentChunk.String()))
	}

	return chunks
}
