package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type axisPtsY struct {
	//position of the axis point values in the deposit structure(description of sequence of elements in the data record).
	//If the Alternate option is used with FNC_VALUES, the position parameter determines the order of values and axis points.
	position      uint16
	positionSet   bool
	datatype      dataTypeEnum
	datatypeSet   bool
	indexIncr     indexOrderEnum
	indexIncrSet  bool
	addressing    addrTypeEnum
	addressingSet bool
}

func parseAxisPtsY(tok *tokenGenerator) (axisPtsY, error) {
	apY := axisPtsY{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("axisPtsY could not be parsed")
			break forLoop
		} else if !apY.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisPtsY position could not be parsed")
				break forLoop
			}
			apY.position = uint16(buf)
			apY.positionSet = true
			log.Info().Msg("axisPtsY position successfully parsed")
		} else if !apY.datatypeSet {
			apY.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsY datatype could not be parsed")
				break forLoop
			}
			apY.datatypeSet = true
			log.Info().Msg("axisPtsY datatype successfully parsed")
		} else if !apY.indexIncrSet {
			apY.indexIncr, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsY indexIncr could not be parsed")
				break forLoop
			}
			apY.indexIncrSet = true
			log.Info().Msg("axisPtsY indexIncr successfully parsed")
		} else if !apY.addressingSet {
			apY.addressing, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsY addressing could not be parsed")
				break forLoop
			}
			apY.addressingSet = true
			log.Info().Msg("axisPtsY addressing successfully parsed")
			break forLoop
		}
	}
	return apY, err
}
