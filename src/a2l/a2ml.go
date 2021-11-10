package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type a2ml struct {
	formatSpecification    string
	formatSpecificationSet bool
}

func parseA2ML(tok *tokenGenerator) (a2ml, error) {
	a2ml := a2ml{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("a2ml could not be parsed")
			break forLoop
		} else if tok.current() == endA2mlToken {
			a2ml.formatSpecificationSet = true
			break forLoop
		} else if !a2ml.formatSpecificationSet {
			a2ml.formatSpecification = a2ml.formatSpecification + spaceToken + tok.current()
		}
	}
	return a2ml, err
}
