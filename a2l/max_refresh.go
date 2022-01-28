package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type MaxRefresh struct {
	scalingUnit    uint16
	scalingUnitSet bool
	rate           uint32
	rateSet        bool
}

func parseMaxRefresh(tok *tokenGenerator) (MaxRefresh, error) {
	mr := MaxRefresh{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("maxRefresh could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("maxRefresh could not be parsed")
			break forLoop
		} else if !mr.scalingUnitSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("maxRefresh scalingUnit could not be parsed")
				break forLoop
			}
			mr.scalingUnit = uint16(buf)
			mr.scalingUnitSet = true
			log.Info().Msg("maxRefresh scalingUnit successfully parsed")
		} else if !mr.rateSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 32)
			if err != nil {
				log.Err(err).Msg("maxRefresh rate could not be parsed")
				break forLoop
			}
			mr.rate = uint32(buf)
			mr.rateSet = true
			log.Info().Msg("maxRefresh rate successfully parsed")
			break forLoop
		}
	}
	return mr, err
}
