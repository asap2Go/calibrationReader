package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type axisRescaleY struct {
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

func parseAxisRescaleY(tok *tokenGenerator) (axisRescaleY, error) {
	arY := axisRescaleY{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("AxisRescaleY could not be parsed")
			break forLoop
		} else if !arY.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("AxisRescaleY position could not be parsed")
				break forLoop
			} else {
				arY.position = uint16(buf)
				arY.positionSet = true
				log.Info().Msg("AxisRescaleY position successfully parsed")
			}
		} else if !arY.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("AxisRescaleY datatype could not be parsed")
				break forLoop
			} else {
				arY.datatype = buf
				arY.datatypeSet = true
				log.Info().Msg("AxisRescaleY datatype successfully parsed")
			}
		} else if !arY.maxNumberOfRescalePairsSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("AxisRescaleY maxNumberOfRescalePairs could not be parsed")
				break forLoop
			} else {
				arY.maxNumberOfRescalePairs = uint16(buf)
				arY.maxNumberOfRescalePairsSet = true
				log.Info().Msg("AxisRescaleY maxNumberOfRescalePairs successfully parsed")
			}
		} else if !arY.indexIncrSet {
			var buf IndexOrderEnum
			buf, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("AxisRescaleY indexIncr could not be parsed")
				break forLoop
			} else {
				arY.indexIncr = buf
				arY.indexIncrSet = true
				log.Info().Msg("AxisRescaleY indexIncr successfully parsed")
			}
		} else if !arY.adressingSet {
			var buf AddrTypeEnum
			buf, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("AxisRescaleY adressing could not be parsed")
				break forLoop
			} else {
				arY.adressing = buf
				arY.adressingSet = true
				log.Info().Msg("AxisRescaleY adressing successfully parsed")
				break forLoop
			}
		}
	}
	return arY, err
}
