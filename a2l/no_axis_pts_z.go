package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type NoAxisPtsZ struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
	Value       uint16
	ValueSet    bool
}

func parseNoAxisPtsZ(tok *tokenGenerator) (NoAxisPtsZ, error) {
	napz := NoAxisPtsZ{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("noAxisPtsz could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("noAxisPtsZ could not be parsed")
			break forLoop
		} else if !napz.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noAxisPtsz position could not be parsed")
				break forLoop
			}
			napz.Position = uint16(buf)
			napz.PositionSet = true
			log.Info().Msg("noAxisPtsz position successfully parsed")
		} else if !napz.DatatypeSet {
			napz.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noAxisPtsz datatype could not be parsed")
				break forLoop
			}
			napz.DatatypeSet = true
			log.Info().Msg("noAxisPtsz datatype successfully parsed")
			break forLoop
		}
	}
	return napz, err
}
