package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type compuVTab struct {
	name                string
	nameSet             bool
	longIdentifier      string
	longIdentifierSet   bool
	conversionType      conversionTypeEnum
	conversionTypeSet   bool
	numberValuePairs    uint16
	numberValuePairsSet bool
	inVal               []float64
	inValSet            bool
	outVal              []string
	outValSet           bool
	defaultValue        defaultValue
}

func parseCompuVtab(tok *tokenGenerator) (compuVTab, error) {
	cvt := compuVTab{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case defaultValueToken:
			cvt.defaultValue, err = parseDefaultValue(tok)
			if err != nil {
				log.Err(err).Msg("compuVtab defaultValue could not be parsed")
				break forLoop
			}
			log.Info().Msg("compuVtab defaultValue successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("compuVtab could not be parsed")
				break forLoop
			} else if tok.current() == endCompuVtabToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("compuVTab could not be parsed")
				break forLoop
			} else if !cvt.nameSet {
				cvt.name = tok.current()
				cvt.nameSet = true
				log.Info().Msg("compuVtab name successfully parsed")
			} else if !cvt.longIdentifierSet {
				cvt.longIdentifier = tok.current()
				cvt.longIdentifierSet = true
				log.Info().Msg("compuVtab longIdentifier successfully parsed")
			} else if !cvt.conversionTypeSet {
				cvt.conversionType, err = parseConversionTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("compuVtab conversionType could not be parsed")
					break forLoop
				}
				cvt.conversionTypeSet = true
				log.Info().Msg("compuVtab conversionType successfully parsed")
			} else if !cvt.numberValuePairsSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 16)
				if err != nil {
					log.Err(err).Msg("compuVtab numberValuePairs could not be parsed")
					break forLoop
				}
				cvt.numberValuePairs = uint16(buf)
				cvt.numberValuePairsSet = true
			} else if !cvt.inValSet || !cvt.outValSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("compuVtab inVal could not be parsed")
					break forLoop
				}
				cvt.inVal = append(cvt.inVal, buf)
				log.Info().Msg("compuVtab inVal successfully parsed")
				if uint16(len(cvt.inVal)) == cvt.numberValuePairs {
					cvt.inValSet = true
					log.Info().Msg("compuVtab inVal successfully parsed")
				}
				tok.next()
				cvt.outVal = append(cvt.outVal, tok.current())
				log.Info().Msg("compuVtab outVal successfully parsed")
				if uint16(len(cvt.outVal)) == cvt.numberValuePairs {
					cvt.outValSet = true
				}
			}
		}
	}
	return cvt, err
}
