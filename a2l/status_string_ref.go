package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type statusStringRef struct {
	conversionTable    string
	conversionTableSet bool
}

func parseStatusStringRef(tok *tokenGenerator) (statusStringRef, error) {
	ssr := statusStringRef{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("statusStringRef could not be parsed")
	} else if !ssr.conversionTableSet {
		ssr.conversionTable = tok.current()
		ssr.conversionTableSet = true
		log.Info().Msg("statusStringRef conversionTable successfully parsed")
	}
	return ssr, err
}
