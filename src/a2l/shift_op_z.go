package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type shiftOpZ struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseShiftOpZ(tok *tokenGenerator) (shiftOpZ, error) {
	soz := shiftOpZ{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("shiftOpz could not be parsed")
			break forLoop
		} else if !soz.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("shiftOpz position could not be parsed")
				break forLoop
			}
			soz.position = uint16(buf)
			soz.positionSet = true
			log.Info().Msg("shiftOpz position successfully parsed")
		} else if !soz.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("shiftOpz datatype could not be parsed")
				break forLoop
			}
			soz.datatype = buf
			soz.datatypeSet = true
			log.Info().Msg("shiftOpz datatype successfully parsed")
			break forLoop
		}
	}
	return soz, err
}
