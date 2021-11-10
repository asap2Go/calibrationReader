package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type axisRescaleZ struct {
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

func parseAxisRescaleZ(tok *tokenGenerator) (axisRescaleZ, error) {
	arZ := axisRescaleZ{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("AxisRescaleZ could not be parsed")
			break forLoop
		} else if !arZ.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("AxisRescaleZ position could not be parsed")
				break forLoop
			} else {
				arZ.position = uint16(buf)
				arZ.positionSet = true
				log.Info().Msg("AxisRescaleZ position successfully parsed")
			}
		} else if !arZ.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("AxisRescaleZ datatype could not be parsed")
				break forLoop
			} else {
				arZ.datatype = buf
				arZ.datatypeSet = true
				log.Info().Msg("AxisRescaleZ datatype successfully parsed")
			}
		} else if !arZ.maxNumberOfRescalePairsSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("AxisRescaleZ maxNumberOfRescalePairs could not be parsed")
				break forLoop
			} else {
				arZ.maxNumberOfRescalePairs = uint16(buf)
				arZ.maxNumberOfRescalePairsSet = true
				log.Info().Msg("AxisRescaleZ maxNumberOfRescalePairs successfully parsed")
			}
		} else if !arZ.indexIncrSet {
			var buf IndexOrderEnum
			buf, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("AxisRescaleZ indexIncr could not be parsed")
				break forLoop
			} else {
				arZ.indexIncr = buf
				arZ.indexIncrSet = true
				log.Info().Msg("AxisRescaleZ indexIncr successfully parsed")
			}
		} else if !arZ.adressingSet {
			var buf AddrTypeEnum
			buf, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("AxisRescaleZ adressing could not be parsed")
				break forLoop
			} else {
				arZ.adressing = buf
				arZ.adressingSet = true
				log.Info().Msg("AxisRescaleZ adressing successfully parsed")
				break forLoop
			}
		}
	}
	return arZ, err
}
