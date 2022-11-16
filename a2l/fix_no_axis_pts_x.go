package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type FixNoAxisPtsX struct {
	NumberOfAxisPoints    uint16
	NumberOfAxisPointsSet bool
}

func parseFixNoAxisPtsX(tok *tokenGenerator) (FixNoAxisPtsX, error) {
	fnapx := FixNoAxisPtsX{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("fixNoAxisPtsx could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("fixNoAxisPtsX could not be parsed")
	} else if !fnapx.NumberOfAxisPointsSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("fixNoAxisPtsx numberOfAxisPoints could not be parsed")
		}
		fnapx.NumberOfAxisPoints = uint16(buf)
		fnapx.NumberOfAxisPointsSet = true
		log.Info().Msg("fixNoAxisPtsx numberOfAxisPoints successfully parsed")
	}
	return fnapx, err
}
