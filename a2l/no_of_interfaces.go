package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type noOfInterfaces struct {
	num    uint16
	numSet bool
}

func parseNoOfInterfaces(tok *tokenGenerator) (noOfInterfaces, error) {
	noi := noOfInterfaces{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("noOfInterfaces could not be parsed: unexpected end of file")
	} else if !noi.numSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("attribute num could not be parsed")
		}
		noi.num = uint16(buf)
		noi.numSet = true
		log.Info().Msg("noOfInterfaces num successfully parsed")
	}
	return noi, err
}
