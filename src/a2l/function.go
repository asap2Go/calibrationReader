package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type function struct {
	name              string
	nameSet           bool
	longIdentifier    string
	longIdentifierSet bool
	annotation        []annotation
	defCharacteristic defCharacteristic
	functionVersion   functionVersion
	ifData            []IfData
	inMeasurement     inMeasurement
	locMeasurement    locMeasurement
	outMeasurement    outMeasurement
	refCharacteristic refCharacteristic
	subFunction       subFunction
}

func parseFunction(tok *tokenGenerator) (function, error) {
	f := function{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginAnnotationToken:
			var buf annotation
			buf, err = parseAnnotation(tok)
			if err != nil {
				log.Err(err).Msg("functionVersion annotation could not be parsed")
				break forLoop
			}
			f.annotation = append(f.annotation, buf)
			log.Info().Msg("function annotation successfully parsed")
		case beginDefCharacteristicToken:
			f.defCharacteristic, err = parseDefCharacteristic(tok)
			if err != nil {
				log.Err(err).Msg("functionVersion defCharacteristic could not be parsed")
				break forLoop
			}
			log.Info().Msg("function defCharacteristic successfully parsed")
		case functionVersionToken:
			f.functionVersion, err = parseFunctionVersion(tok)
			if err != nil {
				log.Err(err).Msg("functionVersion functionVersion could not be parsed")
				break forLoop
			}
			log.Info().Msg("function functionVersion successfully parsed")
		case beginIfDataToken:
			var buf IfData
			buf, err = parseIfData(tok)
			if err != nil {
				log.Err(err).Msg("functionVersion ifData could not be parsed")
				break forLoop
			}
			f.ifData = append(f.ifData, buf)
			log.Info().Msg("function ifData successfully parsed")
		case beginInMeasurementToken:
			f.inMeasurement, err = parseInMeasurement(tok)
			if err != nil {
				log.Err(err).Msg("functionVersion inMeasurement could not be parsed")
				break forLoop
			}
			log.Info().Msg("function inMeasurement successfully parsed")
		case beginLocMeasurementToken:
			f.locMeasurement, err = parseLocMeasurement(tok)
			if err != nil {
				break forLoop
			}
			log.Info().Msg("function locMeasurement successfully parsed")
		case beginOutMeasurementToken:
			f.outMeasurement, err = parseOutMeasurement(tok)
			if err != nil {
				log.Err(err).Msg("functionVersion outMeasurement could not be parsed")
				break forLoop
			}
			log.Info().Msg("function outMeasurement successfully parsed")
		case beginRefCharacteristicToken:
			f.refCharacteristic, err = parseRefCharacteristic(tok)
			if err != nil {
				log.Err(err).Msg("functionVersion refCharacteristic could not be parsed")
				break forLoop
			}
			log.Info().Msg("function refCharacteristic successfully parsed")
		case beginSubFunctionToken:
			f.subFunction, err = parseSubFunction(tok)
			if err != nil {
				log.Err(err).Msg("functionVersion subFunction could not be parsed")
				break forLoop
			}
			log.Info().Msg("function subFunction successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("function could not be parsed")
				break forLoop
			} else if tok.current() == endFunctionToken {
				break forLoop
			} else if !f.nameSet {
				f.name = tok.current()
				f.nameSet = true
				log.Info().Msg("function name successfully parsed")
			} else if !f.longIdentifierSet {
				f.longIdentifier = tok.current()
				f.longIdentifierSet = true
				log.Info().Msg("function longIdentifier successfully parsed")
			}
		}
	}
	return f, err
}
