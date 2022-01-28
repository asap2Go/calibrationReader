package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type fixAxisParList struct {
	axisPtsValue    []float64
	axisPtsValueSet bool
}

func parseFixAxisParList(tok *tokenGenerator) (fixAxisParList, error) {
	fapl := fixAxisParList{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("fixAxisParList could not be parsed")
			break forLoop
		} else if tok.current() == endFixAxisParListToken {
			fapl.axisPtsValueSet = true
			log.Info().Msg("fixAxisParList axisPtsValue successfully parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("fixAxisParList could not be parsed")
			break forLoop
		} else if !fapl.axisPtsValueSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("attribute axisPtsValue could not be parsed")
				break forLoop
			}
			fapl.axisPtsValue = append(fapl.axisPtsValue, buf)
		}
	}
	return fapl, err
}
