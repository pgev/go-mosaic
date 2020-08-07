package scout

import (
	"context"

	p2ppeer "github.com/libp2p/go-libp2p-core/peer"
	lnd "github.com/mosaicdao/go-mosaic/landscape"
	"github.com/mosaicdao/go-mosaic/libs/service"
	thr "github.com/mosaicdao/go-mosaic/threads"
)

type Scout interface {
	service.Service
}

type scout struct {
	service.BaseService

	cancel context.CancelFunc

	landscape lnd.Landscape
	threads   thr.ThreadsNetwork
	peerID    p2ppeer.ID
	policy    []thr.BoardID
}

func NewScout(
	ctx context.Context,
	landscape lnd.Landscape,
	threads thr.ThreadsNetwork,
) (
	Scout, error,
) {
	s := &scout{
		landscape: landscape,
		threads:   threads,
		peerID:    threads.Host().ID(),
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

			case *lnd.PeerInfoUpdateEvent:
			}
		}
	}

	//old stuff
	// boardIDs := s.landscape.GetAssignments(s.peerID)

	// var peers []p2ppeer.AddrInfo
	// var peerIDtoPeer map[p2ppeer.ID]p2peer.AddrInfo

	// for _, boardID := range boardIDs {
	// 	newPeers = s.landscape.GetPeers(boardID)
	// 	for _, peer := range newPeers
	// }
}

func (s *scout) deriveBoardFilter() lnd.SubscriptionOption {
	// TODO: expand on policy
	p := make([]thr.BoardID, len(s.policy))
	// TODO: should we deep copy?
	copy(p, s.policy)
	return lnd.WithBoardFilter(p)
}
