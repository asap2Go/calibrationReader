package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type StepSize struct {
	stepSize    float64
	stepSizeSet bool
}

func parseStepSize(tok *tokenGenerator) (StepSize, error) {
	ss := StepSize{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("stepSize could not be parsed")
	} else if !ss.stepSizeSet {
		var buf float64
		buf, err = strconv.ParseFloat(tok.current(), 64)
		if err != nil {
			log.Err(err).Msg("stepSize stepSize could not be parsed")
		}
		ss.stepSize = buf
		ss.stepSizeSet = true
		log.Info().Msg("stepSize stepSize successfully parsed")
	}
	return ss, err
}
