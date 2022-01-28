package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type varCriterion struct {
	name                       string
	nameSet                    bool
	longIdentifier             string
	longIdentifierSet          bool
	value                      []string
	valueSet                   bool
	varMeasurement             varMeasurement
	varSelectionCharacteristic varSelectionCharacteristic
}

func parseVarCriterion(tok *tokenGenerator) (varCriterion, error) {
	vc := varCriterion{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case varMeasurementToken:
			vc.varMeasurement, err = parseVarMeasurement(tok)
			if err != nil {
				log.Err(err).Msg("varCriterion varMeasurement could not be parsed")
				break forLoop
			}
			log.Info().Msg("varCriterion varMeasurement successfully parsed")
		case varSelectionCharacteristicToken:
			vc.varSelectionCharacteristic, err = parseVarSelectionCharacteristic(tok)
			if err != nil {
				log.Err(err).Msg("varCriterion varSelectionCharacteristic could not be parsed")
				break forLoop
			}
			log.Info().Msg("varCriterion varSelectionCharacteristic successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("varCriterion could not be parsed")
				break forLoop
			} else if tok.current() == endVarCriterionToken {
				vc.valueSet = true
				log.Info().Msg("varCriterion value successfully parsed")
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("varCriterion could not be parsed")
				break forLoop
			} else if !vc.nameSet {
				vc.name = tok.current()
				vc.nameSet = true
				log.Info().Msg("varCriterion name successfully parsed")
			} else if !vc.longIdentifierSet {
				vc.longIdentifier = tok.current()
				vc.longIdentifierSet = true
				log.Info().Msg("varCriterion longIdentifier successfully parsed")
			} else if !vc.valueSet {
				vc.value = append(vc.value, tok.current())
			}
		}
	}
	return vc, err
}
