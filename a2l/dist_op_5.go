package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
DistOp5
uint Position Position of the distance operand in the deposit structure.
datatype Datatype Data type of the distance operand.
Description:
Description of the distance operand in the deposit structure to compute the axis points for
fixed characteristic curves and fixed characteristic maps (see also keyword
FIX_AXIS_PAR_DIST). The axis points distribution for fixed characteristic curves or fixed
characteristic maps is derived from the two 'offset' and 'distance' parameters as follows:
Xi = Offset + (i - 1)*Distance i = { 1...numberofaxispts }
or
Yk = Offset + (k - 1)* Distance k = { 1...numberofaxispts }
or
Zm = Offset + (m - 1)* Distance m = { 1...numberofaxispts }
or
Z4n = Offset + (n - 1)* Distance n = { 1...numberofaxispts }
or
Z5o = Offset + (o - 1)* Distance o = { 1...numberofaxispts }
*/
type DistOp5 struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseDistOp5(tok *tokenGenerator) (DistOp5, error) {
	do := DistOp5{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("distOp5 could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("distOp5 could not be parsed")
			break forLoop
		} else if !do.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("distOp5 position could not be parsed")
				break forLoop
			}
			do.Position = uint16(buf)
			do.PositionSet = true
			log.Info().Msg("distOp5 position successfully parsed")
		} else if !do.DatatypeSet {
			do.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("distOp5 datatype could not be parsed")
				break forLoop
			}
			do.DatatypeSet = true
			log.Info().Msg("distOp5 datatype successfully parsed")
			break forLoop
		}
	}
	return do, err
}
