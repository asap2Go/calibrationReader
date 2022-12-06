package calibrationReader

import (
	"github.com/asap2Go/calibrationReader/a2l"
)

type CharacteristicValues struct {
	characteristic      *a2l.Characteristic
	recordLayout        *a2l.RecordLayout
	AxisXValues         interface{}
	AxisYValues         interface{}
	AxisZValues         interface{}
	Axis4Values         interface{}
	Axis5Values         interface{}
	distOpXValue        interface{}
	distOpYValue        interface{}
	distOpZValue        interface{}
	distOp4Value        interface{}
	distOp5Value        interface{}
	identificationValue interface{}
	noAxisPtsXValue     interface{}
	noAxisPtsYValue     interface{}
	noAxisPtsZValue     interface{}
	noAxisPts4Value     interface{}
	noAxisPts5Value     interface{}
	noRescaleXValue     interface{}
	offsetXValue        interface{}
	offsetYValue        interface{}
	offsetZValue        interface{}
	offset4Value        interface{}
	offset5Value        interface{}
	shiftOpXValue       interface{}
	shiftOpYValue       interface{}
	shiftOpZValue       interface{}
	shiftOp4Value       interface{}
	shiftOp5Value       interface{}
	ValuesDec           interface{}
	ValuesPhy           interface{}
}

func (cv *CharacteristicValues) getCharacteristicValueDecimal() (interface{}, error) {
	return nil, nil
}

func (cv *CharacteristicValues) getCharacteristicValuePhysical() (interface{}, error) {
	return nil, nil
}
