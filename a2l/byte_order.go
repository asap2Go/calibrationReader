package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type ByteOrder struct {
	ByteOrder    ByteOrderEnum
	ByteOrderSet bool
}

func parseByteOrder(tok *tokenGenerator) (ByteOrder, error) {
	bo := ByteOrder{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("byteOrder could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("byteOrder could not be parsed")
	} else if !bo.ByteOrderSet {
		bo.ByteOrder, err = parseByteOrderEnum(tok)
		if err != nil {
			log.Err(err).Msg("byteOrder byteOrder could not be parsed")
		}
		bo.ByteOrderSet = true
		log.Info().Msg("byteOrder byteOrder successfully parsed")
	}
	return bo, err
}
