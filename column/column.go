package column

import (
	"context"

	"github.com/mosaicdao/go-mosaic/types"
)

// Column is the central coordination object for members of a column
//   1. to publish their subjective view on the input batches from the past users,
//   2. to find an objective view among the members (a threshold of signatures from members),
//   3. to perform a computational task set by the gate (on the objective view of the input), and
//   4. to output a batch with the resulting key:value pairs
type Column struct {
	Members map[types.Address]*Member
}

func NewColumn(ctx context.Context) (*Column, error) {
	return nil, nil
}

// NewColumnFromAddress creates a new column from a given address.
func NewColumnFromAddress() (*Column, error) {
	return nil, nil
}

func newColumn() (*Column, error) {
	return nil, nil
}
