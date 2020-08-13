package scout

import (
	"context"

	p2ppeer "github.com/libp2p/go-libp2p-core/peer"
	lnd "github.com/mosaicdao/go-mosaic/landscape"
	"github.com/mosaicdao/go-mosaic/libs/service"
	brd "github.com/mosaicdao/go-mosaic/boards"
)

type Scout interface {
	service.Service
}

type scout struct {
	service.BaseService

	cancel context.CancelFunc

	landscape lnd.Landscape
	boards    brd.BoardsManager
	peerID    p2ppeer.ID
	policy    []brd.BoardID
}

func NewScout(
	ctx context.Context,
	landscape lnd.Landscape,
	boards brd.BoardsManager,
) (
	Scout, error,
) {
	s := &scout{
		landscape: landscape,
		boards:    boards,
		peerID:    boards.Host().ID(),
		policy:    nil,
	}
	s.BaseService = *service.NewBaseService("scout", s)
	return s, nil
}

func (s *scout) OnStart() error {

	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	err := s.updatePolicy()
	if err != nil {
		s.cancel()
		return err
	}

	go s.doScouting(ctx)

	return nil
}

func (s *scout) OnStop() {
	s.cancel()
}

//------------------------------------------------------------------------------
// Private functions

func (s *scout) updatePolicy() error {
	// TODO: this is really stupid, we re just appending for now
	newBoards, err := s.landscape.GetAssignments(s.peerID)
	if err != nil {
		return err
	}
	s.policy = append(s.policy, newBoards...)
	return nil
}

func (s *scout) doScouting(ctx context.Context) {

	boardFilter := s.deriveBoardFilter()
	listener := s.landscape.Subscribe(ctx, boardFilter)

	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-listener:
			if !ok {
				return
			}
			switch event.(type) {
			case *lnd.SourceChangeEvent:
				// TODO: continue to use events to trigger actions on boards
			case *lnd.PeerInfoUpdateEvent:
			}
		}
	}
}

func (s *scout) deriveBoardFilter() lnd.SubscriptionOption {
	// TODO: expand on policy
	p := make([]brd.BoardID, len(s.policy))
	// TODO: should we deep copy?
	copy(p, s.policy)
	return lnd.WithBoardFilter(p)
}
