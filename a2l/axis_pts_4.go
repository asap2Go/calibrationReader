package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type AxisPts4 struct {
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

func parseAxisPts4(tok *tokenGenerator) (AxisPts4, error) {
	ap4 := AxisPts4{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("axisPts4 could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("axisPts4 could not be parsed")
			break forLoop
		} else if !ap4.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisPts4 position could not be parsed")
				break forLoop
			}
			ap4.Position = uint16(buf)
			ap4.PositionSet = true
			log.Info().Msg("axisPts4 position successfully parsed")
		} else if !ap4.DatatypeSet {
			ap4.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPts4 datatype could not be parsed")
				break forLoop
			}
			ap4.DatatypeSet = true
			log.Info().Msg("axisPts4 datatype successfully parsed")
		} else if !ap4.IndexIncrSet {
			ap4.IndexIncr, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPts4 indexIncr could not be parsed")
				break forLoop
			}
			ap4.IndexIncrSet = true
			log.Info().Msg("axisPts4 indexIncr successfully parsed")
		} else if !ap4.AddressingSet {
			ap4.Addressing, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPts4 addressing could not be parsed")
				break forLoop
			}
			ap4.AddressingSet = true
			log.Info().Msg("axisPts4 addressing successfully parsed")
			break forLoop
		}
	}
	return ap4, err
}
