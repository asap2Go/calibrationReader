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

func (*AsciiCharacteristic) getValue(*map[uint32]byte) interface{} {
	return ""
}

type CurveCharacteristic struct {
	Characteristic a2l.Characteristic
	value          [][]interface{}
}

func (*CurveCharacteristic) getValue(*map[uint32]byte) interface{} {
	return ""
}

type CuboidCharacteristic struct {
	Characteristic a2l.Characteristic
	value          [][][]interface{}
}

func (*CuboidCharacteristic) getValue(*map[uint32]byte) interface{} {
	return ""
}

type Cube4Characteristic struct {
	Characteristic a2l.Characteristic
	value          [][][][]interface{}
}

func (*Cube4Characteristic) getValue(*map[uint32]byte) interface{} {
	return ""
}

type Cube5Characteristic struct {
	Characteristic a2l.Characteristic
	value          [][][][][]interface{}
}

func (*Cube5Characteristic) getValue(*map[uint32]byte) interface{} {
	return ""
}

type ValBlkCharacteristic struct {
	Characteristic a2l.Characteristic
	value          []interface{}
}

func (*ValBlkCharacteristic) getValue(*map[uint32]byte) interface{} {
	return ""
}

type ValueCharacteristic struct {
	Characteristic a2l.Characteristic
	value          interface{}
}

func (*ValueCharacteristic) getValue(*map[uint32]byte) interface{} {
	return ""
}

type DerivedCharacteristic struct {
	Characteristic a2l.Characteristic
	value          interface{}
}

func (*DerivedCharacteristic) getValue(*map[uint32]byte) interface{} {
	return ""
}

type extendedSICharacteristic struct {
	Characteristic a2l.Characteristic
	value          interface{}
}

func (*extendedSICharacteristic) getValue(*map[uint32]byte) interface{} {
	return ""
}
