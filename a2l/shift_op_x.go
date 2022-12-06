package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
ShiftOpX is the shift operand in the deposit structure to compute the axis points for
fixed characteristic curves and fixed characteristic maps (see also keyword
FIX_AXIS_PAR).
The axis points distribution for fixed characteristic curves or fixed
characteristic maps is derived from the two 'offset' and 'shift' parameters as follows:

	for i = { 1...numberofaxispts }
	Xi = Offset + (i - 1)*2^Shift
*/
type ShiftOpX struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseShiftOpX(tok *tokenGenerator) (ShiftOpX, error) {
	sox := ShiftOpX{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("shiftOpx could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("shiftOpX could not be parsed")
			break forLoop
		} else if !sox.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("shiftOpx position could not be parsed")
				break forLoop
			}
			sox.Position = uint16(buf)
			sox.PositionSet = true
			log.Info().Msg("shiftOpx position successfully parsed")
		} else if !sox.DatatypeSet {
			sox.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("shiftOpx datatype could not be parsed")
				break forLoop
			}
			sox.DatatypeSet = true
			log.Info().Msg("shiftOpx datatype successfully parsed")
			break forLoop
		}
	}
	return sox, err
}
