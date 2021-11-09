package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type formulaInv struct {
	gx    string
	gxSet bool
}

func parseFormulaInv(tok *tokenGenerator) (formulaInv, error) {
	fi := formulaInv{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("formulaInv could not be parsed")
	} else if !fi.gxSet {
		fi.gx = tok.current()
		fi.gxSet = true
			log.Info().Msg("formulaInv gx successfully parsed")
	}
	return fi, err
}
