# Chapter 10 Example: Retrieval Augmented Generation (RAG) with Genkit Go

This example shows how to build a complete Retrieval Augmented Generation (RAG) system using Genkit Go. It shows how to index PDF documents, create vector embeddings, store them locally, and perform semantic search to enhance AI responses with relevant context.

## Features

- **PDF Document Processing**: Extracts and processes text from PDF files
- **Vector Embeddings**: Uses OpenAI's text-embedding-3-large model for semantic embeddings
- **Local Vector Storage**: Stores embeddings locally using localvec plugin
- **Semantic Search**: Retrieves relevant document chunks based on user queries
- **RAG Implementation**: Combines retrieved context with LLM responses
- **Two-Phase Architecture**: Separate indexing and retrieval flows

## System Components

### 1. Indexer Flow
Processes PDF documents and creates searchable vector embeddings:
- Reads PDF files (default: Arduino report)
- Extracts and chunks text content
- Generates embeddings using OpenAI
- Stores vectors in local database

### 2. Retrieval Flow
Performs semantic search and generates contextualized responses:
- Takes user queries as input
- Searches vector database for relevant content
- Retrieves top-K most similar documents
- Generates AI responses enhanced with retrieved context

## AI Flows

### 1. Indexer Flow
**Endpoint:** `POST /indexerFlow`

Indexes documents for search. Use this to process and store PDF content.

**Request Body:**
```json
{
  "pdfPath": "path/to/document.pdf"  // Optional, defaults to Arduino report
}
```

**Example:**
```bash
curl -X POST http://localhost:9090/indexerFlow \
  -H "Content-Type: application/json" \
  -d '{"pdfPath": "internal/docs/arduino_report.pdf"}'
```

### 2. Retrieval Flow
**Endpoint:** `POST /retrievalFlow`

Searches indexed documents and generates contextual responses.

**Request Body:**
```json
{
  "query": "Your question here",
  "k": 5  // Optional, number of results to retrieve (default: 5)
}
```

**Example:**
```bash
curl -X POST http://localhost:9090/retrievalFlow \
  -H "Content-Type: application/json" \
  -d '{"query": "What are the key features of Arduino?", "k": 3}'
```

## Setup and Running

### Prerequisites
- Go 1.24 or later
- OpenAI API key

### 1. Environment Setup
Set your OpenAI API key:
```bash
export OPENAI_API_KEY="your-openai-api-key-here"
```

Optionally set a custom port (defaults to 9090):
```bash
export PORT="3000"
```

### 2. Install Dependencies
```bash
go mod download
```

### 3. Run the Application
```bash
go run main.go
```

## Usage Workflow

### Step 1: Index Documents
First, run the indexer to process and store your documents:
```bash
curl -X POST http://localhost:9090/indexerFlow \
  -H "Content-Type: application/json" \
  -d '{}'
```

### Step 2: Query the System
Once indexing is complete, you can query the system:
```bash
curl -X POST http://localhost:9090/retrievalFlow \
  -H "Content-Type: application/json" \
  -d '{"query": "How does Arduino work?"}'
```

## Technical Details

### Vector Storage
- Uses `localvec` plugin for local vector storage
- Embeddings stored in `.genkit/indexes` directory
- Supports efficient similarity search

### Embedding Model
- OpenAI `text-embedding-3-large` (3072 dimensions)
- High-quality semantic embeddings
- Optimized for retrieval tasks

### Document Processing
- PDF text extraction and chunking
- Maintains document context and metadata
- Handles large documents efficiently

## Project Structure

```
chapter-10/
├── main.go                 # Application entry point
├── internal/
│   ├── docs/              # Document storage
│   │   └── arduino_report.pdf
│   ├── flows/             # Genkit flows
│   │   ├── indexer.go     # Document indexing flow
│   │   └── retrieve.go    # Search and retrieval flow
│   └── rag/               # RAG utilities
└── .genkit/
    └── indexes/           # Vector storage (created at runtime)
```

## Notes

- Run the indexer flow before using the retrieval flow
- The example uses an Arduino report as sample data
- Vector embeddings are persisted locally for reuse
- The system supports multiple document formats through PDF processing
