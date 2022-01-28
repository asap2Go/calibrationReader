package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type shiftOpY struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseShiftOpY(tok *tokenGenerator) (shiftOpY, error) {
	soy := shiftOpY{}
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
		} else if !soy.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("shiftOpY position could not be parsed")
				break forLoop
			}
			soy.position = uint16(buf)
			soy.positionSet = true
			log.Info().Msg("shiftOpY position successfully parsed")
		} else if !soy.datatypeSet {
			soy.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("shiftOpY datatype could not be parsed")
				break forLoop
			}
			soy.datatypeSet = true
			log.Info().Msg("shiftOpY datatype successfully parsed")
			break forLoop
		}
	}
	return soy, err
}
