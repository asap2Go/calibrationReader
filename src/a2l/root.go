package a2l

import "github.com/rs/zerolog/log"

type rootKeyword struct {
	value    bool
	valueSet bool
}

func parseRoot(tok *tokenGenerator) (rootKeyword, error) {
	r := rootKeyword{}
	var err error
	if !r.valueSet {
		r.value = true
		r.valueSet = true
			log.Info().Msg("root value successfully parsed")
	}
	return r, err
}
