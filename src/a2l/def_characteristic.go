package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type defCharacteristic struct {
	identifier    []string
	identifierSet bool
}

func parseDefCharacteristic(tok *tokenGenerator) (defCharacteristic, error) {
	dc := defCharacteristic{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("defCharacteristic could not be parsed")
			break forLoop
		} else if tok.current() == endDefCharacteristicToken {
			dc.identifierSet = true
				log.Info().Msg("defCharacteristic identifier successfully parsed")
			break forLoop
		} else if !dc.identifierSet {
			dc.identifier = append(dc.identifier, tok.current())
		}
	}
	return dc, err
}
