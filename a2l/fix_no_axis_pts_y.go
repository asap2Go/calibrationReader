package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type FixNoAxisPtsY struct {
	NumberOfAxisPoints    uint16
	NumberOfAxisPointsSet bool
}

func parseFixNoAxisPtsY(tok *tokenGenerator) (FixNoAxisPtsY, error) {
	fnapy := FixNoAxisPtsY{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("fixNoAxisPtsy could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("fixNoAxisPtsY could not be parsed")
	} else if !fnapy.NumberOfAxisPointsSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("fixNoAxisPtsy numberOfAxisPoints could not be parsed")
		}
		fnapy.NumberOfAxisPoints = uint16(buf)
		fnapy.NumberOfAxisPointsSet = true
		log.Info().Msg("fixNoAxisPtsy numberOfAxisPoints successfully parsed")
	}
	return fnapy, err
}
