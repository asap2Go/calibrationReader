package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type fixNoAxisPtsY struct {
	numberOfAxisPoints    uint16
	numberOfAxisPointsSet bool
}

func parseFixNoAxisPtsY(tok *tokenGenerator) (fixNoAxisPtsY, error) {
	fnapy := fixNoAxisPtsY{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("fixNoAxisPtsy could not be parsed")
	} else if !fnapy.numberOfAxisPointsSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
				log.Err(err).Msg("fixNoAxisPtsy numberOfAxisPoints could not be parsed")
		}
		fnapy.numberOfAxisPoints = uint16(buf)
		fnapy.numberOfAxisPointsSet = true
			log.Info().Msg("fixNoAxisPtsy numberOfAxisPoints successfully parsed")
	}
	return fnapy, err
}
