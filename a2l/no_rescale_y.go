package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type noRescaleY struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseNoRescaleY(tok *tokenGenerator) (noRescaleY, error) {
	nry := noRescaleY{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("noRescaley could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("noRescaleY could not be parsed")
			break forLoop
		} else if !nry.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noRescaley position could not be parsed")
				break forLoop
			}
			nry.position = uint16(buf)
			nry.positionSet = true
			log.Info().Msg("noRescaley position successfully parsed")
		} else if !nry.datatypeSet {
			nry.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noRescaley datatype could not be parsed")
				break forLoop
			}
			nry.datatypeSet = true
			log.Info().Msg("noRescaley datatype successfully parsed")
			break forLoop
		}
	}
	return nry, err
}
