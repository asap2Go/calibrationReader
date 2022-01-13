package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type GuardRails struct {
	value    bool
	valueSet bool
}

func parseGuardRails(tok *tokenGenerator) (GuardRails, error) {
	gr := GuardRails{}
	var err error
	tok.next()
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
