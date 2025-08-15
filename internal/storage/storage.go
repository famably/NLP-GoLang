package storage

type Document struct {
    ID      string
    GroupID string
    Content string
}

type Storage interface {
    StoreDocument(groupID, content string) (string, error)
    GetDocumentsByGroup(groupID string) ([]Document, error)
    DocumentExists(id string) bool
}