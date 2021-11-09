package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
)

type bitMask struct {
	mask    string //uint32
	maskSet bool
}

func parseBitMask(tok *tokenGenerator) (bitMask, error) {
	bm := bitMask{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("bitMask could not be parsed")
	} else if !bm.maskSet {
		bm.mask = tok.current()
		bm.maskSet = true
			log.Info().Msg("bitMask mask successfully parsed")
	}
	return bm, err
}
