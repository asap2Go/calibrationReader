package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
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
			log.Err(err).Msg("srcAddry could not be parsed")
			break forLoop
		} else if !say.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("srcAddry position could not be parsed")
				break forLoop
			}
			say.position = uint16(buf)
			say.positionSet = true
			log.Info().Msg("srcAddry position successfully parsed")
		} else if !say.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("srcAddry datatype could not be parsed")
				break forLoop
			}
			say.datatype = buf
			say.datatypeSet = true
			log.Info().Msg("srcAddry datatype successfully parsed")
			break forLoop
		}
	}
	return say, err
}
