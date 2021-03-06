package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type fncValues struct {
	//position of table values (function values) in the deposit structure (description of sequence of elements in the data record).
	position    uint16
	positionSet bool
	//data type of the table values
	datatype    dataTypeEnum
	datatypeSet bool
	//for characteristic maps, this attribute is used to describe how the 2-dimensional table values are mapped onto the 1-dimensional address space
	indexMode      indexModeEnum
	indexModeSet   bool
	addresstype    addrTypeEnum
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
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
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
			fv.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("fncValues datatype could not be parsed")
				break forLoop
			}
			fv.datatypeSet = true
			log.Info().Msg("fncValues datatype successfully parsed")
		} else if !fv.indexModeSet {
			fv.indexMode, err = parseIndexModeEnum(tok)
			if err != nil {
				log.Err(err).Msg("fncValues indexMode could not be parsed")
				break forLoop
			}
			fv.indexModeSet = true
			log.Info().Msg("fncValues indexMode successfully parsed")
		} else if !fv.addresstypeSet {
			fv.addresstype, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("fncValues addresstype could not be parsed")
				break forLoop
			}
			fv.addresstypeSet = true
			log.Info().Msg("fncValues addresstype successfully parsed")
			break forLoop
		}
	}
	return fv, err
}
