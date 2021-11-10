package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type shiftOpX struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseShiftOpX(tok *tokenGenerator) (shiftOpX, error) {
	sox := shiftOpX{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("shiftOpx could not be parsed")
			break forLoop
		} else if !sox.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("shiftOpx position could not be parsed")
				break forLoop
			}
			sox.position = uint16(buf)
			sox.positionSet = true
			log.Info().Msg("shiftOpx position successfully parsed")
		} else if !sox.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("shiftOpx datatype could not be parsed")
				break forLoop
			}
			sox.datatype = buf
			sox.datatypeSet = true
			log.Info().Msg("shiftOpx datatype successfully parsed")
			break forLoop
		}
	}
	return sox, err
}
