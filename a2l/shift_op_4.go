package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type shiftOp4 struct {
	position    uint16
	positionSet bool
	datatype    DataTypeEnum
	datatypeSet bool
}

func parseShiftOp4(tok *tokenGenerator) (shiftOp4, error) {
	so4 := shiftOp4{}
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
		} else if !so4.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("shiftOp4 position could not be parsed")
				break forLoop
			}
			so4.position = uint16(buf)
			so4.positionSet = true
			log.Info().Msg("shiftOp4 position successfully parsed")
		} else if !so4.datatypeSet {
			so4.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("shiftOp4 datatype could not be parsed")
				break forLoop
			}
			so4.datatypeSet = true
			log.Info().Msg("shiftOp4 datatype successfully parsed")
			break forLoop
		}
	}
	return so4, err
}
