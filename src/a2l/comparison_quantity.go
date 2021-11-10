package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type comparisonQuantity struct {
	name    string
	nameSet bool
}

func parseComparisonQuantity(tok *tokenGenerator) (comparisonQuantity, error) {
	cq := comparisonQuantity{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("comparisonQuantity name could not be parsed")
	} else if !cq.nameSet {
		cq.name = tok.current()
		cq.nameSet = true
		log.Info().Msg("comparisonQuantity name successfully parsed")
	}
	return cq, err
}
