package boards

// TopicID defines type of id of topic.
type TopicID byte

// BoardID defines representation of Board identifier
type BoardID string

type Message struct {
	Source       *Source
	TopicID      TopicID
	PayloadBytes []byte
}
