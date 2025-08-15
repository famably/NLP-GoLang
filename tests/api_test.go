package tests

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/famably/chiplens/internal/api"
    "github.com/famably/chiplens/internal/llm"
    "github.com/famably/chiplens/internal/storage"
)

func TestAPIEndpoints(t *testing.T) {
    // Setup
    store := storage.NewInMemoryStorage()
    mockLLM := llm.NewMockClient()
    handler := api.NewHandler(store, mockLLM)
    
    router := gin.Default()
    api.SetupRoutes(router, handler)
    
    t.Run("Store and retrieve documents", func(t *testing.T) {
        // Store document
        w := httptest.NewRecorder()
        req, _ := http.NewRequest("POST", "/groups/test-group/documents", strings.NewReader(`{"content": "test content"}`))
        router.ServeHTTP(w, req)
        
        assert.Equal(t, http.StatusCreated, w.Code)
        
        // Retrieve documents
        w = httptest.NewRecorder()
        req, _ = http.NewRequest("GET", "/groups/test-group/documents", nil)
        router.ServeHTTP(w, req)
        
        assert.Equal(t, http.StatusOK, w.Code)
        assert.Contains(t, w.Body.String(), "test content")
    })
    
    t.Run("Summarize documents", func(t *testing.T) {
        // First store a document
        w := httptest.NewRecorder()
        req, _ := http.NewRequest("POST", "/groups/summary-group/documents", strings.NewReader(`{"content": "summary test"}`))
        router.ServeHTTP(w, req)
        
        // Get summary
        w = httptest.NewRecorder()
        req, _ = http.NewRequest("GET", "/groups/summary-group/summary", nil)
        router.ServeHTTP(w, req)
        
        assert.Equal(t, http.StatusOK, w.Code)
        assert.Contains(t, w.Body.String(), "Mock summary")
    })
}