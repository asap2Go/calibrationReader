package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

/*
ECU_CALIBRATION_OFFSET is used to describe a fixed address offset when accessing
characteristics in the control unit due to
-near pointers in calibration objects. Some record layouts include near pointers
inside a calibration objects from which the calibration system has to compute the
absolute values by adding the ECU_CALIBRATION_OFFSET (CDAMOS)
-variant coding. Some ECU projects include multiple data sets for different engine or
vehicle projects served by one common ECU. By using the
ECU_CALIBRATION_OFFSET, a selection for project base address can be made
*/
type ecuCalibrationOffset struct {
	offset    string
	offsetSet bool
}

func parseEcuCalibrationOffset(tok *tokenGenerator) (ecuCalibrationOffset, error) {
	eco := ecuCalibrationOffset{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("ecuCalibrationOffset could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("ecuCalibrationOffset could not be parsed")
	} else if !eco.offsetSet {
		eco.offset = tok.current()
		eco.offsetSet = true
		log.Info().Msg("ecuCalibrationOffset offset successfully parsed")
	}
	return eco, err
}
