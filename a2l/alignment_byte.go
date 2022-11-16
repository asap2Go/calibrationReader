package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

// alignmentByte is necessary because in complex objects (maps and axis) the alignment of a value may not coincide with the bitwidth of a value.
// This keyword is used to define the alignment in the case of bytes.
type alignmentByte struct {
	//AlignmentBorder describes the border at which the value is aligned to, i.e.
	//its memory address must be dividable by the value AlignmentBorder.
	AlignmentBorder    uint16
	AlignmentBorderSet bool
}

func parseAlignmentByte(tok *tokenGenerator) (alignmentByte, error) {
	ab := alignmentByte{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("alignmentByte could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("alignmentByte could not be parsed")
	} else if !ab.AlignmentBorderSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("alignmentByte alignmentBorder could not be parsed")
		}
		ab.AlignmentBorder = uint16(buf)
		ab.AlignmentBorderSet = true
		log.Info().Msg("alignmentByte alignmentBorder successfully parsed")
	}
	return ab, err
}
