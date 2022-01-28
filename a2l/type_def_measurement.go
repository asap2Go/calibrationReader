package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type typeDefMeasurement struct {
	name              string
	nameSet           bool
	longIdentifier    string
	longIdentifierSet bool
	datatype          dataTypeEnum
	datatypeSet       bool
	conversion        string
	conversionSet     bool
	resolution        uint16
	resolutionSet     bool
	accuracy          float64
	accuracySet       bool
	lowerLimit        float64
	lowerLimitSet     bool
	upperLimit        float64
	upperLimitSet     bool
	bitMask           bitMask
	bitOperation      bitOperation
	byteOrder         byteOrder
	discrete          discreteKeyword
	errorMask         errorMask
	format            format
	layout            layout
	matrixDim         matrixDim
	physUnit          physUnit
}

func parseTypeDefMeasurement(tok *tokenGenerator) (typeDefMeasurement, error) {
	tdm := typeDefMeasurement{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case bitMaskToken:
			tdm.bitMask, err = parseBitMask(tok)
			if err != nil {
				log.Err(err).Msg("typeDefMeasurement bitMask could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefMeasurement bitMask successfully parsed")
		case beginBitOperationToken:
			tdm.bitOperation, err = parseBitOperation(tok)
			if err != nil {
				log.Err(err).Msg("typeDefMeasurement bitOperation could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefMeasurement bitOperation successfully parsed")
		case byteOrderToken:
			tdm.byteOrder, err = parseByteOrder(tok)
			if err != nil {
				log.Err(err).Msg("typeDefMeasurement byteOrder could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefMeasurement byteOrder successfully parsed")
		case discreteToken:
			tdm.discrete, err = parseDiscrete(tok)
			if err != nil {
				log.Err(err).Msg("typeDefMeasurement discrete could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefMeasurement discrete successfully parsed")
		case errorMaskToken:
			tdm.errorMask, err = parseErrorMask(tok)
			if err != nil {
				log.Err(err).Msg("typeDefMeasurement errorMask could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefMeasurement errorMask successfully parsed")
		case formatToken:
			tdm.format, err = parseFormat(tok)
			if err != nil {
				log.Err(err).Msg("typeDefMeasurement format could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefMeasurement format successfully parsed")
		case layoutToken:
			tdm.layout, err = parseLayout(tok)
			if err != nil {
				log.Err(err).Msg("typeDefMeasurement layout could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefMeasurement layout successfully parsed")
		case matrixDimToken:
			tdm.matrixDim, err = parseMatrixDim(tok)
			if err != nil {
				log.Err(err).Msg("typeDefMeasurement matrixDim could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefMeasurement matrixDim successfully parsed")
			log.Info().Str("current token", tok.current()).Msg("typeDefMeasurement current token:")
		case physUnitToken:
			tdm.physUnit, err = parsePhysUnit(tok)
			if err != nil {
				log.Err(err).Msg("typeDefMeasurement physUnit could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefMeasurement physUnit successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("typeDefMeasurement could not be parsed")
				break forLoop
			} else if tok.current() == endTypeDefMeasurementToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("typeDefMeasurement could not be parsed")
				break forLoop
			} else if !tdm.nameSet {
				tdm.name = tok.current()
				tdm.nameSet = true
				log.Info().Msg("typeDefMeasurement name successfully parsed")
			} else if !tdm.longIdentifierSet {
				tdm.longIdentifier = tok.current()
				tdm.longIdentifierSet = true
				log.Info().Msg("typeDefMeasurement longIdentifier successfully parsed")
			} else if !tdm.datatypeSet {
				tdm.datatype, err = parseDataTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("typeDefMeasurement datatype could not be parsed")
					break forLoop
				}
				log.Info().Msg("typeDefMeasurement datatype successfully parsed")
				tdm.datatypeSet = true
				log.Info().Msg("typeDefMeasurement datatype successfully parsed")
			} else if !tdm.conversionSet {
				tdm.conversion = tok.current()
			} else if !tdm.resolutionSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 16)
				if err != nil {
					log.Err(err).Msg("typeDefMeasurement resolution could not be parsed")
					break forLoop
				}
				tdm.resolution = uint16(buf)
				tdm.resolutionSet = true
				log.Info().Msg("typeDefMeasurement resolution successfully parsed")
			} else if !tdm.accuracySet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("typeDefMeasurement accuracy could not be parsed")
					break forLoop
				}
				tdm.accuracy = buf
				tdm.accuracySet = true
				log.Info().Msg("typeDefMeasurement accuracy successfully parsed")
			} else if !tdm.lowerLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("typeDefMeasurement lowerLimit could not be parsed")
					break forLoop
				}
				tdm.lowerLimit = buf
				tdm.lowerLimitSet = true
				log.Info().Msg("typeDefMeasurement lowerLimit successfully parsed")
			} else if !tdm.upperLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("typeDefMeasurement upperLimit could not be parsed")
					break forLoop
				}
				tdm.upperLimit = buf
				tdm.upperLimitSet = true
				log.Info().Msg("typeDefMeasurement upperLimit successfully parsed")
			}
		}
	}
	return tdm, err
}
