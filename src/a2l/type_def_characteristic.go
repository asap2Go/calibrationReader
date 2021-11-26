package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type typeDefCharacteristic struct {
	name              string
	nameSet           bool
	LongIdentifier    string
	longIdentifierSet bool
	Type              typeEnum
	TypeSet           bool
	maxDiff           float64
	maxDiffSet        bool
	conversion        string
	conversionSet     bool
	lowerLimit        float64
	lowerLimitSet     bool
	upperLimit        float64
	upperLimitSet     bool
	axisDescr         []axisDescr
	bitMask           bitMask
	byteOrder         byteOrder
	discrete          discreteKeyword
	encoding          encodingEnum
	extendedLimits    extendedLimits
	format            format
	matrixDim         matrixDim
	number            Number
	physUnit          physUnit
	stepSize          StepSize
}

func parseTypeDefCharacteristic(tok *tokenGenerator) (typeDefCharacteristic, error) {
	tdc := typeDefCharacteristic{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginAxisDescrToken:
			var buf axisDescr
			buf, err = parseAxisDescr(tok)
			if err != nil {
				log.Err(err).Msg("typeDefCharacteristic axisDescr could not be parsed")
				break forLoop
			}
			tdc.axisDescr = append(tdc.axisDescr, buf)
			log.Info().Msg("typeDefCharacteristic axisDescr successfully parsed")
		case bitMaskToken:
			tdc.bitMask, err = parseBitMask(tok)
			if err != nil {
				log.Err(err).Msg("typeDefCharacteristic bitMask could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefCharacteristic bitMask successfully parsed")
		case byteOrderToken:
			tdc.byteOrder, err = parseByteOrder(tok)
			if err != nil {
				log.Err(err).Msg("typeDefCharacteristic byteOrder could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefCharacteristic byteOrder successfully parsed")
		case discreteToken:
			tdc.discrete, err = parseDiscrete(tok)
			if err != nil {
				log.Err(err).Msg("typeDefCharacteristic discrete could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefCharacteristic discrete successfully parsed")
		case displayIdentifierToken:
			tdc.encoding, err = parseEncodingEnum(tok)
			if err != nil {
				log.Err(err).Msg("typeDefCharacteristic encoding could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefCharacteristic displayIdentifier successfully parsed")
		case extendedLimitsToken:
			tdc.extendedLimits, err = parseExtendedLimits(tok)
			if err != nil {
				log.Err(err).Msg("typeDefCharacteristic bufExtendedLimits could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefCharacteristic extendedLimits successfully parsed")
		case formatToken:
			tdc.format, err = parseFormat(tok)
			if err != nil {
				log.Err(err).Msg("typeDefCharacteristic format could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefCharacteristic format successfully parsed")
		case matrixDimToken:
			tdc.matrixDim, err = parseMatrixDim(tok)
			if err != nil {
				log.Err(err).Msg("typeDefCharacteristic matrixDim could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefCharacteristic matrixDim successfully parsed")
		case numberToken:
			tdc.number, err = parseNumber(tok)
			if err != nil {
				log.Err(err).Msg("typeDefCharacteristic number could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefCharacteristic number successfully parsed")
		case physUnitToken:
			tdc.physUnit, err = parsePhysUnit(tok)
			if err != nil {
				log.Err(err).Msg("typeDefCharacteristic physUnit could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefCharacteristic physUnit successfully parsed")
		case stepSizeToken:
			tdc.stepSize, err = parseStepSize(tok)
			if err != nil {
				log.Err(err).Msg("typeDefCharacteristic stepSize could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefCharacteristic stepSize successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("typeDefCharacteristic could not be parsed")
				break forLoop
			} else if tok.current() == endTypeDefCharacteristicToken {
				break forLoop
			} else if !tdc.nameSet {
				tdc.name = tok.current()
				tdc.nameSet = true
				log.Info().Msg("typeDefCharacteristic name successfully parsed")
			} else if !tdc.longIdentifierSet {
				tdc.LongIdentifier = tok.current()
				tdc.longIdentifierSet = true
				log.Info().Msg("typeDefCharacteristic longIdentifier successfully parsed")
			} else if !tdc.TypeSet {
				tdc.Type, err = parseTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("typeDefCharacteristic type could not be parsed")
					break forLoop
				}
				tdc.TypeSet = true
				log.Info().Msg("typeDefCharacteristic type successfully parsed")
			} else if !tdc.maxDiffSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("typeDefCharacteristic maxDiff could not be parsed")
					break forLoop
				}
				tdc.maxDiff = buf
				tdc.maxDiffSet = true
				log.Info().Msg("typeDefCharacteristic maxDiff successfully parsed")
			} else if !tdc.conversionSet {
				tdc.conversion = tok.current()
				tdc.conversionSet = true
				log.Info().Msg("typeDefCharacteristic conversion successfully parsed")
			} else if !tdc.lowerLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("typeDefCharacteristic lowerLimit could not be parsed")
					break forLoop
				}
				tdc.lowerLimit = buf
				tdc.lowerLimitSet = true
				log.Info().Msg("typeDefCharacteristic lowerLimit successfully parsed")
			} else if !tdc.upperLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("typeDefCharacteristic upperLimit could not be parsed")
					break forLoop
				}
				tdc.upperLimit = buf
				tdc.upperLimitSet = true
				log.Info().Msg("typeDefCharacteristic upperLimit successfully parsed")
			}
		}
	}
	return tdc, err
}
