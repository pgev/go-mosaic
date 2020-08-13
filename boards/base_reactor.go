package boards

import (
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type BaseReactor struct {
	service.BaseService

	boardID BoardID
	Switch  *Switch
}

func NewBaseReactor(
	name string, impl service.Servicable, boardID BoardID,
) *BaseReactor {
	return &BaseReactor{
		BaseService: *service.NewBaseService(name, impl),
		boardID:     boardID,
		Switch:      nil,
	}
}

func (*BaseReactor) InitSource(*Source) error { return nil }
func (*BaseReactor) AddSource(*Source) error  { return nil }

func (*BaseReactor) ReceiveMsg(*Message) {}

func (self *BaseReactor) Board() BoardID  { return self.boardID }
func (*BaseReactor) GetTopics() []TopicID { return nil }

func (self *BaseReactor) SetSwitch(sw *Switch) { self.Switch = sw }
