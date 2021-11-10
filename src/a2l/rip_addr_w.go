package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type ripAddrW struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseRipAddrW(tok *tokenGenerator) (ripAddrW, error) {
	raw := ripAddrW{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("ripAddrw could not be parsed")
			break forLoop
		} else if !raw.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("ripAddrw position could not be parsed")
				break forLoop
			}
			raw.position = uint16(buf)
			raw.positionSet = true
			log.Info().Msg("ripAddrw position successfully parsed")
		} else if !raw.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("ripAddrw datatype could not be parsed")
				break forLoop
			}
			raw.datatype = buf
			raw.datatypeSet = true
			log.Info().Msg("ripAddrw datatype successfully parsed")
			break forLoop
		}
	}
	return raw, err
}
