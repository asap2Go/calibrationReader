package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type transformerInObject struct {
	identifier    []string
	identifierSet bool
}

func parseTransformerInObject(tok *tokenGenerator) (transformerInObject, error) {
	tio := transformerInObject{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("transformerInObject could not be parsed")
			break forLoop
		} else if tok.current() == endTransformerInObjectsToken {
			tio.identifierSet = true
			log.Info().Msg("transformerInObject identifier successfully parsed")
			break forLoop
		} else if !tio.identifierSet {
			tio.identifier = append(tio.identifier, tok.current())
		}
	}
	return tio, err
}
