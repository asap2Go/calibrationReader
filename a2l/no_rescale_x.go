package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type noRescaleX struct {
	position    uint16
	positionSet bool
	datatype    DataTypeEnum
	datatypeSet bool
}

func parseNoRescaleX(tok *tokenGenerator) (noRescaleX, error) {
	nrx := noRescaleX{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("noRescalex could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("noRescaleX could not be parsed")
			break forLoop
		} else if !nrx.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noRescalex position could not be parsed")
				break forLoop
			}
			nrx.position = uint16(buf)
			nrx.positionSet = true
			log.Info().Msg("noRescalex position successfully parsed")
			break forLoop
		} else if !nrx.datatypeSet {
			nrx.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noRescalex datatype could not be parsed")
				break forLoop
			}
			nrx.datatypeSet = true
			log.Info().Msg("noRescalex datatype successfully parsed")
		}
	}
	return nrx, err
}
