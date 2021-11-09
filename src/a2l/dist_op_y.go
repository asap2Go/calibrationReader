package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type distOpY struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseDistOpY(tok *tokenGenerator) (distOpY, error) {
	doy := distOpY{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("distOpy could not be parsed")
			break forLoop
		} else if !doy.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
					log.Err(err).Msg("distOpy position could not be parsed")
				break forLoop
			}
			doy.position = uint16(buf)
			doy.positionSet = true
				log.Info().Msg("distOpy position successfully parsed")
		} else if !doy.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
					log.Err(err).Msg("distOpy datatype could not be parsed")
				break forLoop
			}
			doy.datatype = buf
			doy.datatypeSet = true
				log.Info().Msg("distOpy datatype successfully parsed")
			break forLoop
		}
	}
	return doy, err
}
