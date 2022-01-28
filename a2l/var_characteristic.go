package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type varCharacteristic struct {
	name                  string
	nameSet               bool
	identCriterionName    []string
	identCriterionNameSet bool
	varAddress            []varAddress
}

func parseVarCharacteristic(tok *tokenGenerator) (varCharacteristic, error) {
	vc := varCharacteristic{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginVarAddressToken:
			var buf varAddress
			buf, err = parseVarAddress(tok)
			if err != nil {
				log.Err(err).Msg("varCharacteristic varAddress could not be parsed")
				break forLoop
			}
			vc.varAddress = append(vc.varAddress, buf)
			log.Info().Msg("varCharacteristic varAddress successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("varCharacteristic could not be parsed")
				break forLoop
			} else if tok.current() == endVarCharacteristicToken {
				vc.identCriterionNameSet = true
				log.Info().Msg("varCharacteristic identCriterionName successfully parsed")
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("varCharacteristic could not be parsed")
				break forLoop
			} else if !vc.nameSet {
				vc.name = tok.current()
				vc.nameSet = true
				log.Info().Msg("varCharacteristic name successfully parsed")
			} else if !vc.identCriterionNameSet {
				vc.identCriterionName = append(vc.identCriterionName, tok.current())
			}
		}
	}
	return vc, err
}
