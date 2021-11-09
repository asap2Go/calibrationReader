package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type compuMethod struct {
	name              string
	nameSet           bool
	longIdentifier    string
	longIdentifierSet bool
	conversionType    ConversionTypeEnum
	conversionTypeSet bool
	format            string
	formatSet         bool
	unit              string
	unitSet           bool
	coeffs            coeffs
	coeffsLinear      coeffsLinear
	compuTabRef       compuTabRef
	formula           []Formula
	refUnit           refUnit
	statusStringRef   statusStringRef
}

func parseCompuMethod(tok *tokenGenerator) (compuMethod, error) {
	cm := compuMethod{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case coeffsToken:
			cm.coeffs, err = parseCoeffs(tok)
			if err != nil {
					log.Err(err).Msg("compuMethod coeffs could not be parsed")
				break forLoop
			}
				log.Info().Msg("compuMethod coeffs successfully parsed")
		case coeffsLinearToken:
			cm.coeffsLinear, err = parseCoeffsLinear(tok)
			if err != nil {
					log.Err(err).Msg("compuMethod coeffsLinear could not be parsed")
				break forLoop
			}
				log.Info().Msg("compuMethod coeffsLinear successfully parsed")
		case compuTabRefToken:
			cm.compuTabRef, err = parseCompuTabRef(tok)
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
			cm.formula = append(cm.formula, buf)
				log.Info().Msg("compuMethod formula successfully parsed")
		case refUnitToken:
			cm.refUnit, err = parseRefUnit(tok)
			if err != nil {
					log.Err(err).Msg("compuMethod refUnit could not be parsed")
				break forLoop
			}
				log.Info().Msg("compuMethod refUnit successfully parsed")
		case statusStringRefToken:
			cm.statusStringRef, err = parseStatusStringRef(tok)
			if err != nil {
					log.Err(err).Msg("compuMethod statusStringRef could not be parsed")
				break forLoop
			}
				log.Info().Msg("compuMethod statusStringRef successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
					log.Err(err).Msg("compuMethod could not be parsed")
				break forLoop
			} else if tok.current() == endCompuMethodToken {
				break forLoop
			} else if !cm.nameSet {
				cm.name = tok.current()
				cm.nameSet = true
					log.Info().Msg("compuMethod name successfully parsed")
			} else if !cm.longIdentifierSet {
				cm.longIdentifier = tok.current()
				cm.longIdentifierSet = true
					log.Info().Msg("compuMethod longIdentifier successfully parsed")
			} else if !cm.conversionTypeSet {
				cm.conversionType, err = parseConversionTypeEnum(tok)
				if err != nil {
						log.Err(err).Msg("compuMethod conversionType could not be parsed")
					break forLoop
				}
				cm.conversionTypeSet = true
					log.Info().Msg("compuMethod conversionType successfully parsed")
			} else if !cm.formatSet {
				cm.format = tok.current()
				cm.formatSet = true
					log.Info().Msg("compuMethod format successfully parsed")
			} else if !cm.unitSet {
				cm.unit = tok.current()
				cm.unitSet = true
					log.Info().Msg("compuMethod unit successfully parsed")
			}
		}
	}
	return cm, err
}
