package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type format struct {
	formatString    string
	formatStringSet bool
}

func parseFormat(tok *tokenGenerator) (format, error) {
	f := format{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("format could not be parsed")
	} else if !f.formatStringSet {
		f.formatString = tok.current()
		f.formatStringSet = true
			log.Info().Msg("format formatString successfully parsed")
	}
	return f, err
}
