package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type DefaultValueNumeric struct {
	DisplayValue    float64
	DisplayValueSet bool
}

func parseDefaultValueNumeric(tok *tokenGenerator) (DefaultValueNumeric, error) {
	dvn := DefaultValueNumeric{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("defaultValueNumeric could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("defaultValueNumeric could not be parsed")
	} else if !dvn.DisplayValueSet {
		var buf float64
		buf, err = strconv.ParseFloat(tok.current(), 64)
		if err != nil {
			log.Err(err).Msg("attribute displayValue could not be parsed")
		}
		dvn.DisplayValue = buf
		dvn.DisplayValueSet = true
		log.Info().Msg("defaultValueNumeric displayValue successfully parsed")
	}
	return dvn, err
}
