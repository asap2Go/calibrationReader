package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type inMeasurement struct {
	identifier    []string
	identifierSet bool
}

func parseInMeasurement(tok *tokenGenerator) (inMeasurement, error) {
	im := inMeasurement{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("inMeasurement could not be parsed")
			break forLoop
		} else if tok.current() == endInMeasurementToken {
			im.identifierSet = true
				log.Info().Msg("inMeasurement identifier successfully parsed")
			break forLoop
		} else if !im.identifierSet {
			im.identifier = append(im.identifier, tok.current())
		}
	}
	return im, err
}
