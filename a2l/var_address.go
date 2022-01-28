package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type varAddress struct {
	address    []string
	addressSet bool
}

func parseVarAddress(tok *tokenGenerator) (varAddress, error) {
	va := varAddress{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("varAddress could not be parsed")
			break forLoop
		} else if tok.current() == endVarAddressToken {
			va.addressSet = true
			log.Info().Msg("varAddress address successfully parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("varAddress could not be parsed")
			break forLoop
		} else if !va.addressSet {
			va.address = append(va.address, tok.current())
		}
	}
	return va, err
}
