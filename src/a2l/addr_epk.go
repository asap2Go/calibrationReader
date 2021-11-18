package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type addrEpk struct {
	//address contains the Address of the EPROM identifier
	address    uint32
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
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 32)
		if err != nil {
			log.Err(err).Msg("addrEpk address could not be parsed")
		}
		ae.address = uint32(buf)
		ae.addressSet = true
		log.Info().Msg("addrEpk address successfully parsed")
	}
	return ae, err
}
