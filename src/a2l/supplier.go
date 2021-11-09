package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type supplier struct {
	manufacturer    string
	manufacturerSet bool
}

func parseSupplier(tok *tokenGenerator) (supplier, error) {
	s := supplier{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("supplier could not be parsed")
	} else if !s.manufacturerSet {
		s.manufacturer = tok.current()
		s.manufacturerSet = true
			log.Info().Msg("supplier manufacturer successfully parsed")
	}
	return s, err
}
