package calibrationReader

import (
	"errors"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
)

// getDistOpX retrieves the identifier in the deposit structure according to its layout
// specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getIdentification(rl *a2l.RecordLayout, curPos *uint32) (interface{}, error) {
	if !rl.Identification.DatatypeSet {
		err := errors.New("identification datatype not set")
		log.Err(err).Msg("could not retrieve identification value")
		return nil, err
	}
	val, err := cd.getValue(curPos, rl.Identification.Datatype, rl)
	if err != nil {
		log.Err(err).Msg("could not retrieve identification value")
		return nil, err
	}
	*curPos += uint32(rl.Identification.Datatype.GetDatatypeLength())
	return val, err
}
