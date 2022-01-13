package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type Monotony struct {
	monotony    monotonyTypeEnum
	monotonySet bool
}

func parseMonotony(tok *tokenGenerator) (Monotony, error) {
	m := Monotony{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("monotony could not be parsed")
	} else if !m.monotonySet {
		m.monotony, err = parseMonotonyTypeEnum(tok)
		if err != nil {
			log.Err(err).Msg("monotony monotony could not be parsed")
		}
		m.monotonySet = true
		log.Info().Msg("monotony monotony successfully parsed")
	}
	return m, err
}
