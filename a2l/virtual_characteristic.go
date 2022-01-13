package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type VirtualCharacteristic struct {
	formula           string
	formulaSet        bool
	characteristic    []string
	characteristicSet bool
}

func parseVirtualCharacteristic(tok *tokenGenerator) (VirtualCharacteristic, error) {
	vc := VirtualCharacteristic{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("virtualCharacteristic could not be parsed")
			break forLoop
		} else if tok.current() == endVirtualCharacteristicToken {
			vc.characteristicSet = true
			log.Info().Msg("virtualCharacteristic characteristic successfully parsed")
			break forLoop
		} else if !vc.formulaSet {
			vc.formula = tok.current()
			vc.formulaSet = true
			log.Info().Msg("virtualCharacteristic formula successfully parsed")
		} else if !vc.characteristicSet {
			vc.characteristic = append(vc.characteristic, tok.current())
		}
	}
	return vc, err
}
