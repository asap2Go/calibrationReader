package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
ShiftOpY is the shift operand in the deposit structure to compute the axis points for
fixed characteristic curves and fixed characteristic maps (see also keyword
FIX_AXIS_PAR).
The axis points distribution for fixed characteristic curves or fixed
characteristic maps is derived from the two 'offset' and 'shift' parameters as follows:

	for i = { 1...numberofaxispts }
	Xi = Offset + (i - 1)*2^Shift
*/
type ShiftOpY struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseShiftOpY(tok *tokenGenerator) (ShiftOpY, error) {
	soy := ShiftOpY{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("shiftOpY could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("shiftOpY could not be parsed")
			break forLoop
		} else if !soy.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("shiftOpY position could not be parsed")
				break forLoop
			}
			soy.Position = uint16(buf)
			soy.PositionSet = true
			log.Info().Msg("shiftOpY position successfully parsed")
		} else if !soy.DatatypeSet {
			soy.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("shiftOpY datatype could not be parsed")
				break forLoop
			}
			soy.DatatypeSet = true
			log.Info().Msg("shiftOpY datatype successfully parsed")
			break forLoop
		}
	}
	return soy, err
}
