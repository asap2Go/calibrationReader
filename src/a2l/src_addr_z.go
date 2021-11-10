package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type srcAddrZ struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseSrcAddrZ(tok *tokenGenerator) (srcAddrZ, error) {
	saz := srcAddrZ{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("srcAddrz could not be parsed")
			break forLoop
		} else if !saz.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("srcAddrz position could not be parsed")
				break forLoop
			}
			saz.position = uint16(buf)
			saz.positionSet = true
			log.Info().Msg("srcAddrz position successfully parsed")
		} else if !saz.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("srcAddrz datatype could not be parsed")
				break forLoop
			}
			saz.datatype = buf
			saz.datatypeSet = true
			log.Info().Msg("srcAddrz datatype successfully parsed")
			break forLoop
		}
	}
	return saz, err
}
