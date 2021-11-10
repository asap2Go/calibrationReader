package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type offsetX struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseOffsetX(tok *tokenGenerator) (offsetX, error) {
	ox := offsetX{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("offsetx could not be parsed")
			break forLoop
		} else if !ox.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("offsetx position could not be parsed")
				break forLoop
			}
			ox.position = uint16(buf)
			ox.positionSet = true
			log.Info().Msg("offsetx position successfully parsed")
		} else if !ox.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("offsetx datatype could not be parsed")
				break forLoop
			}
			ox.datatype = buf
			ox.datatypeSet = true
			log.Info().Msg("offsetx datatype successfully parsed")
			break forLoop
		}
	}
	return ox, err
}
