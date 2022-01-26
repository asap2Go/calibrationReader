package a2l

import "github.com/rs/zerolog/log"

type staticAddressOffsetsKeyword struct {
	value    bool
	valueSet bool
}

func parseStaticAddressOffsets(tok *tokenGenerator) (staticAddressOffsetsKeyword, error) {
	sao := staticAddressOffsetsKeyword{}
	var err error
	if !sao.valueSet {
		sao.value = true
		sao.valueSet = true
		log.Info().Msg("StaticAddressOffsets value successfully parsed")
	}
	return sao, err
}
