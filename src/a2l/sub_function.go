package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type subFunction struct {
	identifier    []string
	identifierSet bool
}

func parseSubFunction(tok *tokenGenerator) (subFunction, error) {
	sf := subFunction{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("subFunction could not be parsed")
			break forLoop
		} else if tok.current() == endSubFunctionToken {
			sf.identifierSet = true
				log.Info().Msg("subFunction identifier successfully parsed")
			break forLoop
		} else if !sf.identifierSet {
			sf.identifier = append(sf.identifier, tok.current())
		}
	}
	return sf, err
}
