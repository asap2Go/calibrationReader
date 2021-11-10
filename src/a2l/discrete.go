package a2l

import "github.com/rs/zerolog/log"

type discreteKeyword struct {
	value    bool
	valueSet bool
}

func parseDiscrete(tok *tokenGenerator) (discreteKeyword, error) {
	d := discreteKeyword{}
	var err error
	if !d.valueSet {
		d.value = true
		d.valueSet = true
		log.Info().Msg("discrete value successfully parsed")
	}
	return d, err
}
