package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type ExtendedLimits struct {
	lowerLimit    float64
	lowerLimitSet bool
	upperLimit    float64
	upperLimitSet bool
}

func parseExtendedLimits(tok *tokenGenerator) (ExtendedLimits, error) {
	el := ExtendedLimits{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("extendedLimits could not be parsed")
			break forLoop
		} else if !el.lowerLimitSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("extendedLimits lowerLimit could not be parsed")
				break forLoop
			}
			el.lowerLimit = buf
			el.lowerLimitSet = true
			log.Info().Msg("extendedLimits lowerLimit successfully parsed")
		} else if !el.upperLimitSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("extendedLimits upperLimit could not be parsed")
				break forLoop
			}
			el.upperLimit = buf
			el.upperLimitSet = true
			log.Info().Msg("extendedLimits upperLimit successfully parsed")
			break forLoop
		}
	}
	return el, err
}
