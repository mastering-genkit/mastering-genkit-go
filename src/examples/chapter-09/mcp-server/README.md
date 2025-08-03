# Chapter 9 Example: MCP (Model Context Protocol) Server with File System Tools

This example demonstrates how to create a Genkit Go application that implements an MCP (Model Context Protocol) server, exposing AI tools via standardized protocol for use by MCP clients like Claude Desktop, VS Code, and other compatible applications.

## Features

- Implements MCP (Model Context Protocol) server using Genkit Go
- Exposes file system management tools via MCP protocol
- Demonstrates proper Go project structure with internal packages
- Includes comprehensive system information gathering
- Provides error handling and logging
- Serves tools via stdio interface for MCP client integration

## Tools Available

### 1. List Directories Tool
Lists all directories in a specified directory path.

**Input:**
- `directory`: Directory path to list contents of

**Features:**
- Filters results to show only directories (not files)
- Comprehensive error handling for invalid paths
- Returns empty result with error message if no directories found

### 2. Create Directory Tool
Creates new directories with specified names or paths.

**Input:**
- `directory`: Directory name or path to create

**Features:**
- Supports creating nested directory structures using `os.MkdirAll`
- Uses safe permissions (0755)
- Comprehensive error handling and validation
- Input validation to prevent empty directory names

### 3. System Information Tool
Gathers comprehensive system information based on specified criteria and formatting options.

**Input:**
- `info_types`: Array of information types to gather (`env`, `workdir`, `hostname`, `user`)
- `env_vars`: Optional array of specific environment variables to retrieve
- `format`: Output format (`json`, `text`, or `summary`)
- `include_empty`: Whether to include empty/unset values
- `max_length`: Maximum length for each value (0 for no limit)

**Supported Info Types:**
- `env`: Environment variables (common ones or specified list)
- `workdir`: Current working directory
- `hostname`: System hostname
- `user`: Current user (with Windows fallback support)

**Output Formats:**
- `json`: Structured JSON output
- `text`: Plain text format
- `summary`: Text with summary header

## MCP Server Architecture

The server is built using Genkit's MCP plugin and follows these patterns:

1. **Server Creation**: Uses the custom `NewMCPServer` wrapper for simplified setup
2. **Tool Registration**: Tools are registered and exposed automatically via MCP protocol
3. **Stdio Interface**: Serves via standard input/output for MCP client integration
4. **Protocol Compliance**: Fully compliant with MCP specification

## Running the MCP Server

### Prerequisites

1. Ensure Go 1.24.5 or later is installed

### Starting the Server

1. Navigate to the server directory:
   ```bash
   cd src/examples/chapter-09/mcp-server
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the MCP server:
   ```bash
   go run main.go
   ```

The server will start and listen for MCP protocol messages via stdio. It's designed to be used with MCP clients rather than direct HTTP requests.

## Development and Debugging

### Inspecting the Server

You can inspect the MCP server capabilities using the MCP inspector:

```bash
npx @modelcontextprotocol/inspector go run main.go
```

This will open a web interface where you can:
- View available tools and their schemas
- Test tool execution interactively
- Debug MCP protocol communications
- Validate server compliance

## Integrating with MCP Clients

### Claude Desktop Integration

Add to your Claude Desktop configuration file (`claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "genkit-file-system": {
      "command": "go",
      "args": ["run", "main.go"],
      "cwd": "/path/to/your/mcp-server"
    }
  }
}
```

## Tool Usage Examples

Once connected via an MCP client, you can use natural language to interact with the tools:

### Directory Operations
- "List all directories in /Users/username/Documents"
- "Create a new directory called 'my-project'"
- "Make a directory structure 'app/src/components'"

### System Information
- "Get system information including environment and working directory"
- "Show me the current user and hostname"
- "Get the PATH environment variable in JSON format"
- "Get system info but limit each value to 50 characters"

## Project Structure

```
mcp-server/
├── main.go                 # Entry point and server setup
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── README.md              # This file
└── internal/
    ├── mcp/
    │   └── server.go      # MCP server wrapper utilities
    └── tools/
        └── directories.go  # Tool implementations
```

## Key Components

### MCP Server Wrapper (`internal/mcp/server.go`)
Provides a simplified interface for creating MCP servers with sensible defaults and easy tool registration.

### Tool Implementations (`internal/tools/directories.go`)
Contains all tool implementations:
- File system operations (list, create directories)
- System information gathering
- Proper input validation and error handling

### Main Application (`main.go`)
- Initializes Genkit
- Creates and configures the MCP server
- Registers tools and starts the stdio interface

## Dependencies

- **Genkit Go**: Core framework for AI applications
- **MCP Plugin**: Implements Model Context Protocol server capabilities