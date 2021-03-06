package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type fixNoAxisPtsZ struct {
	numberOfAxisPoints    uint16
	numberOfAxisPointsSet bool
}

func parseFixNoAxisPtsZ(tok *tokenGenerator) (fixNoAxisPtsZ, error) {
	fnapz := fixNoAxisPtsZ{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("fixNoAxisPtsz could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("fixNoAxisPtsZ could not be parsed")
	} else if !fnapz.numberOfAxisPointsSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("fixNoAxisPtsz numberOfAxisPoints could not be parsed")
		}
		fnapz.numberOfAxisPoints = uint16(buf)
		fnapz.numberOfAxisPointsSet = true
		log.Info().Msg("fixNoAxisPtsz numberOfAxisPoints successfully parsed")
	}
	return fnapz, err
}
