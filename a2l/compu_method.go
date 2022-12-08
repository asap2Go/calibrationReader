package a2l

import (
	"errors"
	"math"

	"github.com/rs/zerolog/log"
)

type compuMethod struct {
	name              string
	nameSet           bool
	longIdentifier    string
	longIdentifierSet bool
	conversionType    conversionTypeEnum
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

func (cm *compuMethod) convDecToPhy(dec float64) (float64, error) {
	var err error
	switch cm.conversionType {
	case Identical:
		return dec, nil
	case Form:
		//implement later.
	case Linear:
		if !(cm.coeffsLinear.aSet && cm.coeffsLinear.bSet) {
			err = errors.New("coeffsLinear not set in compuMethod: " + cm.name)
			log.Err(err).Msg("decimal value could not be converted")
			return 0, err
		}
		return cm.coeffsLinear.a*dec + cm.coeffsLinear.b, err
	case RatFunc:
		if !(cm.coeffs.aSet && cm.coeffs.bSet && cm.coeffs.cSet && cm.coeffs.dSet && cm.coeffs.eSet && cm.coeffs.fSet) {
			err = errors.New("coeffs not set in compuMethod: " + cm.name)
			log.Err(err).Msg("decimal value could not be converted")
			return 0, err
		}
		phy, err := cm.calcRatFunc(dec)
		if err != nil {
			log.Err(err).Msg("decimal value could not be converted")
			return phy, err
		}
	case TabIntp:
	case TabNointp:
	case TabVerb:
	default:
		err = errors.New("conversion Type undefined in compuMethod: " + cm.name)
		log.Err(err).Msg("decimal value could not be converted")
		return 0, err
	}
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
				err = errors.New("unedecpected end of file")
				log.Err(err).Msg("compuMethod could not be parsed")
				break forLoop
			} else if tok.current() == endCompuMethodToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unedecpected token " + tok.current())
				log.Err(err).Msg("compuMethod could not be parsed")
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

func (cm *compuMethod) calcRatFunc(dec float64) (float64, error) {
	//following formula defines f(Physical) = Decimal
	//y = (axx + bx + c) / (dxx + ex + f)
	//inverted fi(Decimal) = Physical
	//y = (e dec - b)/(2 (a - d dec)) ± sqrt((e dec - b)^2 - 4 (d dec - a) (f dec - c))/(2 (a - d dec))
	firstDivisor := (2 * (cm.coeffs.a - cm.coeffs.d*dec))
	if firstDivisor == 0 {
		err = errors.New("rationality function cannot be computed(zero divisor) for compuMethod: " + cm.name)
		log.Err(err).Msg("decimal value could not be converted")
		return 0, err
	}
	secondDivisorPositive := (2 * (cm.coeffs.a - cm.coeffs.d*dec)) +
		math.Sqrt(math.Pow((cm.coeffs.e*dec-cm.coeffs.b), 2)-4*(cm.coeffs.d*dec-cm.coeffs.a)*(cm.coeffs.f*dec-cm.coeffs.c))/firstDivisor
	secondDivisorNegative := (2 * (cm.coeffs.a - cm.coeffs.d*dec)) -
		math.Sqrt(math.Pow((cm.coeffs.e*dec-cm.coeffs.b), 2)-4*(cm.coeffs.d*dec-cm.coeffs.a)*(cm.coeffs.f*dec-cm.coeffs.c))/firstDivisor

	if secondDivisorPositive != 0 && secondDivisorNegative != 0 {
		plusVal := (cm.coeffs.e*dec - cm.coeffs.b) / secondDivisorPositive
		minusVal := (cm.coeffs.e*dec - cm.coeffs.b) / secondDivisorNegative
		testVal := (cm.coeffs.a*plusVal*plusVal + cm.coeffs.b*plusVal + cm.coeffs.c) / (cm.coeffs.d*plusVal*plusVal + cm.coeffs.e*plusVal + cm.coeffs.f)
		if testVal == dec {
			return plusVal, err
		} else {
			return minusVal, err
		}
	} else if secondDivisorPositive != 0 {
		plusVal := (cm.coeffs.e*dec - cm.coeffs.b) / secondDivisorPositive
		return plusVal, err
	} else if secondDivisorNegative != 0 {
		minusVal := (cm.coeffs.e*dec - cm.coeffs.b) / secondDivisorNegative
		return minusVal, err
	} else {
		err = errors.New("rationality function cannot be computed(zero divisor) for compuMethod: " + cm.name)
		log.Err(err).Msg("decimal value could not be converted")
		return 0, err
	}

}
