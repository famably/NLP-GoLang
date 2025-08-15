package main

import (
    "log"
    
    "github.com/gin-gonic/gin"
    "github.com/famably/chiplens/internal/api"
    "github.com/famably/chiplens/internal/llm"
    "github.com/famably/chiplens/internal/storage"
)

func main() {
    // Initialize dependencies
    store := storage.NewInMemoryStorage()
    llmClient := llm.NewMockClient()
    
    // Create API handler
    handler := api.NewHandler(store, llmClient)
    
    // Setup router
    router := gin.Default()
    api.SetupRoutes(router, handler)
    
    // Start server
    log.Println("Server starting on :8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}