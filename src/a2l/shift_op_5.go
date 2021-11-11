package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type shiftOp5 struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseShiftOp5(tok *tokenGenerator) (shiftOp5, error) {
	so5 := shiftOp5{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("shiftOp5 could not be parsed")
			break forLoop
		} else if !so5.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("shiftOp5 position could not be parsed")
				break forLoop
			}
			so5.position = uint16(buf)
			so5.positionSet = true
			log.Info().Msg("shiftOp5 position successfully parsed")
		} else if !so5.datatypeSet {
			so5.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("shiftOp5 datatype could not be parsed")
				break forLoop
			}
			so5.datatypeSet = true
			log.Info().Msg("shiftOp5 datatype successfully parsed")
			break forLoop
		}
	}
	return so5, err
}
