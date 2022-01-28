package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type customer struct {
	customer    string
	customerSet bool
}

func parseCustomer(tok *tokenGenerator) (customer, error) {
	c := customer{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("customer could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("customer could not be parsed")
	} else if !c.customerSet {
		c.customer = tok.current()
		c.customerSet = true
		log.Info().Msg("customer customer successfully parsed")
	}
	return c, err
}
