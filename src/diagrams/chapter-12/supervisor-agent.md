```mermaid
graph TD
    A[User Input] --> B[Supervisor Agent]
    B --> C[Research Agent]
    B --> D[Analysis Agent]
    B --> E[Writing Agent]
    C --> B
    D --> B
    E --> B
    B --> F[Final Response]
```