package calibrationReader

import (
	"errors"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
)

// getDistOpX retrieves the identifier in the deposit structure according to its layout
// specified within the record layout and their values as calibrated in the hex file
func (cd *CalibrationData) getReserved(rl *a2l.RecordLayout, curPos *uint32) error {
	if !rl.Reserved.DataSizeSet {
		err := errors.New("reserved datasize not set")
		log.Err(err).Msg("could not retrieve reserved datasize")
		return err
	}
	//Value of reserved is not relevant. Only its datasize and the resulting offset for other datastructres are necessary
	*curPos += uint32(rl.Reserved.DataSize.GetDataSizeLength())
	return nil
}
