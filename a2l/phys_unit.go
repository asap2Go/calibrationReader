package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type physUnit struct {
	unit    string
	unitSet bool
}

func parsePhysUnit(tok *tokenGenerator) (physUnit, error) {
	pu := physUnit{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("physUnit could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("physUnit could not be parsed")
	} else if !pu.unitSet {
		pu.unit = tok.current()
		pu.unitSet = true
		log.Info().Msg("physUnit unit successfully parsed")
	}
	return pu, err
}
