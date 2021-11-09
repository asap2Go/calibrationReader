package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type cpuType struct {
	cpu    string
	cpuSet bool
}

func parseCpuType(tok *tokenGenerator) (cpuType, error) {
	ct := cpuType{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("cpuType could not be parsed")
	} else if !ct.cpuSet {
		ct.cpu = tok.current()
		ct.cpuSet = true
			log.Info().Msg("cpuType cpu successfully parsed")
	}
	return ct, err
}
