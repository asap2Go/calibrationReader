package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type CompuTabRef struct {
	ConversionTable    string
	ConversionTableSet bool
}

func parseCompuTabRef(tok *tokenGenerator) (CompuTabRef, error) {
	ctr := CompuTabRef{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("compuTabRef could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("compuTabRef could not be parsed")
	} else if !ctr.ConversionTableSet {
		ctr.ConversionTable = tok.current()
		ctr.ConversionTableSet = true
	}
	return ctr, err
}
