package main

import (
	"context"
	"log"
	mcpinternal "mastering-genkit-go/example/chapter-09/mcp-server/internal/mcp"
	"mastering-genkit-go/example/chapter-09/mcp-server/internal/tools"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

func main() {
	ctx := context.Background()

	// Initialize Genkit with the OpenAI plugin with your OpenAI API key.
	g, err := genkit.Init(ctx)
	if err != nil {
		log.Fatalf("could not initialize Genkit: %v", err)
	}

	// Create the MCP server with the file system server
	// This server will expose tools via the Model Context Protocol
	// and allow clients to connect and use them.
	mcpServer := mcpinternal.NewMCPServer(g, "file-system", "1.0.0", []ai.Tool{
		tools.NewListDirectories(g),
		tools.NewCreateDirectory(g),
		tools.NewSystemInfo(g),
	})
	// Start the MCP server to expose tools via the Model Context Protocol
	log.Println("Starting MCP server...")
	err = mcpServer.ServeStdio(ctx)
	if err != nil {
		log.Fatalf("could not start MCP server: %v", err)
	}
}
