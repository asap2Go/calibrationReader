package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type srcAddrX struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseSrcAddrX(tok *tokenGenerator) (srcAddrX, error) {
	sax := srcAddrX{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("srcAddrx could not be parsed")
			break forLoop
		} else if !sax.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("srcAddrx position could not be parsed")
				break forLoop
			}
			sax.position = uint16(buf)
			sax.positionSet = true
			log.Info().Msg("srcAddrx position successfully parsed")
		} else if !sax.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("srcAddrx datatype could not be parsed")
				break forLoop
			}
			sax.datatype = buf
			sax.datatypeSet = true
			log.Info().Msg("srcAddrx datatype successfully parsed")
			break forLoop
		}
	}
	return sax, err
}
