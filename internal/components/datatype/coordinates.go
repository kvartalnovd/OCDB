package datatype

import (
	"fmt"

	"github.com/kvartalnovd/OCDB/internal/components/timeline"
	"github.com/kvartalnovd/OCDB/internal/components/timeline/timelineint"
)

type CoordsOpts struct {
	coords   []uint32
	timeline timeline.Timeline
}

type Coordinates struct {
	x uint32
	y uint32
	z uint32
	t timeline.Timeline
}

func NewCoordinates(opts *CoordsOpts) *Coordinates {
	var x, y, z uint32 = 0, 0, 0

	if len(opts.coords) >= 3 {
		z = opts.coords[2]
	}

	if len(opts.coords) >= 2 {
		y = opts.coords[1]
	}

	if len(opts.coords) >= 1 {
		x = opts.coords[0]
	}

	return &Coordinates{x, y, z, *timeline.NewTimeline()}
}

func (c *Coordinates) X() uint32               { return c.x }
func (c *Coordinates) Y() uint32               { return c.y }
func (c *Coordinates) Z() uint32               { return c.z }
func (c *Coordinates) T() timelineint.Timeline { return c.t }

func (c Coordinates) String() string {
	return fmt.Sprintf("(%d, %d, %d)", c.x, c.y, c.z)
}
