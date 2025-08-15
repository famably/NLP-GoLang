package llm

type Client interface {
    Summarize(documents []string) (string, error)
}