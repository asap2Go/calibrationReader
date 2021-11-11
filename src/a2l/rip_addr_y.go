package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type ripAddrY struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseRipAddrY(tok *tokenGenerator) (ripAddrY, error) {
	ray := ripAddrY{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("ripAddry could not be parsed")
			break forLoop
		} else if !ray.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("ripAddry position could not be parsed")
				break forLoop
			}
			ray.position = uint16(buf)
			ray.positionSet = true
			log.Info().Msg("ripAddry position successfully parsed")
		} else if !ray.datatypeSet {
			ray.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("ripAddry datatype could not be parsed")
				break forLoop
			}
			ray.datatypeSet = true
			log.Info().Msg("ripAddry datatype successfully parsed")
			break forLoop
		}
	}
	return ray, err
}
