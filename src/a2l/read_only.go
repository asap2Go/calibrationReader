package a2l

import "github.com/rs/zerolog/log"

type readOnlyKeyword struct {
	value    bool
	valueSet bool
}

func parseReadOnly(tok *tokenGenerator) (readOnlyKeyword, error) {
	ro := readOnlyKeyword{}
	var err error
	if !ro.valueSet {
		ro.value = true
		ro.valueSet = true
		log.Info().Msg("readOnly value successfully parsed")
	}
	return ro, err
}
