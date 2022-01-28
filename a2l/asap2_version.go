package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*asap2Version upgrade number will be incremented if additional functionality is implemented to ASAM MCD-2 MC standard
which has no effect on existing applications (compatible modifications).
The version number will be incremented in case if incompatible modifications.
The ASAP2_VERSION keyword is mandatory and expected before the keyword PROJECT.*/
type asap2Version struct {
	versionNo    uint16
	versionNoSet bool
	upgradeNo    uint16
	upgradeNoSet bool
}

func parseASAP2Version(tok *tokenGenerator) (asap2Version, error) {
	av := asap2Version{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("asap2Version could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("asap2Version could not be parsed")
			break forLoop
		} else if !av.versionNoSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("asap2Version versionNo could not be parsed")
				break forLoop
			}
			av.versionNo = uint16(buf)
			av.versionNoSet = true
			log.Info().Msg("asap2Version versionNo successfully parsed")
		} else if !av.upgradeNoSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("asap2Version upgradeNo could not be parsed")
				break forLoop
			}
			av.upgradeNo = uint16(buf)
			av.upgradeNoSet = true
			log.Info().Msg("asap2Version upgradeNo successfully parsed")
			break forLoop
		}
	}
	return av, err
}
