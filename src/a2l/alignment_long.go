package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type alignmentLong struct {
	alignmentBorder    uint16
	alignmentBorderSet bool
}

func parseAlignmentLong(tok *tokenGenerator) (alignmentLong, error) {
	al := alignmentLong{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("alignmentLong could not be parsed")
	} else if !al.alignmentBorderSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("alignmentLong alignmentBorder could not be parsed")
		}
		al.alignmentBorder = uint16(buf)
		al.alignmentBorderSet = true
		log.Info().Msg("alignmentLong alignmentBorder successfully parsed")
	}
	return al, err
}
