package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gorilla/mux"
)

type Response struct {
    Message   string    `json:"message"`
    Timestamp time.Time `json:"timestamp"`
    Version   string    `json:"version"`
}

type HealthResponse struct {
    Status    string    `json:"status"`
    Timestamp time.Time `json:"timestamp"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{
        Message:   "Hello, World from Leapcell!",
        Timestamp: time.Now(),
        Version:   "1.0.0",
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    response := HealthResponse{
        Status:    "healthy",
        Timestamp: time.Now(),
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func main() {
    // Get port from environment variable, default to 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Create router
    router := mux.NewRouter()

    // Define routes
    router.HandleFunc("/", helloHandler).Methods("GET")
    router.HandleFunc("/health", healthHandler).Methods("GET")

    // Add middleware for logging
    router.Use(loggingMiddleware)

    fmt.Printf("Server starting on port %s...\n", port)
    fmt.Printf("Visit http://localhost:%s for Hello World\n", port)
    fmt.Printf("Visit http://localhost:%s/health for health check\n", port)

    // Start server
    log.Fatal(http.ListenAndServe(":"+port, router))
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
    })
}
