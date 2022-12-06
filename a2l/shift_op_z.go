package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
ShiftOpZ is the shift operand in the deposit structure to compute the axis points for
fixed characteristic curves and fixed characteristic maps (see also keyword
FIX_AXIS_PAR).
The axis points distribution for fixed characteristic curves or fixed
characteristic maps is derived from the two 'offset' and 'shift' parameters as follows:

	for i = { 1...numberofaxispts }
	Xi = Offset + (i - 1)*2^Shift
*/
type ShiftOpZ struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseShiftOpZ(tok *tokenGenerator) (ShiftOpZ, error) {
	soz := ShiftOpZ{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("shiftOpz could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("shiftOpZ could not be parsed")
			break forLoop
		} else if !soz.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("shiftOpz position could not be parsed")
				break forLoop
			}
			soz.Position = uint16(buf)
			soz.PositionSet = true
			log.Info().Msg("shiftOpz position successfully parsed")
		} else if !soz.DatatypeSet {
			soz.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("shiftOpz datatype could not be parsed")
				break forLoop
			}
			soz.DatatypeSet = true
			log.Info().Msg("shiftOpz datatype successfully parsed")
			break forLoop
		}
	}
	return soz, err
}
