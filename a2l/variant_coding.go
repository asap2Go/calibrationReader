package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type variantCoding struct {
	varCharacteristic []varCharacteristic
	varCriterion      []varCriterion
	varForbiddenComb  []varForbiddenComb
	varNaming         varNaming
	varSeparator      varSeparator
	varSeparatorSet   bool
}

func parseVariantCoding(tok *tokenGenerator) (variantCoding, error) {
	vc := variantCoding{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginVarCharacteristicToken:
			var buf varCharacteristic
			buf, err = parseVarCharacteristic(tok)
			if err != nil {
				log.Err(err).Msg("variantCoding varCharacteristic could not be parsed")
				break forLoop
			}
			vc.varCharacteristic = append(vc.varCharacteristic, buf)
			log.Info().Msg("variantCoding varCharacteristic successfully parsed")
		case beginVarCriterionToken:
			var buf varCriterion
			buf, err = parseVarCriterion(tok)
			if err != nil {
				log.Err(err).Msg("variantCoding varCriterion could not be parsed")
				break forLoop
			}
			vc.varCriterion = append(vc.varCriterion, buf)
			log.Info().Msg("variantCoding varCriterion successfully parsed")
		case beginVarForbiddenCombToken:
			var buf varForbiddenComb
			buf, err = parseVarForbiddenComb(tok)
			if err != nil {
				log.Err(err).Msg("variantCoding varForbiddenComb could not be parsed")
				break forLoop
			}
			vc.varForbiddenComb = append(vc.varForbiddenComb, buf)
			log.Info().Msg("variantCoding varForbiddenComb successfully parsed")
		case varNamingToken:
			vc.varNaming, err = parseVarNaming(tok)
			if err != nil {
				log.Err(err).Msg("variantCoding varNaming could not be parsed")
				break forLoop
			}
			log.Info().Msg("variantCoding varNaming successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("variantCoding could not be parsed")
				break forLoop
			} else if tok.current() == endVariantCodingToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("variantCoding could not be parsed")
				break forLoop
			} else if !vc.varSeparatorSet {
				vc.varSeparator, err = parseVarSeparator(tok)
				if err != nil {
					log.Err(err).Msg("variantCoding varSeparator could not be parsed")
					break forLoop
				}
				vc.varSeparatorSet = true
				log.Info().Msg("variantCoding varSeparator successfully parsed")
			}
		}
	}
	return vc, err
}
