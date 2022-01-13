package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type ecuAddressExtension struct {
	extension    int16
	extensionSet bool
}

func parseECUAddressExtension(tok *tokenGenerator) (ecuAddressExtension, error) {
	eae := ecuAddressExtension{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("ecuAddressExtension could not be parsed")
	} else if !eae.extensionSet {
		var buf int64
		buf, err = strconv.ParseInt(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("ecuAddressExtension extension could not be parsed")
		}
		eae.extension = int16(buf)
		eae.extensionSet = true
		log.Info().Msg("ecuAddressExtension extension successfully parsed")
	}
	return eae, err
}
