package threads

// TopicID defines type of id of topic.
type TopicID byte

// BoardID defines representation of Board identifier
type BoardID string

type Message struct {
	Sender       *Sender
	BoardID      BoardID
	TopicID      TopicID
	PayloadBytes []byte
}
