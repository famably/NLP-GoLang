package llm

import "strconv"

type MockClient struct{}

func NewMockClient() *MockClient {
    return &MockClient{}
}

func (c *MockClient) Summarize(documents []string) (string, error) {
    if len(documents) == 0 {
        return "", nil
    }
    // Use strconv to properly convert int to string
    return "Mock summary of " + strconv.Itoa(len(documents)) + " documents", nil
}