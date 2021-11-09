package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type errorMask struct {
	mask    uint32
	maskSet bool
}

func parseErrorMask(tok *tokenGenerator) (errorMask, error) {
	em := errorMask{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("errorMask could not be parsed")
	} else if !em.maskSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 32)
		if err != nil {
				log.Err(err).Msg("errorMask mask could not be parsed")
		}
		em.mask = uint32(buf)
		em.maskSet = true
			log.Info().Msg("errorMask mask successfully parsed")
	}
	return em, err
}
