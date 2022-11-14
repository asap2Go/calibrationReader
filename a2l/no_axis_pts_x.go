package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type noAxisPtsX struct {
	position    uint16
	positionSet bool
	datatype    DataTypeEnum
	datatypeSet bool
}

func parseNoAxisPtsX(tok *tokenGenerator) (noAxisPtsX, error) {
	napx := noAxisPtsX{}
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
		} else if !napx.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noAxisPtsx position could not be parsed")
				break forLoop
			}
			napx.position = uint16(buf)
			napx.positionSet = true
			log.Info().Msg("noAxisPtsx position successfully parsed")
		} else if !napx.datatypeSet {
			napx.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noAxisPtsx datatype could not be parsed")
				break forLoop
			}
			napx.datatypeSet = true
			log.Info().Msg("noAxisPtsx datatype successfully parsed")
			break forLoop
		}
	}
	return napx, err
}
