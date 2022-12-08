package calibrationReader

import (
	"errors"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
)

//check access type. DIRECT is the most used. Just read value from a given address.
//in case other access types are set this gets more complicated as either offsets or pointers are leveraged to
//define the position of the calibration objects.

// getAxisPointsX retrieves Axis Points according to their layout specified within the record layout and their values as calibrated in the hex file
func (cv *CharacteristicValues) getAxisPointsX(cd *CalibrationData, rl *a2l.RecordLayout, curPos *uint32) ([]float64, error) {
	var val []float64
	var bufByte []byte
	var bufFloat float64
	var err error
	var i int64
	var noAxisPts int64
	if rl.FixNoAxisPtsX.NumberOfAxisPointsSet {
		noAxisPts = int64(rl.FixNoAxisPtsX.NumberOfAxisPoints)
	} else if !rl.FixNoAxisPtsX.NumberOfAxisPointsSet {
		if cv.noAxisPtsXValue <= 0 {
			err = errors.New("number of axisPts is smaller or equal to zero")
			log.Err(err).Msg("could not retrieve NoAxisPointsX value")
			return nil, err
		}
		noAxisPts = cv.noAxisPtsXValue
	} else {
		err = errors.New("number of axis points could not be determined")
		log.Err(err).Msg("could not convert axisPointsX values")
		return nil, err
	}
	for i = 0; i < noAxisPts; i++ {
		bufByte, err = cd.getValue(curPos, rl.AxisPtsX.Datatype, rl)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPointsX value")
			return nil, err
		}
		bufFloat, err = cd.convertByteSliceToDatatype(bufByte, rl.AxisPtsX.Datatype)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPointsX value")
			return nil, err
		}
		val = append(val, bufFloat)
		*curPos += uint32(rl.AxisPtsX.Datatype.GetDatatypeLength())
	}
	return val, err
}

// getAxisPointsY retrieves Axis Points according to their layout specified within the record layout and their values as calibrated in the hex file
func (cv *CharacteristicValues) getAxisPointsY(cd *CalibrationData, rl *a2l.RecordLayout, curPos *uint32) ([]float64, error) {
	var val []float64
	var bufByte []byte
	var bufFloat float64
	var err error
	var i int64
	var noAxisPts int64
	if rl.FixNoAxisPtsY.NumberOfAxisPointsSet {
		noAxisPts = int64(rl.FixNoAxisPtsY.NumberOfAxisPoints)
	} else if !rl.FixNoAxisPtsY.NumberOfAxisPointsSet {
		if cv.noAxisPtsYValue <= 0 {
			err = errors.New("number of axisPts is smaller or equal to zero")
			log.Err(err).Msg("could not retrieve NoAxisPointsY value")
			return nil, err
		}
		noAxisPts = cv.noAxisPtsYValue
	} else {
		err = errors.New("number of axis points could not be determined")
		log.Err(err).Msg("could not convert axisPointsY values")
		return nil, err
	}
	for i = 0; i < noAxisPts; i++ {
		bufByte, err = cd.getValue(curPos, rl.AxisPtsY.Datatype, rl)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPointsY value")
			return nil, err
		}
		bufFloat, err = cd.convertByteSliceToDatatype(bufByte, rl.AxisPtsY.Datatype)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPointsY value")
			return nil, err
		}
		val = append(val, bufFloat)
		*curPos += uint32(rl.AxisPtsY.Datatype.GetDatatypeLength())
	}
	return val, err
}

