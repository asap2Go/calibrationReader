package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type curveAxisRef struct {
	curveAxis    string
	curveAxisSet bool
}

func parseCurveAxisRef(tok *tokenGenerator) (curveAxisRef, error) {
	car := curveAxisRef{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("curveAxisRef could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("curveAxisRef could not be parsed")
	} else if !car.curveAxisSet {
		car.curveAxis = tok.current()
		car.curveAxisSet = true
		log.Info().Msg("curveAxisRef curveAxis successfully parsed")
	}
	return car, err
}
