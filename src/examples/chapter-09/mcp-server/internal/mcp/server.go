package mcp

import (
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/mcp"
)

// NewMCPServer creates a new MCP server that exposes tools via the Model Context Protocol.
// This is a convenience wrapper around mcp.NewMCPServer that provides
// a simplified interface for creating MCP servers in your applications.
//
// Parameters:
//   - g: The Genkit instance to use for server creation
//   - name: A unique identifier for this MCP server instance
//   - version: Version string for the server (e.g., "1.0.0")
//   - tools: Optional slice of tools to expose. If nil, all defined tools are auto-exposed
//
// Example usage:
//
//	// Option 1: Auto-expose all defined tools
//	server := NewMCPServer(g, "genkit-calculator", "1.0.0", nil)
//
//	// Option 2: Expose only specific tools
//	server := NewMCPServer(g, "genkit-calculator", "1.0.0", []ai.Tool{addTool, multiplyTool})
//
//	// Start the MCP server
//	log.Println("Starting MCP server...")
//	if err := server.ServeStdio(ctx); err != nil {
//	    log.Fatal(err)
//	}
//
// Returns:
//   - *mcp.GenkitMCPServer: The configured MCP server instance that can be used to serve tools
func NewMCPServer(g *genkit.Genkit, name string, version string, tools []ai.Tool) *mcp.GenkitMCPServer {
	// Set default version if empty
	if version == "" {
		version = "1.0.0"
	}

	// Create server options
	options := mcp.MCPServerOptions{
		Name:    name,
		Version: version,
	}

	// Add tools if provided
	if tools != nil {
		options.Tools = tools
	}

	// Create and return the MCP server
	return mcp.NewMCPServer(g, options)
}
