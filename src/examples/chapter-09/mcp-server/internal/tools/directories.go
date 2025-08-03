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

// NewSystemInfo creates a tool that gathers system information based on complex criteria.
func NewSystemInfo(genkitClient *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(
		genkitClient,
		"systemInfo",
		"Gathers system information based on specified criteria and formatting options",
		func(ctx *ai.ToolContext, input struct {
			InfoTypes    []string `json:"info_types" jsonschema_description:"Types of info to gather: 'env', 'workdir', 'hostname', 'user'"`
			EnvVars      []string `json:"env_vars" jsonschema_description:"Specific environment variables to retrieve"`
			Format       string   `json:"format" jsonschema_description:"Output format: 'json', 'text', or 'summary'"`
			IncludeEmpty bool     `json:"include_empty" jsonschema_description:"Whether to include empty/unset values"`
			MaxLength    int      `json:"max_length" jsonschema_description:"Maximum length for each value (0 for no limit)"`
		}) (string, error) {
			log.Printf("Tool 'systemInfo' called with types: %v, format: %s", input.InfoTypes, input.Format)

			if len(input.InfoTypes) == 0 {
				return "Error: At least one info type must be specified", nil
			}

			var results []string

			for _, infoType := range input.InfoTypes {
				switch infoType {
				case "env":
					envInfo := "Environment Variables:\n"
					if len(input.EnvVars) > 0 {
						// Get specific environment variables
						for _, envVar := range input.EnvVars {
							value := os.Getenv(envVar)
							if !input.IncludeEmpty && value == "" {
								continue
							}
							if input.MaxLength > 0 && len(value) > input.MaxLength {
								value = value[:input.MaxLength] + "..."
							}
							envInfo += fmt.Sprintf("  %s=%s\n", envVar, value)
						}
					} else {
						// Get common environment variables
						commonVars := []string{"HOME", "USER", "PATH", "SHELL", "PWD"}
						for _, envVar := range commonVars {
							value := os.Getenv(envVar)
							if !input.IncludeEmpty && value == "" {
								continue
							}
							if input.MaxLength > 0 && len(value) > input.MaxLength {
								value = value[:input.MaxLength] + "..."
							}
							envInfo += fmt.Sprintf("  %s=%s\n", envVar, value)
						}
					}
					results = append(results, envInfo)

				case "workdir":
					workdir, err := os.Getwd()
					if err != nil {
						results = append(results, fmt.Sprintf("Working Directory: Error - %v\n", err))
					} else {
						if input.MaxLength > 0 && len(workdir) > input.MaxLength {
							workdir = workdir[:input.MaxLength] + "..."
						}
						results = append(results, fmt.Sprintf("Working Directory: %s\n", workdir))
					}

				case "hostname":
					hostname, err := os.Hostname()
					if err != nil {
						results = append(results, fmt.Sprintf("Hostname: Error - %v\n", err))
					} else {
						results = append(results, fmt.Sprintf("Hostname: %s\n", hostname))
					}

				case "user":
					user := os.Getenv("USER")
					if user == "" {
						user = os.Getenv("USERNAME") // Windows fallback
					}
					if !input.IncludeEmpty && user == "" {
						continue
					}
					results = append(results, fmt.Sprintf("Current User: %s\n", user))

				default:
					results = append(results, fmt.Sprintf("Unknown info type: %s\n", infoType))
				}
			}

			// Format output based on requested format
			output := ""
			switch input.Format {
			case "json":
				output = "{\n"
				for i, result := range results {
					output += fmt.Sprintf("  \"result_%d\": %q", i+1, result)
					if i < len(results)-1 {
						output += ","
					}
					output += "\n"
				}
				output += "}"
			case "summary":
				output = fmt.Sprintf("System Information Summary (%d items):\n", len(results))
				for _, result := range results {
					output += result
				}
			default: // "text" or any other value
				for _, result := range results {
					output += result
				}
			}

			return output, nil
		})
}
