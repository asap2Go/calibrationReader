package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
Offset4 is the description of the 'offset' parameter in the deposit structure to compute the axis points for
fixed characteristic curves and fixed characteristic maps (see also keyword
FIX_AXIS_PAR). The axis points for fixed characteristic curves or fixed characteristic
maps are derived from the two 'offset' and 'shift' parameters as follows:
Xi = Offset + (i - 1)*2Shift i = { 1...numberofaxispts }
*/
type Offset4 struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseOffset4(tok *tokenGenerator) (Offset4, error) {
	o4 := Offset4{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("offset4 could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("offset4 could not be parsed")
			break forLoop
		} else if !o4.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("offset4 position could not be parsed")
				break forLoop
			}
			o4.Position = uint16(buf)
			o4.PositionSet = true
			log.Info().Msg("offset4 position successfully parsed")
		} else if !o4.DatatypeSet {
			o4.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("offset4 datatype could not be parsed")
				break forLoop
			}
			o4.DatatypeSet = true
			log.Info().Msg("offset4 datatype successfully parsed")
			break forLoop
		}
	}
	return o4, err
}
