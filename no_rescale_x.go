package calibrationReader

import (
	"errors"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
)

// getNoRescaleX retrieves the number of Rescale X-axis pairs according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getNoRescaleX(rl *a2l.RecordLayout, curPos *uint32) (int64, error) {
	if !rl.NoRescaleX.DatatypeSet {
		err := errors.New("noRescaleX datatype not set")
		log.Err(err).Msg("could not retrieve noRescaleX value")
		return 0, err
	}
	bufBytes, err := cd.getValue(curPos, rl.NoRescaleX.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve noRescaleX value")
		return 0, err
	}
	val, err := cd.convertByteSliceToDatatype(bufBytes, rl.NoRescaleX.Datatype)
	if err != nil {
		log.Err(err).Msg("could not retrieve noRescaleX value")
		return 0, err
	}
	*curPos += uint32(rl.NoRescaleX.Datatype.GetDatatypeLength())
	return int64(val), err
}
