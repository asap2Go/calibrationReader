package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type characteristic struct {
	Name                    string
	nameSet                 bool
	LongIdentifier          string
	longIdentifierSet       bool
	Type                    typeEnum
	TypeSet                 bool
	address                 string
	addressSet              bool
	deposit                 string
	depositSet              bool
	maxDiff                 float64
	maxDiffSet              bool
	conversion              string
	conversionSet           bool
	lowerLimit              float64
	lowerLimitSet           bool
	upperLimit              float64
	upperLimitSet           bool
	annotation              []annotation
	axisDescr               []axisDescr
	bitMask                 bitMask
	byteOrder               byteOrder
	calibrationAccess       calibrationAccessEnum
	comparisonQuantity      comparisonQuantity
	dependentCharacteristic []DependentCharacteristic
	discrete                discreteKeyword
	displayIdentifier       DisplayIdentifier
	ecuAddressExtension     ecuAddressExtension
	extendedLimits          extendedLimits
	format                  format
	functionList            []FunctionList
	guardRails              GuardRails
	ifData                  []IfData
	mapList                 []MapList
	matrixDim               matrixDim
	maxRefresh              MaxRefresh
	number                  Number
	physUnit                physUnit
	readOnly                readOnlyKeyword
	refMemorySegment        refMemorySegment
	stepSize                StepSize
	symbolLink              symbolLink
	virtualCharacteristic   []VirtualCharacteristic
}

