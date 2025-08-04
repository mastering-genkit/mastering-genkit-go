package tools

import "github.com/firebase/genkit/go/ai"

func ConvertToolsToToolRefs(tools []ai.Tool) []ai.ToolRef {
	// Convert []ai.Tool to []ai.ToolRef
	toolRefs := make([]ai.ToolRef, len(tools))
	for i, tool := range tools {
		toolRefs[i] = tool
	}
	return toolRefs
}
