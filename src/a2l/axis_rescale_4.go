package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type axisRescale4 struct {
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

func parseAxisRescale4(tok *tokenGenerator) (axisRescale4, error) {
	ar4 := axisRescale4{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("axisRescale4 could not be parsed")
			break forLoop
		} else if !ar4.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisRescale4 position could not be parsed")
				break forLoop
			}
			ar4.position = uint16(buf)
			ar4.positionSet = true
			log.Info().Msg("axisRescale4 position successfully parsed")
		} else if !ar4.datatypeSet {
			ar4.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisRescale4 datatype could not be parsed")
				break forLoop
			}
			ar4.datatypeSet = true
			log.Info().Msg("axisRescale4 datatype successfully parsed")
		} else if !ar4.maxNumberOfRescalePairsSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisRescale4 maxNumberOfRescalePairs could not be parsed")
				break forLoop
			}
			ar4.maxNumberOfRescalePairs = uint16(buf)
			ar4.maxNumberOfRescalePairsSet = true
			log.Info().Msg("axisRescale4 maxNumberOfRescalePairs successfully parsed")
		} else if !ar4.indexIncrSet {
			ar4.indexIncr, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisRescale4 indexIncr could not be parsed")
				break forLoop
			}
			ar4.indexIncrSet = true
			log.Info().Msg("axisRescale4 indexIncr successfully parsed")
		} else if !ar4.adressingSet {
			ar4.adressing, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisRescale4 adressing could not be parsed")
				break forLoop
			}
			ar4.adressingSet = true
			log.Info().Msg("axisRescale4 adressing successfully parsed")
			break forLoop
		}
	}
	return ar4, err
}
