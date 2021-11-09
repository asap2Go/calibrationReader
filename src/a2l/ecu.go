package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type ecu struct {
	controlUnit    string
	controlUnitSet bool
}

func parseEcu(tok *tokenGenerator) (ecu, error) {
	e := ecu{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("ecu could not be parsed")
	} else if !e.controlUnitSet {
		e.controlUnit = tok.current()
		e.controlUnitSet = true
			log.Info().Msg("ecu controlUnit successfully parsed")
	}
	return e, err
}
