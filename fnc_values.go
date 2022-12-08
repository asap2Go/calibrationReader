package calibrationReader

import (
	"errors"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
)

// getFncValues1D retrieves Axis Points according to their layout specified within the record layout and their values as calibrated in the hex file
func (cv *CharacteristicValues) getFncValues1D(cd *CalibrationData, rl *a2l.RecordLayout, curPos *uint32) (float64, []byte, error) {
	var bufByte []byte
	var bufFloat float64
	var err error

	bufByte, err = cd.getValue(curPos, rl.FncValues.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve FncValues values")
		return 0.0, bufByte, err
	}
	bufFloat, err = cd.convertByteSliceToDatatype(bufByte, rl.FncValues.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve FncValues values")
		return 0.0, bufByte, err
	}
	*curPos += uint32(rl.FncValues.Datatype.GetDatatypeLength())
	return bufFloat, bufByte, err
}

// getFncValues2D retrieves Axis Points according to their layout specified within the record layout and their values as calibrated in the hex file
func (cv *CharacteristicValues) getFncValues2D(cd *CalibrationData, rl *a2l.RecordLayout, curPos *uint32) ([]float64, error) {
	var val []float64
	var bufByte []byte
	var bufFloat float64
	var err error
	var i int64
	var noFncValues int64
	if rl.FixNoAxisPtsX.NumberOfAxisPointsSet {
		noFncValues = int64(rl.FixNoAxisPtsX.NumberOfAxisPoints)
	} else if !rl.FixNoAxisPtsX.NumberOfAxisPointsSet {
		if cv.characteristic.MatrixDim.DimX <= 0 {
			err = errors.New("number of FncValues is smaller or equal to zero")
			log.Err(err).Msg("could not retrieve NoFncValues value")
			return nil, err
		}
		noFncValues = cv.noFncValuesXValue
	} else {
		err = errors.New("number of FncValues could not be determined")
		log.Err(err).Msg("could not convert FncValues values")
		return nil, err
	}
	for i = 0; i < noFncValues; i++ {
		bufByte, err = cd.getValue(curPos, rl.AxisPtsX.Datatype, rl)
		if err != nil {
			log.Err(err).Msg("could not retrieve FncValues values")
			return nil, err
		}
		bufFloat, err = cd.convertByteSliceToDatatype(bufByte, rl.AxisPtsX.Datatype)
		if err != nil {
			log.Err(err).Msg("could not retrieve FncValues values")
			return nil, err
		}
		val = append(val, bufFloat)
		*curPos += uint32(rl.AxisPtsX.Datatype.GetDatatypeLength())
	}
	return val, err
}
