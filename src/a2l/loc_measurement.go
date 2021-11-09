package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type locMeasurement struct {
	identifier    []string
	identifierSet bool
}

func parseLocMeasurement(tok *tokenGenerator) (locMeasurement, error) {
	lm := locMeasurement{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("locMeasurement could not be parsed")
			break forLoop
		} else if tok.current() == endLocMeasurementToken {
			lm.identifierSet = true
				log.Info().Msg("locMeasurement identifier successfully parsed")
			break forLoop
		} else if !lm.identifierSet {
			lm.identifier = append(lm.identifier, tok.current())
		}
	}
	return lm, err
}
