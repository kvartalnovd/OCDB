package coordsint

import (
	"github.com/kvartalnovd/OCDB/internal/components/timeline/timelineint"
)

type Coords interface {
	X() uint32
	Y() uint32
	Z() uint32
	T() timelineint.Timeline
}
