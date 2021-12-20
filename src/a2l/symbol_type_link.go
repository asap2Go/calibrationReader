package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type symbolTypeLink struct {
	symbolName string
}

func parseSymbolTypeLink(tok *tokenGenerator) (symbolTypeLink, error) {
	var err error
	stl := symbolTypeLink{}
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("symbolLink could not be parsed")
	} else {
		stl.symbolName = tok.current()
		log.Info().Msg("symbolTypeLink symbolName successfully parsed")
	}
	return stl, err
}
