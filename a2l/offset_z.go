package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
OffsetZ is the description of the 'offset' parameter in the deposit structure to compute the axis points for
fixed characteristic curves and fixed characteristic maps (see also keyword
FIX_AXIS_PAR). The axis points for fixed characteristic curves or fixed characteristic
maps are derived from the two 'offset' and 'shift' parameters as follows:
Xi = Offset + (i - 1)*2Shift i = { 1...numberofaxispts }
*/
type OffsetZ struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseOffsetZ(tok *tokenGenerator) (OffsetZ, error) {
	oz := OffsetZ{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("offsetz could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("offsetZ could not be parsed")
			break forLoop
		} else if !oz.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("offsetz position could not be parsed")
				break forLoop
			}
			oz.Position = uint16(buf)
			oz.PositionSet = true
			log.Info().Msg("offsetz position successfully parsed")
		} else if !oz.DatatypeSet {
			oz.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("offsetz datatype could not be parsed")
				break forLoop
			}
			oz.DatatypeSet = true
			log.Info().Msg("offsetz datatype successfully parsed")
			break forLoop
		}
	}
	return oz, err
}
