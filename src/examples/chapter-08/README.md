# Chapter 8 Example: Directory Management and Date Tools with AI

This example demonstrates how to create a Genkit Go application that uses AI tools to generate content while providing useful utilities like directory management (listing and creating directories) and current date/time information.

## Features

- Uses OpenAI for content generation
- Implements custom tools for directory management operations
- Implements a date tool that provides current date/time in various formats
- Demonstrates proper Go project structure and conventions
- Includes error handling and logging
- Provides two specialized flows for different use cases

## Tools Available

### 1. List Directories Tool
Lists all directories in the current working directory.

### 2. Create Directory Tool
Creates new directories with specified names or paths. Features:
- Supports creating nested directory structures
- Uses safe permissions (0755)
- Comprehensive error handling
- Input validation

### 3. Current Date Tool
Provides the current date and time in a specified format. Supports:
- Named formats: `RFC3339` (default), `Kitchen`, `Stamp`, `DateTime`, `DateOnly`, `TimeOnly`, `RFC822`, `RFC1123`
- Custom Go time formats (e.g., `2006-01-02 15:04:05`)
- JSON input: `{"format": "Kitchen"}` or simple string input

## API Endpoints

### 1. Operating System Flow
**Endpoint:** `POST /operatingSystemFlow`

General-purpose flow that includes all available tools (directory listing, directory creation, and date/time). Use this for general OS operations.

## Running the Example

1. Set up your Google AI API key:
   ```bash
   export GOOGLE_AI_API_KEY="your-api-key-here"
   ```

2. Optionally set a custom port (defaults to 9090):
   ```bash
   export PORT="3000"
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. The server will start and you can make POST requests to the available endpoints:
   - `POST /operatingSystemFlow` - General OS operations (all tools available)

## Example Usage

### Creating Directories
Send a POST request to either endpoint with requests like:
- `"Create a directory called 'test'"`
- `"Make a new folder named 'projects'"`
- `"Create the directory structure 'app/src/components'"`

### Listing Directories
Send requests like:
- `"List all directories in the current folder"`
- `"Show me what directories are here"`

### Getting Current Date/Time
Send requests like:
- `"What's the current date and time?"`
- `"Give me the date in RFC3339 format"`

## cURL Examples

Here are practical examples of calling the flow using cURL commands:

### Create a Directory
```bash
curl -X POST http://127.0.0.1:9090/operatingSystemFlow \
  -H "Content-Type: application/json" \
  -d '{"data":"Create a directory called test-folder"}'
```

### List Directories
```bash
curl -X POST http://127.0.0.1:9090/operatingSystemFlow \
  -H "Content-Type: application/json" \
  -d '{"data":"List all directories in folder /"}'
```

### Get Current Date
```bash
curl -X POST http://127.0.0.1:9090/operatingSystemFlow \
  -H "Content-Type: application/json" \
  -d '{"data":"What is the current date and time?"}'
```

## Running with Genkit CLI
To run this example using the Genkit CLI, follow these steps:
```bash
genkit flow:run operatingSystemFlow '"List all directories in folder /"'
```

## Project Structure

- `main.go` - Entry point and server setup with both flows
- `internal/flows/os.go` - General operating system flow definition
- `internal/tools/directories.go` - Tools for directory listing and creation
- `internal/tools/date.go` - Tool implementation for date/time formatting