package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type fncValues struct {
	position       uint16
	positionSet    bool
	datatype       dataTypeEnum
	datatypeSet    bool
	indexMode      IndexModeEnum
	indexModeSet   bool
	addresstype    AddrTypeEnum
	addresstypeSet bool
}

func parseFncValues(tok *tokenGenerator) (fncValues, error) {
	fv := fncValues{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("fncValues could not be parsed")
			break forLoop
		} else if !fv.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("fncValues position could not be parsed")
				break forLoop
			}
			fv.position = uint16(buf)
			fv.positionSet = true
			log.Info().Msg("fncValues position successfully parsed")
		} else if !fv.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("fncValues datatype could not be parsed")
				break forLoop
			}
			fv.datatype = buf
			fv.datatypeSet = true
			log.Info().Msg("fncValues datatype successfully parsed")
		} else if !fv.indexModeSet {
			var buf IndexModeEnum
			buf, err = parseIndexModeEnum(tok)
			if err != nil {
				log.Err(err).Msg("fncValues indexMode could not be parsed")
				break forLoop
			}
			fv.indexMode = buf
			fv.indexModeSet = true
			log.Info().Msg("fncValues indexMode successfully parsed")
		} else if !fv.addresstypeSet {
			var buf AddrTypeEnum
			buf, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("fncValues addresstype could not be parsed")
				break forLoop
			}
			fv.addresstype = buf
			fv.addresstypeSet = true
			log.Info().Msg("fncValues addresstype successfully parsed")
			break forLoop
		}
	}
	return fv, err
}
