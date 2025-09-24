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

## Docker Deployment

### Build Docker Image

```bash
docker build -t go-hello-world .
```

### Run Docker Container

```bash
docker run -p 8080:8080 go-hello-world
```

## Leapcell Deployment

This application is ready for deployment on Leapcell. The Dockerfile uses multi-stage builds for optimal image size and includes all necessary configurations.

### Deployment Steps

1. Push your code to a Git repository
2. Connect your repository to Leapcell
3. Leapcell will automatically detect the Dockerfile and build your application
4. The application will be available on the provided Leapcell URL

### Configuration

The application automatically uses the `PORT` environment variable provided by Leapcell. No additional configuration is required.

## Project Structure

```
.
├── main.go          # Main application code
├── go.mod           # Go module definition
├── go.sum           # Go module checksums
├── Dockerfile       # Docker configuration
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

- [Gorilla Mux](https://github.com/gorilla/mux) - HTTP router and URL matcher
