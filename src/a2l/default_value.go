package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type defaultValue struct {
	displayString    string
	displayStringSet bool
}

func parseDefaultValue(tok *tokenGenerator) (defaultValue, error) {
	dv := defaultValue{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("defaultValue could not be parsed")
	} else if !dv.displayStringSet {
		dv.displayString = tok.current()
		dv.displayStringSet = true
		log.Info().Msg("defaultValue displayString successfully parsed")
	}
	return dv, err
}
