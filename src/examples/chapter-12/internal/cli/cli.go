package cli

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/firebase/genkit/go/ai"
	aix "github.com/firebase/genkit/go/ai/exp"
)

// errQuit signals that the user asked to leave the application.
var errQuit = errors.New("quit")

// CLI is an interactive terminal front end for a Genkit agent.
type CLI struct {
	agent   *aix.Agent[any]
	scanner *bufio.Scanner
}

// New creates a CLI for the given agent.
func New(a *aix.Agent[any]) *CLI {
	return &CLI{
		agent:   a,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

// Run starts the interactive loop. Each iteration of the outer loop is one
// agent session; "clear" ends the current session and starts a fresh one.
func (c *CLI) Run(ctx context.Context) error {
	fmt.Println("🤖 AI Agent")
	fmt.Println("Type your message and press Enter. Type 'quit', 'exit', or 'bye' to exit.")
	fmt.Println("Type 'clear' to start a new conversation.")
	fmt.Println("Type 'history' to see the conversation history.")
	fmt.Println()

	for {
		err := c.runSession(ctx)
		if errors.Is(err, errQuit) {
			fmt.Println("Goodbye! 👋")
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println("✅ Conversation cleared. Starting a new session.")
	}
}

// runSession opens a bidirectional connection to the agent and exchanges
// turns until the user clears the conversation (nil) or quits (errQuit).
func (c *CLI) runSession(ctx context.Context) error {
	// Connect without a session ID starts a brand-new session
	conn, err := c.agent.Connect(ctx)
	if err != nil {
		return fmt.Errorf("connect to agent: %w", err)
	}

	var (
		lastSnapshot string // snapshot saved at the end of the last turn
		result       error
	)
loop:
	for {
		fmt.Print("You: ")
		if !c.scanner.Scan() {
			result = errQuit
			break
		}
		input := strings.TrimSpace(c.scanner.Text())

		// Handle special commands
		switch strings.ToLower(input) {
		case "quit", "exit", "bye":
			result = errQuit
			break loop
		case "clear":
			break loop
		case "history":
			c.showHistory(ctx, lastSnapshot)
			continue
		case "":
			continue
		}

		// Send the user message and stream the agent's reply
		if err := conn.SendText(input); err != nil {
			result = fmt.Errorf("send message: %w", err)
			break
		}
		end, err := streamTurn(conn)
		if err != nil {
			result = err
			break
		}
		if end.SnapshotID != "" {
			lastSnapshot = end.SnapshotID
		}
		if end.FinishReason == aix.AgentFinishReasonFailed {
			// A failed turn ends the invocation; Output reports the error.
			break
		}
	}

	// Output closes the connection and waits for the final result.
	out, err := conn.Output()
	if err != nil && !errors.Is(err, context.Canceled) {
		fmt.Printf("❌ Error: %v\n", err)
	}
	if out != nil && out.Error != nil {
		fmt.Printf("❌ Agent failed (%s): %s\n", out.Error.Status, out.Error.Message)
	}
	return result
}

// streamTurn prints the agent's reply as it streams, including the tools it
// calls, and returns once the turn has ended.
func streamTurn(conn *aix.AgentConnection[any]) (*aix.TurnEnd, error) {
	fmt.Print("🤖 AI: ")
	for chunk, err := range conn.Receive() {
		if err != nil {
			return nil, fmt.Errorf("stream turn: %w", err)
		}
		if mc := chunk.ModelChunk; mc != nil {
			fmt.Print(mc.Text())
			for _, p := range mc.Content {
				if p.IsToolRequest() && !p.ToolRequest.Partial {
					fmt.Printf("\n   🔧 %s %s\n", p.ToolRequest.Name, toJSON(p.ToolRequest.Input))
				}
			}
		}
		if chunk.TurnEnd != nil {
			fmt.Print("\n\n")
			return chunk.TurnEnd, nil
		}
	}
	return nil, errors.New("connection closed before the turn ended")
}

// showHistory prints the conversation stored in a session snapshot.
func (c *CLI) showHistory(ctx context.Context, snapshotID string) {
	if snapshotID == "" {
		fmt.Println("📝 No conversation history yet.")
		return
	}
	snap, err := c.agent.GetSnapshot(ctx, snapshotID)
	if err != nil {
		fmt.Printf("❌ Error reading snapshot: %v\n", err)
		return
	}

	fmt.Println("📝 Conversation History:")
	fmt.Println(strings.Repeat("-", 50))
	for i, msg := range snap.State.Messages {
		switch msg.Role {
		case ai.RoleUser:
			fmt.Printf("%d. You: %s\n", i+1, msg.Text())
		case ai.RoleModel:
			for _, p := range msg.Content {
				if p.IsToolRequest() {
					fmt.Printf("%d. 🔧 Tool call: %s %s\n", i+1, p.ToolRequest.Name, toJSON(p.ToolRequest.Input))
				}
			}
			if text := msg.Text(); text != "" {
				fmt.Printf("%d. 🤖 AI: %s\n", i+1, text)
			}
		case ai.RoleTool:
			for _, p := range msg.Content {
				if p.IsToolResponse() {
					fmt.Printf("%d. 🔧 Tool result: %s %s\n", i+1, p.ToolResponse.Name, toJSON(p.ToolResponse.Output))
				}
			}
		}
	}
	fmt.Println(strings.Repeat("-", 50))
}

func toJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%v", v)
	}
	return string(b)
}
