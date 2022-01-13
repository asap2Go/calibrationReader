package a2l

import "github.com/rs/zerolog/log"

type readWriteKeyword struct {
	value    bool
	valueSet bool
}

func parseReadWrite(tok *tokenGenerator) (readWriteKeyword, error) {
	rw := readWriteKeyword{}
	var err error
	if !rw.valueSet {
		rw.value = true
		rw.valueSet = true
		log.Info().Msg("readWrite value successfully parsed")
	}
	return rw, err
}
