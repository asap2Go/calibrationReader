package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type arraySize struct {
	number    uint16
	numberSet bool
}

func parseArraySize(tok *tokenGenerator) (arraySize, error) {
	as := arraySize{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("arraySize could not be parsed")
	} else if !as.numberSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
				log.Err(err).Msg("arraySize number could not be parsed")
		}
		as.number = uint16(buf)
		as.numberSet = true
			log.Info().Msg("arraySize number successfully parsed")
	}
	return as, err
}
