package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type DependentCharacteristic struct {
	formula           string
	formulaSet        bool
	characteristic    []string
	characteristicSet bool
}

func parseDependentCharacteristic(tok *tokenGenerator) (DependentCharacteristic, error) {
	dc := DependentCharacteristic{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("dependentCharacteristic could not be parsed")
			break forLoop
		} else if tok.current() == endDependentCharacteristicToken {
			dc.characteristicSet = true
			log.Info().Msg("dependentCharacteristic characteristic successfully parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("dependentCharacteristic could not be parsed")
			break forLoop
		} else if !dc.formulaSet {
			dc.formula = tok.current()
			dc.formulaSet = true
			log.Info().Msg("dependentCharacteristic formula successfully parsed")
		} else if !dc.characteristicSet {
			dc.characteristic = append(dc.characteristic, tok.current())
		}

	}
	return dc, err
}
