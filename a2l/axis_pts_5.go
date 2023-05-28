package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type AxisPts5 struct {
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

func parseAxisPts5(tok *tokenGenerator) (AxisPts5, error) {
	ap5 := AxisPts5{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("axisPts5 could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("axisPts5 could not be parsed")
			break forLoop
		} else if !ap5.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisPts5 position could not be parsed")
				break forLoop
			}
			ap5.Position = uint16(buf)
			ap5.PositionSet = true
			log.Info().Msg("axisPts5 position successfully parsed")
		} else if !ap5.DatatypeSet {
			ap5.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPts5 datatype could not be parsed")
				break forLoop
			}
			ap5.DatatypeSet = true
			log.Info().Msg("axisPts5 datatype successfully parsed")
		} else if !ap5.IndexIncrSet {
			ap5.IndexIncr, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPts5 indexIncr could not be parsed")
				break forLoop
			}
			ap5.IndexIncrSet = true
			log.Info().Msg("axisPts5 indexIncr successfully parsed")
		} else if !ap5.AddressingSet {
			ap5.Addressing, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPts5 addressing could not be parsed")
				break forLoop
			}
			ap5.AddressingSet = true
			log.Info().Msg("axisPts5 addressing successfully parsed")
			break forLoop
		}
	}
	return ap5, err
}
