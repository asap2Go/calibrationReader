package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type MaxGrad struct {
	maxGradient    float64
	maxGradientSet bool
}

func parseMaxGrad(tok *tokenGenerator) (MaxGrad, error) {
	mg := MaxGrad{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("maxGrad could not be parsed")
	} else if !mg.maxGradientSet {
		var buf float64
		buf, err = strconv.ParseFloat(tok.current(), 64)
		if err != nil {
			log.Err(err).Msg("maxGrad maxGradient could not be parsed")
		}
		mg.maxGradient = buf
		mg.maxGradientSet = true
		log.Info().Msg("maxGrad maxGradient successfully parsed")
	}
	return mg, err
}
