package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

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
