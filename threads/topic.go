package threads

// TopicID represents a type of a Topic.
type TopicID byte

// Topic groups messages communicated between members and past users.
type Topic struct {
	ID TopicID
}
