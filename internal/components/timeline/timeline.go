package timeline

import (
	"fmt"

	"github.com/google/uuid"
)

type Timeline struct {
	tag uuid.UUID
}

func NewTimeline() *Timeline {
	tagId := uuid.New()
	return &Timeline{
		tag: tagId,
	}
}

func (tl Timeline) Tag() uuid.UUID { return tl.tag }

func (tl Timeline) String() string {
	return fmt.Sprintf("#%s", tl.tag)
}
