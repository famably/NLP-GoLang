# Chip AI Analysis Service

A Go REST API for storing and summarizing documents using LLM.

## Features

- 📝 Store documents grouped by ID (in-memory storage)
- 🔍 Retrieve all documents for a group
- ✨ Generate AI-powered summaries (mock or real LLM)
- 🧩 Clean architecture with mockable components
- 🐳 Docker support
- 🧪 Comprehensive tests

## Technical Implementation

### Core Components
| Component | Implementation Choice | Rationale |
|-----------|----------------------|-----------|
| **Storage** | In-memory with `sync.RWMutex` | Thread-safe operations without persistence complexity |
| **LLM** | Mocked client (`internal/llm/mock_client.go`) | Enables testing without API dependencies | 
| **HTTP** | Gin framework | Lightweight with middleware support |
| **Testing** | 100% handler coverage | Verifies all API contracts |

## Setup Instructions

### 1. Prerequisites
```bash
go version 1.25.0

docker 24.0+
```

### 2. Development & Testing
```bash
# Run the API locally (hot-reload for development)
go run main.go

# Run tests (with verbose output)
go test -v ./...

# Build executable (outputs to bin/chip-api)
go build -o bin/chip-api main.go
```

### 3. Docker Workflow
```bash
# Build Docker image (tags as 'chiplens')
docker build -t chip-api .

# Run container (maps port 8080)
docker run -p 8080:8080 chip-api
```

### 4. Makefile Shortcuts
```bash
# Build, test, and run using Makefile
make run     # Run the application in development mode
make test    # Runs tests
make build   # Compiles binary to bin/chip-api
make docker  # Builds Docker image