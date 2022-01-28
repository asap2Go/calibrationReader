package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type siExponents struct {
	length               int16
	lengthSet            bool
	mass                 int16
	massSet              bool
	time                 int16
	timeSet              bool
	electricCurrent      int16
	electricCurrentSet   bool
	temperature          int16
	temperatureSet       bool
	amountOfSubstance    int16
	amountOfSubstanceSet bool
	luminousIntensity    int16
	luminousIntensitySet bool
}

func parseSiExponents(tok *tokenGenerator) (siExponents, error) {
	se := siExponents{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("siExponents could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("siExponents could not be parsed")
			break forLoop
		} else if !se.lengthSet {
			var buf int64
			buf, err = strconv.ParseInt(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("siExponents length could not be parsed")
				break forLoop
			}
			se.length = int16(buf)
			se.lengthSet = true
			log.Info().Msg("siExponents length successfully parsed")
		} else if !se.massSet {
			var buf int64
			buf, err = strconv.ParseInt(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("siExponents mass could not be parsed")
				break forLoop
			}
			se.mass = int16(buf)
			se.massSet = true
			log.Info().Msg("siExponents mass successfully parsed")
		} else if !se.timeSet {
			var buf int64
			buf, err = strconv.ParseInt(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("siExponents time could not be parsed")
				break forLoop
			}
			se.time = int16(buf)
			se.timeSet = true
			log.Info().Msg("siExponents time successfully parsed")
		} else if !se.electricCurrentSet {
			var buf int64
			buf, err = strconv.ParseInt(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("siExponents electricCurrent could not be parsed")
				break forLoop
			}
			se.electricCurrent = int16(buf)
			se.electricCurrentSet = true
			log.Info().Msg("siExponents electricCurrent successfully parsed")
		} else if !se.temperatureSet {
			var buf int64
			buf, err = strconv.ParseInt(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("siExponents temperature could not be parsed")
				break forLoop
			}
			se.temperature = int16(buf)
			se.temperatureSet = true
			log.Info().Msg("siExponents temperature successfully parsed")
		} else if !se.amountOfSubstanceSet {
			var buf int64
			buf, err = strconv.ParseInt(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("siExponents amountOfSubstance could not be parsed")
				break forLoop
			}
			se.amountOfSubstance = int16(buf)
			se.amountOfSubstanceSet = true
			log.Info().Msg("siExponents amountOfSubstance successfully parsed")
		} else if !se.luminousIntensitySet {
			var buf int64
			buf, err = strconv.ParseInt(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("siExponents luminousIntensity could not be parsed")
				break forLoop
			}
			se.luminousIntensity = int16(buf)
			se.luminousIntensitySet = true
			log.Info().Msg("siExponents luminousIntensity successfully parsed")
			break forLoop
		}
	}
	return se, err
}
