package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type customerNo struct {
	number    string
	numberSet bool
}

func parseCustomerNo(tok *tokenGenerator) (customerNo, error) {
	cn := customerNo{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("customerNo could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("customerNo could not be parsed")
	} else if !cn.numberSet {
		cn.number = tok.current()
		cn.numberSet = true
		log.Info().Msg("customerNo number successfully parsed")
	}
	return cn, err
}
