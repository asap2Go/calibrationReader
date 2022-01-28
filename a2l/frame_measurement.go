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
			//there can be multiple identifiers
			//The end of the identifier listing is end the of the containing FRAME structure or the beginning of a IfData Structure
			fm.identifierSet = true
			log.Info().Msg("frameMeasurement identifier successfully parsed")
			//go back one token so the parseFrame Method is able to detect them
			tok.previous()
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("frameMeasurement could not be parsed")
			break forLoop
		} else if !fm.identifierSet {
			fm.identifier = append(fm.identifier, tok.current())
		}
	}
	return fm, err
}
