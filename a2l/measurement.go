package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type measurement struct {
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
	modelLink           modelLink
	physUnit            physUnit
	readWrite           readWriteKeyword
	refMemorySegment    refMemorySegment
	symbolLink          symbolLink
	virtual             virtual
}

func parseMeasurement(tok *tokenGenerator) (measurement, error) {
	m := measurement{}
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
			m.arraySize, err = parseArraySize(tok)
			if err != nil {
				log.Err(err).Msg("measurement arraySize could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement arraySize successfully parsed")
		case bitMaskToken:
			m.bitMask, err = parseBitMask(tok)
			if err != nil {
				log.Err(err).Msg("measurement bitMask could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement bitMask successfully parsed")
		case beginBitOperationToken:
			m.bitOperation, err = parseBitOperation(tok)
			if err != nil {
				log.Err(err).Msg("measurement bitOperation could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement bitOperation successfully parsed")
		case byteOrderToken:
			m.byteOrder, err = parseByteOrder(tok)
			if err != nil {
				log.Err(err).Msg("measurement byteOrder could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement byteOrder successfully parsed")
		case discreteToken:
			m.discrete, err = parseDiscrete(tok)
			if err != nil {
				log.Err(err).Msg("measurement discrete could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement discrete successfully parsed")
		case displayIdentifierToken:
			m.displayIdentifier, err = parseDisplayIdentifier(tok)
			if err != nil {
				log.Err(err).Msg("measurement displayIdentifier could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement displayIdentifier successfully parsed")
		case ecuAddressToken:
			m.ecuAddress, err = parseEcuAddress(tok)
			if err != nil {
				log.Err(err).Msg("measurement ecuAddress could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement ecuAddress successfully parsed")
		case ecuAddressExtensionToken:
			m.ecuAddressExtension, err = parseECUAddressExtension(tok)
			if err != nil {
				log.Err(err).Msg("measurement ecuAddressExtension could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement ecuAddressExtension successfully parsed")
		case errorMaskToken:
			m.errorMask, err = parseErrorMask(tok)
			if err != nil {
				log.Err(err).Msg("measurement errorMask could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement errorMask successfully parsed")
		case formatToken:
			m.format, err = parseFormat(tok)
			if err != nil {
				log.Err(err).Msg("measurement format could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement format successfully parsed")
		case beginFunctionListToken:
			m.functionList, err = parseFunctionList(tok)
			if err != nil {
				log.Err(err).Msg("measurement functionList could not be parsed")
				break forLoop
			}
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
			m.layout, err = parseLayout(tok)
			if err != nil {
				log.Err(err).Msg("measurement layout could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement layout successfully parsed")
		case matrixDimToken:
			m.matrixDim, err = parseMatrixDim(tok)
			if err != nil {
				log.Err(err).Msg("measurement matrixDim could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement matrixDim successfully parsed")
			log.Info().Str("current token", tok.current()).Msg("measurement current token:")
		case maxRefreshToken:
			m.maxRefresh, err = parseMaxRefresh(tok)
			if err != nil {
				log.Err(err).Msg("measurement maxRefresh could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement maxRefresh successfully parsed")
		case modelLinkToken:
			m.modelLink, err = parseModelLink(tok)
			if err != nil {
				log.Err(err).Msg("measurement modelLink could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement modelLink successfully parsed")
		case physUnitToken:
			m.physUnit, err = parsePhysUnit(tok)
			if err != nil {
				log.Err(err).Msg("measurement physUnit could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement physUnit successfully parsed")
		case readWriteToken:
			m.readWrite, err = parseReadWrite(tok)
			if err != nil {
				log.Err(err).Msg("measurement readWrite could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement readWrite successfully parsed")
		case refMemorySegmentToken:
			m.refMemorySegment, err = parseRefMemorySegment(tok)
			if err != nil {
				log.Err(err).Msg("measurement refMemorySegment could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement refMemorySegment successfully parsed")
		case symbolLinkToken:
			m.symbolLink, err = parseSymbolLink(tok)
			if err != nil {
				log.Err(err).Msg("measurement symbolLink could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement symbolLink successfully parsed")
		case beginVirtualToken:
			m.virtual, err = parseVirtual(tok)
			if err != nil {
				log.Err(err).Msg("measurement virtual could not be parsed")
				break forLoop
			}
			log.Info().Msg("measurement virtual successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("measurement could not be parsed")
				break forLoop
			} else if tok.current() == endMeasurementToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("measurement could not be parsed")
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
				m.datatype, err = parseDataTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("measurement datatype could not be parsed")
					break forLoop
				}
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
