package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type dataSize struct {
	size    uint16
	sizeSet bool
}

func parseDataSize(tok *tokenGenerator) (dataSize, error) {
	ds := dataSize{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("dataSize could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("dataSize could not be parsed")
	} else if !ds.sizeSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("dataSize size could not be parsed")
		}
		ds.size = uint16(buf)
		ds.sizeSet = true
		log.Info().Msg("dataSize size successfully parsed")
	}
	return ds, err
}
