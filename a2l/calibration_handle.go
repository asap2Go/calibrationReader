package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type calibrationHandle struct {
	handle                []string
	handleSet             bool
	calibrationHandleText calibrationHandleText
}

func parseCalibrationHandle(tok *tokenGenerator) (calibrationHandle, error) {
	ch := calibrationHandle{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case calibrationHandleTextToken:
			ch.calibrationHandleText, err = parseCalibrationHandleText(tok)
			if err != nil {
				log.Err(err).Msg("calibrationHandle calibrationHandleText could not be parsed")
				break forLoop
			}
			log.Info().Msg("calibrationHandle calibrationHandleText successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("calibrationHandle could not be parsed")
				break forLoop
			} else if tok.current() == endCalibrationHandleToken {
				ch.handleSet = true
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("calibrationHandle could not be parsed")
				break forLoop
			} else if !ch.handleSet {
				ch.handle = append(ch.handle, tok.current())
				log.Info().Msg("calibrationHandle handle successfully parsed")
			}
		}
	}
	return ch, err
}
