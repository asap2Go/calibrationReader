package calibrationReader

import (
	"errors"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
)

// getAxisPointsX retrieves Axis Points according to their layout specified within the record layout and their values as calibrated in the hex file
func (cv *CharacteristicValues) getAxisPointsX(cd *CalibrationData, rl *a2l.RecordLayout, curPos *uint32) (interface{}, error) {
	var val []interface{}
	var buf interface{}
	var err error
	var i uint32
	var noAxisPts uint32
	if rl.FixNoAxisPtsX.NumberOfAxisPointsSet {
		noAxisPts = uint32(rl.FixNoAxisPtsX.NumberOfAxisPoints)
	} else if !rl.FixNoAxisPtsX.NumberOfAxisPointsSet {
		buf, isInt := cv.noAxisPtsXValue.(int)
		if !isInt {
			err = errors.New("could not convert number of axisPts datatype to int")
			log.Err(err).Msg("could not retrieve NoAxisPointsX value")
			return nil, err
		}
		noAxisPts = uint32(buf)
	} else {
		err = errors.New("number of axis points could not be determined")
		log.Err(err).Msg("could not convert axisPointsX values")
		return nil, err
	}
	for i = 0; i < noAxisPts; i++ {
		buf, err = cd.getValue(curPos, rl.AxisPtsX.Datatype, rl)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPointsX value")
			return nil, err
		}
		val = append(val, buf)
		*curPos += uint32(rl.AxisPtsX.Datatype.GetDatatypeLength())
	}
	return val, err
}

// getAxisPointsY retrieves Axis Points according to their layout specified within the record layout and their values as calibrated in the hex file
func (cv *CharacteristicValues) getAxisPointsY(cd *CalibrationData, rl *a2l.RecordLayout, curPos *uint32) (interface{}, error) {
	var val []interface{}
	var buf interface{}
	var err error
	var i uint32
	var noAxisPts uint32
	if rl.FixNoAxisPtsY.NumberOfAxisPointsSet {
		noAxisPts = uint32(rl.FixNoAxisPtsY.NumberOfAxisPoints)
	} else if !rl.FixNoAxisPtsY.NumberOfAxisPointsSet {
		buf, isInt := cv.noAxisPtsYValue.(int)
		if !isInt {
			err = errors.New("could not convert number of axisPts datatype to int")
			log.Err(err).Msg("could not retrieve NoAxisPointsY value")
			return nil, err
		}
		noAxisPts = uint32(buf)
	} else {
		err = errors.New("number of axis points could not be determined")
		log.Err(err).Msg("could not convert axisPointsY values")
		return nil, err
	}
	for i = 0; i < noAxisPts; i++ {
		buf, err = cd.getValue(curPos, rl.AxisPtsY.Datatype, rl)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPointsY value")
			return nil, err
		}
		val = append(val, buf)
		*curPos += uint32(rl.AxisPtsY.Datatype.GetDatatypeLength())
	}
	return val, err
}

// getAxisPointsZ retrieves Axis Points according to their layout specified within the record layout and their values as calibrated in the hex file
func (cv *CharacteristicValues) getAxisPointsZ(cd *CalibrationData, rl *a2l.RecordLayout, curPos *uint32) (interface{}, error) {
	var val []interface{}
	var buf interface{}
	var err error
	var i uint32
	var noAxisPts uint32
	if rl.FixNoAxisPtsZ.NumberOfAxisPointsSet {
		noAxisPts = uint32(rl.FixNoAxisPtsZ.NumberOfAxisPoints)
	} else if !rl.FixNoAxisPtsZ.NumberOfAxisPointsSet {
		buf, isInt := cv.noAxisPtsZValue.(int)
		if !isInt {
			err = errors.New("could not convert number of axisPts datatype to int")
			log.Err(err).Msg("could not retrieve NoAxisPointsZ value")
			return nil, err
		}
		noAxisPts = uint32(buf)
	} else {
		err = errors.New("number of axis points could not be determined")
		log.Err(err).Msg("could not convert axisPointsZ values")
		return nil, err
	}
	for i = 0; i < noAxisPts; i++ {
		buf, err = cd.getValue(curPos, rl.AxisPtsZ.Datatype, rl)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPointsZ value")
			return nil, err
		}
		val = append(val, buf)
		*curPos += uint32(rl.AxisPtsZ.Datatype.GetDatatypeLength())
	}
	return val, err
}

// getAxisPoints4 retrieves Axis Points according to their layout specified within the record layout and their values as calibrated in the hex file
func (cv *CharacteristicValues) getAxisPoints4(cd *CalibrationData, rl *a2l.RecordLayout, curPos *uint32) (interface{}, error) {
	var val []interface{}
	var buf interface{}
	var err error
	var i uint32
	var noAxisPts uint32
	if rl.FixNoAxisPts4.NumberOfAxisPointsSet {
		noAxisPts = uint32(rl.FixNoAxisPts4.NumberOfAxisPoints)
	} else if !rl.FixNoAxisPts4.NumberOfAxisPointsSet {
		buf, isInt := cv.noAxisPts4Value.(int)
		if !isInt {
			err = errors.New("could not convert number of axisPts datatype to int")
			log.Err(err).Msg("could not retrieve NoAxisPoints4 value")
			return nil, err
		}
		noAxisPts = uint32(buf)
	} else {
		err = errors.New("number of axis points could not be determined")
		log.Err(err).Msg("could not convert axisPoints4 values")
		return nil, err
	}
	for i = 0; i < noAxisPts; i++ {
		buf, err = cd.getValue(curPos, rl.AxisPts4.Datatype, rl)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPoints4 value")
			return nil, err
		}
		val = append(val, buf)
		*curPos += uint32(rl.AxisPts4.Datatype.GetDatatypeLength())
	}
	return val, err
}

// getAxisPoints5 retrieves Axis Points according to their layout specified within the record layout and their values as calibrated in the hex file
func (cv *CharacteristicValues) getAxisPoints5(cd *CalibrationData, rl *a2l.RecordLayout, curPos *uint32) (interface{}, error) {
	var val []interface{}
	var buf interface{}
	var err error
	var i uint32
	var noAxisPts uint32
	if rl.FixNoAxisPts5.NumberOfAxisPointsSet {
		noAxisPts = uint32(rl.FixNoAxisPts5.NumberOfAxisPoints)
	} else if !rl.FixNoAxisPts5.NumberOfAxisPointsSet {
		buf, isInt := cv.noAxisPts5Value.(int)
		if !isInt {
			err = errors.New("could not convert number of axisPts datatype to int")
			log.Err(err).Msg("could not retrieve NoAxisPoints5 value")
			return nil, err
		}
		noAxisPts = uint32(buf)
	} else {
		err = errors.New("number of axis points could not be determined")
		log.Err(err).Msg("could not convert axisPoints5 values")
		return nil, err
	}
	for i = 0; i < noAxisPts; i++ {
		buf, err = cd.getValue(curPos, rl.AxisPts5.Datatype, rl)
		if err != nil {
			log.Err(err).Msg("could not retrieve axisPoints5 value")
			return nil, err
		}
		val = append(val, buf)
		*curPos += uint32(rl.AxisPts5.Datatype.GetDatatypeLength())
	}
	return val, err
}
