package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type MapList struct {
	name    []string
	nameSet bool
}

func parseMapList(tok *tokenGenerator) (MapList, error) {
	l := MapList{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("mapList: could not be parsed")
			break forLoop
		} else if tok.current() == endMapListToken {
			l.nameSet = true
			log.Info().Msg("mapList name successfully parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("mapList could not be parsed")
			break forLoop
		} else if !l.nameSet {
			l.name = append(l.name, tok.current())
		}
	}
	return l, err
}
