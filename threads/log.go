package threads

type Log interface {
}

type Message struct {
	TopicID TopicID
	Payload []byte
}
