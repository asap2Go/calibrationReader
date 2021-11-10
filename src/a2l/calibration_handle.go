package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type calibrationHandle struct {
	handle                int32
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
			} else {
				log.Info().Msg("calibrationHandle calibrationHandleText successfully parsed")
			}
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("calibrationHandle could not be parsed")
				break forLoop
			} else if tok.current() == endCalibrationHandleToken {
				break forLoop
			} else if !ch.handleSet {
				var buf int64
				buf, err = strconv.ParseInt(tok.current(), 10, 32)
				if err != nil {
					log.Err(err).Msg("calibrationHandle handle could not be parsed")
					break forLoop
				} else {
					ch.handle = int32(buf)
					ch.handleSet = true
					log.Info().Msg("calibrationHandle handle successfully parsed")
				}
			}
		}
	}
	return ch, err
}
