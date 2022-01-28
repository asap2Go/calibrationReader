package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type unitConversion struct {
	gradient    float64
	gradientSet bool
	offset      float64
	offsetSet   bool
}

func parseUnitConversion(tok *tokenGenerator) (unitConversion, error) {
	uc := unitConversion{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("unitConversion could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("unitConversion could not be parsed")
			break forLoop
		} else if !uc.gradientSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("unitConversion gradient could not be parsed")
				break forLoop
			}
			uc.gradient = buf
			uc.gradientSet = true
			log.Info().Msg("unitConversion gradient successfully parsed")
		} else if !uc.offsetSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("unitConversion offset could not be parsed")
				break forLoop
			}
			uc.offset = buf
			uc.offsetSet = true
			log.Info().Msg("unitConversion offset successfully parsed")
			break forLoop
		}
	}
	return uc, err
}
