package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

//alignmentFloat32Ieee is necessary because in complex objects (maps and axis) the alignment of a value may not coincide with the bitwidth of a value.
//This keyword is used to define the alignment in the case of 32bit floats.
type alignmentFloat32Ieee struct {
	//alignmentBorder describes the border at which the value is aligned to, i.e.
	//its memory address must be dividable by the value AlignmentBorder.
	alignmentBorder    uint16
	alignmentBorderSet bool
}

func parseAlignmentFloat32Ieee(tok *tokenGenerator) (alignmentFloat32Ieee, error) {
	af32 := alignmentFloat32Ieee{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("alignmentFloat32Ieee could not be parsed")
	} else if !af32.alignmentBorderSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("alignmentFloat32Ieee alignmentBorder could not be parsed")
		}
		af32.alignmentBorder = uint16(buf)
		af32.alignmentBorderSet = true
		log.Info().Msg("alignmentFloat32Ieee alignmentBorder successfully parsed")
	}
	return af32, err
}
