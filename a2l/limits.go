package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type limits struct {
	lowerLimit    float64
	lowerLimitSet bool
	upperLimit    float64
	upperLimitSet bool
}

func parseLimits(tok *tokenGenerator) (limits, error) {
	l := limits{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("limits could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("limits could not be parsed")
			break forLoop
		} else if !l.lowerLimitSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("limits lowerLimit could not be parsed")
				break forLoop
			}
			l.lowerLimit = buf
			l.lowerLimitSet = true
			log.Info().Msg("limits lowerLimit successfully parsed")
		} else if !l.upperLimitSet {
			var buf float64
			buf, err = strconv.ParseFloat(tok.current(), 64)
			if err != nil {
				log.Err(err).Msg("limits upperLimit could not be parsed")
				break forLoop
			}
			l.upperLimit = buf
			l.upperLimitSet = true
			log.Info().Msg("limits upperLimit successfully parsed")
			break forLoop
		}
	}
	return l, err
}
