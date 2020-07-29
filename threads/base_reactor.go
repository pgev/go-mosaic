package threads

import (
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type BaseReactor struct {
	service.BaseService

	board  *Board
	Switch *Switch
}

func NewBaseReactor(
	name string, impl service.Servicable, board *Board,
) *BaseReactor {
	return &BaseReactor{
		BaseService: *service.NewBaseService(name, impl),
		board:       board,
		Switch:      nil,
	}
}

func (*BaseReactor) InitSource(source Source) (Source, error) { return source, nil }
func (*BaseReactor) AddSource(Source) error                   { return nil }

func (*BaseReactor) ReceiveMsg(*Message) {}

func (r *BaseReactor) Board() *Board     { return r.board }
func (*BaseReactor) GetTopics() []*Topic { return nil }

func (r *BaseReactor) SetSwitch(sw *Switch) { r.Switch = sw }
