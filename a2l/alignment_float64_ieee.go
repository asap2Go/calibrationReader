package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

// alignmentFloat64Ieee is necessary because in complex objects (maps and axis) the alignment of a value may not coincide with the bitwidth of a value.
// This keyword is used to define the alignment in the case of 64bit floats.
type alignmentFloat64Ieee struct {
	//AlignmentBorder describes the border at which the value is aligned to, i.e.
	//its memory address must be dividable by the value AlignmentBorder.
	AlignmentBorder    uint16
	AlignmentBorderSet bool
}

func parseAlignmentFloat64Ieee(tok *tokenGenerator) (alignmentFloat64Ieee, error) {
	af64 := alignmentFloat64Ieee{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("alignmentFloat64Ieee could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("alignmentFloat64Ieee could not be parsed")
	} else if !af64.AlignmentBorderSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("alignmentFloat64Ieee alignmentBorder could not be parsed")
		}
		af64.AlignmentBorder = uint16(buf)
		af64.AlignmentBorderSet = true
		log.Info().Msg("alignmentFloat64Ieee alignmentBorder successfully parsed")
	}
	return af64, err
}
