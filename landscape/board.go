package landscape

import (
	"github.com/mosaicdao/go-mosaic/threads"
)

type SourceChange struct {
	BoardID        threads.BoardID
	addedSources   []threads.LogID
	removedSources []threads.LogID
}

type BoardLog struct {
	BoardID threads.BoardID
}
