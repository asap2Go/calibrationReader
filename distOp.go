package calibrationReader

import (
	"errors"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
)

// getDistOpX retrieves the distance operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getDistOpX(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.DistOpX.DatatypeSet {
		err := errors.New("distOpX datatype not set")
		log.Err(err).Msg("could not retrieve distOpX value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.DistOpX.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve distOpX value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.DistOpX.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve distOpX value")
		return 0, err
	}
	*curPos += uint32(rl.DistOpX.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getDistOpY retrieves the distance operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getDistOpY(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.DistOpY.DatatypeSet {
		err := errors.New("distOpY datatype not set")
		log.Err(err).Msg("could not retrieve distOpY value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.DistOpY.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve distOpY value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.DistOpY.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve distOpY value")
		return 0, err
	}
	*curPos += uint32(rl.DistOpY.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getDistOpZ retrieves the distance operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getDistOpZ(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.DistOpZ.DatatypeSet {
		err := errors.New("distOpZ datatype not set")
		log.Err(err).Msg("could not retrieve distOpZ value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.DistOpZ.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve distOpZ value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.DistOpZ.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve distOpZ value")
		return 0, err
	}
	*curPos += uint32(rl.DistOpZ.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getDistOp4 retrieves the distance operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getDistOp4(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.DistOp4.DatatypeSet {
		err := errors.New("distOp4 datatype not set")
		log.Err(err).Msg("could not retrieve distOp4 value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.DistOp4.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve distOp4 value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.DistOp4.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve distOp4 value")
		return 0, err
	}
	*curPos += uint32(rl.DistOp4.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getDistOp5 retrieves the distance operator according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getDistOp5(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.DistOp5.DatatypeSet {
		err := errors.New("distOp5 datatype not set")
		log.Err(err).Msg("could not retrieve distOp5 value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.DistOp5.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve distOp5 value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.DistOp5.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve distOp5 value")
		return 0, err
	}
	*curPos += uint32(rl.DistOp5.Datatype.GetDatatypeLength())
	return int64(val), err
}
