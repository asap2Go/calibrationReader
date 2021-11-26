package a2l

import "github.com/rs/zerolog/log"

type consistentExchangeKeyword struct {
	value    bool
	valueSet bool
}

func parseConsistentExchangeKeyword(tok *tokenGenerator) (consistentExchangeKeyword, error) {
	ce := consistentExchangeKeyword{}
	var err error
	if !ce.valueSet {
		ce.value = true
		ce.valueSet = true
		log.Info().Msg("consistentExchange value successfully parsed")
	}
	return ce, err
}
