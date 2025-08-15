package api

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
    "github.com/famably/chiplens/internal/llm"
    "github.com/famably/chiplens/internal/storage"
)

type Handler struct {
    storage storage.Storage
    llm     llm.Client
}

func NewHandler(storage storage.Storage, llm llm.Client) *Handler {
    return &Handler{
        storage: storage,
        llm:     llm,
    }
}

func (h *Handler) StoreDocument(c *gin.Context) {
    groupID := c.Param("group_id")
    
    var req struct {
        Content string `json:"content"`
    }
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    id, err := h.storage.StoreDocument(groupID, req.Content)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) GetDocuments(c *gin.Context) {
    groupID := c.Param("group_id")
    
    docs, err := h.storage.GetDocumentsByGroup(groupID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, docs)
}

func (h *Handler) SummarizeDocuments(c *gin.Context) {
    groupID := c.Param("group_id")
    
    docs, err := h.storage.GetDocumentsByGroup(groupID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    if len(docs) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "no documents found for group"})
        return
    }
    
    var contents []string
    for _, doc := range docs {
        contents = append(contents, doc.Content)
    }
    
    summary, err := h.llm.Summarize(contents)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate summary"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"summary": summary})
}