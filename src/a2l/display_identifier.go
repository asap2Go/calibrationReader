package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type DisplayIdentifier struct {
	displayName    string
	displayNameSet bool
}

func parseDisplayIdentifier(tok *tokenGenerator) (DisplayIdentifier, error) {
	di := DisplayIdentifier{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("displayIdentifier could not be parsed")
	} else if !di.displayNameSet {
		di.displayName = tok.current()
		di.displayNameSet = true
			log.Info().Msg("displayIdentifier displayName successfully parsed")
	}
	return di, err
}
