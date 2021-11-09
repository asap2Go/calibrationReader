package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type a2mlVersion struct {
	versionNo    uint16
	versionNoSet bool
	upgradeNo    uint16
	upgradeNoSet bool
}

func parseA2MLVersion(tok *tokenGenerator) (a2mlVersion, error) {
	av := a2mlVersion{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("a2mlVersion could not be parsed")
			break forLoop
		} else if !av.versionNoSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
					log.Err(err).Msg("a2mlVersion versionNo could not be parsed")
				break forLoop
			}
			av.versionNo = uint16(buf)
			av.versionNoSet = true
				log.Info().Msg("a2mlVersion versionNo successfully parsed")
		} else if !av.upgradeNoSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
					log.Err(err).Msg("a2mlVersion upgradeNo could not be parsed")
				break forLoop
			}
			av.upgradeNo = uint16(buf)
			av.upgradeNoSet = true
				log.Info().Msg("a2mlVersion upgradeNo successfully parsed")
			break forLoop
		}
	}
	return av, err
}
