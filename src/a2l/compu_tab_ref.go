package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type compuTabRef struct {
	conversionTable    string
	conversionTableSet bool
}

func parseCompuTabRef(tok *tokenGenerator) (compuTabRef, error) {
	ctr := compuTabRef{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("compuTabRef could not be parsed")
	} else if !ctr.conversionTableSet {
		ctr.conversionTable = tok.current()
		ctr.conversionTableSet = true
	}
	return ctr, err
}
