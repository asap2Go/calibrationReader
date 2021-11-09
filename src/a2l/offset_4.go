package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type offset4 struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseOffset4(tok *tokenGenerator) (offset4, error) {
	o4 := offset4{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("offset4 could not be parsed")
			break forLoop
		} else if !o4.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
					log.Err(err).Msg("offset4 position could not be parsed")
				break forLoop
			}
			o4.position = uint16(buf)
			o4.positionSet = true
				log.Info().Msg("offset4 position successfully parsed")
		} else if !o4.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
					log.Err(err).Msg("offset4 datatype could not be parsed")
				break forLoop
			}
			o4.datatype = buf
			o4.datatypeSet = true
				log.Info().Msg("offset4 datatype successfully parsed")
			break forLoop
		}
	}
	return o4, err
}
