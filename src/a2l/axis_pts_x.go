package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type axisPtsX struct {
	//position of the axis point values in the deposit structure(description of sequence of elements in the data record).
	//If the Alternate option is used with FNC_VALUES, the position parameter determines the order of values and axis points.
	position      uint16
	positionSet   bool
	datatype      dataTypeEnum
	datatypeSet   bool
	indexIncr     IndexOrderEnum
	indexIncrSet  bool
	addressing    AddrTypeEnum
	addressingSet bool
}

func parseAxisPtsX(tok *tokenGenerator) (axisPtsX, error) {
	apX := axisPtsX{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("axisPtsX could not be parsed")
			break forLoop
		} else if !apX.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisPtsX position could not be parsed")
				break forLoop
			}
			apX.position = uint16(buf)
			apX.positionSet = true
			log.Info().Msg("axisPtsX position successfully parsed")
		} else if !apX.datatypeSet {
			apX.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsX datatype could not be parsed")
				break forLoop
			}
			apX.datatypeSet = true
			log.Info().Msg("axisPtsX datatype successfully parsed")
		} else if !apX.indexIncrSet {
			apX.indexIncr, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsX indexIncr could not be parsed")
				break forLoop
			}
			apX.indexIncrSet = true
			log.Info().Msg("axisPtsX indexIncr successfully parsed")
		} else if !apX.addressingSet {
			apX.addressing, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsX addressing could not be parsed")
				break forLoop
			}
			apX.addressingSet = true
			log.Info().Msg("axisPtsX addressing successfully parsed")
			break forLoop
		}
	}
	return apX, err
}
