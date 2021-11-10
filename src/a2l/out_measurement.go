package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type outMeasurement struct {
	identifier    []string
	identifierSet bool
}

func parseOutMeasurement(tok *tokenGenerator) (outMeasurement, error) {
	om := outMeasurement{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("outMeasurement could not be parsed")
			break forLoop
		} else if tok.current() == endOutMeasurementToken {
			om.identifierSet = true
			log.Info().Msg("outMeasurement identifier successfully parsed")
			break forLoop
		} else if !om.identifierSet {
			om.identifier = append(om.identifier, tok.current())
		}
	}
	return om, err
}
