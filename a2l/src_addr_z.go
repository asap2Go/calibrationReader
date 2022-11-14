package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type srcAddrZ struct {
	position    uint16
	positionSet bool
	datatype    DataTypeEnum
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
			log.Err(err).Msg("srcAddrZ could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("srcAddrZ could not be parsed")
			break forLoop
		} else if !saz.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("srcAddrZ position could not be parsed")
				break forLoop
			}
			saz.position = uint16(buf)
			saz.positionSet = true
			log.Info().Msg("srcAddrZ position successfully parsed")
		} else if !saz.datatypeSet {
			saz.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("srcAddrZ datatype could not be parsed")
				break forLoop
			}
			saz.datatypeSet = true
			log.Info().Msg("srcAddrZ datatype successfully parsed")
			break forLoop
		}
	}
	return saz, err
}
