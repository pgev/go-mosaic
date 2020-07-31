package threads

import (
	"crypto/sha256"
	"fmt"

	"github.com/mosaicdao/go-mosaic/libs/service"
)

// Switch defines a switch struct.
type Switch struct {
	service.BaseService

	reactorsByLocus map[string]Reactor
	reactorsByBoard map[string][]Reactor
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
	boardId := reactor.BoardId()

	for _, topicId := range reactor.GetTopicIds() {
		locus := hashLocus(boardId, topicId)

		// No two reactors can share the same topic.
		if sw.reactorsByLocus[locus] != nil {
			panic(
				fmt.Sprintf(
					"There is already a reactor (%v) registered for the board/topic pair %X/%X",
					sw.reactorsByLocus[locus],
					boardId,
					topicId,
				),
			)
		}

		sw.reactorsByLocus[locus] = reactor
	}

	boardIdHash := hashBoardId(boardId)
	sw.reactorsByBoard[boardIdHash] = append(sw.reactorsByBoard[boardIdHash], reactor)

	reactor.SetSwitch(sw)
}

func (sw *Switch) AddSource(source *Source) {

	boardIDHash := hashBoardId(source.BoardId)

	for _, reactor := range sw.reactorsByBoard[boardIDHash] {
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

	for _, reactor := range sw.reactorsByBoard[boardIDHash] {
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

func hashLocus(boardId BoardId, topicId TopicId) string {
	h := sha256.Sum256(append(boardId, byte(topicId)))
	return string(h[:])
}

func hashBoardId(boardId BoardId) string {
	h := sha256.Sum256(boardId)
	return string(h[:])
}
