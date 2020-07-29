package threads

import (
	"crypto/sha256"
	"fmt"

	"github.com/mosaicdao/go-mosaic/libs/service"
)

// Switch defines a switch struct.
type Switch struct {
	service.BaseService

	reactorsByBoardMsgStream map[string]Reactor
	reactorsByBoard          map[string][]Reactor
	reactors                 []Reactor

	sources []Source
}

func (sw *Switch) OnStart() error {
	// starts reactors
	for _, reactor := range sw.reactors {
		err := reactor.Start()
		if err != nil {
			panic(
				fmt.Sprintf(
					"Reactor '%v'failed to start: '%v'",
					reactor,
					err,
				),
			)
		}
	}

	// starts sources
	for _, source := range sw.sources {
		err := source.Start()
		if err != nil {
			panic(
				fmt.Sprintf(
					"Source '%v'failed to start: '%v'",
					source,
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

	// stops sources
	for _, source := range sw.sources {
		source.Stop()
	}
}

// AddReactor adds a reactor to the switch.
func (sw *Switch) AddReactor(reactor Reactor) {
	boardID := reactor.Board().ID

	for _, topic := range reactor.GetTopics() {
		topicID := topic.ID

		boardMsgStreamHash := hashBoardMsgStream(boardID, topicID)

		// No two reactors can share the same topic.
		if sw.reactorsByBoardMsgStream[boardMsgStreamHash] != nil {
			panic(
				fmt.Sprintf(
					"There is already a reactor (%v) registered for the board/topic pair %X/%X",
					sw.reactorsByBoardMsgStream[boardMsgStreamHash],
					boardID,
					topicID,
				),
			)
		}

		sw.reactorsByBoardMsgStream[boardMsgStreamHash] = reactor
	}

	boardIDHash := hashBoardID(boardID)
	sw.reactorsByBoard[boardIDHash] = append(sw.reactorsByBoard[boardIDHash], reactor)

	reactor.SetSwitch(sw)
}

func (sw *Switch) AddSource(source Source) {

	boardIDHash := hashBoardID(source.Board().ID)

	for _, reactor := range sw.reactorsByBoard[boardIDHash] {
		err, source := reactor.InitSource(source)
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

	for _, reactor := range sw.reactorsByBoard[boardIDHash] {
		err := reactor.AddSource(source)
		if err != nil {
			panic(
				fmt.Sprintf(
					"Reactor '%v'failed to add source '%v'",
					reactor,
					source,
				),
			)
		}
	}

	source.InitReactors(sw.reactorsByBoardMsgStream)

	sw.sources = append(sw.sources, source)
}

func hashBoardMsgStream(boardID BoardID, topicID TopicID) string {
	h := sha256.Sum256(append(boardID, topicID...))
	return string(h[:])
}

func hashBoardID(boardID BoardID) string {
	h := sha256.Sum256(boardID)
	return string(h[:])
}
