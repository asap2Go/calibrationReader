package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type distOpX struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseDistOpX(tok *tokenGenerator) (distOpX, error) {
	dox := distOpX{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("distOpx could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("distOpX could not be parsed")
			break forLoop
		} else if !dox.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("distOpx position could not be parsed")
				break forLoop
			}
			dox.position = uint16(buf)
			dox.positionSet = true
			log.Info().Msg("distOpx position successfully parsed")
		} else if !dox.datatypeSet {
			dox.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("distOpx datatype could not be parsed")
				break forLoop
			}
			dox.datatypeSet = true
			log.Info().Msg("distOpx datatype successfully parsed")
			break forLoop
		}
	}
	return dox, err
}
