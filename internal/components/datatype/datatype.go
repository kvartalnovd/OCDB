package datatype

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kvartalnovd/OCDB/internal/components/datatype/coordsint"
	"github.com/kvartalnovd/OCDB/internal/components/datatype/datatypeint"
)

const (
	ZAxisDataTypes = 4

	YAxisString = 56
	YAxisVChar  = 13
)

type OCDBDataOpts struct {
	Name        string
	Coords      coordsint.Coords
	DataModules []datatypeint.DBDataType
}

type OCDBData struct {
	name        string
	coords      coordsint.Coords
	dataModules []datatypeint.DBDataType
}

func NewOCDBData(opts *OCDBDataOpts) *OCDBData {
	return &OCDBData{
		coords:      opts.Coords, // NewCoordinates(x, yAxis, ZAxisDataTypes)
		name:        opts.Name,
		dataModules: opts.DataModules,
	}
}

func (d OCDBData) Name() string { return d.name }

func (d OCDBData) DataModules() []datatypeint.DBDataType {
	return d.dataModules
}

func (d OCDBData) Coords() coordsint.Coords {
	return d.coords
}

func (d OCDBData) ShowStoreStructure() string {

	modulesCoordsStruct := make(map[uint32]map[uint32][]uint32, len(d.dataModules))

	for _, module := range d.dataModules {

		yAxis, ok := modulesCoordsStruct[module.Coords().Z()]
		if !ok {
			yAxis = make(map[uint32][]uint32, 0)
		}

		xAxis, ok := yAxis[module.Coords().Y()]
		if !ok {
			xAxis = make([]uint32, 0)
		}

		xAxis = append(xAxis, module.Coords().X())

		yAxis[module.Coords().Y()] = xAxis
		modulesCoordsStruct[module.Coords().Z()] = yAxis
	}

	result := ""
	for z, yBlock := range modulesCoordsStruct {

		result = fmt.Sprintf("%s%d", result, z)

		for y, xBlock := range yBlock {
			result = fmt.Sprintf("%s:%d:%s", result, y, strings.Join(map2(xBlock, func(x uint32) string { return strconv.Itoa(int(x)) }), " "))
		}

		result = fmt.Sprintf("%s; \n %v", result, []byte(result))
	}

	return result
}

func map2[T, U any](data []T, f func(T) U) []U {

	res := make([]U, 0, len(data))

	for _, e := range data {
		res = append(res, f(e))
	}

	return res
}
