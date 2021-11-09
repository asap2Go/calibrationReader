package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type offset5 struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseOffset5(tok *tokenGenerator) (offset5, error) {
	o5 := offset5{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("offset5 could not be parsed")
			break forLoop
		} else if !o5.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
					log.Err(err).Msg("offset5 position could not be parsed")
				break forLoop
			}
			o5.position = uint16(buf)
			o5.positionSet = true
				log.Info().Msg("offset5 position successfully parsed")
		} else if !o5.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
					log.Err(err).Msg("offset5 datatype could not be parsed")
				break forLoop
			}
			o5.datatype = buf
			o5.datatypeSet = true
				log.Info().Msg("offset5 datatype successfully parsed")
			break forLoop
		}
	}
	return o5, err
}
