package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type distOp5 struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseDistOp5(tok *tokenGenerator) (distOp5, error) {
	do := distOp5{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("distOp5 could not be parsed")
			break forLoop
		} else if !do.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
					log.Err(err).Msg("distOp5 position could not be parsed")
				break forLoop
			}
			do.position = uint16(buf)
			do.positionSet = true
				log.Info().Msg("distOp5 position successfully parsed")
		} else if !do.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
					log.Err(err).Msg("distOp5 datatype could not be parsed")
				break forLoop
			}
			do.datatype = buf
			do.datatypeSet = true
				log.Info().Msg("distOp5 datatype successfully parsed")
			break forLoop
		}
	}
	return do, err
}
