package threads

import (
	"fmt"

	"github.com/mosaicdao/go-mosaic/libs/service"
)

// Switch defines a switch struct.
type Switch struct {
	service.BaseService

	topics            []*Topic
	reactors          map[string]Reactor
	reactorsByTopicID map[TopicID]Reactor
}

// OnStart implements Servicable.OnStart() by starting the switch and all registered reactors.
// The function starts all registered reactors sequentially and returns an error and stops
// if one fails to start. It does not stops the ones already started.
func (sw *Switch) OnStart() error {
	// starts reactors
	for _, reactor := range sw.reactors {
		err := reactor.Start()
		if err != nil {
			return fmt.Errorf("failed to start %v: %w", reactor, err)
		}
	}

	return nil
}

// OnStop implements Servicable.OnStop() by stopping the switch and all registered reactors.
func (sw *Switch) OnStop() {
	// stops reactors
	for _, reactor := range sw.reactors {
		reactor.Stop()
	}
}

// AddReactor adds a reactor to the switch.
// The function requires there is no reactor with the same name.
// The function updates a mapping from a topic id to a reactor based on the reactor's topics.
// The function updates a mapping from the given reactor name to the reactor.
// The function requires that no two reactors can share the same topic.
// The function sets the current object as a switch to the given reactor.
func (sw *Switch) AddReactor(name string, reactor Reactor) {
	if sw.reactors[name] != nil {
		panic(
			fmt.Sprintf(
				"There is already a reactor (%v) registered with the same name %v",
				sw.reactors[name],
				name,
			),
		)
	}

	for _, topic := range reactor.GetTopics() {
		topicID := topic.ID

		// No two reactors can share the same topic.
		if sw.reactorsByTopicID[topicID] != nil {
			panic(
				fmt.Sprintf(
					"There is already a reactor (%v) registered for the topic %X",
					sw.reactorsByTopicID[topicID],
					topicID,
				),
			)
		}

		sw.topics = append(sw.topics, topic)
		sw.reactorsByTopicID[topicID] = reactor
	}

	sw.reactors[name] = reactor
	reactor.SetSwitch(sw)
}

// RemoveReactor removes the given reactor from the switch.
// The function requires that there the given reactor is registered under the given name.
// The function updates a mapping from a topic id to a reactor based on the reactor's topics.
// The function sets the given reactor's switch to nil.
func (sw *Switch) RemoveReactor(name string, reactor Reactor) {
	if sw.reactors[name] == nil {
		panic(
			fmt.Sprintf(
				"There is no reactor with the given name %v",
				name,
			),
		)
	}

	if sw.reactors[name] != reactor {
		panic(
			fmt.Sprintf(
				"There is a different reactor (%v) registered with the given name (%v)",
				sw.reactors[name],
				name,
			),
		)
	}

	for _, topic := range reactor.GetTopics() {
		// removes topic
		for i := 0; i < len(sw.topics); i++ {
			if topic.ID == sw.topics[i].ID {
				sw.topics = append(sw.topics[:i], sw.topics[i+1:]...)
				break
			}
		}
		delete(sw.reactorsByTopicID, topic.ID)
	}
	delete(sw.reactors, name)
	reactor.SetSwitch(nil)
}

// Reactors returns a mapping of reactors by a registered name.
func (sw *Switch) Reactors() map[string]Reactor {
	return sw.reactors
}

// Reactor returns a registered reactor by a name or nil if there is no one.
func (sw *Switch) Reactor(name string) Reactor {
	return sw.reactors[name]
}
