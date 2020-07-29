package threads

// TopicID defines type of id of topic.
type TopicID []byte

type Topic struct {
	ID TopicID
}

// BoardID defines type of id of board.
type BoardID []byte

type Board struct {
	ID BoardID
}

type Message struct {
	Sender       *Sender
	Board        *Board
	Topic        *Topic
	PayloadBytes []byte
}

func NewMessage(
	sender *Sender, board *Board, topic *Topic, payloadBytes []byte,
) *Message {
	return &Message{
		Sender:       sender,
		Board:        board,
		Topic:        topic,
		PayloadBytes: payloadBytes,
	}
}
