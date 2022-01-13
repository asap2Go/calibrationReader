package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type varMeasurement struct {
	name    string
	nameSet bool
}

func parseVarMeasurement(tok *tokenGenerator) (varMeasurement, error) {
	vm := varMeasurement{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("varMeasurement could not be parsed")
	} else if !vm.nameSet {
		vm.name = tok.current()
		vm.nameSet = true
		log.Info().Msg("varMeasurement name successfully parsed")
	}
	return vm, err
}
