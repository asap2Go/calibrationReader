package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type phoneNo struct {
	telnum    string
	telnumSet bool
}

func parsePhoneNo(tok *tokenGenerator) (phoneNo, error) {
	pn := phoneNo{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("phoneNo could not be parsed")
	} else if !pn.telnumSet {
		pn.telnum = tok.current()
		pn.telnumSet = true
			log.Info().Msg("phoneNo telnum successfully parsed")
	}
	return pn, err
}
