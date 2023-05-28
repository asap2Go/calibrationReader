package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type AxisPtsZ struct {
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

func parseAxisPtsZ(tok *tokenGenerator) (AxisPtsZ, error) {
	apZ := AxisPtsZ{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("axisPtsZ could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("axisPtsZ could not be parsed")
			break forLoop
		} else if !apZ.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisPtsZ position could not be parsed")
				break forLoop
			}
			apZ.Position = uint16(buf)
			apZ.PositionSet = true
			log.Info().Msg("axisPtsZ position successfully parsed")
		} else if !apZ.DatatypeSet {
			apZ.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsZ datatype could not be parsed")
				break forLoop
			}
			apZ.DatatypeSet = true
			log.Info().Msg("axisPtsZ datatype successfully parsed")
		} else if !apZ.IndexIncrSet {
			apZ.IndexIncr, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsZ indexIncr could not be parsed")
				break forLoop
			}
			apZ.IndexIncrSet = true
			log.Info().Msg("axisPtsZ indexIncr successfully parsed")
		} else if !apZ.AddressingSet {
			apZ.Addressing, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsZ addressing could not be parsed")
				break forLoop
			}
			apZ.AddressingSet = true
			log.Info().Msg("axisPtsZ addressing successfully parsed")
			break forLoop
		}
	}
	return apZ, err
}
