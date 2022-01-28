package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type refGroup struct {
	identifier    []string
	identifierSet bool
}

func parseRefGroup(tok *tokenGenerator) (refGroup, error) {
	rg := refGroup{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("refGroup could not be parsed")
			break forLoop
		} else if tok.current() == endRefGroupToken {
			rg.identifierSet = true
			log.Info().Msg("refGroup identifier successfully parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("refGroup could not be parsed")
			break forLoop
		} else if !rg.identifierSet {
			rg.identifier = append(rg.identifier, tok.current())
		}
	}
	return rg, err
}
