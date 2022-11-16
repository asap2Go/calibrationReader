package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type FixNoAxisPtsZ struct {
	NumberOfAxisPoints    uint16
	NumberOfAxisPointsSet bool
}

func parseFixNoAxisPtsZ(tok *tokenGenerator) (FixNoAxisPtsZ, error) {
	fnapz := FixNoAxisPtsZ{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("fixNoAxisPtsz could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("fixNoAxisPtsZ could not be parsed")
	} else if !fnapz.NumberOfAxisPointsSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("fixNoAxisPtsz numberOfAxisPoints could not be parsed")
		}
		fnapz.NumberOfAxisPoints = uint16(buf)
		fnapz.NumberOfAxisPointsSet = true
		log.Info().Msg("fixNoAxisPtsz numberOfAxisPoints successfully parsed")
	}
	return fnapz, err
}
