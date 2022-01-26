package a2l

import "github.com/rs/zerolog/log"

type SignExtendKeyword struct {
	value    bool
	valueSet bool
}

func parseSignExtend(tok *tokenGenerator) (SignExtendKeyword, error) {
	se := SignExtendKeyword{}
	var err error
	if !se.valueSet {
		se.value = true
		se.valueSet = true
		log.Info().Msg("signExtend value successfully parsed")
	}
	return se, err
}
