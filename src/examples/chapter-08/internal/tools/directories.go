package tools

import (
	"fmt"
	"log"
	"os"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

type ListDirectoryInput struct {
	Directory string `json:"directory" jsonschema_description:"Directory to list contents of"`
}

// NewListDirectories creates a tool that lists directories in the current working directory.
func NewListDirectories(genkitClient *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(
		genkitClient,
		"listDirectories",
		"Lists all directories in the current working directory.",
		func(ctx *ai.ToolContext, input ListDirectoryInput) ([]string, error) {
			log.Printf("Tool 'listDirectories' called with input: %s", input)

			// Read directory entries
			entries, err := os.ReadDir(input.Directory)
			if err != nil {
				log.Printf("Error reading directory: %v", err)
				return nil, fmt.Errorf("Error reading directory '%s': %v", input.Directory, err)
			}

			// Filter for directories only
			var directories []string
			for _, entry := range entries {
				if entry.IsDir() {
					directories = append(directories, entry.Name())
				}
			}

			if len(directories) == 0 {
				return nil, fmt.Errorf("No directories found in current working directory")
			}

			return directories, nil
		})
}

type CreateDirectoryInput struct {
	Directory string `json:"directory" jsonschema_description:"Directory to create"`
}

// NewCreateDirectory creates a tool that creates a new directory.
func NewCreateDirectory(genkitClient *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(
		genkitClient,
		"createDirectory",
		"Creates a new directory with the specified name. Input should be the directory name or path.",
		func(ctx *ai.ToolContext, input CreateDirectoryInput) (string, error) {
			log.Printf("Tool 'createDirectory' called with input: %s", input)

			if input.Directory == "" {
				return "Error: Directory name cannot be empty", nil
			}

			// Create the directory
			err := os.MkdirAll(input.Directory, 0755)
			if err != nil {
				log.Printf("Error creating directory '%s': %v", input.Directory, err)
				return "Error: Could not create directory '" + input.Directory + "': " + err.Error(), nil
			}

			log.Printf("Successfully created directory: %s", input.Directory)
			return "Successfully created directory: " + input.Directory, nil
		})
}
