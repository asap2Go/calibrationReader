package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type fixAxisPar struct {
	offset       int16
	offsetSet    bool
	shift        int16
	shiftSet     bool
	numberapo    uint16
	numberapoSet bool
}

func parseFixAxisPar(tok *tokenGenerator) (fixAxisPar, error) {
	fap := fixAxisPar{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("fixAxisPar could not be parsed")
			break forLoop
		} else if !fap.offsetSet {
			var buf int64
			buf, err = strconv.ParseInt(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("fixAxisPar offset could not be parsed")
				break forLoop
			}
			fap.offset = int16(buf)
			fap.offsetSet = true
			log.Info().Msg("fixAxisPar offset successfully parsed")
		} else if !fap.shiftSet {
			var buf int64
			buf, err = strconv.ParseInt(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("fixAxisPar shift could not be parsed")
				break forLoop
			}
			fap.shift = int16(buf)
			fap.shiftSet = true
			log.Info().Msg("fixAxisPar shift successfully parsed")
		} else if !fap.numberapoSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("fixAxisPar numberapo could not be parsed")
				break forLoop
			}
			fap.numberapo = uint16(buf)
			fap.numberapoSet = true
			log.Info().Msg("fixAxisPar numberapo successfully parsed")
			break forLoop
		}
	}
	return fap, err
}
