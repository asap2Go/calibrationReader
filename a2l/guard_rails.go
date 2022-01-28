package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type guardRailsKeyword struct {
	value    bool
	valueSet bool
}

func parseGuardRails(tok *tokenGenerator) (guardRailsKeyword, error) {
	gr := guardRailsKeyword{}
	var err error
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("guardRails could not be parsed: unexpected end of file")
	} else if !gr.valueSet {
		gr.value = true
		gr.valueSet = true
		log.Info().Msg("guardRails value successfully parsed")
	}
	return gr, err
}
