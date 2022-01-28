package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type Number struct {
	number    uint16
	numberSet bool
}

func parseNumber(tok *tokenGenerator) (Number, error) {
	n := Number{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("number could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("number could not be parsed")
	} else if !n.numberSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("number number could not be parsed")
		}
		n.number = uint16(buf)
		n.numberSet = true
		log.Info().Msg("number number successfully parsed")
	}
	return n, err
}
