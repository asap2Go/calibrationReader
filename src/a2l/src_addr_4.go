package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type srcAddr4 struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseSrcAddr4(tok *tokenGenerator) (srcAddr4, error) {
	sa4 := srcAddr4{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("srcAddr4 could not be parsed")
			break forLoop
		} else if !sa4.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("srcAddr4 position could not be parsed")
				break forLoop
			}
			sa4.position = uint16(buf)
			sa4.positionSet = true
			log.Info().Msg("srcAddr4 position successfully parsed")
		} else if !sa4.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("srcAddr4 datatype could not be parsed")
				break forLoop
			}
			sa4.datatype = buf
			sa4.datatypeSet = true
			log.Info().Msg("srcAddr4 datatype successfully parsed")
			break forLoop
		}
	}
	return sa4, err
}
