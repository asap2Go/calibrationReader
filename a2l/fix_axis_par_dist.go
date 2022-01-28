package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type fixAxisParDist struct {
	offset       int16
	offsetSet    bool
	distance     int16
	distanceSet  bool
	numberapo    uint16
	numberapoSet bool
}

func parseFixAxisParDist(tok *tokenGenerator) (fixAxisParDist, error) {
	fapd := fixAxisParDist{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("fixAxisParDist could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("fixAxisParDist could not be parsed")
			break forLoop
		} else if !fapd.offsetSet {
			var buf int64
			buf, err = strconv.ParseInt(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("fixAxisParDist offset could not be parsed")
				break forLoop
			}
			fapd.offset = int16(buf)
			fapd.offsetSet = true
			log.Info().Msg("fixAxisParDist offset successfully parsed")
		} else if !fapd.distanceSet {
			var buf int64
			buf, err = strconv.ParseInt(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("fixAxisParDist distance could not be parsed")
				break forLoop
			}
			fapd.distance = int16(buf)
			fapd.distanceSet = true
			log.Info().Msg("fixAxisParDist distance successfully parsed")
		} else if !fapd.numberapoSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("fixAxisParDist numberapo could not be parsed")
				break forLoop
			}
			fapd.numberapo = uint16(buf)
			fapd.numberapoSet = true
			log.Info().Msg("fixAxisParDist numberapo successfully parsed")
			break forLoop
		}
	}
	return fapd, err
}
