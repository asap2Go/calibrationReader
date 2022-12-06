package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
ShiftOp4 is the shift operand in the deposit structure to compute the axis points for
fixed characteristic curves and fixed characteristic maps (see also keyword
FIX_AXIS_PAR).
The axis points distribution for fixed characteristic curves or fixed
characteristic maps is derived from the two 'offset' and 'shift' parameters as follows:

	for i = { 1...numberofaxispts }
	Xi = Offset + (i - 1)*2^Shift
*/
type ShiftOp4 struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseShiftOp4(tok *tokenGenerator) (ShiftOp4, error) {
	so4 := ShiftOp4{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("shiftOp4 could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("shiftOp4 could not be parsed")
			break forLoop
		} else if !so4.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("shiftOp4 position could not be parsed")
				break forLoop
			}
			so4.Position = uint16(buf)
			so4.PositionSet = true
			log.Info().Msg("shiftOp4 position successfully parsed")
		} else if !so4.DatatypeSet {
			so4.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("shiftOp4 datatype could not be parsed")
				break forLoop
			}
			so4.DatatypeSet = true
			log.Info().Msg("shiftOp4 datatype successfully parsed")
			break forLoop
		}
	}
	return so4, err
}
