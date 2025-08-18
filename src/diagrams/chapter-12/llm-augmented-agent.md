```mermaid
graph TD
    A[User Input] --> B[LLM Agent]
    B --> C[Memory System]
    B --> D[Tool Integration]
    C --> B
    D --> B
```