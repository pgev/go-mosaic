package threads

import (
	"fmt"

	"github.com/mosaicdao/go-mosaic/libs/service"
)

// Switch defines a switch struct.
type Switch struct {
	service.BaseService

	topics            []Topic
	reactorsByTopicID map[TopicID]Reactor
	databuses         []Databus
}

// OnStart implements Servicable.OnStart() by starting the switch, reactors and databuses.
// The function starts all registered reactors sequentially and returns an error and stops
// if one fails to start. It does not stops the ones already started.
// The function starts all registered databuses sequentially and returns an error and stops
// if one fails to start. It does not stops the ones already started.
func (sw *Switch) OnStart() error {
	// starts reactors
	for _, reactor := range sw.reactorsByTopicID {
		err := reactor.Start()
		if err != nil {
			return fmt.Errorf("Failed to start reactor %v: %w", reactor, err)
		}
	}

	// starts databuses
	for _, databus := range sw.databuses {
		err := databus.Start()
		if err != nil {
			return fmt.Errorf("Failed to start databus %v: %w", databus, err)
		}
	}

	return nil
}

// OnStop implements Servicable.OnStop() by stopping the switch, registered reactors and databuses.
func (sw *Switch) OnStop() {
	// stops reactors
	for _, reactor := range sw.reactorsByTopicID {
		reactor.Stop()
	}

	// stops member message dispatcher
	for _, databus := range sw.databuses {
		databus.Stop()
	}
}

// AddReactor adds a reactor to the switch.
// The function updates a mapping from a topic id to a reactor based on the reactor's topics.
// The function requires that no two reactors can share the same topic.
// The function sets the current object as a switch to the given reactor.
func (sw *Switch) AddReactor(reactor Reactor) {
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

	reactor.SetSwitch(sw)
}

// RemoveReactor removes the given reactor from the switch.
// The function updates a mapping from a topic id to a reactor based on the reactor's topics.
// The function sets the given reactor's switch to nil.
func (sw *Switch) RemoveReactor(name string, reactor Reactor) {
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

	reactor.SetSwitch(nil)
}

// AddDatabus ...
// @todo: no new reactor can be added after adding first database
// @todo: can we add new database if switch is running (thread safety!)
func (sw *Switch) AddDatabus(databus Databus) {
	databus.InitReactors(sw.reactorsByTopicID)

	for _, reactor := range sw.reactorsByTopicID {
		databus = reactor.InitDatabus(databus)
	}

	for _, reactor := range sw.reactorsByTopicID {
		reactor.AddDatabus(databus)
	}

	sw.databuses = append(sw.databuses, databus)
}

func (sw *Switch) RemoveDatabus(databus Databus) {
	for _, reactor := range sw.reactorsByTopicID {
		reactor.RemoveDatabus(databus)
	}

	databus.Stop()

	for i := 0; i < len(sw.databuses); i++ {
		if databus.compareByID(sw.databuses[i]) {
			sw.databuses = append(sw.databuses[:i], sw.databuses[i+1:]...)
		}
	}
}
