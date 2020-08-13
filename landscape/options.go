package landscape

import (
	"github.com/mosaicdao/go-mosaic/boards"
)

// SubscriptionFilter
type SubscriptionFilter struct {
	boardIDs []boards.BoardID
}

type SubscriptionOption func(*SubscriptionFilter)

func WithBoardFilter(boardIDs []boards.BoardID) SubscriptionOption {
	return func(subFilter *SubscriptionFilter) {
		subFilter.boardIDs = append(subFilter.boardIDs, boardIDs...)
	}
}
