package a2l

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"strings"

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
		ae.address, err = parseHexAddressToUint32(tok.current())
		if err != nil {
			log.Err(err).Msg("addrEpk address could not be parsed")
		}
		ae.addressSet = true
		log.Info().Msg("addrEpk address successfully parsed")
	}
	return ae, err
}

func parseHexAddressToUint32(str string) (uint32, error) {
	var byteSlice []byte
	var err error
	var address uint32
	str = strings.ReplaceAll(str, "0x", "")
	byteSlice, err = hex.DecodeString(str)
	if err != nil {
		log.Err(err)
		return 0, err
	}
	//convert bytes to uint16, add index
	address = binary.BigEndian.Uint32(byteSlice)
	return address, err
}
