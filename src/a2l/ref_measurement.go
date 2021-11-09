package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type refMeasurement struct {
	identifier    []string
	identifierSet bool
}

func parseRefMeasurement(tok *tokenGenerator) (refMeasurement, error) {
	rm := refMeasurement{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("refMeasurement could not be parsed")
			break forLoop
		} else if tok.current() == endRefMeasurementToken {
			rm.identifierSet = true
				log.Info().Msg("refMeasurement identifier successfully parsed")
			break forLoop
		} else if !rm.identifierSet {
			rm.identifier = append(rm.identifier, tok.current())
		}
	}
	return rm, err
}
