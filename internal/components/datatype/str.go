package datatype

import "fmt"

const OCDBStringName = "OCDB-String"

type OCDBString struct {
	// storeData []OcdbVChar
	value string
	OCDBData
}

func NewOCDBString(value string) *OCDBString {

	coords := []uint32{uint32(randInt(30, 50)), YAxisString, ZAxisDataTypes}
	ocdbData := NewOCDBData(
		&OCDBDataOpts{
			Name:   OCDBStringName,
			Coords: NewCoordinates(&CoordsOpts{coords: coords}),
		},
	)

	storeStructure := transformToStoreData(value)
	for _, char := range storeStructure {
		ocdbData.dataModules = append(ocdbData.dataModules, char)
	}

	return &OCDBString{
		value: value,
		// storeData: storeStructure,
		OCDBData: *ocdbData,
	}
}

func (str *OCDBString) String() string {
	return fmt.Sprintf(" <%s> { %s } %s \n %s \n _______\n \n%v\n", str.name, str.value, str.coords, str.ShowStoreStructure(), str.dataModules)
}

func transformToStoreData(value string) []OcdbVChar {
	blocksize := 5
	sliceLen := len(value) / blocksize
	blocks := make([][5]byte, sliceLen)

	for i := 0; i < sliceLen; i++ {

		from, to := i*blocksize, (i*blocksize)+blocksize
		strBlock := value[from:to]

		copy(blocks[i][:], strBlock)
	}

	vChars := make([]OcdbVChar, len(blocks))
	for i, block := range blocks {
		vChars[i] = *NewOcdbVChar(block)
	}

	return vChars
}
