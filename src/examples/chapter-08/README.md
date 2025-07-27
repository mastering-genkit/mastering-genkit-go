# Chapter 8 Example: Directory Listing and Date Tools with AI

This example demonstrates how to create a Genkit Go application that uses AI tools to generate content while providing useful utilities like directory listing and current date/time information.

## Features

- Uses Google AI (Gemini 2.5 Flash) for content generation
- Implements a custom tool that lists directories in the current working directory
- Implements a date tool that provides current date/time in various formats
- Demonstrates proper Go project structure and conventions
- Includes error handling and logging

## Tools Available

### 1. List Directories Tool
Lists all directories in the current working directory.

### 2. Current Date Tool
Provides the current date and time in a specified format. Supports:
- Named formats: `RFC3339` (default), `Kitchen`, `Stamp`, `DateTime`, `DateOnly`, `TimeOnly`, `RFC822`, `RFC1123`
- Custom Go time formats (e.g., `2006-01-02 15:04:05`)
- JSON input: `{"format": "Kitchen"}` or simple string input

## Running the Example

1. Set up your Google AI API key:
   ```bash
   export GOOGLE_AI_API_KEY="your-api-key-here"
   ```

2. Optionally set a custom port (defaults to 8080):
   ```bash
   export PORT="3000"
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. The server will start and you can make POST requests to `/operatingSystemFlow` with a topic to generate content using the available tools.

## Project Structure

- `main.go` - Entry point and server setup
- `internal/flows/os.go` - Flow definition for content generation
- `internal/tools/directories.go` - Tool implementation for directory listing
- `internal/tools/date.go` - Tool implementation for date/time formatting