package mcp

import (
	"github.com/firebase/genkit/go/plugins/mcp"
)

// NewFilesystemServerConfig creates an MCPServerConfig for the filesystem MCP server.
// This server provides filesystem operations like reading/writing files, creating directories,
// searching files, and managing file metadata.
//
// The filesystem server requires at least one allowed directory to operate.
// You can specify multiple directories that the server will have access to.
//
// Parameters:
//   - name: A unique identifier for this server instance (e.g., "filesystem-server")
//   - allowedDirs: One or more directory paths that the server is allowed to access.
//     All filesystem operations will be restricted to these directories.
//
// Example usage:
//
//	config := NewFilesystemServerConfig("filesystem-server", "/Users/username/Documents", "/tmp")
//
// The server provides tools for:
// - Reading text and media files
// - Writing and editing files
// - Creating and listing directories
// - Moving files and directories
// - Searching files
// - Getting file metadata
//
// For more information, see: https://www.npmjs.com/package/@modelcontextprotocol/server-filesystem
func NewFilesystemServerConfig(name string, allowedDirs ...string) mcp.MCPServerConfig {
	// Build args with the package name and allowed directories
	args := []string{"-y", "@modelcontextprotocol/server-filesystem"}
	args = append(args, allowedDirs...)

	return mcp.MCPServerConfig{
		Name: name,
		Config: mcp.MCPClientOptions{
			Name: name,
			Stdio: &mcp.StdioConfig{
				Command: "npx",
				Args:    args,
			},
		},
	}
}
