package a2l

import "github.com/rs/zerolog/log"

type SignExtend struct {
	value    bool
	valueSet bool
}

func parseSignExtend(tok *tokenGenerator) (SignExtend, error) {
	se := SignExtend{}
	var err error
	if !se.valueSet {
		se.value = true
		se.valueSet = true
			log.Info().Msg("signExtend value successfully parsed")
	}
	return se, err
}
