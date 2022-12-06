package calibrationReader

import (
	"errors"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
)

// getShiftOpX retrieves the shift operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getShiftOpX(rl *a2l.RecordLayout, curPos *uint32) (interface{}, error) {
	if !rl.ShiftOpX.DatatypeSet {
		err := errors.New("shiftX datatype not set")
		log.Err(err).Msg("could not retrieve shiftX value")
		return nil, err
	}
	val, err := cd.getValue(curPos, rl.ShiftOpX.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve shiftX value")
		return nil, err
	}
	*curPos += uint32(rl.ShiftOpX.Datatype.GetDatatypeLength())
	return val, err
}

// getShiftOpY retrieves the shift operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getShiftOpY(rl *a2l.RecordLayout, curPos *uint32) (interface{}, error) {
	if !rl.ShiftOpY.DatatypeSet {
		err := errors.New("shiftY datatype not set")
		log.Err(err).Msg("could not retrieve shiftY value")
		return nil, err
	}
	val, err := cd.getValue(curPos, rl.ShiftOpY.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve shiftY value")
		return nil, err
	}
	*curPos += uint32(rl.ShiftOpY.Datatype.GetDatatypeLength())
	return val, err
}

// getShiftOpZ retrieves the shift operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getShiftOpZ(rl *a2l.RecordLayout, curPos *uint32) (interface{}, error) {
	if !rl.ShiftOpZ.DatatypeSet {
		err := errors.New("shiftZ datatype not set")
		log.Err(err).Msg("could not retrieve shiftZ value")
		return nil, err
	}
	val, err := cd.getValue(curPos, rl.ShiftOpZ.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve shiftZ value")
		return nil, err
	}
	*curPos += uint32(rl.ShiftOpZ.Datatype.GetDatatypeLength())
	return val, err
}

// getShiftOp4 retrieves the shift operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getShiftOp4(rl *a2l.RecordLayout, curPos *uint32) (interface{}, error) {
	if !rl.ShiftOp4.DatatypeSet {
		err := errors.New("shift4 datatype not set")
		log.Err(err).Msg("could not retrieve shift4 value")
		return nil, err
	}
	val, err := cd.getValue(curPos, rl.ShiftOp4.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve shift4 value")
		return nil, err
	}
	*curPos += uint32(rl.ShiftOp4.Datatype.GetDatatypeLength())
	return val, err
}

// getShiftOp5 retrieves the shift operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getShiftOp5(rl *a2l.RecordLayout, curPos *uint32) (interface{}, error) {
	if !rl.ShiftOp5.DatatypeSet {
		err := errors.New("shift5 datatype not set")
		log.Err(err).Msg("could not retrieve shift5 value")
		return nil, err
	}
	val, err := cd.getValue(curPos, rl.ShiftOp5.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve shift5 value")
		return nil, err
	}
	*curPos += uint32(rl.ShiftOp5.Datatype.GetDatatypeLength())
	return val, err
}
