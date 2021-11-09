package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type LeftShift struct {
	bitcount    uint32
	bitcountSet bool
}

func parseLeftShift(tok *tokenGenerator) (LeftShift, error) {
	ls := LeftShift{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("leftShift could not be parsed")
	} else if !ls.bitcountSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 32)
		if err != nil {
				log.Err(err).Msg("leftShift bitcount could not be parsed")
		}
		ls.bitcount = uint32(buf)
		ls.bitcountSet = true
			log.Info().Msg("leftShift bitcount successfully parsed")
	}
	return ls, err
}
