package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

// alignmentInt64 is necessary because in complex objects (maps and axis) the alignment of a value may not coincide with the bitwidth of a value.
// This keyword is used to define the alignment in the case of 64bit integers.
type alignmentInt64 struct {
	//AlignmentBorder describes the border at which the value is aligned to, i.e.
	//its memory address must be dividable by the value AlignmentBorder.
	AlignmentBorder    uint16
	AlignmentBorderSet bool
}

func parseAlignmentInt64(tok *tokenGenerator) (alignmentInt64, error) {
	ai64 := alignmentInt64{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("alignmentInt64 could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("alignmentInt64 could not be parsed")
	} else if !ai64.AlignmentBorderSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("alignmentInt64 alignmentBorder could not be parsed")
		}
		ai64.AlignmentBorder = uint16(buf)
		ai64.AlignmentBorderSet = true
		log.Info().Msg("alignmentInt64 alignmentBorder successfully parsed")
	}
	return ai64, err
}
