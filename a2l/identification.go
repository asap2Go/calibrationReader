package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

// Identification defines an 'identifier' in an adjustable object.
type Identification struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseIdentification(tok *tokenGenerator) (Identification, error) {
	i := Identification{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("identification could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("identification could not be parsed")
			break forLoop
		} else if !i.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("identification position could not be parsed")
				break forLoop
			}
			i.Position = uint16(buf)
			i.PositionSet = true
			log.Info().Msg("identification position successfully parsed")
		} else if !i.DatatypeSet {
			i.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("identification position could not be parsed")
				break forLoop
			}
			i.DatatypeSet = true
			log.Info().Msg("identification datatype successfully parsed")
			break forLoop
		}
	}
	return i, err
}
