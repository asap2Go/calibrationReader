package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type calibrationHandleText struct {
	text    string
	textSet bool
}

func parseCalibrationHandleText(tok *tokenGenerator) (calibrationHandleText, error) {
	cht := calibrationHandleText{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("calibrationHandleText could not be parsed")
	} else if !cht.textSet {
		cht.text = tok.current()
		cht.textSet = true
		log.Info().Msg("calibrationHandleText text successfully parsed")
	}
	return cht, err
}
