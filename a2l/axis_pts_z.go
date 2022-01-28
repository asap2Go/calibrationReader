package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type axisPtsZ struct {
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

func parseAxisPtsZ(tok *tokenGenerator) (axisPtsZ, error) {
	apZ := axisPtsZ{}
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
		} else if !apZ.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisPtsZ position could not be parsed")
				break forLoop
			}
			apZ.position = uint16(buf)
			apZ.positionSet = true
			log.Info().Msg("axisPtsZ position successfully parsed")
		} else if !apZ.datatypeSet {
			apZ.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsZ datatype could not be parsed")
				break forLoop
			}
			apZ.datatypeSet = true
			log.Info().Msg("axisPtsZ datatype successfully parsed")
		} else if !apZ.indexIncrSet {
			apZ.indexIncr, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsZ indexIncr could not be parsed")
				break forLoop
			}
			apZ.indexIncrSet = true
			log.Info().Msg("axisPtsZ indexIncr successfully parsed")
		} else if !apZ.addressingSet {
			apZ.addressing, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPtsZ addressing could not be parsed")
				break forLoop
			}
			apZ.addressingSet = true
			log.Info().Msg("axisPtsZ addressing successfully parsed")
			break forLoop
		}
	}
	return apZ, err
}
