package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type version struct {
	versionIdentifier    string
	versionIdentifierSet bool
}

func parseVersion(tok *tokenGenerator) (version, error) {
	v := version{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("version could not be parsed")
	} else if !v.versionIdentifierSet {
		v.versionIdentifier = tok.current()
		v.versionIdentifierSet = true
			log.Info().Msg("version versionIdentifier successfully parsed")
	}
	return v, err
}
