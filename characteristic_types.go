package calibrationReader

import "github.com/asap2Go/calibrationReader/a2l"

/*implements the different types of characteristics that can be parsed from a2l
and be filled with byte values from the hex-file.*/

type CharacteristicType interface {
	getValue() interface{}
}

type AsciiCharacteristic struct {
	Characteristic a2l.Characteristic
	value          string
}

func (t *AsciiCharacteristic) getValue(*map[uint32]byte) interface{} {
	return t.value
}

type CurveCharacteristic struct {
	Characteristic a2l.Characteristic
	value          [][]interface{}
}

func (t *CurveCharacteristic) getValue(*map[uint32]byte) interface{} {
	return t.value
}

type CuboidCharacteristic struct {
	Characteristic a2l.Characteristic
	value          [][][]interface{}
}

func (t *CuboidCharacteristic) getValue(*map[uint32]byte) interface{} {
	return t.value
}

type Cube4Characteristic struct {
	Characteristic a2l.Characteristic
	value          [][][][]interface{}
}

func (t *Cube4Characteristic) getValue(*map[uint32]byte) interface{} {
	return t.value
}

type Cube5Characteristic struct {
	Characteristic a2l.Characteristic
	value          [][][][][]interface{}
}

func (t *Cube5Characteristic) getValue(*map[uint32]byte) interface{} {
	return t.value
}

type ValBlkCharacteristic struct {
	Characteristic a2l.Characteristic
	value          []interface{}
}

func (t *ValBlkCharacteristic) getValue(*map[uint32]byte) interface{} {
	return t.value
}

type ValueCharacteristic struct {
	Characteristic a2l.Characteristic
	value          interface{}
}

func (t *ValueCharacteristic) getValue(*map[uint32]byte) interface{} {
	return t.value
}

type DerivedCharacteristic struct {
	Characteristic a2l.Characteristic
	value          interface{}
}

func (t *DerivedCharacteristic) getValue(*map[uint32]byte) interface{} {
	return t.value
}

type extendedSICharacteristic struct {
	Characteristic a2l.Characteristic
	value          interface{}
}

func (t *extendedSICharacteristic) getValue(*map[uint32]byte) interface{} {
	return t.value
}
