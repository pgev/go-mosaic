package landscape

import (
	"github.com/mosaicdao/go-mosaic/threads"
)

// SubscriptionFilter
type SubscriptionFilter struct {
	boardIDs []threads.BoardID
}

type SubscriptionOption func(*SubscriptionFilter)

func WithBoardFilter(boardIDs []threads.BoardID) SubscriptionOption {
	return func(subFilter *SubscriptionFilter) {
		subFilter.boardIDs = append(subFilter.boardIDs, boardIDs...)
	}
}
