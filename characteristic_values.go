package calibrationReader

import (
	"sync"

	"github.com/asap2Go/calibrationReader/a2l"
)

type CharacteristicValues struct {
	characteristic      *a2l.Characteristic
	recordLayout        *a2l.RecordLayout
	AxisXValues         []float64
	AxisYValues         []float64
	AxisZValues         []float64
	Axis4Values         []float64
	Axis5Values         []float64
	distOpXValue        int64
	distOpYValue        int64
	distOpZValue        int64
	distOp4Value        int64
	distOp5Value        int64
	fncValues           interface{}
	identificationValue interface{}
	noAxisPtsXValue     int64
	noAxisPtsYValue     int64
	noAxisPtsZValue     int64
	noAxisPts4Value     int64
	noAxisPts5Value     int64
	noRescaleXValue     int64
	offsetXValue        int64
	offsetYValue        int64
	offsetZValue        int64
	offset4Value        int64
	offset5Value        int64
	shiftOpXValue       int64
	shiftOpYValue       int64
	shiftOpZValue       int64
	shiftOp4Value       int64
	shiftOp5Value       int64
}

func NewCharacteristicValues(characteristic *a2l.Characteristic, recordLayout *a2l.RecordLayout) *CharacteristicValues {
	return &CharacteristicValues{
		characteristic: characteristic,
		recordLayout:   recordLayout,
	}
}

func (cv CharacteristicValues) computeValues(cd *CalibrationData, valueChannel chan interface{}, errChannel chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	return
}
