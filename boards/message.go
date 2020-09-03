package boards

// TopicID defines type of id of topic.
type TopicID byte

type Message struct {
	Source       *Source
	TopicID      TopicID
	PayloadBytes []byte
}
