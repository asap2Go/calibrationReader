package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type reserved struct {
	position    uint16
	positionSet bool
	dataSize    dataSizeEnum
	dataSizeSet bool
}

func parseReserved(tok *tokenGenerator) (reserved, error) {
	r := reserved{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("reserved could not be parsed")
			break forLoop
		} else if !r.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
					log.Err(err).Msg("reserved position could not be parsed")
				break forLoop
			}
			r.position = uint16(buf)
			r.positionSet = true
				log.Info().Msg("reserved position successfully parsed")
		} else if !r.dataSizeSet {
			var buf dataSizeEnum
			buf, err = parseDataSizeEnum(tok)
			if err != nil {
					log.Err(err).Msg("reserved dataSize could not be parsed")
				break forLoop
			}
			r.dataSize = buf
			r.dataSizeSet = true
				log.Info().Msg("reserved dataSize successfully parsed")
			break forLoop
		}
	}
	return r, err
}
