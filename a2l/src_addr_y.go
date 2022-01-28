package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type srcAddrY struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseSrcAddrY(tok *tokenGenerator) (srcAddrY, error) {
	say := srcAddrY{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("srcAddrY could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("srcAddrY could not be parsed")
			break forLoop
		} else if !say.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("srcAddrY position could not be parsed")
				break forLoop
			}
			say.position = uint16(buf)
			say.positionSet = true
			log.Info().Msg("srcAddrY position successfully parsed")
		} else if !say.datatypeSet {
			say.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("srcAddrY datatype could not be parsed")
				break forLoop
			}
			say.datatypeSet = true
			log.Info().Msg("srcAddrY datatype successfully parsed")
			break forLoop
		}
	}
	return say, err
}
