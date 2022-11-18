package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type srcAddrX struct {
	position    uint16
	positionSet bool
	datatype    DataTypeEnum
	datatypeSet bool
}

// srcAddrX is the description of the address of the input quantity in an adjustable object
func parseSrcAddrX(tok *tokenGenerator) (srcAddrX, error) {
	sax := srcAddrX{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("srcAddrX could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("srcAddrX could not be parsed")
			break forLoop
		} else if !sax.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("srcAddrX position could not be parsed")
				break forLoop
			}
			sax.position = uint16(buf)
			sax.positionSet = true
			log.Info().Msg("srcAddrX position successfully parsed")
		} else if !sax.datatypeSet {
			sax.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("srcAddrX datatype could not be parsed")
				break forLoop
			}
			sax.datatypeSet = true
			log.Info().Msg("srcAddrX datatype successfully parsed")
			break forLoop
		}
	}
	return sax, err
}
