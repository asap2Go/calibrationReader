package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
)

type ecuAddress struct {
	address    string
	addressSet bool
}

func parseEcuAddress(tok *tokenGenerator) (ecuAddress, error) {
	ea := ecuAddress{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("ecuAddress could not be parsed")
	} else if !ea.addressSet {
		ea.address = tok.current()
		ea.addressSet = true
			log.Info().Msg("ecuAddress address successfully parsed")
	}
	return ea, err
}
