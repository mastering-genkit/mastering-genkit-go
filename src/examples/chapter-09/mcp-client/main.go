package main

import (
	"context"
	"log"
	"mastering-genkit-go/example/chapter-09/mcp-client/internal/flows"
	"mastering-genkit-go/example/chapter-09/mcp-client/internal/tools"

	mcpinternal "mastering-genkit-go/example/chapter-09/mcp-client/internal/mcp"
	"net/http"
	"os"

	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/mcp"
	"github.com/firebase/genkit/go/plugins/server"
	bedrock "github.com/xavidop/genkit-aws-bedrock-go"
)

func main() {
	ctx := context.Background()

	bedrockPlugin := &bedrock.Bedrock{
		Region: "us-east-1",
	}

	// Initialize Genkit with the AWS Bedrock plugin and Haiku model.
	g := genkit.Init(ctx,
		genkit.WithPlugins(bedrockPlugin),
		genkit.WithDefaultModel("bedrock/anthropic.claude-3-haiku-20240307-v1:0"), // Set default model

	)

	bedrock.DefineCommonModels(bedrockPlugin, g)

	// Get the MCPClient for file operations
	mcpFileSystem := mcpinternal.NewFilesystemServerConfig("file-system", "./")

	// Create the MCP manager with the file system server
	manager, err := mcpinternal.NewMCPManagerWrapper("my-manager", "1.0.0", []mcp.MCPServerConfig{
		mcpFileSystem,
	})
	if err != nil {
		log.Fatalf("Failed to create MCP manager: %v", err)
	}

	// Get all active tools from the MCP manager
	toolList, err := manager.GetActiveTools(ctx, g)
	if err != nil {
		log.Fatalf("Failed to get active tools: %v", err)
	}

	operatingSystemFlow := flows.NewOperatingSystemFlow(g, tools.ConvertToolsToToolRefs(toolList))

	mux := http.NewServeMux()
	mux.HandleFunc("POST /operatingSystemFlow", genkit.Handler(operatingSystemFlow))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Starting server on 127.0.0.1:%s", port)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+port, mux))
}
