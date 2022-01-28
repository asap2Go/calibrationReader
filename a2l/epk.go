package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type epk struct {
	identifier    string
	identifierSet bool
}

func parseEpk(tok *tokenGenerator) (epk, error) {
	e := epk{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("epk could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("epk could not be parsed")
	} else if !e.identifierSet {
		e.identifier = tok.current()
		e.identifierSet = true
		log.Info().Msg("epk identifier successfully parsed")
	}
	return e, err
}
