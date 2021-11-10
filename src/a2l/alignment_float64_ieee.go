package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type alignmentFloat64Ieee struct {
	alignmentBorder    uint16
	alignmentBorderSet bool
}

func parseAlignmentFloat64Ieee(tok *tokenGenerator) (alignmentFloat64Ieee, error) {
	af64 := alignmentFloat64Ieee{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("alignmentFloat64Ieee could not be parsed")
	} else if !af64.alignmentBorderSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("alignmentFloat64Ieee alignmentBorder could not be parsed")
		}
		af64.alignmentBorder = uint16(buf)
		af64.alignmentBorderSet = true
		log.Info().Msg("alignmentFloat64Ieee alignmentBorder successfully parsed")
	}
	return af64, err
}
