package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type alignmentInt64 struct {
	alignmentBorder    uint16
	alignmentBorderSet bool
}

func parseAlignmentInt64(tok *tokenGenerator) (alignmentInt64, error) {
	ai64 := alignmentInt64{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("alignmentInt64 could not be parsed")
	} else if !ai64.alignmentBorderSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("alignmentInt64 alignmentBorder could not be parsed")
		}
		ai64.alignmentBorder = uint16(buf)
		ai64.alignmentBorderSet = true
		log.Info().Msg("alignmentInt64 alignmentBorder successfully parsed")
	}
	return ai64, err
}
