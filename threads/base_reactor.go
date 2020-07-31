package threads

import (
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type BaseReactor struct {
	service.BaseService

	boardId BoardId
	Switch  *Switch
}

func NewBaseReactor(
	name string, impl service.Servicable, boardId BoardId,
) *BaseReactor {
	return &BaseReactor{
		BaseService: *service.NewBaseService(name, impl),
		boardId:     boardId,
		Switch:      nil,
	}
}

func (*BaseReactor) InitSource(*Source) error { return nil }
func (*BaseReactor) AddSource(*Source) error  { return nil }

func (*BaseReactor) ReceiveMsg(*Message) {}

func (self *BaseReactor) Board() BoardId  { return self.boardId }
func (*BaseReactor) GetTopics() []TopicId { return nil }

func (self *BaseReactor) SetSwitch(sw *Switch) { self.Switch = sw }
