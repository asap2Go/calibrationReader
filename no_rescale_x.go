package calibrationReader

import (
	"errors"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
)

// getNoRescaleX retrieves the number of Rescale X-axis pairs according to its layout specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getNoRescaleX(rl *a2l.RecordLayout, curPos *uint32) (interface{}, error) {
	if !rl.NoRescaleX.DatatypeSet {
		err := errors.New("noRescaleX datatype not set")
		log.Err(err).Msg("could not retrieve NoRescaleX value")
		return nil, err
	}
	val, err := cd.getValue(curPos, rl.NoRescaleX.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve noRescaleX value")
		return nil, err
	}
	*curPos += uint32(rl.NoRescaleX.Datatype.GetDatatypeLength())
	return val, err
}
