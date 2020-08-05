package landscape

import (
	"github.com/mosaicdao/go-mosaic/threads"
)

type SubscriptionFilter struct {
	boardIDs []threads.BoardID
}

type SubscriptionOption func(*SubscriptionFilter)

func WithBoardFilter(boardID threads.BoardID) SubscriptionOption {
	return func(subFilter *SubscriptionFilter) {
		subFilter.boardIDs = append(subFilter.boardIDs, boardID)
	}
}
