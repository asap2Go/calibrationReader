package calibrationReader

import (
	"errors"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
)

// getAxisPtsX retrieves the number of X-axis points according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getNoAxisPtsX(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.NoAxisPtsX.DatatypeSet {
		err := errors.New("noAxisPtsX datatype not set")
		log.Err(err).Msg("could not retrieve noAxisPtsX value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.NoAxisPtsX.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve noAxisPtsX value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.NoAxisPtsX.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve noAxisPtsX value")
		return 0, err
	}
	*curPos += uint32(rl.NoAxisPtsX.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getAxisPtsY retrieves the number of X-axis points according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getNoAxisPtsY(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.NoAxisPtsY.DatatypeSet {
		err := errors.New("noAxisPtsY datatype not set")
		log.Err(err).Msg("could not retrieve noAxisPtsY value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.NoAxisPtsY.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve noAxisPtsY value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.NoAxisPtsY.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve noAxisPtsY value")
		return 0, err
	}
	*curPos += uint32(rl.NoAxisPtsY.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getAxisPtsZ retrieves the number of X-axis points according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getNoAxisPtsZ(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.NoAxisPtsZ.DatatypeSet {
		err := errors.New("noAxisPtsZ datatype not set")
		log.Err(err).Msg("could not retrieve noAxisPtsZ value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.NoAxisPtsZ.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve noAxisPtsZ value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.NoAxisPtsZ.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve noAxisPtsZ value")
		return 0, err
	}
	*curPos += uint32(rl.NoAxisPtsZ.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getAxisPts4 retrieves the number of X-axis points according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getNoAxisPts4(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.NoAxisPts4.DatatypeSet {
		err := errors.New("noAxisPts4 datatype not set")
		log.Err(err).Msg("could not retrieve noAxisPts4 value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.NoAxisPts4.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve noAxisPts4 value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.NoAxisPts4.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve noAxisPts4 value")
		return 0, err
	}
	*curPos += uint32(rl.NoAxisPts4.Datatype.GetDatatypeLength())
	return int64(val), err
}

// getAxisPts5 retrieves the number of X-axis points according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getNoAxisPts5(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.NoAxisPts5.DatatypeSet {
		err := errors.New("noAxisPts5 datatype not set")
		log.Err(err).Msg("could not retrieve noAxisPts5 value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.NoAxisPts5.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve noAxisPts5 value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.NoAxisPts5.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve noAxisPts5 value")
		return 0, err
	}
	*curPos += uint32(rl.NoAxisPts5.Datatype.GetDatatypeLength())
	return int64(val), err
}
