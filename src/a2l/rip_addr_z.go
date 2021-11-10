package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type ripAddrZ struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
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
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("ripAddrz datatype could not be parsed")
				break forLoop
			}
			raz.datatype = buf
			raz.datatypeSet = true
			log.Info().Msg("ripAddrz datatype successfully parsed")
			break forLoop
		}
	}
	return raz, err
}
