package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type addrEpk struct {
	//address contains the Address of the EPROM identifier
	address    string
	addressSet bool
}

func parseAddrEpk(tok *tokenGenerator) (addrEpk, error) {
	ae := addrEpk{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("addrEpk could not be parsed")
	} else if !ae.addressSet {
		ae.address = tok.current()
		ae.addressSet = true
		log.Info().Msg("addrEpk address successfully parsed")
	}
	return ae, err
}
