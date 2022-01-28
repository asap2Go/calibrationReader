package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type fixNoAxisPts4 struct {
	numberOfAxisPoints    uint16
	numberOfAxisPointsSet bool
}

func parseFixNoAxisPts4(tok *tokenGenerator) (fixNoAxisPts4, error) {
	fnap := fixNoAxisPts4{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("fixNoAxisPts4 could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("fixNoAxisPts4 could not be parsed")
	} else if !fnap.numberOfAxisPointsSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("fixNoAxisPts4 numberOfAxisPoints could not be parsed")
		}
		fnap.numberOfAxisPoints = uint16(buf)
		fnap.numberOfAxisPointsSet = true
		log.Info().Msg("fixNoAxisPts4 numberOfAxisPoints successfully parsed")
	}
	return fnap, err
}
