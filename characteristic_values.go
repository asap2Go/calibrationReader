package calibrationReader

import (
	"github.com/asap2Go/calibrationReader/a2l"
)

type CharacteristicValues struct {
	characteristic *a2l.Characteristic
	recordLayout   *a2l.RecordLayout
	AxisX          interface{}
	AxisY          interface{}
	AxisZ          interface{}
	Axis4          interface{}
	Axis5          interface{}
	valuesBin      interface{}
	valuesHex      interface{}
	valuesDec      interface{}
	valuesPhy      interface{}
}

func (cv *CharacteristicValues) getCharacteristicValueDecimal() (interface{}, error) {
	return nil, nil
}

func (cv *CharacteristicValues) getCharacteristicValuePhysical() (interface{}, error) {
	return nil, nil
}

func (cv *CharacteristicValues) getCharacteristicValueHex() (interface{}, error) {
	return nil, nil
}
