package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type varAddress struct {
	address    []uint32
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
		} else if !va.addressSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 32)
			if err != nil {
					log.Err(err).Msg("varAddress address could not be parsed")
				break forLoop
			}
			va.address = append(va.address, uint32(buf))
		}
	}
	return va, err
}
