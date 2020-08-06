package threads

import (
	"fmt"

	"github.com/mosaicdao/go-mosaic/libs/service"
)

// Switch defines a switch struct.
type Switch struct {
	service.BaseService

	reactorsByLocus map[string]Reactor
	reactorsByBoard map[BoardID][]Reactor
	reactors        []Reactor
}

func (sw *Switch) OnStart() error {
	// starts reactors
	for _, reactor := range sw.reactors {
		err := reactor.Start()
		if err != nil {
			panic(
				fmt.Sprintf(
					"Reactor '%v' failed to start: '%v'",
					reactor,
					err,
				),
			)
		}
	}

	return nil
}

func (sw *Switch) OnStop() {
	// stops reactors
	for _, reactor := range sw.reactors {
		reactor.Stop()
	}
}

// AddReactor adds a reactor to the switch.
func (sw *Switch) AddReactor(reactor Reactor) {
	boardID := reactor.BoardID()

	for _, topicID := range reactor.GetTopicIDs() {
		locus := hashLocus(boardID, topicID)

		// No two reactors can share the same topic.
		if sw.reactorsByLocus[locus] != nil {
			panic(
				fmt.Sprintf(
					"There is already a reactor (%v) registered for the board/topic pair %X/%X",
					sw.reactorsByLocus[locus],
					boardID,
					topicID,
				),
			)
		}

		sw.reactorsByLocus[locus] = reactor
	}

	sw.reactorsByBoard[boardID] = append(sw.reactorsByBoard[boardID], reactor)

	reactor.SetSwitch(sw)
}

func (sw *Switch) AddSource(source *Source) {

	for _, reactor := range sw.reactorsByBoard[source.BoardID] {
		err := reactor.InitSource(source)
		if err != nil {
			panic(
				fmt.Sprintf(
					"Reactor '%v' failed to init source '%v'",
					reactor,
					source,
				),
			)
		}
	}

	for _, reactor := range sw.reactorsByBoard[source.BoardID] {
		err := reactor.AddSource(source)
		if err != nil {
			panic(
				fmt.Sprintf(
					"Reactor '%v' failed to add source '%v'",
					reactor,
					source,
				),
			)
		}
	}
}

func hashLocus(boardID BoardID, topicID TopicID) string {
	return string(boardID) + "/" + string(topicID)
}
