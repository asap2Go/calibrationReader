package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type varForbiddenComb struct {
	criterionName     []string
	criterionNameSet  bool
	criterionValue    []string
	criterionValueSet bool
}

func parseVarForbiddenComb(tok *tokenGenerator) (varForbiddenComb, error) {
	vfc := varForbiddenComb{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("varForbiddenComb could not be parsed")
			break forLoop
		} else if tok.current() == endVarForbiddenCombToken {
			vfc.criterionNameSet = true
			vfc.criterionValueSet = true
			log.Info().Msg("varForbiddenComb criterionName successfully parsed")
			log.Info().Msg("varForbiddenComb criterionValue successfully parsed")
			break forLoop
		} else if !vfc.criterionNameSet || !vfc.criterionValueSet {
			vfc.criterionName = append(vfc.criterionName, tok.current())
			if tok.next() != emptyToken {
				vfc.criterionValue = append(vfc.criterionValue, tok.current())
			} else {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("varForbiddenComb could not be parsed")
				break forLoop
			}
		}
	}
	return vfc, err
}
