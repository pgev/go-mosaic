package boardstore

import (
	"github.com/mosaicdao/go-mosaic/boards"
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type BoardStore interface {
	service.Service

	BoardMetadata

	ViewBook
	SliceBook

	// Boards returns all of the board IDs stored
	Boards() boards.BoardIDSlice
}

type ViewBook interface {

}

type SliceBook interface {

}

type BoardMetadata interface {

}
