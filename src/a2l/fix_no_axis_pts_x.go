package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type fixNoAxisPtsX struct {
	numberOfAxisPoints    uint16
	numberOfAxisPointsSet bool
}

func parseFixNoAxisPtsX(tok *tokenGenerator) (fixNoAxisPtsX, error) {
	fnapx := fixNoAxisPtsX{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("fixNoAxisPtsx could not be parsed")
	} else if !fnapx.numberOfAxisPointsSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
				log.Err(err).Msg("fixNoAxisPtsx numberOfAxisPoints could not be parsed")
		}
		fnapx.numberOfAxisPoints = uint16(buf)
		fnapx.numberOfAxisPointsSet = true
			log.Info().Msg("fixNoAxisPtsx numberOfAxisPoints successfully parsed")
	}
	return fnapx, err
}
