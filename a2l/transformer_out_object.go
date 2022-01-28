package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type transformerOutObject struct {
	identifier    []string
	identifierSet bool
}

func parseTransformerOutObject(tok *tokenGenerator) (transformerOutObject, error) {
	tio := transformerOutObject{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("transformerOutObject could not be parsed")
			break forLoop
		} else if tok.current() == endTransformerOutObjectsToken {
			tio.identifierSet = true
			log.Info().Msg("transformerOutObject identifier successfully parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("transformerOutObject could not be parsed")
			break forLoop
		} else if !tio.identifierSet {
			tio.identifier = append(tio.identifier, tok.current())
		}
	}
	return tio, err
}
