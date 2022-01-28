package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type defaultValueNumeric struct {
	displayValue    float64
	displayValueSet bool
}

func parseDefaultValueNumeric(tok *tokenGenerator) (defaultValueNumeric, error) {
	dvn := defaultValueNumeric{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("defaultValueNumeric could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("defaultValueNumeric could not be parsed")
	} else if !dvn.displayValueSet {
		var buf float64
		buf, err = strconv.ParseFloat(tok.current(), 64)
		if err != nil {
			log.Err(err).Msg("attribute displayValue could not be parsed")
		}
		dvn.displayValue = buf
		dvn.displayValueSet = true
		log.Info().Msg("defaultValueNumeric displayValue successfully parsed")
	}
	return dvn, err
}
