package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type SystemConstant struct {
	Name     string
	NameSet  bool
	Value    string
	ValueSet bool
}

func parseSystemConstant(tok *tokenGenerator) (SystemConstant, error) {
	sc := SystemConstant{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("systemConstant: " + sc.Name + " could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("systemConstant could not be parsed")
			break forLoop
		} else if !sc.NameSet {
			sc.Name = tok.current()
			sc.NameSet = true
			log.Info().Msg("systemConstant name successfully parsed")
		} else if !sc.ValueSet {
			sc.Value = tok.current()
			sc.ValueSet = true
			log.Info().Msg("systemConstant value successfully parsed")
			break forLoop
		}
	}
	return sc, err
}
