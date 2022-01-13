package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type ripAddr5 struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseRipAddr5(tok *tokenGenerator) (ripAddr5, error) {
	ra5 := ripAddr5{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("ripAddr5 could not be parsed")
			break forLoop
		} else if !ra5.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("ripAddr5 position could not be parsed")
				break forLoop
			}
			ra5.position = uint16(buf)
			ra5.positionSet = true
			log.Info().Msg("ripAddr5 position successfully parsed")
		} else if !ra5.datatypeSet {
			ra5.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("ripAddr5 datatype could not be parsed")
				break forLoop
			}
			ra5.datatypeSet = true
			log.Info().Msg("ripAddr5 datatype successfully parsed")
			break forLoop
		}
	}
	return ra5, err
}
