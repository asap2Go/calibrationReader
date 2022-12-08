package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type CompuTab struct {
	Name                string
	NameSet             bool
	longIdentifier      string
	longIdentifierSet   bool
	ConversionType      conversionTypeEnum
	ConversionTypeSet   bool
	NumberValuePairs    uint16
	NumberValuePairsSet bool
	InVal               []float64
	InValSet            bool
	OutVal              []float64
	OutValSet           bool
	DefaultValue        DefaultValue
	DefaultValueNumeric DefaultValueNumeric
}

func parseCompuTab(tok *tokenGenerator) (CompuTab, error) {
	ct := CompuTab{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case defaultValueToken:
			ct.DefaultValue, err = parseDefaultValue(tok)
			if err != nil {
				log.Err(err).Msg("compuTab defaultValue could not be parsed")
				break forLoop
			}
			log.Info().Msg("compuTab defaultValue successfully parsed")
		case defaultValueNumericToken:
			ct.DefaultValueNumeric, err = parseDefaultValueNumeric(tok)
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
			} else if !ct.NameSet {
				ct.Name = tok.current()
				ct.NameSet = true
				log.Info().Msg("compuTab name successfully parsed")
			} else if !ct.longIdentifierSet {
				ct.longIdentifier = tok.current()
				ct.longIdentifierSet = true
				log.Info().Msg("compuTab longIdentifier successfully parsed")
			} else if !ct.ConversionTypeSet {
				ct.ConversionType, err = parseConversionTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("compuTab conversionType could not be parsed")
					break forLoop
				}
				ct.ConversionTypeSet = true
				log.Info().Msg("compuTab conversionType successfully parsed")
			} else if !ct.NumberValuePairsSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 16)
				if err != nil {
					log.Err(err).Msg("compuTab numberValuePairs could not be parsed")
					break forLoop
				}
				ct.NumberValuePairs = uint16(buf)
				ct.NumberValuePairsSet = true
				log.Info().Msg("compuTab numberValuePairs successfully parsed")
			} else if !ct.InValSet || !ct.OutValSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("compuTab inVal could not be parsed")
					break forLoop
				}
				ct.InVal = append(ct.InVal, buf)
				log.Info().Msg("compuTab inVal successfully parsed")
				if uint16(len(ct.InVal)) == ct.NumberValuePairs {
					ct.InValSet = true
					log.Info().Msg("compuTab inVal successfully parsed")
				}
				buf, err = strconv.ParseFloat(tok.next(), 64)
				if err != nil {
					log.Err(err).Msg("compuTab outVal could not be parsed")
					break forLoop
				}
				ct.OutVal = append(ct.OutVal, buf)
				log.Info().Msg("compuTab outVal successfully parsed")
				if uint16(len(ct.OutVal)) == ct.NumberValuePairs {
					ct.OutValSet = true
				}
			}
		}
	}
	return ct, err
}
