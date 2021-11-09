package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type RightShift struct {
	bitcount    uint32
	bitcountSet bool
}

func parseRightShift(tok *tokenGenerator) (RightShift, error) {
	rs := RightShift{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("rightShift could not be parsed")
	} else if !rs.bitcountSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 32)
		if err != nil {
				log.Err(err).Msg("rightShift bitcount could not be parsed")
		}
		rs.bitcount = uint32(buf)
		rs.bitcountSet = true
			log.Info().Msg("rightShift bitcount successfully parsed")
	}
	return rs, err
}
