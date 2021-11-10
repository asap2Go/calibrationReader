package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type unit struct {
	name              string
	nameSet           bool
	longIdentifier    string
	longIdentifierSet bool
	display           string
	displaySet        bool
	Type              typeEnum
	TypeSet           bool
	refUnit           refUnit
	siExponents       siExponents
	unitConversion    unitConversion
}

func parseUnit(tok *tokenGenerator) (unit, error) {
	u := unit{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case refUnitToken:
			var buf refUnit
			buf, err = parseRefUnit(tok)
			if err != nil {
				log.Err(err).Msg("unit refUnit could not be parsed")
				break forLoop
			}
			u.refUnit = buf
			log.Info().Msg("unit refUnit successfully parsed")
		case siExponentsToken:
			var buf siExponents
			buf, err = parseSiExponents(tok)
			if err != nil {
				log.Err(err).Msg("unit siExponents could not be parsed")
				break forLoop
			}
			u.siExponents = buf
			log.Info().Msg("unit siExponents successfully parsed")
		case unitConversionToken:
			var buf unitConversion
			buf, err = parseUnitConversion(tok)
			if err != nil {
				log.Err(err).Msg("unit unitConversion could not be parsed")
				break forLoop
			}
			u.unitConversion = buf
			log.Info().Msg("unit unitConversion successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("unit could not be parsed")
				break forLoop
			} else if tok.current() == endUnitToken {
				break forLoop
			} else if !u.nameSet {
				u.name = tok.current()
				u.nameSet = true
				log.Info().Msg("unit name successfully parsed")
			} else if !u.longIdentifierSet {
				u.longIdentifier = tok.current()
				u.nameSet = true
				log.Info().Msg("unit name successfully parsed")
			} else if !u.displaySet {
				u.display = tok.current()
				u.displaySet = true
				log.Info().Msg("unit display successfully parsed")
			} else if !u.TypeSet {
				var buf typeEnum
				buf, err = parseTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("unit type could not be parsed")
					break forLoop
				}
				u.Type = buf
				u.TypeSet = true
				log.Info().Msg("unit TYPE successfully parsed")
			}
		}
	}
	return u, err
}
