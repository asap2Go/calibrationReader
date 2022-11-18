package calibrationReader

import (
	"github.com/asap2Go/calibrationReader/a2l"
)

type CharacteristicValues struct {
	characteristic  *a2l.Characteristic
	recordLayout    *a2l.RecordLayout
	AxisXValues     interface{}
	AxisYValues     interface{}
	AxisZValues     interface{}
	Axis4Values     interface{}
	Axis5Values     interface{}
	DistOpXValue    interface{}
	DistOpYValue    interface{}
	DistOpZValue    interface{}
	DistOp4Value    interface{}
	DistOp5Value    interface{}
	Identification  interface{}
	NoAxisPtsXValue interface{}
	NoAxisPtsYValue interface{}
	NoAxisPtsZValue interface{}
	NoAxisPts4Value interface{}
	NoAxisPts5Value interface{}
	NoRescaleXValue interface{}
	valuesBin       interface{}
	valuesHex       interface{}
	valuesDec       interface{}
	valuesPhy       interface{}
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
