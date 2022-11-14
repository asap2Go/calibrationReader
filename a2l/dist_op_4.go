package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type distOp4 struct {
	position    uint16
	positionSet bool
	datatype    DataTypeEnum
	datatypeSet bool
}

func parseDistOp4(tok *tokenGenerator) (distOp4, error) {
	do := distOp4{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("distOp4 could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("distOp4 could not be parsed")
			break forLoop
		} else if !do.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("distOp4 position could not be parsed")
				break forLoop
			}
			do.position = uint16(buf)
			do.positionSet = true
			log.Info().Msg("distOp4 position successfully parsed")
		} else if !do.datatypeSet {
			do.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("distOp4 datatype could not be parsed")
				break forLoop
			}
			do.datatypeSet = true
			log.Info().Msg("distOp4 datatype successfully parsed")
			break forLoop
		}
	}
	return do, err
}
