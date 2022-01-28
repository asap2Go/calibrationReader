package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type FunctionList struct {
	name    []string
	nameSet bool
}

func parseFunctionList(tok *tokenGenerator) (FunctionList, error) {
	fl := FunctionList{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("functionList could not be parsed")
			break forLoop
		} else if tok.current() == endFunctionListToken {
			fl.nameSet = true
			log.Info().Msg("functionList name successfully parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("functionList could not be parsed")
			break forLoop
		} else if !fl.nameSet {
			fl.name = append(fl.name, tok.current())
		}
	}
	return fl, err
}
