package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type DefaultValue struct {
	DisplayString    string
	DisplayStringSet bool
}

func parseDefaultValue(tok *tokenGenerator) (DefaultValue, error) {
	dv := DefaultValue{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("defaultValue could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("defaultValue could not be parsed")
	} else if !dv.DisplayStringSet {
		dv.DisplayString = tok.current()
		dv.DisplayStringSet = true
		log.Info().Msg("defaultValue displayString successfully parsed")
	}
	return dv, err
}
