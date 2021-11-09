package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type ecuCalibrationOffset struct {
	offset    int32
	offsetSet bool
}

func parseEcuCalibrationOffset(tok *tokenGenerator) (ecuCalibrationOffset, error) {
	eco := ecuCalibrationOffset{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("ecuCalibrationOffset could not be parsed")
	} else if !eco.offsetSet {
		var buf int64
		buf, err = strconv.ParseInt(tok.current(), 10, 32)
		if err != nil {
				log.Err(err).Msg("ecuCalibrationOffset offset could not be parsed")
		}
		eco.offset = int32(buf)
		eco.offsetSet = true
			log.Info().Msg("ecuCalibrationOffset offset successfully parsed")
	}
	return eco, err
}
