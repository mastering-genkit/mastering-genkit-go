package tools

import (
	"log"
	"os"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

// NewListDirectories creates a tool that lists directories in the current working directory.
func NewListDirectories(genkitClient *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(
		genkitClient,
		"listDirectories",
		"Lists all directories in the current working directory.",
		func(ctx *ai.ToolContext, input string) (string, error) {
			log.Printf("Tool 'listDirectories' called with input: %s", input)

			// Get current working directory
			pwd, err := os.Getwd()
			if err != nil {
				log.Printf("Error getting working directory: %v", err)
				return "Error: Could not get current directory", nil
			}

			// Read directory entries
			entries, err := os.ReadDir(pwd)
			if err != nil {
				log.Printf("Error reading directory: %v", err)
				return "Error: Could not read directory contents", nil
			}

			// Filter for directories only
			var directories []string
			for _, entry := range entries {
				if entry.IsDir() {
					directories = append(directories, entry.Name())
				}
			}

			if len(directories) == 0 {
				return "No directories found in current working directory", nil
			}

			result := "Directories found:\n"
			for _, dir := range directories {
				result += "- " + dir + "\n"
			}

			return result, nil
		})
}
