package a2l

import "github.com/rs/zerolog/log"

type staticRecordLayoutKeyword struct {
	value    bool
	valueSet bool
}

func parseStaticRecordLayout(tok *tokenGenerator) (staticRecordLayoutKeyword, error) {
	srl := staticRecordLayoutKeyword{}
	var err error
	if !srl.valueSet {
		srl.value = true
		srl.valueSet = true
			log.Info().Msg("staticRecordLayout value successfully parsed")
	}
	return srl, err
}
