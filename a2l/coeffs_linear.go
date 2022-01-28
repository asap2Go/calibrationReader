package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type coeffsLinear struct {
	a    float64
	aSet bool
	b    float64
	bSet bool
}

func parseCoeffsLinear(tok *tokenGenerator) (coeffsLinear, error) {
	cl := coeffsLinear{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("coeffsLinear could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("coeffsLinear could not be parsed")
			break forLoop
		} else if !cl.aSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffsLinear a could not be parsed")
				break forLoop
			}
			cl.a = buf
			cl.aSet = true
			log.Info().Msg("coeffsLinear a successfully parsed")
		} else if !cl.bSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffsLinear b could not be parsed")
				break forLoop
			}
			cl.b = buf
			cl.bSet = true
			log.Info().Msg("coeffsLinear b successfully parsed")
			break forLoop
		}
	}
	return cl, err
}
