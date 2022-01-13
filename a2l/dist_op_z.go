package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type distOpZ struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseDistOpZ(tok *tokenGenerator) (distOpZ, error) {
	doz := distOpZ{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("distOpz could not be parsed")
			break forLoop
		} else if !doz.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("distOpz position could not be parsed")
				break forLoop
			}
			doz.position = uint16(buf)
			doz.positionSet = true
			log.Info().Msg("distOpz position successfully parsed")
		} else if !doz.datatypeSet {
			doz.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("distOpz datatype could not be parsed")
				break forLoop
			}
			doz.datatypeSet = true
			log.Info().Msg("distOpz datatype successfully parsed")
			break forLoop
		}
	}
	return doz, err
}
