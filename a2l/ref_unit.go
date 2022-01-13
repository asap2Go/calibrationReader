package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type refUnit struct {
	unit    string
	unitSet bool
}

func parseRefUnit(tok *tokenGenerator) (refUnit, error) {
	ru := refUnit{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("refUnit could not be parsed")
	} else if !ru.unitSet {
		ru.unit = tok.current()
		ru.unitSet = true
		log.Info().Msg("refUnit unit successfully parsed")
	}
	return ru, err
}
