# Chapter 12 Example: AI Agent

This example demonstrates how to build an AI agent with conversation history. The tool integrates with Genkit Go to provide an interactive chat experience through the terminal.

## Features

- **Interactive CLI Interface**: Chat with the AI agent directly from your terminal
- **Conversation History**: Maintains context across multiple messages in a session
- **Command Support**: Built-in commands for managing the conversation
- **Real-time Responses**: Uses Genkit Generate to process messages and provide responses
- **Local AI Model**: Uses Ollama with the gemma3n:e4b model for privacy and offline operation

## Prerequisites

Before running this example, ensure you have:

1. **Go 1.24+** installed
2. **Ollama** installed and running locally
3. **gemma3n:e4b model** downloaded in Ollama

### Setting up Ollama

1. Install Ollama from [https://ollama.ai](https://ollama.ai)
2. Start Ollama (it typically runs on http://localhost:11434)
3. Download the required model:
   ```bash
   ollama pull gemma3n:e4b
   ```
4. Verify the model is available:
   ```bash
   ollama list
   ```

## Running the Example

1. **Navigate to the project directory:**
   ```bash
   cd src/examples/chapter-12
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the AI agent:**
   ```bash
   go run .
   ```

## Using the AI Agent

Once the agent starts, you'll see a welcome message and be able to interact with it:

```
🤖 AI Agent
Type your message and press Enter. Type 'quit', 'exit', or 'bye' to exit.
Type 'clear' to clear conversation history.
Type 'history' to see the conversation history.

You: Hello, how are you?
🤖 AI: Hello! I'm doing well, thank you for asking. I'm here and ready to help you with any questions or tasks you might have. How are you doing today?

You: What's the weather like?
🤖 AI: I don't have access to real-time weather information since I'm a local AI model without internet connectivity. To get current weather information, you would need to:

1. Check a weather website like Weather.com or AccuWeather
2. Use a weather app on your phone
3. Ask a voice assistant with internet access
4. Look outside your window for current local conditions

Is there anything else I can help you with that doesn't require real-time data?

You: history
📝 Conversation History:
--------------------------------------------------
1. You: Hello, how are you?
2. 🤖 AI: Hello! I'm doing well, thank you for asking...
3. You: What's the weather like?
4. 🤖 AI: I don't have access to real-time weather information...
--------------------------------------------------

You: clear
✅ Conversation history cleared.

You: quit
Goodbye! 👋
```

## Available Commands

- **Regular messages**: Type any message to chat with the AI
- **`quit`**, **`exit`**, or **`bye`**: Exit the application
- **`clear`**: Clear the conversation history and start fresh
- **`history`**: Display the current conversation history
- **Empty line**: Continue to next prompt (no action)

## Architecture

The project is structured as follows:

```
├── main.go                    # Entry point - sets up Genkit and starts AI agent
├── internal/
│   ├── cli/
│   │   └── agent.go          # CLI interface and user interaction logic
│   └── flows/
│       └── chat.go           # Genkit flow for processing chat messages
├── go.mod                    # Go module definition
└── README.md                 # This file
```

### Components

1. **`main.go`**: Initializes Genkit with the Ollama plugin and starts the AI agent
2. **`cli/agent.go`**: Handles user input, manages conversation history, and provides the interactive interface
3. **`flows/chat.go`**: Defines the Genkit flow that processes messages and maintains conversation context

## How It Works

1. **Initialization**: The application starts by initializing Genkit with the Ollama plugin
2. **Flow Creation**: A chat flow is created that can process messages with conversation history
3. **CLI Loop**: The agent enters an interactive loop:
   - Reads user input from the terminal
   - Processes special commands (quit, clear, history)
   - Sends regular messages to the chat flow with conversation history
   - Displays the AI response
   - Updates the conversation history
4. **Context Preservation**: Each message includes the full conversation history, allowing the AI to understand context and provide relevant responses
