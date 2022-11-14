package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type axisPts4 struct {
	//position of the axis point values in the deposit structure(description of sequence of elements in the data record).
	//If the Alternate option is used with FNC_VALUES, the position parameter determines the order of values and axis points.
	position      uint16
	positionSet   bool
	datatype      DataTypeEnum
	datatypeSet   bool
	indexIncr     indexOrderEnum
	indexIncrSet  bool
	addressing    addrTypeEnum
	addressingSet bool
}

func parseAxisPts4(tok *tokenGenerator) (axisPts4, error) {
	ap4 := axisPts4{}
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
		} else if !ap4.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisPts4 position could not be parsed")
				break forLoop
			}
			ap4.position = uint16(buf)
			ap4.positionSet = true
			log.Info().Msg("axisPts4 position successfully parsed")
		} else if !ap4.datatypeSet {
			ap4.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPts4 datatype could not be parsed")
				break forLoop
			}
			ap4.datatypeSet = true
			log.Info().Msg("axisPts4 datatype successfully parsed")
		} else if !ap4.indexIncrSet {
			ap4.indexIncr, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPts4 indexIncr could not be parsed")
				break forLoop
			}
			ap4.indexIncrSet = true
			log.Info().Msg("axisPts4 indexIncr successfully parsed")
		} else if !ap4.addressingSet {
			ap4.addressing, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPts4 addressing could not be parsed")
				break forLoop
			}
			ap4.addressingSet = true
			log.Info().Msg("axisPts4 addressing successfully parsed")
			break forLoop
		}
	}
	return ap4, err
}
