package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type NoAxisPtsX struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
	Value       uint16
	ValueSet    bool
}

func parseNoAxisPtsX(tok *tokenGenerator) (NoAxisPtsX, error) {
	napx := NoAxisPtsX{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("noAxisPtsx could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("noAxisPtsX could not be parsed")
			break forLoop
		} else if !napx.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noAxisPtsx position could not be parsed")
				break forLoop
			}
			napx.Position = uint16(buf)
			napx.PositionSet = true
			log.Info().Msg("noAxisPtsx position successfully parsed")
		} else if !napx.DatatypeSet {
			napx.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noAxisPtsx datatype could not be parsed")
				break forLoop
			}
			napx.DatatypeSet = true
			log.Info().Msg("noAxisPtsx datatype successfully parsed")
			break forLoop
		}
	}
	return napx, err
}
