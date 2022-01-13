package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type subGroup struct {
	identifier    []string
	identifierSet bool
}

func parseSubGroup(tok *tokenGenerator) (subGroup, error) {
	sg := subGroup{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("subGroup could not be parsed")
			break forLoop
		} else if tok.current() == endSubGroupToken {
			sg.identifierSet = true
			log.Info().Msg("subGroup identifier successfully parsed")
			break forLoop
		} else if !sg.identifierSet {
			sg.identifier = append(sg.identifier, tok.current())
		}
	}
	return sg, err
}
