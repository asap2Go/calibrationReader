package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type alignmentWord struct {
	alignmentBorder    uint16
	alignmentBorderSet bool
}

func parseAlignmentWord(tok *tokenGenerator) (alignmentWord, error) {
	aw := alignmentWord{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("alignmentWord could not be parsed")
	} else if !aw.alignmentBorderSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
				log.Err(err).Msg("alignmentWord alignmentBorder could not be parsed")
		}
			aw.alignmentBorder = uint16(buf)
			aw.alignmentBorderSet = true
				log.Info().Msg("alignmentWord alignmentBorder successfully parsed")
	}
	return aw, err
}
