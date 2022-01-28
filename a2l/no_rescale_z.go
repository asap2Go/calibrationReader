package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type noRescaleZ struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseNoRescaleZ(tok *tokenGenerator) (noRescaleZ, error) {
	nrz := noRescaleZ{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("noRescalez could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("noRescaleZ could not be parsed")
			break forLoop
		} else if !nrz.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noRescalez position could not be parsed")
				break forLoop
			}
			nrz.position = uint16(buf)
			nrz.positionSet = true
			log.Info().Msg("noRescalez position successfully parsed")
		} else if !nrz.datatypeSet {
			nrz.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noRescalez datatype could not be parsed")
				break forLoop
			}
			nrz.datatypeSet = true
			log.Info().Msg("noRescalez datatype successfully parsed")
			break forLoop
		}
	}
	return nrz, err
}
