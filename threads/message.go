package threads

// TopicId defines type of id of topic.
type TopicId byte

// BoardId defines type of id of board.
type BoardId []byte

type Message struct {
	Sender       *Sender
	BoardId      BoardId
	TopicId      TopicId
	PayloadBytes []byte
}
