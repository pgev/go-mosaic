package landscape

import (
	"github.com/mosaicdao/go-mosaic/boards"
)

type SourceChange struct {
	BoardID        boards.BoardID
	addedSources   []boards.LogID
	removedSources []boards.LogID
}

type BoardLog struct {
	BoardID boards.BoardID
}
