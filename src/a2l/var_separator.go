package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type varSeparator struct {
	separator    string
	separatorSet bool
}

func parseVarSeparator(tok *tokenGenerator) (varSeparator, error) {
	vs := varSeparator{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("varSeparator could not be parsed")
	} else if !vs.separatorSet {
		vs.separator = tok.current()
		vs.separatorSet = true
			log.Info().Msg("varSeparator separator successfully parsed")
	}
	return vs, err
}
