package cli

import (
	"bufio"
	"context"
	"fmt"
	"mastering-genkit-go/example/chapter-12/internal/flows"
	"os"
	"strings"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
)

// Agent represents the CLI agent with conversation history
type Agent struct {
	chatFlow *core.Flow[flows.ChatRequest, flows.ChatResponse, struct{}]
	history  []*ai.Message
}

// NewAgent creates a new CLI agent
func NewAgent(chatFlow *core.Flow[flows.ChatRequest, flows.ChatResponse, struct{}]) *Agent {
	return &Agent{
		chatFlow: chatFlow,
		history:  make([]*ai.Message, 0),
	}
}

// Run starts the interactive CLI session
func (a *Agent) Run(ctx context.Context) error {
	fmt.Println("🤖 AI Agent")
	fmt.Println("Type your message and press Enter. Type 'quit', 'exit', or 'bye' to exit.")
	fmt.Println("Type 'clear' to clear conversation history.")
	fmt.Println("Type 'history' to see the conversation history.")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("You: ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		// Handle special commands
		switch strings.ToLower(input) {
		case "quit", "exit", "bye":
			fmt.Println("Goodbye! 👋")
			return nil
		case "clear":
			a.clearHistory()
			fmt.Println("✅ Conversation history cleared.")
			continue
		case "history":
			a.showHistory()
			continue
		case "":
			continue
		}

		// Process the user input through the chat flow
		response, err := a.processMessage(ctx, input)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		// Display the AI response
		fmt.Printf("🤖 AI: %s", response.Response)

		// Update local history with the response
		a.history = response.History
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	return nil
}

// processMessage sends a message to the chat flow and returns the response
func (a *Agent) processMessage(ctx context.Context, message string) (*flows.ChatResponse, error) {
	request := flows.ChatRequest{
		Message: message,
		History: a.history,
	}

	// Execute the chat flow
	response, err := a.chatFlow.Run(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("chat flow execution failed: %w", err)
	}

	return &response, nil
}

// clearHistory clears the conversation history
func (a *Agent) clearHistory() {
	a.history = make([]*ai.Message, 0)
}

// showHistory displays the current conversation history
func (a *Agent) showHistory() {
	if len(a.history) == 0 {
		fmt.Println("📝 No conversation history yet.")
		return
	}

	fmt.Println("📝 Conversation History:")
	fmt.Println(strings.Repeat("-", 50))

	for i, msg := range a.history {
		var role string
		switch msg.Role {
		case ai.RoleUser:
			role = "You"
		case ai.RoleModel:
			role = "🤖 AI"
		default:
			role = "System"
		}

		// Extract text content from message parts
		var text string
		if len(msg.Content) > 0 && msg.Content[0] != nil {
			// ai.Part has a Text field (not method)
			text = msg.Content[0].Text
		}

		fmt.Printf("%d. %s: %s\n", i+1, role, text)
	}
	fmt.Println(strings.Repeat("-", 50))
}
