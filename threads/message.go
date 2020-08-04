package threads

// TopicID defines type of id of topic.
type TopicID byte

// BoardID defines type of id of board.
type BoardID []byte

type Message struct {
	Sender       *Sender
	BoardID      BoardID
	TopicID      TopicID
	PayloadBytes []byte
}
