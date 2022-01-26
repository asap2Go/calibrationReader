package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type errorMask struct {
	mask    string
	maskSet bool
}

func parseErrorMask(tok *tokenGenerator) (errorMask, error) {
	em := errorMask{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("errorMask could not be parsed")
	} else if !em.maskSet {
		em.mask = tok.current()
		em.maskSet = true
		log.Info().Msg("errorMask mask successfully parsed")
	}
	return em, err
}
