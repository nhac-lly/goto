# Go Hello World for Leapcell

A simple Go HTTP server that returns "Hello World" responses in JSON format, designed for deployment on Leapcell.

## Features

- RESTful HTTP server with JSON responses
- Health check endpoint
- Request logging middleware
- Containerized with Docker
- Environment-based port configuration
- Optimized for cloud deployment

## API Endpoints

- `GET /` - Returns a hello world message with timestamp
- `GET /health` - Health check endpoint

## Local Development

### Prerequisites

- Go 1.21 or higher
- Docker (optional, for containerized deployment)

### Running Locally

1. Clone or download this project
2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the server:
   ```bash
   go run main.go
   ```

4. The server will start on port 8080 by default. Visit:
   - http://localhost:8080 for the hello world message
   - http://localhost:8080/health for health check

### Environment Variables

- `PORT` - Server port (default: 8080)

## Leapcell Deployment

This application is ready for deployment on Leapcell using their native Go runtime support. No Docker required!

### Deployment Steps

1. Push your code to a Git repository
2. Connect your repository to Leapcell
3. Leapcell will automatically detect the `leapcell.yaml` configuration and build your Go application
4. The application will be available on the provided Leapcell URL

### Configuration

The project includes a `leapcell.yaml` file that tells Leapcell how to build and run your Go application:

```yaml
name: go-hello-world
runtime: go
build:
  command: go build -o main .
run:
  command: ./main
env:
  PORT: 8080
```

The application automatically uses the `PORT` environment variable provided by Leapcell.

## Project Structure

```
.
├── main.go          # Main application code
├── go.mod           # Go module definition
├── leapcell.yaml    # Leapcell deployment configuration
└── README.md        # This file
```

## Response Format

### Hello World Endpoint (`/`)
```json
{
  "message": "Hello, World from Leapcell!",
  "timestamp": "2023-12-07T10:30:00Z",
  "version": "1.0.0"
}
```

### Health Check Endpoint (`/health`)
```json
{
  "status": "healthy",
  "timestamp": "2023-12-07T10:30:00Z"
}
```

## Dependencies

This project uses only Go's standard library, making it lightweight and easy to deploy without external dependencies.
