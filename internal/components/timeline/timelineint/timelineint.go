package timelineint

import "github.com/google/uuid"

type Timeline interface {
	Tag() uuid.UUID
	String() string
}
