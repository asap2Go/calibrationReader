package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type CoeffsLinear struct {
	A    float64
	ASet bool
	B    float64
	BSet bool
}

func parseCoeffsLinear(tok *tokenGenerator) (CoeffsLinear, error) {
	cl := CoeffsLinear{}
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
		} else if !cl.ASet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffsLinear a could not be parsed")
				break forLoop
			}
			cl.A = buf
			cl.ASet = true
			log.Info().Msg("coeffsLinear a successfully parsed")
		} else if !cl.BSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffsLinear b could not be parsed")
				break forLoop
			}
			cl.B = buf
			cl.BSet = true
			log.Info().Msg("coeffsLinear b successfully parsed")
			break forLoop
		}
	}
	return cl, err
}
