package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type frameMeasurement struct {
	identifier    []string
	identifierSet bool
}

func parseFrameMeasurement(tok *tokenGenerator) (frameMeasurement, error) {
	fm := frameMeasurement{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("frameMeasurement could not be parsed")
			break forLoop
		} else if tok.current() == endFrameToken || tok.current() == beginIfDataToken {
			fm.identifierSet = true
				log.Info().Msg("frameMeasurement identifier successfully parsed")
			break forLoop
		} else if !fm.identifierSet {
			fm.identifier = append(fm.identifier, tok.current())
		}
	}
	return fm, err
}
