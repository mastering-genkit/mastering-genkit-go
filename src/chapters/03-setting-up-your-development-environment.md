# Setting Up Your Development Environment

## Introduction

Getting started with Genkit Go requires more than just installing a package. You need to understand how Genkit's unique architecture influences your development setup, project structure, and workflow. Unlike traditional Go libraries that integrate seamlessly into existing projects, Genkit Go brings its own developer tools, runtime requirements, and architectural patterns that fundamentally shape how you build AI applications.

This chapter guides you through setting up a production-ready Genkit Go development environment. You'll learn not just the "how" but the "why" behind each configuration choice, understand the Developer UI's role in your workflow, and establish project structures that scale from prototypes to production systems.

## Prerequisites

Before setting up your Genkit Go environment, ensure you have:

- **Go 1.24 or later**: Genkit Go leverages modern Go features for type safety and performance
- **Node.js 20+ and npm**: Required for the Genkit CLI and Developer UI
- **Git**: For version control and accessing example repositories
- **Basic understanding of Go modules**: Familiarity with `go mod` commands and dependency management

## Installing Genkit CLI

The Genkit CLI is your primary interface for development, providing the Developer UI, flow testing, and deployment utilities. While Genkit Go applications can run without the CLI in production, it's essential for development.

```bash
# Install Genkit CLI globally
npm install -g genkit-cli

# Verify installation
genkit --version
```

The CLI installation provides several key commands:

- `genkit start -- <command to run your code>`: Launches the Developer UI and connects to your running application
- `genkit flow:run <flowName>`: Run a specified flow. Your runtime must already be running in a separate terminal with the `GENKIT_ENV=dev` environment variable set.
- `genkit eval:flow <flowName>`: Evaluate a specific flow. Your runtime must already be running in a separate terminal with the `GENKIT_ENV=dev` environment variable set.

## Understanding the Genkit Developer UI

![](../images/developer-ui.png)

AI application development presents unique challenges that traditional debugging tools weren't designed to handle. When working with LLMs, you face:

- **Non-deterministic outputs**: The same prompt can produce different responses each time, making it impossible to write traditional unit tests with expected outputs
- **Black-box behavior**: Understanding why an LLM generated a specific response is often opaque, unlike stepping through deterministic code
- **Complex interaction chains**: AI applications often involve multiple model calls, tool invocations, and context management that are difficult to trace
- **Cost and latency concerns**: Each model call has financial and performance implications that need monitoring

Genkit's Developer UI addresses these AI-specific challenges by providing purpose-built tools for AI development. Rather than trying to force traditional debugging paradigms onto probabilistic systems, it embraces the unique nature of AI applications with visual inspection of model interactions, flow execution tracking, and real-time observability.

### Architecture Deep Dive

Looking at the source code implementation:

```go
// Reflection server starts on port 3100 by default
addr := "127.0.0.1:3100"
if os.Getenv("GENKIT_REFLECTION_PORT") != "" {
    addr = "127.0.0.1:" + os.Getenv("GENKIT_REFLECTION_PORT")
}
```

