package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type functionVersion struct {
	versionIdentifier    string
	versionIdentifierSet bool
}

func parseFunctionVersion(tok *tokenGenerator) (functionVersion, error) {
	fv := functionVersion{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("functionVersion could not be parsed")
	} else if !fv.versionIdentifierSet {
		fv.versionIdentifier = tok.current()
		fv.versionIdentifierSet = true
		log.Info().Msg("functionVersion versionIdentifier successfully parsed")
	}
	return fv, err
}
