package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type DistOpZ struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseDistOpZ(tok *tokenGenerator) (DistOpZ, error) {
	doz := DistOpZ{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("distOpz could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("distOpZ could not be parsed")
			break forLoop
		} else if !doz.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("distOpz position could not be parsed")
				break forLoop
			}
			doz.Position = uint16(buf)
			doz.PositionSet = true
			log.Info().Msg("distOpz position successfully parsed")
		} else if !doz.DatatypeSet {
			doz.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("distOpz datatype could not be parsed")
				break forLoop
			}
			doz.DatatypeSet = true
			log.Info().Msg("distOpz datatype successfully parsed")
			break forLoop
		}
	}
	return doz, err
}
