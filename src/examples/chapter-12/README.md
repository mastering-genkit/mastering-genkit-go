# Chapter 12 Example: AI Agent

This example demonstrates how to build an AI agent with Genkit Go's built-in **Agents API**. The agent keeps its conversation in a session, calls tools, and streams its responses to an interactive terminal UI.

## Features

- **Built-in Genkit Agent**: Defined with `genkitx.DefineAgent` and an inline prompt
- **Conversation Memory**: The session state (messages, tool calls and tool results) is saved as a snapshot after every turn in Genkit's built-in in-memory session store
- **Tool Calling**: The agent can call a `getCurrentDateTime` tool
- **Streaming**: Responses and tool calls are streamed through a bidirectional connection (`Connect`)
- **Interactive CLI Interface**: Chat with the AI agent directly from your terminal
- **Local AI Model**: Uses Ollama with the `gemma4:e4b` model for privacy and offline operation

## Prerequisites

Before running this example, ensure you have:

1. **Go 1.27+** installed
2. **Ollama** installed and running locally
3. **gemma4:e4b model** downloaded in Ollama (it supports tool calling)

### Setting up Ollama

1. Install Ollama from [https://ollama.com](https://ollama.com)
2. Start Ollama (it typically runs on http://localhost:11434)
3. Download the required model:
   ```bash
   ollama pull gemma4:e4b
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

   Or with the Genkit Developer UI, to inspect the traces of every turn:
   ```bash
   genkit start -- go run .
   ```

## Using the AI Agent

Once the agent starts, you'll see a welcome message and be able to interact with it:

```
🤖 AI Agent
Type your message and press Enter. Type 'quit', 'exit', or 'bye' to exit.
Type 'clear' to start a new conversation.
Type 'history' to see the conversation history.

You: my name is xavi
🤖 AI: Hello Xavi! It's nice to meet you.

You: what is my name?
🤖 AI: Your name is Xavi.

You: What time is it?
🤖 AI: 
   🔧 getCurrentDateTime {}
It is currently Monday, October 5, 2026, at 17:47:30 CEST.

You: history
📝 Conversation History:
--------------------------------------------------
1. You: my name is xavi
2. 🤖 AI: Hello Xavi! It's nice to meet you.
3. You: what is my name?
4. 🤖 AI: Your name is Xavi.
5. You: What time is it?
6. 🔧 Tool call: getCurrentDateTime {}
7. 🔧 Tool result: getCurrentDateTime "Mon, 05 Oct 2026 17:47:30 CEST"
8. 🤖 AI: It is currently Monday, October 5, 2026, at 17:47:30 CEST.
--------------------------------------------------
You: clear
✅ Conversation cleared. Starting a new session.
You: what is my name?
🤖 AI: I do not have access to your personal information, so I do not know your name.

You: quit
Goodbye! 👋
```

## Available Commands

- **Regular messages**: Type any message to chat with the AI
- **`quit`**, **`exit`**, or **`bye`**: Exit the application
- **`clear`**: Start a new session with an empty conversation
- **`history`**: Display the conversation stored in the current session
- **Empty line**: Continue to next prompt (no action)

## Architecture

The project is structured as follows:

```
├── main.go                    # Entry point - sets up Genkit, the session store and the agent
├── internal/
│   ├── agent/
│   │   └── assistant.go      # Agent definition and tools
│   └── cli/
│       └── cli.go            # CLI interface: connection, streaming and commands
├── go.mod                    # Go module definition
└── README.md                 # This file
```

### Components

1. **`main.go`**: Initializes Genkit with the Ollama plugin and the experimental APIs (`genkit.WithExperimental()`), creates the in-memory session store and starts the CLI
2. **`agent/assistant.go`**: Defines the `getCurrentDateTime` tool and the `assistant` agent with `genkitx.DefineAgent`
3. **`cli/cli.go`**: Opens a bidirectional connection to the agent, streams each turn, and reads the conversation history back from the session store

## How It Works

1. **Initialization**: Genkit is initialized with the Ollama plugin and the experimental APIs enabled
2. **Agent Definition**: The agent declares its model, system prompt and tools in an `aix.InlinePrompt`, and uses `localstore.NewInMemorySessionStore` as its session store
3. **CLI Loop**: The CLI calls `agent.Connect` to open a session and then:
   - Reads user input from the terminal
   - Processes special commands (quit, clear, history)
   - Sends regular messages with `conn.SendText`
   - Streams the response and any tool calls with `conn.Receive` until the turn ends
4. **Context Preservation**: The agent appends every message to its session state and saves a snapshot after each turn, so the model always sees the full conversation. `clear` closes the connection and opens a new session with an empty history.
