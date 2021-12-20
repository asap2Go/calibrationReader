package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type calibrationHandle struct {
	handle                []uint32
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
			} else if !ch.handleSet {
				var buf uint32
				buf, err = parseHexAddressToUint32(tok.current())
				if err != nil {
					log.Err(err).Msg("calibrationHandle handle could not be parsed")
					break forLoop
				}
				ch.handle = append(ch.handle, buf)
				log.Info().Msg("calibrationHandle handle successfully parsed")
			}
		}
	}
	return ch, err
}
