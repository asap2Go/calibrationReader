package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type calibrationMethod struct {
	method            string
	methodSet         bool
	version           uint32
	versionSet        bool
	calibrationHandle []calibrationHandle
}

func parseCalibrationMethod(tok *tokenGenerator) (calibrationMethod, error) {
	cm := calibrationMethod{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginCalibrationHandleToken:
			var buf calibrationHandle
			buf, err = parseCalibrationHandle(tok)
			if err != nil {
				log.Err(err).Msg("calibrationMethod calibrationHandle could not be parsed")
				break forLoop
			}
			cm.calibrationHandle = append(cm.calibrationHandle, buf)
			log.Info().Msg("calibrationMethod calibrationHandle successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("calibrationMethod could not be parsed")
				break forLoop
			} else if tok.current() == endCalibrationMethodToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("calibrationMethod could not be parsed")
				break forLoop
			} else if !cm.methodSet {
				cm.method = tok.current()
				cm.methodSet = true
				log.Info().Msg("calibrationMethod method successfully parsed")
			} else if !cm.versionSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 32)
				if err != nil {
					log.Err(err).Msg("calibrationMethod version could not be parsed")
					break forLoop
				}
				cm.version = uint32(buf)
				cm.versionSet = true
				log.Info().Msg("calibrationMethod version successfully parsed")
			}
		}
	}
	return cm, err
}
