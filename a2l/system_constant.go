package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type SystemConstant struct {
	name     string
	nameSet  bool
	value    string
	valueSet bool
}

func parseSystemConstant(tok *tokenGenerator) (SystemConstant, error) {
	sc := SystemConstant{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("systemConstant: " + sc.name + " could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("systemConstant could not be parsed")
			break forLoop
		} else if !sc.nameSet {
			sc.name = tok.current()
			sc.nameSet = true
			log.Info().Msg("systemConstant name successfully parsed")
		} else if !sc.valueSet {
			sc.value = tok.current()
			sc.valueSet = true
			log.Info().Msg("systemConstant value successfully parsed")
			break forLoop
		}
	}
	return sc, err
}
