package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type Coeffs struct {
	A    float64
	ASet bool
	B    float64
	BSet bool
	C    float64
	CSet bool
	D    float64
	DSet bool
	E    float64
	ESet bool
	F    float64
	FSet bool
}

func parseCoeffs(tok *tokenGenerator) (Coeffs, error) {
	co := Coeffs{}
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
		} else if !co.ASet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffs a could not be parsed")
				break forLoop
			}
			co.A = buf
			co.ASet = true
			log.Info().Msg("coeffs a successfully parsed")
		} else if !co.BSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffs b could not be parsed")
				break forLoop
			}
			co.B = buf
			co.BSet = true
			log.Info().Msg("coeffs b successfully parsed")
		} else if !co.CSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffs c could not be parsed")
				break forLoop
			}
			co.C = buf
			co.CSet = true
			log.Info().Msg("coeffs c successfully parsed")
		} else if !co.DSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffs d could not be parsed")
				break forLoop
			}
			co.D = buf
			co.DSet = true
			log.Info().Msg("coeffs d successfully parsed")
		} else if !co.ESet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffs e could not be parsed")
				break forLoop
			}
			co.E = buf
			co.ESet = true
			log.Info().Msg("coeffs e successfully parsed")
		} else if !co.FSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("coeffs f could not be parsed")
				break forLoop
			}
			co.F = buf
			co.FSet = true
			log.Info().Msg("coeffs f successfully parsed")
			break forLoop
		}
	}
	return co, err
}
