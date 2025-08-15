package storage

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"sync"
	"time"
)

func generateID() string {
	b := make([]byte, 8) // Generates 16-character hex string
	if _, err := rand.Read(b); err != nil {
		// Fallback if crypto/rand fails
		return "fallback-" + strconv.FormatInt(time.Now().UnixNano(), 36)
	}
	return hex.EncodeToString(b)
}

type InMemoryStorage struct {
    documents map[string]Document
    mu        sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
    return &InMemoryStorage{
        documents: make(map[string]Document),
    }
}

func (s *InMemoryStorage) StoreDocument(groupID, content string) (string, error) {
    doc := Document{
        ID:      generateID(),
        GroupID: groupID,
        Content: content,
    }
    
    s.mu.Lock()
    defer s.mu.Unlock()
    s.documents[doc.ID] = doc
    
    return doc.ID, nil
}

func (s *InMemoryStorage) GetDocumentsByGroup(groupID string) ([]Document, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    var docs []Document
    for _, doc := range s.documents {
        if doc.GroupID == groupID {
            docs = append(docs, doc)
        }
    }
    
    return docs, nil
}

func (s *InMemoryStorage) DocumentExists(id string) bool {
    s.mu.RLock()
    defer s.mu.RUnlock()
    _, exists := s.documents[id]
    return exists
}