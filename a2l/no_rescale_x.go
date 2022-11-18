package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

// NoRescaleX defines the actual number of rescaling axis point value pairs.
type NoRescaleX struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseNoRescaleX(tok *tokenGenerator) (NoRescaleX, error) {
	nrx := NoRescaleX{}
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
		} else if !nrx.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noRescalex position could not be parsed")
				break forLoop
			}
			nrx.Position = uint16(buf)
			nrx.PositionSet = true
			log.Info().Msg("noRescalex position successfully parsed")
			break forLoop
		} else if !nrx.DatatypeSet {
			nrx.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noRescalex datatype could not be parsed")
				break forLoop
			}
			nrx.DatatypeSet = true
			log.Info().Msg("noRescalex datatype successfully parsed")
		}
	}
	return nrx, err
}
