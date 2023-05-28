package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type AxisPtsY struct {
	//Position of the axis point values in the deposit structure(description of sequence of elements in the data record).
	//If the Alternate option is used with FNC_VALUES, the Position parameter determines the order of values and axis points.
	Position      uint16
	PositionSet   bool
	Datatype      DataTypeEnum
	DatatypeSet   bool
	IndexIncr     indexOrderEnum
	IndexIncrSet  bool
	Addressing    AddrTypeEnum
	AddressingSet bool
	Values        interface{}
	ValuesSet     bool
}

func parseAxisPtsY(tok *tokenGenerator) (AxisPtsY, error) {
	apY := AxisPtsY{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("axisPtsY could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("axisPtsY could not be parsed")
			break forLoop
		} else if !apY.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisPtsY position could not be parsed")
				break forLoop
			}
			apY.Position = uint16(buf)
			apY.PositionSet = true
			log.Info().Msg("axisPtsY position successfully parsed")
		} else if !apY.DatatypeSet {
			apY.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsY datatype could not be parsed")
				break forLoop
			}
			apY.DatatypeSet = true
			log.Info().Msg("axisPtsY datatype successfully parsed")
		} else if !apY.IndexIncrSet {
			apY.IndexIncr, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsY indexIncr could not be parsed")
				break forLoop
			}
			apY.IndexIncrSet = true
			log.Info().Msg("axisPtsY indexIncr successfully parsed")
		} else if !apY.AddressingSet {
			apY.Addressing, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsY addressing could not be parsed")
				break forLoop
			}
			apY.AddressingSet = true
			log.Info().Msg("axisPtsY addressing successfully parsed")
			break forLoop
		}
	}
	return apY, err
}
