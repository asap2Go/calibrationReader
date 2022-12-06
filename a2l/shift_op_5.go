package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
ShiftOp5 is the shift operand in the deposit structure to compute the axis points for
fixed characteristic curves and fixed characteristic maps (see also keyword
FIX_AXIS_PAR).
The axis points distribution for fixed characteristic curves or fixed
characteristic maps is derived from the two 'offset' and 'shift' parameters as follows:

	for i = { 1...numberofaxispts }
	Xi = Offset + (i - 1)*2^Shift
*/
type ShiftOp5 struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseShiftOp5(tok *tokenGenerator) (ShiftOp5, error) {
	so5 := ShiftOp5{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("shiftOp5 could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("shiftOp5 could not be parsed")
			break forLoop
		} else if !so5.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("shiftOp5 position could not be parsed")
				break forLoop
			}
			so5.Position = uint16(buf)
			so5.PositionSet = true
			log.Info().Msg("shiftOp5 position successfully parsed")
		} else if !so5.DatatypeSet {
			so5.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("shiftOp5 datatype could not be parsed")
				break forLoop
			}
			so5.DatatypeSet = true
			log.Info().Msg("shiftOp5 datatype successfully parsed")
			break forLoop
		}
	}
	return so5, err
}
