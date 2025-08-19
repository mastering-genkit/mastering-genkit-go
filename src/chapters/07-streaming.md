# Streaming Flows

## Introduction

The transformation from batch to streaming represents one of the most fundamental shifts in how we design interactive systems. Consider the difference between watching a master craftsman work versus simply receiving the finished product. There's something deeply human about understanding through observation of process, not just outcomes. We learn by watching the steps unfold, seeing the decisions being made, understanding the reasoning as it happens. This principle applies directly to AI systems - users don't just want answers, they want to witness the thinking unfold.

Streaming transforms AI interactions from transactional to conversational. Instead of waiting for complete responses, users see thoughts forming in real-time, creating a sense of collaboration rather than computation. This shift has profound implications for system architecture, resource management, and user experience design.

This chapter explores streaming through multiple lenses:

- The cognitive and perceptual foundations of why streaming improves user experience
- Deep architectural analysis of streaming protocols and their trade-offs
- Internal implementation details of Genkit's streaming architecture

## Prerequisites

Before diving into streaming, you should be comfortable with:

- Building basic Flows (Chapter 6)
- Understanding Genkit's generation patterns (Chapter 4)
- HTTP server implementation in Go

## The Psychology of Progressive Disclosure

The effectiveness of streaming isn't accidental - it taps into fundamental aspects of human cognition and perception. Research in cognitive psychology reveals that humans process information more efficiently when it arrives in progressive chunks rather than all at once, a phenomenon related to the concept of "chunking" in working memory.

Consider the difference between reading a book page by page versus having the entire contents dumped into your mind instantly. The sequential, progressive nature of streaming aligns with how we naturally process complex information. This isn't just about perceived performance - though that matters enormously - it's about cognitive compatibility.

Furthermore, streaming provides users with agency. The ability to stop, redirect, or modify a request mid-stream transforms the interaction from a passive waiting experience into an active dialogue. This psychological shift from "waiting for results" to "participating in generation" fundamentally changes user engagement patterns.

## Implementing Streaming Flows in Genkit

Genkit provides streaming capabilities through `DefineStreamingFlow`, which extends the standard Flow pattern with real-time data delivery. The key difference is the addition of a stream callback that sends partial results as they become available.

```go
genkit.DefineStreamingFlow(g, "streamingFlow",
    func(ctx context.Context, request string, stream func(context.Context, string) error) (string, error) {
        // Input: what we're processing
        // Stream: incremental results sent via callback
        // Output: final complete result
    })
```

The three type parameters serve distinct purposes:

- **Input Type**: What the flow receives from the client
- **Stream Type**: The format of incremental chunks sent during processing
- **Output Type**: The complete final response

This design allows flexibility - clients can choose to consume the stream for real-time updates or wait for the complete response. The stream callback sends data immediately while the function continues processing, enabling responsive user interfaces without sacrificing the ability to get complete results.

## How Genkit Streams: Server-Sent Events (SSE)

Genkit uses Server-Sent Events (SSE) for streaming responses. While WebSockets might seem like an obvious choice for real-time communication, SSE is actually better suited for AI text generation where data flows in one direction: from server to client.

### WebSocket vs SSE: A Detailed Comparison

| Aspect | WebSocket | Server-Sent Events (SSE) |
|--------|-----------|--------------------------|
| **Communication** | Bidirectional (full-duplex) | Unidirectional (server → client) |
| **Protocol** | Custom protocol over TCP | Standard HTTP |
| **Data Format** | Text and binary | Text only |
| **Complexity** | Higher (state management) | Lower (stateless) |
| **Use Case** | Real-time bidirectional | Server push notifications |
| **Memory Usage** | 2-3x higher | Baseline |
| **Infrastructure** | Special handling required | Standard HTTP infra |

For AI streaming, SSE's simplicity and unidirectional nature align perfectly with the use case: the client sends a prompt, and the server streams back a response.

### Why SSE Works Better for AI

Looking at modern LLM chat applications like ChatGPT, Claude, and Gemini, we can observe how streaming responses have become the standard for AI interactions. Their real-time text generation demonstrates why SSE is particularly effective for AI chat applications:

**Perfect Fit for the Use Case**: AI text generation is unidirectional - the client sends a prompt, the server streams back a response. WebSocket's bidirectional capabilities add complexity without providing value for this specific pattern.

**Proven Pattern**: The widespread adoption of streaming in AI chat applications shows its effectiveness for user experience. Users can read and process responses as they're generated, rather than waiting for complete responses.

**Infrastructure Compatibility**: SSE works seamlessly with existing HTTP infrastructure - CDNs, load balancers, and proxies all handle it without special configuration. This matters when deploying globally distributed AI services.
