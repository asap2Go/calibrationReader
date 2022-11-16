package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type AxisPtsX struct {
	//Position of the axis point values in the deposit structure(description of sequence of elements in the data record).
	//If the Alternate option is used with FNC_VALUES, the Position parameter determines the order of values and axis points.
	Position      uint16
	PositionSet   bool
	Datatype      DataTypeEnum
	DatatypeSet   bool
	IndexIncr     indexOrderEnum
	IndexIncrSet  bool
	Addressing    addrTypeEnum
	AddressingSet bool
	Values        interface{}
	ValuesSet     bool
}

func parseAxisPtsX(tok *tokenGenerator) (AxisPtsX, error) {
	apX := AxisPtsX{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("axisPtsX could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("axisPtsX could not be parsed")
			break forLoop
		} else if !apX.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisPtsX position could not be parsed")
				break forLoop
			}
			apX.Position = uint16(buf)
			apX.PositionSet = true
			log.Info().Msg("axisPtsX position successfully parsed")
		} else if !apX.DatatypeSet {
			apX.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsX datatype could not be parsed")
				break forLoop
			}
			apX.DatatypeSet = true
			log.Info().Msg("axisPtsX datatype successfully parsed")
		} else if !apX.IndexIncrSet {
			apX.IndexIncr, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsX indexIncr could not be parsed")
				break forLoop
			}
			apX.IndexIncrSet = true
			log.Info().Msg("axisPtsX indexIncr successfully parsed")
		} else if !apX.AddressingSet {
			apX.Addressing, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsX addressing could not be parsed")
				break forLoop
			}
			apX.AddressingSet = true
			log.Info().Msg("axisPtsX addressing successfully parsed")
			break forLoop
		}
	}
	return apX, err
}
