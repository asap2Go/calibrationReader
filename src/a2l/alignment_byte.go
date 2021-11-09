package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type alignmentByte struct {
	alignmentBorder    uint16
	alignmentBorderSet bool
}

func parseAlignmentByte(tok *tokenGenerator) (alignmentByte, error) {
	ab := alignmentByte{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("alignmentByte could not be parsed")
	} else if !ab.alignmentBorderSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
				log.Err(err).Msg("alignmentByte alignmentBorder could not be parsed")
		} 
			ab.alignmentBorder = uint16(buf)
			ab.alignmentBorderSet = true
				log.Info().Msg("alignmentByte alignmentBorder successfully parsed")
	}
	return ab, err
}
