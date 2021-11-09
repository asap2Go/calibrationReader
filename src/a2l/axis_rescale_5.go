package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type axisRescale5 struct {
	position                   uint16
	positionSet                bool
	datatype                   dataTypeEnum
	datatypeSet                bool
	maxNumberOfRescalePairs    uint16
	maxNumberOfRescalePairsSet bool
	indexIncr                  IndexOrderEnum
	indexIncrSet               bool
	adressing                  AddrTypeEnum
	adressingSet               bool
}

func parseAxisRescale5(tok *tokenGenerator) (axisRescale5, error) {
	ar5 := axisRescale5{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("axisRescale5 could not be parsed")
			break forLoop
		} else if !ar5.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
					log.Err(err).Msg("axisRescale5 position could not be parsed")
				break forLoop
			}
			ar5.position = uint16(buf)
			ar5.positionSet = true
				log.Info().Msg("axisRescale5 position successfully parsed")
		} else if !ar5.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
					log.Err(err).Msg("axisRescale5 datatype could not be parsed")
				break forLoop
			}
			ar5.datatype = buf
			ar5.datatypeSet = true
				log.Info().Msg("axisRescale5 datatype successfully parsed")
		} else if !ar5.maxNumberOfRescalePairsSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
					log.Err(err).Msg("axisRescale5 maxNumberOfRescalePairs could not be parsed")
				break forLoop
			}
			ar5.maxNumberOfRescalePairs = uint16(buf)
			ar5.maxNumberOfRescalePairsSet = true
				log.Info().Msg("axisRescale5 maxNumberOfRescalePairs successfully parsed")
		} else if !ar5.indexIncrSet {
			var buf IndexOrderEnum
			buf, err = parseIndexOrderEnum(tok)
			if err != nil {
					log.Err(err).Msg("axisRescale5 indexIncr could not be parsed")
				break forLoop
			}
			ar5.indexIncr = buf
			ar5.indexIncrSet = true
				log.Info().Msg("axisRescale5 indexIncr successfully parsed")
		} else if !ar5.adressingSet {
			var buf AddrTypeEnum
			buf, err = parseAddrTypeEnum(tok)
			if err != nil {
					log.Err(err).Msg("axisRescale5 adressing could not be parsed")
				break forLoop
			}
			ar5.adressing = buf
			ar5.adressingSet = true
				log.Info().Msg("axisRescale5 adressing successfully parsed")
			break forLoop
		}
	}
	return ar5, err
}
