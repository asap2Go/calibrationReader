package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type coeffs struct {
	a    float64
	aSet bool
	b    float64
	bSet bool
	c    float64
	cSet bool
	d    float64
	dSet bool
	e    float64
	eSet bool
	f    float64
	fSet bool
}

func parseCoeffs(tok *tokenGenerator) (coeffs, error) {
	co := coeffs{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("coeffs could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("coeffs could not be parsed")
			break forLoop
		} else if !co.aSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffs a could not be parsed")
				break forLoop
			}
			co.a = buf
			co.aSet = true
			log.Info().Msg("coeffs a successfully parsed")
		} else if !co.bSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffs b could not be parsed")
				break forLoop
			}
			co.b = buf
			co.bSet = true
			log.Info().Msg("coeffs b successfully parsed")
		} else if !co.cSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffs c could not be parsed")
				break forLoop
			}
			co.c = buf
			co.cSet = true
			log.Info().Msg("coeffs c successfully parsed")
		} else if !co.dSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffs d could not be parsed")
				break forLoop
			}
			co.d = buf
			co.dSet = true
			log.Info().Msg("coeffs d successfully parsed")
		} else if !co.eSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffs e could not be parsed")
				break forLoop
			}
			co.e = buf
			co.eSet = true
			log.Info().Msg("coeffs e successfully parsed")
		} else if !co.fSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffs f could not be parsed")
				break forLoop
			}
			co.f = buf
			co.fSet = true
			log.Info().Msg("coeffs f successfully parsed")
			break forLoop
		}
	}
	return co, err
}