// getAxisPointsZ retrieves Axis Points according to their layout specified within the record layout and their values as calibrated in the hex file
func (cv *CharacteristicValues) getAxisPointsZ(cd *CalibrationData, rl *a2l.RecordLayout, curPos *uint32) ([]float64, error) {
	var val []float64
	var bufByte []byte
	var bufFloat float64
	var err error
	var i int64
	var noAxisPts int64
	if rl.FixNoAxisPtsZ.NumberOfAxisPointsSet {
		noAxisPts = int64(rl.FixNoAxisPtsZ.NumberOfAxisPoints)
	} else if !rl.FixNoAxisPtsZ.NumberOfAxisPointsSet {
		if cv.noAxisPtsZValue <= 0 {
			err = errors.New("number of axisPts is smaller or equal to zero")
			log.Err(err).Msg("could not retrieve NoAxisPointsZ value")
			return nil, err
		}
		noAxisPts = cv.noAxisPtsZValue
	} else {
		err = errors.New("number of axis points could not be determined")
		log.Err(err).Msg("could not convert axisPointsZ values")
		return nil, err
	}
	for i = 0; i < noAxisPts; i++ {
		bufByte, err = cd.getValue(curPos, rl.AxisPtsZ.Datatype, rl)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPointsZ value")
			return nil, err
		}
		bufFloat, err = cd.convertByteSliceToDatatype(bufByte, rl.AxisPtsZ.Datatype)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPointsZ value")
			return nil, err
		}
		val = append(val, bufFloat)
		*curPos += uint32(rl.AxisPtsZ.Datatype.GetDatatypeLength())
	}
	return val, err
}

// getAxisPoints4 retrieves Axis Points according to their layout specified within the record layout and their values as calibrated in the hex file
func (cv *CharacteristicValues) getAxisPoints4(cd *CalibrationData, rl *a2l.RecordLayout, curPos *uint32) ([]float64, error) {
	var val []float64
	var bufByte []byte
	var bufFloat float64
	var err error
	var i int64
	var noAxisPts int64
	if rl.FixNoAxisPts4.NumberOfAxisPointsSet {
		noAxisPts = int64(rl.FixNoAxisPts4.NumberOfAxisPoints)
	} else if !rl.FixNoAxisPts4.NumberOfAxisPointsSet {
		if cv.noAxisPts4Value <= 0 {
			err = errors.New("number of axisPts is smaller or equal to zero")
			log.Err(err).Msg("could not retrieve NoAxisPoints4 value")
			return nil, err
		}
		noAxisPts = cv.noAxisPts4Value
	} else {
		err = errors.New("number of axis points could not be determined")
		log.Err(err).Msg("could not convert axisPoints4 values")
		return nil, err
	}
	for i = 0; i < noAxisPts; i++ {
		bufByte, err = cd.getValue(curPos, rl.AxisPts4.Datatype, rl)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPoints4 value")
			return nil, err
		}
		bufFloat, err = cd.convertByteSliceToDatatype(bufByte, rl.AxisPts4.Datatype)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPoints4 value")
			return nil, err
		}
		val = append(val, bufFloat)
		*curPos += uint32(rl.AxisPts4.Datatype.GetDatatypeLength())
	}
	return val, err
}

// getAxisPoints5 retrieves Axis Points according to their layout specified within the record layout and their values as calibrated in the hex file
func (cv *CharacteristicValues) getAxisPoints5(cd *CalibrationData, rl *a2l.RecordLayout, curPos *uint32) ([]float64, error) {
	var val []float64
	var bufByte []byte
	var bufFloat float64
	var err error
	var i int64
	var noAxisPts int64
	if rl.FixNoAxisPts5.NumberOfAxisPointsSet {
		noAxisPts = int64(rl.FixNoAxisPts5.NumberOfAxisPoints)
	} else if !rl.FixNoAxisPts5.NumberOfAxisPointsSet {
		if cv.noAxisPts5Value <= 0 {
			err = errors.New("number of axisPts is smaller or equal to zero")
			log.Err(err).Msg("could not retrieve NoAxisPoints5 value")
			return nil, err
		}
		noAxisPts = cv.noAxisPts5Value
	} else {
		err = errors.New("number of axis points could not be determined")
		log.Err(err).Msg("could not convert axisPoints5 values")
		return nil, err
	}
	for i = 0; i < noAxisPts; i++ {
		bufByte, err = cd.getValue(curPos, rl.AxisPts5.Datatype, rl)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPoints5 value")
			return nil, err
		}
		bufFloat, err = cd.convertByteSliceToDatatype(bufByte, rl.AxisPts5.Datatype)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPoints5 value")
			return nil, err
		}
		val = append(val, bufFloat)
		*curPos += uint32(rl.AxisPts5.Datatype.GetDatatypeLength())
	}
	return val, err
}
