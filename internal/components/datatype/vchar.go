package datatype

import (
	"fmt"
	"hash/fnv"
)

type OcdbVChar struct {
	value [5]byte
	OCDBData
}

func NewOcdbVChar(value [5]byte) *OcdbVChar {
	x := hash(string(value[:]))
	coords := []uint32{x, YAxisVChar, ZAxisDataTypes}
	return &OcdbVChar{
		value: value,
		OCDBData: *NewOCDBData(
			&OCDBDataOpts{
				Name:   "OCDB-V-Char",
				Coords: NewCoordinates(&CoordsOpts{coords: coords}),
			},
		),
	}
}

func (v OcdbVChar) String() string {
	return fmt.Sprintf(" <%s> %s { %s | %v }\n", v.name, v.coords, v.value, v.value)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
