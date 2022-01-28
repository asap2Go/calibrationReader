package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type varSelectionCharacteristic struct {
	name    string
	nameSet bool
}

func parseVarSelectionCharacteristic(tok *tokenGenerator) (varSelectionCharacteristic, error) {
	vsc := varSelectionCharacteristic{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("varSelectionCharacteristic could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("varSelectionCharacteristic could not be parsed")
	} else if !vsc.nameSet {
		vsc.name = tok.current()
		vsc.nameSet = true
		log.Info().Msg("varSelectionCharacteristic name successfully parsed")
	}
	return vsc, err
}
