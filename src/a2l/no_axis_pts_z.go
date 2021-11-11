package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type noAxisPtsZ struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseNoAxisPtsZ(tok *tokenGenerator) (noAxisPtsZ, error) {
	napz := noAxisPtsZ{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("noAxisPtsz could not be parsed")
			break forLoop
		} else if !napz.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noAxisPtsz position could not be parsed")
				break forLoop
			}
			napz.position = uint16(buf)
			napz.positionSet = true
			log.Info().Msg("noAxisPtsz position successfully parsed")
		} else if !napz.datatypeSet {
			napz.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noAxisPtsz datatype could not be parsed")
				break forLoop
			}
			napz.datatypeSet = true
			log.Info().Msg("noAxisPtsz datatype successfully parsed")
			break forLoop
		}
	}
	return napz, err
}
