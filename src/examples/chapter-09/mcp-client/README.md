# Chapter 9 Example: Model Context Protocol (MCP) Client with File System Operations

This example demonstrates how to create a Genkit Go application that integrates with the Model Context Protocol (MCP) to provide AI-powered file system operations. The application uses MCP's filesystem server to enable safe and controlled file operations through AI assistance.

## Features

- Uses OpenAI for content generation
- Integrates with MCP (Model Context Protocol) filesystem server
- Provides AI-powered file system operations through MCP tools
- Demonstrates proper Go project structure and conventions
- Includes error handling and logging
- Secure file system access with configurable allowed directories

## MCP Integration

This example uses the MCP filesystem server to provide safe file system operations. The MCP (Model Context Protocol) allows AI models to interact with external tools and services in a secure and standardized way.

### Available MCP Tools

The filesystem MCP server provides the following capabilities:

- **Reading Files**: Read text and media files from allowed directories
- **Writing Files**: Create and edit files within permitted locations
- **Directory Operations**: Create, list, and manage directories
- **File Management**: Move files and directories, manage metadata
- **File Search**: Search for files within the allowed directory structure
- **Metadata Access**: Get file information like size, permissions, and timestamps

## API Endpoints

### Operating System Flow
**Endpoint:** `POST /operatingSystemFlow`

General-purpose flow that includes all available MCP filesystem tools. Use this for file system operations powered by AI.

## Prerequisites

Before running this example, you need to have Node.js and npm installed, as the MCP filesystem server is an npm package.

1. Install Node.js (if not already installed):
   ```bash
   # macOS with Homebrew
   brew install node
   
   # Or download from https://nodejs.org/
   ```

2. The application will automatically install the MCP filesystem server when needed.

## Running the Example

1. Set up your OpenAI API key:
   ```bash
   export OPENAI_API_KEY="your-api-key-here"
   ```

2. Optionally set a custom port (defaults to 9090):
   ```bash
   export PORT="3000"
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. The server will start and you can make POST requests to the available endpoint:
   - `POST /operatingSystemFlow` - AI-powered file system operations via MCP

## Security Configuration

The MCP filesystem server is configured to only allow access to the root directory (`/`) by default. In a production environment, you should restrict this to specific directories for security:

```go
// Example: Restrict to specific directories
mcpFileSystem := mcpinternal.NewFilesystemServerConfig("file-system", 
    "/Users/username/Documents", 
    "/tmp/safe-directory")
```

## Example Usage

### File Operations
Send POST requests for file-related tasks:
- `"Read the contents of the file example.txt"`
- `"Create a new file called hello.txt with the content 'Hello World'"`
- `"List all files in the current directory"`
- `"Search for all .go files in the project"`


## cURL Example

Here are practical examples of calling the flow using cURL commands:

### Read a File
```bash
curl -X POST http://127.0.0.1:9090/operatingSystemFlow \
  -H "Content-Type: application/json" \
  -d '{"data":"Read the contents of README.md and summarize it"}'
```

### Create a File
```bash
curl -X POST http://127.0.0.1:9090/operatingSystemFlow \
  -H "Content-Type: application/json" \
  -d '{"data":"Create a file called test.txt with the content Hello from MCP!"}'
```

## Invoking the flow using Genkit CLI
You can also invoke the flow using the Genkit CLI:

```bash
genkit flow:run operatingSystemFlow '"List all directories in folder /"'
```