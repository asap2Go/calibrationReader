package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type virtual struct {
	measuringChannel    []string
	measuringChannelSet bool
}

func parseVirtual(tok *tokenGenerator) (virtual, error) {
	v := virtual{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("virtual could not be parsed")
			break forLoop
		} else if tok.current() == endVirtualToken {
			v.measuringChannelSet = true
			log.Info().Msg("virtual measuringChannel successfully parsed")
			break forLoop
		} else if !v.measuringChannelSet {
			v.measuringChannel = append(v.measuringChannel, tok.current())
		}
	}
	return v, err
}
