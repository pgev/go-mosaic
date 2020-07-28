package threads

type TopicID byte

// Topic groups messages communicated between members and past users.
type Topic struct {
	ID TopicID
}