[https://github.com/firebase/genkit/blob/main/go/genkit/reflection.go#L61-L126](https://github.com/firebase/genkit/blob/main/go/genkit/reflection.go#L61-L126)

The reflection server is the bridge between your Go application and the Developer UI. When your application runs in development mode, it automatically starts this server, exposing runtime information about flows, prompts, and model interactions. This design allows the JavaScript-based Developer UI to introspect your Go application without language barriers.

### Starting the Genkit Developer UI for Go

When developing with Genkit Go, you have two options for running your application:

1. Production mode: Run your Go application directly with `go run .`
2. Development mode: Use the Genkit CLI to launch both your application and the Developer UI together

For development, the recommended approach is to use the Genkit CLI to manage your application lifecycle. This ensures proper communication between your Go application and the Developer UI through the reflection server:

```bash
# Start Developer UI with your Go application
genkit start -- go run .

# For applications with specific entry points
genkit start -- go run ./cmd/server
```

This approach ensures proper connection between the Developer UI and your application's reflection server. The Developer UI automatically discovers available flows, monitors executions, and provides real-time feedback.

## Project Structure Best Practices

While there's no single "correct" way to structure a Genkit Go application, certain patterns have proven effective for maintainability and scalability. Here's one approach that aligns well with Go conventions while accommodating Genkit's specific needs:

```text
myapp/
├── main.go                 # Application entry point
├── internal/
│   ├── flows/              # Flow definitions
│   │   ├── flows.go        # Flow registration
│   │   ├── greeting.go     # Individual flow implementations
│   │   └── chat.go
│   ├── prompts/            # Prompt management
│   │   └── templates.go
│   └── tools/              # Custom tool implementations
│       └── calculator.go
├── prompts/                # Dotprompt files (optional)
│   └── greeting.prompt
├── go.mod
├── go.sum
└── README.md
```

This structure balances Go's preference for flat hierarchies with the organizational needs of AI applications. The `internal` package prevents external imports of your application logic, while clear subdirectories help organize different Genkit components. Adapt this structure to your specific needs - smaller projects might need less separation, while larger ones might benefit from additional organization.

For our first application in this chapter, we'll start with a much simpler structure - just a single `main.go` file. This allows us to focus on verifying that our environment is correctly set up before diving into more complex architectural patterns in later chapters.

## First Genkit Go Application

Let's build a simple "Hello Genkit Go" application to verify our environment is properly configured. We'll create a minimal example that demonstrates Genkit's core functionality without overwhelming complexity.

### Step 1: Initialize the Project

```bash
# Create and navigate to project directory
mkdir genkit-intro && cd genkit-intro

# Initialize Go module
go mod init example/genkit-intro

# Install Genkit package
go get github.com/firebase/genkit/go

# Download and verify dependencies
go mod tidy
```

### Step 2: Create a Simple Application

Create a `main.go` file with the following content:

> **Note**: If you encounter import errors when writing the code, run `go mod tidy` to ensure all dependencies are properly downloaded.

```go
package main

import (
    "context"
    "log"

    "github.com/firebase/genkit/go/ai"
    "github.com/firebase/genkit/go/genkit"
    "github.com/firebase/genkit/go/plugins/googlegenai"
)

func main() {
    ctx := context.Background()

    // Initialize Genkit with the Google AI plugin and Gemini 2.0 Flash.
    g, err := genkit.Init(ctx,
        genkit.WithPlugins(&googlegenai.GoogleAI{}),
        genkit.WithDefaultModel("googleai/gemini-2.5-flash"),
    )
    if err != nil {
        log.Fatalf("could not initialize Genkit: %v", err)
    }

    resp, err := genkit.Generate(ctx, g, ai.WithPrompt("Hello Genkit Go!"))
    if err != nil {
        log.Fatalf("could not generate model response: %v", err)
    }

    log.Println(resp.Text())
}
```

This minimal example demonstrates the simplest way to use Genkit Go:

- Initializes Genkit with the Google AI plugin
- Makes a single generation request
- Prints the response

Note that Genkit automatically reads the `GEMINI_API_KEY` from your environment variables, so you don't need to explicitly pass it in your code.

### Step 3: Running the Application

Configure your Gemini API key by setting the `GEMINI_API_KEY` environment variable:

```bash
export GEMINI_API_KEY=<your API key>
```

If you don't already have one, create a key in [Google AI Studio](https://aistudio.google.com/app/apikey). Google AI provides a generous free-of-charge tier and does not require a credit card to get started.

> **Security Warning**: Never commit your API key to version control or expose it publicly. Leaked API keys can be exploited, resulting in unexpected charges or security breaches.

Run the app to see the model response:

```bash
go run .
```

If your API key is not properly set, you'll see an error like:

```text
could not initialize Genkit: genkit.Init: plugin *googlegenai.GoogleAI initialization failed: 
GoogleAI.Init: Google AI requires setting GEMINI_API_KEY or GOOGLE_API_KEY in the environment. 
You can get an API key at https://ai.google.dev
exit status 1
```

>Note that Genkit accepts either `GEMINI_API_KEY` or `GOOGLE_API_KEY` as the environment variable name. Throughout this book, we'll use `GEMINI_API_KEY` for consistency.

Once your API key is properly configured, you should see output similar to:

```text
Hello there! Welcome to the world of **Genkit with Go!**

You've reached the Go implementation of Genkit, Google's open-source framework designed to help you build robust, production-ready AI applications. With Genkit Go, you can:

*   **Orchestrate** complex AI flows.
*   **Define custom actions** and integrate with external tools.
*   **Connect** to various large language models (LLMs) like Gemini, OpenAI, and more.
*   **Develop, test, and observe** your GenAI applications with ease.
*   Leverage the **performance and concurrency** strengths of Go.

How can I help you get started or dive deeper into Genkit Go today? Are you looking to:

*   **Set up your first project?**
*   **Understand core concepts** like flows, actions, or models?
*   **Explore specific integrations** or use cases?
*   **Troubleshoot** something you're working on?

Let me know what's on your mind!
```

### Why Can't We Use the Developer UI Yet?

You might be eager to explore the Genkit Developer UI - and for good reason! It's one of Genkit's most powerful features. However, our simple example above runs once and exits immediately, which isn't compatible with the Developer UI.

To use the Developer UI effectively, your application needs to define at least one **Flow** - a reusable AI workflow that can be called and tested through the UI.

Our current example is just a one-shot script without any Flows defined. If you try `genkit start -- go run .`, it will execute the Generate call and exit immediately - this is expected behavior.

Don't worry though! In Chapter 6 "Building Flows", you'll learn how to create applications that work seamlessly with the Developer UI. Once you have Flows defined, the Developer UI unlocks powerful capabilities:

- Visual inspection of every AI interaction
- Interactive testing with different inputs
- Real-time monitoring of token usage and costs
- Step-by-step debugging of complex AI workflows

For now, let's focus on confirming your environment is properly set up. The fact that you got a response from Gemini means you're ready to move forward!

## Instructing Your AI Copilot

Modern AI development requires your coding assistant to understand Genkit's specific patterns and conventions. Genkit provides comprehensive documentation sets at [llms.txt](https://genkit.dev/llms.txt), each serving different purposes:

- **Complete documentation**: [llms-full.txt](https://genkit.dev/llms-full.txt)
- **Abridged documentation**: [llms-small.txt](https://genkit.dev/llms-small.txt)
- **Building AI Workflows**: [_llms-txt/building-ai-workflows.txt](https://genkit.dev/_llms-txt/building-ai-workflows.txt)
- **Deploying AI Workflows**: [_llms-txt/deploying-ai-workflows.txt](https://genkit.dev/_llms-txt/deploying-ai-workflows.txt)
- **Observability**: [_llms-txt/observing-ai-workflows.txt](https://genkit.dev/_llms-txt/observing-ai-workflows.txt)
- **Writing Plugins**: [_llms-txt/writing-plugins.txt](https://genkit.dev/_llms-txt/writing-plugins.txt)

Configure your preferred AI coding assistant by creating the appropriate instruction file. Here are examples for popular tools:

### Claude Code

Create `CLAUDE.md` in your project root:

```markdown
# Genkit Go Project

Reference official Genkit documentation:
- Complete docs: https://genkit.dev/llms-full.txt
- Building workflows: https://genkit.dev/_llms-txt/building-ai-workflows.txt
- Observability: https://genkit.dev/_llms-txt/observing-ai-workflows.txt
```

### Gemini CLI

Create `GEMINI.md` in your project root:

```markdown
# Genkit Go with Gemini

Reference official documentation:
- Genkit docs: https://genkit.dev/llms-full.txt
- Workflow patterns: https://genkit.dev/_llms-txt/building-ai-workflows.txt
- Plugin docs: https://genkit.dev/_llms-txt/plugin-documentation.txt
```

### Cursor

Create `.cursor/rules/genkit.mdc`:

```markdown
---
description:
globs: *.go
alwaysApply: false
---

This project uses Genkit Go. Reference official documentation:
- Complete reference: https://genkit.dev/llms-full.txt
- Quick reference: https://genkit.dev/llms-small.txt
- Deployment guide: https://genkit.dev/_llms-txt/deploying-ai-workflows.txt
```

### GitHub Copilot

Create `.github/instructions/genkit.instructions.md`:

```markdown
---
applyTo: '**'
---

This project uses Genkit Go. Reference the official documentation:
- Complete docs: https://genkit.dev/llms-full.txt
- Building workflows: https://genkit.dev/_llms-txt/building-ai-workflows.txt
```

### Cline

Create `.clinerules` in your project root:

```text
This project uses Genkit Go. Reference official documentation:
- Complete reference: https://genkit.dev/llms-full.txt
- Building workflows: https://genkit.dev/_llms-txt/building-ai-workflows.txt
- Observability: https://genkit.dev/_llms-txt/observing-ai-workflows.txt
```

### Windsurf

Create `.windsurfrules` in your project root:

```text
This project uses Genkit Go. Reference official documentation:
- Genkit docs: https://genkit.dev/llms-full.txt
- Quick reference: https://genkit.dev/llms-small.txt
- Workflow patterns: https://genkit.dev/_llms-txt/building-ai-workflows.txt
```

By creating these simple instruction files, your AI coding assistant will know to reference the official Genkit documentation when generating code. This helps ensure accuracy and keeps your assistant up-to-date with the latest patterns.

## Key Takeaways

- **Go 1.24+ and Node.js 20+** are required for Genkit Go development
- **Genkit CLI is essential** for development, providing the Developer UI and testing utilities
- **Simple setup process**: Initialize project, install dependencies, create main.go, and run
- **Environment variables matter** - use `GEMINI_API_KEY` and never commit it to version control
- **Developer UI requires Flows** - our simple example doesn't support it yet (covered in Chapter 6)
- **AI copilot instructions** help generate accurate Genkit code by referencing official documentation

## Next Steps

With your development environment configured, Chapter 4 dives into mastering AI generation with Genkit Go. You'll learn:

- Deep understanding of `Generate` vs `GenerateData` internals
- Schema validation mechanisms and custom validators
- Error handling patterns specific to AI interactions
- Advanced prompt management with Dotprompt

Your environment is ready. Let's start building sophisticated AI applications with Genkit Go.
