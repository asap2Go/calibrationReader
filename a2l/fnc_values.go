package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

// FncValues defines the way the data of a specific record layout is stored within the hex file.
type FncValues struct {
	//Position of table values (function values) in the deposit structure (description of sequence of elements in the data record).
	Position    uint16
	PositionSet bool
	//Datatype of the table values
	Datatype    DataTypeEnum
	DatatypeSet bool
	/*IndexMode for characteristic maps, curves and value blocks,
	this field is used to describe how the 2-dimensional table values
	are mapped onto the 1-dimensional address space*/
	IndexMode    indexModeEnum
	IndexModeSet bool
	/*Addresstype defines the addressing of the table values:
	Enumeration for description of the addressing of table
	values or axis point values:
	PBYTE:	The relevant memory location has a 1 byte pointer
			to this table value or axis point value.
	PWORD:	The relevant memory location has a 2 byte pointer
			to this table value or axis point value.
	PLONG:	The relevant memory location has a 4 byte pointer
			to this table value or axis point value.
	PLONGLONG:	The relevant memory location has a 8 byte pointer
			to this table value or axis point value.
	DIRECT:	The relevant memory location has the first table value
			or axis point value, all others follow with incrementing address. */
	Addresstype    AddrTypeEnum
	AddresstypeSet bool
}

func parseFncValues(tok *tokenGenerator) (FncValues, error) {
	fv := FncValues{}
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
		} else if !fv.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("fncValues position could not be parsed")
				break forLoop
			}
			fv.Position = uint16(buf)
			fv.PositionSet = true
			log.Info().Msg("fncValues position successfully parsed")
		} else if !fv.DatatypeSet {
			fv.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("fncValues datatype could not be parsed")
				break forLoop
			}
			fv.DatatypeSet = true
			log.Info().Msg("fncValues datatype successfully parsed")
		} else if !fv.IndexModeSet {
			fv.IndexMode, err = parseIndexModeEnum(tok)
			if err != nil {
				log.Err(err).Msg("fncValues indexMode could not be parsed")
				break forLoop
			}
			fv.IndexModeSet = true
			log.Info().Msg("fncValues indexMode successfully parsed")
		} else if !fv.AddresstypeSet {
			fv.Addresstype, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("fncValues addresstype could not be parsed")
				break forLoop
			}
			fv.AddresstypeSet = true
			log.Info().Msg("fncValues addresstype successfully parsed")
			break forLoop
		}
	}
	return fv, err
}
