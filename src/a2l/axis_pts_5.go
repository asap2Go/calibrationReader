package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type axisPts5 struct {
	position      uint16
	positionSet   bool
	datatype      dataTypeEnum
	datatypeSet   bool
	indexIncr     IndexOrderEnum
	indexIncrSet  bool
	addressing    AddrTypeEnum
	addressingSet bool
}

func parseAxisPts5(tok *tokenGenerator) (axisPts5, error) {
	ap5 := axisPts5{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("axisPts5 could not be parsed")
			break forLoop
		} else if !ap5.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisPts5 position could not be parsed")
				break forLoop
			}
			ap5.position = uint16(buf)
			ap5.positionSet = true
			log.Info().Msg("axisPts5 position successfully parsed")
		} else if !ap5.datatypeSet {
			ap5.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPts5 datatype could not be parsed")
				break forLoop
			}
			ap5.datatypeSet = true
			log.Info().Msg("axisPts5 datatype successfully parsed")
		} else if !ap5.indexIncrSet {
			ap5.indexIncr, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPts5 indexIncr could not be parsed")
				break forLoop
			}
			ap5.indexIncrSet = true
			log.Info().Msg("axisPts5 indexIncr successfully parsed")
		} else if !ap5.addressingSet {
			ap5.addressing, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPts5 addressing could not be parsed")
				break forLoop
			}
			ap5.addressingSet = true
			log.Info().Msg("axisPts5 addressing successfully parsed")
			break forLoop
		}
	}
	return ap5, err
}
