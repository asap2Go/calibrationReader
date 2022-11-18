package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
OffsetY is the description of the 'offset' parameter in the deposit structure to compute the axis points for
fixed characteristic curves and fixed characteristic maps (see also keyword
FIX_AXIS_PAR). The axis points for fixed characteristic curves or fixed characteristic
maps are derived from the two 'offset' and 'shift' parameters as follows:
Xi = Offset + (i - 1)*2Shift i = { 1...numberofaxispts }
*/
type offsetY struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseOffsetY(tok *tokenGenerator) (offsetY, error) {
	oy := offsetY{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("offsety could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("offsetY could not be parsed")
			break forLoop
		} else if !oy.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("offsety position could not be parsed")
				break forLoop
			}
			oy.Position = uint16(buf)
			oy.PositionSet = true
			log.Info().Msg("offsety position successfully parsed")
		} else if !oy.DatatypeSet {
			oy.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("offsety datatype could not be parsed")
				break forLoop
			}
			oy.DatatypeSet = true
			log.Info().Msg("offsety datatype successfully parsed")
			break forLoop
		}
	}
	return oy, err
}
