package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type axisRescaleX struct {
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

func parseAxisRescaleX(tok *tokenGenerator) (axisRescaleX, error) {
	arX := axisRescaleX{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("AxisRescaleX could not be parsed")
			break forLoop
		} else if !arX.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
					log.Err(err).Msg("AxisRescaleX position could not be parsed")
				break forLoop
			}
			arX.position = uint16(buf)
			arX.positionSet = true
				log.Info().Msg("AxisRescaleX position successfully parsed")
		} else if !arX.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
					log.Err(err).Msg("AxisRescaleX datatype could not be parsed")
				break forLoop
			}
			arX.datatype = buf
			arX.datatypeSet = true
				log.Info().Msg("AxisRescaleX datatype successfully parsed")
		} else if !arX.maxNumberOfRescalePairsSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
					log.Err(err).Msg("AxisRescaleX maxNumberOfRescalePairs could not be parsed")
				break forLoop
			}
			arX.maxNumberOfRescalePairs = uint16(buf)
			arX.maxNumberOfRescalePairsSet = true
				log.Info().Msg("AxisRescaleX maxNumberOfRescalePairs successfully parsed")
		} else if !arX.indexIncrSet {
			var buf IndexOrderEnum
			buf, err = parseIndexOrderEnum(tok)
			if err != nil {
					log.Err(err).Msg("AxisRescaleX indexIncr could not be parsed")
				break forLoop
			}
			arX.indexIncr = buf
			arX.indexIncrSet = true
				log.Info().Msg("AxisRescaleX indexIncr successfully parsed")
		} else if !arX.adressingSet {
			var buf AddrTypeEnum
			buf, err = parseAddrTypeEnum(tok)
			if err != nil {
					log.Err(err).Msg("AxisRescaleX adressing could not be parsed")
				break forLoop
			}
			arX.adressing = buf
			arX.adressingSet = true
				log.Info().Msg("AxisRescaleX adressing successfully parsed")
			break forLoop
		}
	}
	return arX, err
}
