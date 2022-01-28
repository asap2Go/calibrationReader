package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type byteOrder struct {
	byteOrder    byteOrderEnum
	byteOrderSet bool
}

func parseByteOrder(tok *tokenGenerator) (byteOrder, error) {
	bo := byteOrder{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("byteOrder could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("byteOrder could not be parsed")
	} else if !bo.byteOrderSet {
		bo.byteOrder, err = parseByteOrderEnum(tok)
		if err != nil {
			log.Err(err).Msg("byteOrder byteOrder could not be parsed")
		}
		bo.byteOrderSet = true
		log.Info().Msg("byteOrder byteOrder successfully parsed")
	}
	return bo, err
}
