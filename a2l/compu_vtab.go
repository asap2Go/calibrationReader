package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type CompuVTab struct {
	Name                string
	NameSet             bool
	LongIdentifier      string
	LongIdentifierSet   bool
	ConversionType      conversionTypeEnum
	ConversionTypeSet   bool
	NumberValuePairs    uint16
	NumberValuePairsSet bool
	InVal               []float64
	InValSet            bool
	OutVal              []string
	OutValSet           bool
	DefaultValue        DefaultValue
}

func parseCompuVtab(tok *tokenGenerator) (CompuVTab, error) {
	cvt := CompuVTab{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case defaultValueToken:
			cvt.DefaultValue, err = parseDefaultValue(tok)
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
			} else if !cvt.NameSet {
				cvt.Name = tok.current()
				cvt.NameSet = true
				log.Info().Msg("compuVtab name successfully parsed")
			} else if !cvt.LongIdentifierSet {
				cvt.LongIdentifier = tok.current()
				cvt.LongIdentifierSet = true
				log.Info().Msg("compuVtab longIdentifier successfully parsed")
			} else if !cvt.ConversionTypeSet {
				cvt.ConversionType, err = parseConversionTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("compuVtab conversionType could not be parsed")
					break forLoop
				}
				cvt.ConversionTypeSet = true
				log.Info().Msg("compuVtab conversionType successfully parsed")
			} else if !cvt.NumberValuePairsSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 16)
				if err != nil {
					log.Err(err).Msg("compuVtab numberValuePairs could not be parsed")
					break forLoop
				}
				cvt.NumberValuePairs = uint16(buf)
				cvt.NumberValuePairsSet = true
			} else if !cvt.InValSet || !cvt.OutValSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("compuVtab inVal could not be parsed")
					break forLoop
				}
				cvt.InVal = append(cvt.InVal, buf)
				log.Info().Msg("compuVtab inVal successfully parsed")
				if uint16(len(cvt.InVal)) == cvt.NumberValuePairs {
					cvt.InValSet = true
					log.Info().Msg("compuVtab inVal successfully parsed")
				}
				tok.next()
				cvt.OutVal = append(cvt.OutVal, tok.current())
				log.Info().Msg("compuVtab outVal successfully parsed")
				if uint16(len(cvt.OutVal)) == cvt.NumberValuePairs {
					cvt.OutValSet = true
				}
			}
		}
	}
	return cvt, err
}
