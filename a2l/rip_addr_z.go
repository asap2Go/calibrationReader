package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type ripAddrZ struct {
	position    uint16
	positionSet bool
	datatype    DataTypeEnum
	datatypeSet bool
}

func parseRipAddrZ(tok *tokenGenerator) (ripAddrZ, error) {
	raz := ripAddrZ{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("ripAddrz could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("ripAddrZ could not be parsed")
			break forLoop
		} else if !raz.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("ripAddrz position could not be parsed")
				break forLoop
			}
			raz.position = uint16(buf)
			raz.positionSet = true
			log.Info().Msg("ripAddrz position successfully parsed")
		} else if !raz.datatypeSet {
			raz.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("ripAddrz datatype could not be parsed")
				break forLoop
			}
			raz.datatypeSet = true
			log.Info().Msg("ripAddrz datatype successfully parsed")
			break forLoop
		}
	}
	return raz, err
}
