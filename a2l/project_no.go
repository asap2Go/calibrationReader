package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type projectNo struct {
	projectNumber    string
	projectNumberSet bool
}

func parseProjectNo(tok *tokenGenerator) (projectNo, error) {
	pn := projectNo{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("projectNo could not be parsed")
	} else if !pn.projectNumberSet {
		pn.projectNumber = tok.current()
		pn.projectNumberSet = true
		log.Info().Msg("projectNo projectNumber successfully parsed")
	}
	return pn, err
}
