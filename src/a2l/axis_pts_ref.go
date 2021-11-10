package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type axisPtsRef struct {
	axisPoints    string
	axisPointsSet bool
}

func parseAxisPtsRef(tok *tokenGenerator) (axisPtsRef, error) {
	apr := axisPtsRef{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("axisPtsRef could not be parsed")
	} else if !apr.axisPointsSet {
		apr.axisPoints = tok.current()
		apr.axisPointsSet = true
		log.Info().Msg("axisPtsRef axisPoints successfully parsed")
	}
	return apr, err
}
