package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type MEASUREMENT struct {
	name                string
	nameSet             bool
	longIdentifier      string
	longIdentifierSet   bool
	datatype            dataTypeEnum
	datatypeSet         bool
	conversion          string
	conversionSet       bool
	resolution          uint16
	resolutionSet       bool
	accuracy            float64
	accuracySet         bool
	lowerLimit          float64
	lowerLimitSet       bool
	upperLimit          float64
	upperLimitSet       bool
	annotation          []annotation
	arraySize           arraySize
	bitMask             bitMask
	bitOperation        bitOperation
	byteOrder           byteOrder
	discrete            discreteKeyword
	displayIdentifier   DisplayIdentifier
	ecuAddress          ecuAddress
	ecuAddressExtension ecuAddressExtension
	errorMask           errorMask
	format              format
	functionList        FunctionList
	ifData              []IfData
	layout              layout
	matrixDim           matrixDim
	maxRefresh          MaxRefresh
	physUnit            physUnit
	readWrite           readWriteKeyword
	refMemorySegment    refMemorySegment
	symbolLink          symbolLink
	virtual             virtual
}

func parseMeasurement(tok *tokenGenerator) (MEASUREMENT, error) {
	m := MEASUREMENT{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginAnnotationToken:
			var buf annotation
			buf, err = parseAnnotation(tok)
			if err != nil {
				log.Err(err).Msg("measurement annotation could not be parsed")
				break forLoop
			}
			m.annotation = append(m.annotation, buf)
			log.Info().Msg("measurement annotation successfully parsed")
		case arraySizeToken:
			var buf arraySize
			buf, err = parseArraySize(tok)
			if err != nil {
				log.Err(err).Msg("measurement arraySize could not be parsed")
				break forLoop
			}
			m.arraySize = buf
			log.Info().Msg("measurement arraySize successfully parsed")
		case bitMaskToken:
			var buf bitMask
			buf, err = parseBitMask(tok)
			if err != nil {
				log.Err(err).Msg("measurement bitMask could not be parsed")
				break forLoop
			}
			m.bitMask = buf
			log.Info().Msg("measurement bitMask successfully parsed")
		case beginBitOperationToken:
			var buf bitOperation
			buf, err = parseBitOperation(tok)
			if err != nil {
				log.Err(err).Msg("measurement bitOperation could not be parsed")
				break forLoop
			}
			m.bitOperation = buf
			log.Info().Msg("measurement bitOperation successfully parsed")
		case byteOrderToken:
			var buf byteOrder
			buf, err = parseByteOrder(tok)
			if err != nil {
				log.Err(err).Msg("measurement byteOrder could not be parsed")
				break forLoop
			}
			m.byteOrder = buf
			log.Info().Msg("measurement byteOrder successfully parsed")
		case discreteToken:
			var buf discreteKeyword
			buf, err = parseDiscrete(tok)
			if err != nil {
				log.Err(err).Msg("measurement discrete could not be parsed")
				break forLoop
			}
			m.discrete = buf
			log.Info().Msg("measurement discrete successfully parsed")
		case displayIdentifierToken:
			var buf DisplayIdentifier
			buf, err = parseDisplayIdentifier(tok)
			if err != nil {
				log.Err(err).Msg("measurement displayIdentifier could not be parsed")
				break forLoop
			}
			m.displayIdentifier = buf
			log.Info().Msg("measurement displayIdentifier successfully parsed")
		case ecuAddressToken:
			var buf ecuAddress
			buf, err = parseEcuAddress(tok)
			if err != nil {
				log.Err(err).Msg("measurement ecuAddress could not be parsed")
				break forLoop
			}
			m.ecuAddress = buf
			log.Info().Msg("measurement ecuAddress successfully parsed")
		case ecuAddressExtensionToken:
			var buf ecuAddressExtension
			buf, err = parseECUAddressExtension(tok)
			if err != nil {
				log.Err(err).Msg("measurement ecuAddressExtension could not be parsed")
				break forLoop
			}
			m.ecuAddressExtension = buf
			log.Info().Msg("measurement ecuAddressExtension successfully parsed")
		case errorMaskToken:
			var buf errorMask
			buf, err = parseErrorMask(tok)
			if err != nil {
				log.Err(err).Msg("measurement errorMask could not be parsed")
				break forLoop
			}
			m.errorMask = buf
			log.Info().Msg("measurement errorMask successfully parsed")
		case formatToken:
			var buf format
			buf, err = parseFormat(tok)
			if err != nil {
				log.Err(err).Msg("measurement format could not be parsed")
				break forLoop
			}
			m.format = buf
			log.Info().Msg("measurement format successfully parsed")
		case beginFunctionListToken:
			var buf FunctionList
			buf, err = parseFunctionList(tok)
			if err != nil {
				log.Err(err).Msg("measurement functionList could not be parsed")
				break forLoop
			}
			m.functionList = buf
			log.Info().Msg("measurement functionList successfully parsed")
		case beginIfDataToken:
			var buf IfData
			buf, err = parseIfData(tok)
			if err != nil {
				log.Err(err).Msg("measurement ifData could not be parsed")
				break forLoop
			}
			m.ifData = append(m.ifData, buf)
			log.Info().Msg("measurement ifData successfully parsed")
		case layoutToken:
			var buf layout
			buf, err = parseLayout(tok)
			if err != nil {
				log.Err(err).Msg("measurement layout could not be parsed")
				break forLoop
			}
			m.layout = buf
			log.Info().Msg("measurement layout successfully parsed")
		case matrixDimToken:
			var buf matrixDim
			buf, err = parseMatrixDim(tok)
			if err != nil {
				log.Err(err).Msg("measurement matrixDim could not be parsed")
				break forLoop
			}
			m.matrixDim = buf
			log.Info().Msg("measurement matrixDim successfully parsed")
			log.Info().Str("current token", tok.current()).Msg("measurement current token:")
		case maxRefreshToken:
			var buf MaxRefresh
			buf, err = parseMaxRefresh(tok)
			if err != nil {
				log.Err(err).Msg("measurement maxRefresh could not be parsed")
				break forLoop
			}
			m.maxRefresh = buf
			log.Info().Msg("measurement maxRefresh successfully parsed")
		case physUnitToken:
			var buf physUnit
			buf, err = parsePhysUnit(tok)
			if err != nil {
				log.Err(err).Msg("measurement physUnit could not be parsed")
				break forLoop
			}
			m.physUnit = buf
			log.Info().Msg("measurement physUnit successfully parsed")
		case readWriteToken:
			var buf readWriteKeyword
			buf, err = parseReadWrite(tok)
			if err != nil {
				log.Err(err).Msg("measurement readWrite could not be parsed")
				break forLoop
			}
			m.readWrite = buf
			log.Info().Msg("measurement readWrite successfully parsed")
		case refMemorySegmentToken:
			var buf refMemorySegment
			buf, err = parseRefMemorySegment(tok)
			if err != nil {
				log.Err(err).Msg("measurement refMemorySegment could not be parsed")
				break forLoop
			}
			m.refMemorySegment = buf
			log.Info().Msg("measurement refMemorySegment successfully parsed")
		case symbolLinkToken:
			var buf symbolLink
			buf, err = parseSymbolLink(tok)
			if err != nil {
				log.Err(err).Msg("measurement symbolLink could not be parsed")
				break forLoop
			}
			m.symbolLink = buf
			log.Info().Msg("measurement symbolLink successfully parsed")
		case beginVirtualToken:
			var buf virtual
			buf, err = parseVirtual(tok)
			if err != nil {
				log.Err(err).Msg("measurement virtual could not be parsed")
				break forLoop
			}
			m.virtual = buf
			log.Info().Msg("measurement virtual successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("measurement could not be parsed")
				break forLoop
			} else if tok.current() == endMeasurementToken {
				break forLoop
			} else if !m.nameSet {
				m.name = tok.current()
				m.nameSet = true
				log.Info().Msg("measurement name successfully parsed")
			} else if !m.longIdentifierSet {
				m.longIdentifier = tok.current()
				m.longIdentifierSet = true
				log.Info().Msg("measurement longIdentifier successfully parsed")
			} else if !m.datatypeSet {
				var buf dataTypeEnum
				buf, err = parseDataTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("measurement datatype could not be parsed")
					break forLoop
				}
				m.datatype = buf
				log.Info().Msg("measurement datatype successfully parsed")
				m.datatypeSet = true
				log.Info().Msg("measurement datatype successfully parsed")
			} else if !m.conversionSet {
				m.conversion = tok.current()
			} else if !m.resolutionSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 16)
				if err != nil {
					log.Err(err).Msg("measurement resolution could not be parsed")
					break forLoop
				}
				m.resolution = uint16(buf)
				m.resolutionSet = true
				log.Info().Msg("measurement resolution successfully parsed")
			} else if !m.accuracySet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("measurement accuracy could not be parsed")
					break forLoop
				}
				m.accuracy = buf
				m.accuracySet = true
				log.Info().Msg("measurement accuracy successfully parsed")
			} else if !m.lowerLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("measurement lowerLimit could not be parsed")
					break forLoop
				}
				m.lowerLimit = buf
				m.lowerLimitSet = true
				log.Info().Msg("measurement lowerLimit successfully parsed")
			} else if !m.upperLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("measurement upperLimit could not be parsed")
					break forLoop
				}
				m.upperLimit = buf
				m.upperLimitSet = true
				log.Info().Msg("measurement upperLimit successfully parsed")
			}
		}
	}
	return m, err
}
