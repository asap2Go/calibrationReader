package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
OffsetX is the description of the 'offset' parameter in the deposit structure to compute the axis points for
fixed characteristic curves and fixed characteristic maps (see also keyword
FIX_AXIS_PAR). The axis points for fixed characteristic curves or fixed characteristic
maps are derived from the two 'offset' and 'shift' parameters as follows:
Xi = Offset + (i - 1)*2Shift i = { 1...numberofaxispts }
*/
type OffsetX struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseOffsetX(tok *tokenGenerator) (OffsetX, error) {
	ox := OffsetX{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("offsetx could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("offsetX could not be parsed")
			break forLoop
		} else if !ox.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("offsetx position could not be parsed")
				break forLoop
			}
			ox.Position = uint16(buf)
			ox.PositionSet = true
			log.Info().Msg("offsetx position successfully parsed")
		} else if !ox.DatatypeSet {
			ox.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("offsetx datatype could not be parsed")
				break forLoop
			}
			ox.DatatypeSet = true
			log.Info().Msg("offsetx datatype successfully parsed")
			break forLoop
		}
	}
	return ox, err
}
