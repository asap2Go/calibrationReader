package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

//alignmentFloat32Ieee is necessary because in complex objects (maps and axis) the alignment of a value may not coincide with the bitwidth of a value.
//This keyword is used to define the alignment in the case of longs.
type alignmentLong struct {
	//alignmentBorder describes the border at which the value is aligned to, i.e.
	//its memory address must be dividable by the value AlignmentBorder.
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
