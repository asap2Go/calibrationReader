package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type CompuMethod struct {
	Name              string
	NameSet           bool
	LongIdentifier    string
	LongIdentifierSet bool
	ConversionType    conversionTypeEnum
	ConversionTypeSet bool
	Format            string
	FormatSet         bool
	Unit              string
	UnitSet           bool
	Coeffs            Coeffs
	CoeffsLinear      CoeffsLinear
	CompuTabRef       CompuTabRef
	Formula           []Formula
	RefUnit           refUnit
	StatusStringRef   StatusStringRef
}

func parseCompuMethod(tok *tokenGenerator) (CompuMethod, error) {
	cm := CompuMethod{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case coeffsToken:
			cm.Coeffs, err = parseCoeffs(tok)
			if err != nil {
				log.Err(err).Msg("compuMethod coeffs could not be parsed")
				break forLoop
			}
			log.Info().Msg("compuMethod coeffs successfully parsed")
		case coeffsLinearToken:
			cm.CoeffsLinear, err = parseCoeffsLinear(tok)
			if err != nil {
				log.Err(err).Msg("compuMethod coeffsLinear could not be parsed")
				break forLoop
			}
			log.Info().Msg("compuMethod coeffsLinear successfully parsed")
		case compuTabRefToken:
			cm.CompuTabRef, err = parseCompuTabRef(tok)
			if err != nil {
				log.Err(err).Msg("compuMethod compuTabRef could not be parsed")
				break forLoop
			}
			log.Info().Msg("compuMethod compuTabRef successfully parsed")
		case beginFormulaToken:
			var buf Formula
			buf, err = parseFormula(tok)
			if err != nil {
				log.Err(err).Msg("compuMethod formula could not be parsed")
				break forLoop
			}
			cm.Formula = append(cm.Formula, buf)
			log.Info().Msg("compuMethod formula successfully parsed")
		case refUnitToken:
			cm.RefUnit, err = parseRefUnit(tok)
			if err != nil {
				log.Err(err).Msg("compuMethod refUnit could not be parsed")
				break forLoop
			}
			log.Info().Msg("compuMethod refUnit successfully parsed")
		case statusStringRefToken:
			cm.StatusStringRef, err = parseStatusStringRef(tok)
			if err != nil {
				log.Err(err).Msg("compuMethod statusStringRef could not be parsed")
				break forLoop
			}
			log.Info().Msg("compuMethod statusStringRef successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unedecpected end of file")
				log.Err(err).Msg("compuMethod could not be parsed")
				break forLoop
			} else if tok.current() == endCompuMethodToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unedecpected token " + tok.current())
				log.Err(err).Msg("compuMethod could not be parsed")
				break forLoop
			} else if !cm.NameSet {
				cm.Name = tok.current()
				cm.NameSet = true
				log.Info().Msg("compuMethod name successfully parsed")
			} else if !cm.LongIdentifierSet {
				cm.LongIdentifier = tok.current()
				cm.LongIdentifierSet = true
				log.Info().Msg("compuMethod longIdentifier successfully parsed")
			} else if !cm.ConversionTypeSet {
				cm.ConversionType, err = parseConversionTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("compuMethod conversionType could not be parsed")
					break forLoop
				}
				cm.ConversionTypeSet = true
				log.Info().Msg("compuMethod conversionType successfully parsed")
			} else if !cm.FormatSet {
				cm.Format = tok.current()
				cm.FormatSet = true
				log.Info().Msg("compuMethod format successfully parsed")
			} else if !cm.UnitSet {
				cm.Unit = tok.current()
				cm.UnitSet = true
				log.Info().Msg("compuMethod unit successfully parsed")
			}
		}
	}
	return cm, err
}
