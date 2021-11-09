package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type offsetZ struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseOffsetZ(tok *tokenGenerator) (offsetZ, error) {
	oz := offsetZ{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("offsetz could not be parsed")
			break forLoop
		} else if !oz.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
					log.Err(err).Msg("offsetz position could not be parsed")
				break forLoop
			}
			oz.position = uint16(buf)
			oz.positionSet = true
				log.Info().Msg("offsetz position successfully parsed")
		} else if !oz.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
					log.Err(err).Msg("offsetz datatype could not be parsed")
				break forLoop
			}
			oz.datatype = buf
			oz.datatypeSet = true
				log.Info().Msg("offsetz datatype successfully parsed")
			break forLoop
		}
	}
	return oz, err
}
