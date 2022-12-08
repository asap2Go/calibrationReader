package calibrationReader

import (
	"errors"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
)

// getShiftOpX retrieves the shift operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getShiftOpX(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.ShiftOpX.DatatypeSet {
		err := errors.New("shiftOpX datatype not set")
		log.Err(err).Msg("could not retrieve shiftOpX value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.ShiftOpX.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve shiftOpX value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.ShiftOpX.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve shiftOpX value")
		return 0, err
	}
	*curPos += uint32(rl.ShiftOpX.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getShiftOpY retrieves the shift operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getShiftOpY(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.ShiftOpY.DatatypeSet {
		err := errors.New("shiftOpY datatype not set")
		log.Err(err).Msg("could not retrieve shiftOpY value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.ShiftOpY.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve shiftOpY value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.ShiftOpY.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve shiftOpY value")
		return 0, err
	}
	*curPos += uint32(rl.ShiftOpY.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getShiftOpZ retrieves the shift operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getShiftOpZ(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.ShiftOpZ.DatatypeSet {
		err := errors.New("shiftOpZ datatype not set")
		log.Err(err).Msg("could not retrieve shiftOpZ value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.ShiftOpZ.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve shiftOpZ value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.ShiftOpZ.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve shiftOpZ value")
		return 0, err
	}
	*curPos += uint32(rl.ShiftOpZ.Datatype.GetDatatypeLength())
	return int64(val), err
}

// / getShiftOp4 retrieves the shift operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getShiftOp4(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.ShiftOp4.DatatypeSet {
		err := errors.New("shiftOp4 datatype not set")
		log.Err(err).Msg("could not retrieve shiftOp4 value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.ShiftOp4.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve shiftOp4 value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.ShiftOp4.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve shiftOp4 value")
		return 0, err
	}
	*curPos += uint32(rl.ShiftOp4.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getShiftOp5 retrieves the shift operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getShiftOp5(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.ShiftOp5.DatatypeSet {
		err := errors.New("shiftOp5 datatype not set")
		log.Err(err).Msg("could not retrieve shiftOp5 value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.ShiftOp5.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve shiftOp5 value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.ShiftOp5.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve shiftOp5 value")
		return 0, err
	}
	*curPos += uint32(rl.ShiftOp5.Datatype.GetDatatypeLength())
	return int64(val), err
}
