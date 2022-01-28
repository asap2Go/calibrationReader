package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type ripAddrX struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseRipAddrX(tok *tokenGenerator) (ripAddrX, error) {
	rax := ripAddrX{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("ripAddrx could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("ripAddrX could not be parsed")
			break forLoop
		} else if !rax.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("ripAddrx position could not be parsed")
				break forLoop
			}
			rax.position = uint16(buf)
			rax.positionSet = true
			log.Info().Msg("ripAddrx position successfully parsed")
		} else if !rax.datatypeSet {
			rax.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("ripAddrx datatype could not be parsed")
				break forLoop
			}
			rax.datatypeSet = true
			log.Info().Msg("ripAddrx datatype successfully parsed")
			break forLoop
		}
	}
	return rax, err
}
