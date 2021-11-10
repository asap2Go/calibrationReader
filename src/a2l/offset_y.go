package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type offsetY struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseOffsetY(tok *tokenGenerator) (offsetY, error) {
	oy := offsetY{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("offsety could not be parsed")
			break forLoop
		} else if !oy.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("offsety position could not be parsed")
				break forLoop
			}
			oy.position = uint16(buf)
			oy.positionSet = true
			log.Info().Msg("offsety position successfully parsed")
		} else if !oy.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("offsety datatype could not be parsed")
				break forLoop
			}
			oy.datatype = buf
			oy.datatypeSet = true
			log.Info().Msg("offsety datatype successfully parsed")
			break forLoop
		}
	}
	return oy, err
}
