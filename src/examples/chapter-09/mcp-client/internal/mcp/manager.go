package mcp

import (
	"fmt"
	"log"

	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/mcp"
)

// NewMCPHostWrapper creates a new MCP host with the provided configuration parameters.
// This is a convenience wrapper around mcp.NewMCPHost that accepts individual parameters
// instead of requiring a full MCPHostOptions struct.
//
// Parameters:
//   - name: A string identifier for your MCP host instance
//   - version: Version number for this host (defaults to "1.0.0" if empty)
//   - mcpServers: An array of MCPServerConfig configurations
//
// Example usage:
//
//	servers := []mcp.MCPServerConfig{
//	    {
//	        Name: "everything-server",
//	        Config: mcp.MCPClientOptions{
//	            Name: "everything-server",
//	            Stdio: &mcp.StdioConfig{
//	                Command: "npx",
//	                Args:    []string{"-y", "@modelcontextprotocol/server-everything"},
//	            },
//	        },
//	    },
//	    {
//	        Name: "mcp-server-time",
//	        Config: mcp.MCPClientOptions{
//	            Name: "mcp-server-time",
//	            Stdio: &mcp.StdioConfig{
//	                Command: "uvx",
//	                Args:    []string{"mcp-server-time"},
//	            },
//	        },
//	    },
//	}
//
//	manager, err := NewMCPHostWrapper(g, "my-app", "1.0.0", servers)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// The returned manager can be used to:
// - Get tools from all connected servers: manager.GetActiveTools(ctx, g)
// - Get prompts from specific servers: manager.GetPrompt(ctx, g, serverName, promptName, args)
// - Connect to additional servers: manager.Connect(name, options)
// - Disconnect from servers: manager.Disconnect(name)
func NewMCPHostWrapper(g *genkit.Genkit, name string, version string, mcpServers []mcp.MCPServerConfig) (*mcp.MCPHost, error) {
	// Set default version if empty
	if version == "" {
		version = "1.0.0"
	}

	// Create MCPManagerOptions from the parameters
	options := mcp.MCPHostOptions{
		Name:       name,
		Version:    version,
		MCPServers: mcpServers,
	}

	// Create the MCP manager with the constructed configuration
	manager, err := mcp.NewMCPHost(g, options)

	if err != nil {
		return nil, fmt.Errorf("failed to create MCP manager: %w", err)
	}

	log.Printf("Successfully created MCP manager '%s' (v%s) with %d servers", name, version, len(mcpServers))
	for _, server := range mcpServers {
		if server.Config.Stdio != nil {
			log.Printf("  - Server: %s (command: %s %v)", server.Name, server.Config.Stdio.Command, server.Config.Stdio.Args)
		} else {
			log.Printf("  - Server: %s", server.Name)
		}
	}

	return manager, nil
}
