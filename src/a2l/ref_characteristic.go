package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type refCharacteristic struct {
	identifier    []string
	identifierSet bool
}

func parseRefCharacteristic(tok *tokenGenerator) (refCharacteristic, error) {
	rc := refCharacteristic{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("refCharacteristic could not be parsed")
			break forLoop
		} else if tok.current() == endRefCharacteristicToken {
			rc.identifierSet = true
				log.Info().Msg("refCharacteristic identifier successfully parsed")
			break forLoop
		} else if !rc.identifierSet {
			rc.identifier = append(rc.identifier, tok.current())
		}
	}
	return rc, err
}
