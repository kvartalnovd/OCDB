package datatypeint

import "github.com/kvartalnovd/OCDB/internal/components/datatype/coordsint"

type DBDataType interface {
	// Name() string
	// Value() any
	// String() string
	Coords() coordsint.Coords
	DataModules() []DBDataType
}
