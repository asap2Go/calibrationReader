package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type alignmentFloat32Ieee struct {
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