func parseCharacteristic(tok *tokenGenerator) (characteristic, error) {
	c := characteristic{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginAnnotationToken:
			var buf annotation
			buf, err = parseAnnotation(tok)
			if err != nil {
				log.Err(err).Msg("characteristic annotation could not be parsed")
				break forLoop
			}
			c.annotation = append(c.annotation, buf)
			log.Info().Msg("characteristic annotation successfully parsed")
		case beginAxisDescrToken:
			var buf axisDescr
			buf, err = parseAxisDescr(tok)
			if err != nil {
				log.Err(err).Msg("characteristic axisDescr could not be parsed")
				break forLoop
			}
			c.axisDescr = append(c.axisDescr, buf)
			log.Info().Msg("characteristic axisDescr successfully parsed")
		case bitMaskToken:
			c.bitMask, err = parseBitMask(tok)
			if err != nil {
				log.Err(err).Msg("characteristic bitMask could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic bitMask successfully parsed")
		case byteOrderToken:
			c.byteOrder, err = parseByteOrder(tok)
			if err != nil {
				log.Err(err).Msg("characteristic byteOrder could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic byteOrder successfully parsed")
		case calibrationAccessToken:
			c.calibrationAccess, err = parseCalibrationAccessEnum(tok)
			if err != nil {
				log.Err(err).Msg("characteristic calibrationAccess could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic calibrationAccess successfully parsed")
		case comparisonQuantityToken:
			c.comparisonQuantity, err = parseComparisonQuantity(tok)
			if err != nil {
				log.Err(err).Msg("characteristic comparisonQuantity could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic comparisonQuantity successfully parsed")
		case beginDependentCharacteristicToken:
			var buf DependentCharacteristic
			buf, err = parseDependentCharacteristic(tok)
			if err != nil {
				log.Err(err).Msg("characteristic dependentCharacteristic could not be parsed")
				break forLoop
			}
			c.dependentCharacteristic = append(c.dependentCharacteristic, buf)
			log.Info().Msg("characteristic dependentCharacteristic successfully parsed")
		case discreteToken:
			c.discrete, err = parseDiscrete(tok)
			if err != nil {
				log.Err(err).Msg("characteristic discrete could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic discrete successfully parsed")
		case displayIdentifierToken:
			c.displayIdentifier, err = parseDisplayIdentifier(tok)
			if err != nil {
				log.Err(err).Msg("characteristic displayIdentifier could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic displayIdentifier successfully parsed")
		case ecuAddressExtensionToken:
			c.ecuAddressExtension, err = parseECUAddressExtension(tok)
			if err != nil {
				log.Err(err).Msg("characteristic ecuAddressExtension could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic ecuAddressExtension successfully parsed")
		case extendedLimitsToken:
			c.extendedLimits, err = parseExtendedLimits(tok)
			if err != nil {
				log.Err(err).Msg("characteristic bufExtendedLimits could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic extendedLimits successfully parsed")
		case formatToken:
			c.format, err = parseFormat(tok)
			if err != nil {
				log.Err(err).Msg("characteristic format could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic format successfully parsed")
		case beginFunctionListToken:
			var buf FunctionList
			buf, err = parseFunctionList(tok)
			if err != nil {
				log.Err(err).Msg("characteristic functionList could not be parsed")
				break forLoop
			}
			c.functionList = append(c.functionList, buf)
			log.Info().Msg("characteristic functionList successfully parsed")
		case guardRailsToken:
			c.guardRails, err = parseGuardRails(tok)
			if err != nil {
				log.Err(err).Msg("characteristic guardRails could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic guardRails successfully parsed")
		case beginIfDataToken:
			var buf IfData
			buf, err = parseIfData(tok)
			if err != nil {
				log.Err(err).Msg("characteristic ifData could not be parsed")
				break forLoop
			}
			c.ifData = append(c.ifData, buf)
			log.Info().Msg("characteristic ifData successfully parsed")
		case beginMapListToken:
			var buf MapList
			buf, err = parseMapList(tok)
			if err != nil {
				log.Err(err).Msg("characteristic mapList could not be parsed")
				break forLoop
			}
			c.mapList = append(c.mapList, buf)
			log.Info().Msg("characteristic mapList successfully parsed")
		case matrixDimToken:
			c.matrixDim, err = parseMatrixDim(tok)
			if err != nil {
				log.Err(err).Msg("characteristic matrixDim could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic matrixDim successfully parsed")
		case maxRefreshToken:
			c.maxRefresh, err = parseMaxRefresh(tok)
			if err != nil {
				log.Err(err).Msg("characteristic maxRefresh could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic maxRefresh successfully parsed")
		case numberToken:
			c.number, err = parseNumber(tok)
			if err != nil {
				log.Err(err).Msg("characteristic number could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic number successfully parsed")
		case physUnitToken:
			c.physUnit, err = parsePhysUnit(tok)
			if err != nil {
				log.Err(err).Msg("characteristic physUnit could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic physUnit successfully parsed")
		case readOnlyToken:
			c.readOnly, err = parseReadOnly(tok)
			if err != nil {
				log.Err(err).Msg("characteristic readOnly could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic readOnly successfully parsed")
		case refMemorySegmentToken:
			c.refMemorySegment, err = parseRefMemorySegment(tok)
			if err != nil {
				log.Err(err).Msg("characteristic refMemorySegment could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic refMemorySegment successfully parsed")
		case stepSizeToken:
			c.stepSize, err = parseStepSize(tok)
			if err != nil {
				log.Err(err).Msg("characteristic stepSize could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic stepSize successfully parsed")
		case symbolLinkToken:
			c.symbolLink, err = parseSymbolLink(tok)
			if err != nil {
				log.Err(err).Msg("characteristic symbolLink could not be parsed")
				break forLoop
			}
			log.Info().Msg("characteristic symbolLink successfully parsed")
		case beginVirtualCharacteristicToken:
			var buf VirtualCharacteristic
			buf, err = parseVirtualCharacteristic(tok)
			if err != nil {
				log.Err(err).Msg("characteristic virtualCharacteristic could not be parsed")
				break forLoop
			}
			c.virtualCharacteristic = append(c.virtualCharacteristic, buf)
			log.Info().Msg("characteristic virtualCharacteristic successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("characteristic could not be parsed")
				break forLoop
			} else if tok.current() == endCharacteristicToken {
				break forLoop
			} else if !c.nameSet {
				c.Name = tok.current()
				c.nameSet = true
				log.Info().Msg("characteristic name successfully parsed")
			} else if !c.longIdentifierSet {
				c.LongIdentifier = tok.current()
				c.longIdentifierSet = true
				log.Info().Msg("characteristic longIdentifier successfully parsed")
			} else if !c.TypeSet {
				c.Type, err = parseTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("characteristic type could not be parsed")
					break forLoop
				}
				c.TypeSet = true
				log.Info().Msg("characteristic type successfully parsed")
			} else if !c.addressSet {
				c.address = tok.current()
				c.addressSet = true
				log.Info().Msg("characteristic address successfully parsed")
			} else if !c.depositSet {
				c.deposit = tok.current()
				c.depositSet = true
				log.Info().Msg("characteristic deposit successfully parsed")
			} else if !c.maxDiffSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("characteristic maxDiff could not be parsed")
					break forLoop
				}
				c.maxDiff = buf
				c.maxDiffSet = true
				log.Info().Msg("characteristic maxDiff successfully parsed")
			} else if !c.conversionSet {
				c.conversion = tok.current()
				c.conversionSet = true
				log.Info().Msg("characteristic conversion successfully parsed")
			} else if !c.lowerLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("characteristic lowerLimit could not be parsed")
					break forLoop
				}
				c.lowerLimit = buf
				c.lowerLimitSet = true
				log.Info().Msg("characteristic lowerLimit successfully parsed")
			} else if !c.upperLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("characteristic upperLimit could not be parsed")
					break forLoop
				}
				c.upperLimit = buf
				c.upperLimitSet = true
				log.Info().Msg("characteristic upperLimit successfully parsed")
			}
		}
	}
	return c, err
}
