package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type compuTab struct {
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
	outVal              []float64
	outValSet           bool
	defaultValue        defaultValue
	defaultValueNumeric defaultValueNumeric
}

func parseCompuTab(tok *tokenGenerator) (compuTab, error) {
	ct := compuTab{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case defaultValueToken:
			ct.defaultValue, err = parseDefaultValue(tok)
			if err != nil {
				log.Err(err).Msg("compuTab defaultValue could not be parsed")
				break forLoop
			}
			log.Info().Msg("compuTab defaultValue successfully parsed")
		case defaultValueNumericToken:
			ct.defaultValueNumeric, err = parseDefaultValueNumeric(tok)
			if err != nil {
				log.Err(err).Msg("compuTab defaultValueNumeric could not be parsed")
				break forLoop
			}
			log.Info().Msg("compuTab defaultValueNumeric successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("compuTab could not be parsed")
				break forLoop
			} else if tok.current() == endCompuTabToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("compuTab could not be parsed")
				break forLoop
			} else if !ct.nameSet {
				ct.name = tok.current()
				ct.nameSet = true
				log.Info().Msg("compuTab name successfully parsed")
			} else if !ct.longIdentifierSet {
				ct.longIdentifier = tok.current()
				ct.longIdentifierSet = true
				log.Info().Msg("compuTab longIdentifier successfully parsed")
			} else if !ct.conversionTypeSet {
				ct.conversionType, err = parseConversionTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("compuTab conversionType could not be parsed")
					break forLoop
				}
				ct.conversionTypeSet = true
				log.Info().Msg("compuTab conversionType successfully parsed")
			} else if !ct.numberValuePairsSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 16)
				if err != nil {
					log.Err(err).Msg("compuTab numberValuePairs could not be parsed")
					break forLoop
				}
				ct.numberValuePairs = uint16(buf)
				ct.numberValuePairsSet = true
				log.Info().Msg("compuTab numberValuePairs successfully parsed")
			} else if !ct.inValSet || !ct.outValSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("compuTab inVal could not be parsed")
					break forLoop
				}
				ct.inVal = append(ct.inVal, buf)
				log.Info().Msg("compuTab inVal successfully parsed")
				if uint16(len(ct.inVal)) == ct.numberValuePairs {
					ct.inValSet = true
					log.Info().Msg("compuTab inVal successfully parsed")
				}
				buf, err = strconv.ParseFloat(tok.next(), 64)
				if err != nil {
					log.Err(err).Msg("compuTab outVal could not be parsed")
					break forLoop
				}
				ct.outVal = append(ct.outVal, buf)
				log.Info().Msg("compuTab outVal successfully parsed")
				if uint16(len(ct.outVal)) == ct.numberValuePairs {
					ct.outValSet = true
				}
			}
		}
	}
	return ct, err
}
