package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type Formula struct {
	fx         []string
	fxSet      bool
	formulaInv formulaInv
}

func parseFormula(tok *tokenGenerator) (Formula, error) {
	f := Formula{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case formulaInvToken:
			f.formulaInv, err = parseFormulaInv(tok)
			if err != nil {
				log.Err(err).Msg("formula formulaInv could not be parsed")
				break forLoop
			}
			log.Info().Msg("formula formulaInv successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("formula could not be parsed")
				break forLoop
			} else if tok.current() == endFormulaToken {
				f.fxSet = true
				log.Info().Msg("formula fx successfully parsed")
				break forLoop
			} else if !f.fxSet {
				f.fx = append(f.fx, tok.current())
			}
		}
	}
	return f, err
}
