package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type typeDefAxis struct {
	name              string
	nameSet           bool
	LongIdentifier    string
	longIdentifierSet bool
	inputQuantity     string
	inputQuantitySet  bool
	recordLayout      string
	recordLayoutSet   bool
	maxDiff           float64
	maxDiffSet        bool
	conversion        string
	conversionSet     bool
	maxAxisPoints     uint16
	maxAxisPointsSet  bool
	lowerLimit        float64
	lowerLimitSet     bool
	upperLimit        float64
	upperLimitSet     bool
	byteOrder         ByteOrder
	deposit           deposit
	extendedLimits    extendedLimits
	format            format
	monotony          Monotony
	physUnit          physUnit
	stepSize          StepSize
}

func parseTypeDefAxis(tok *tokenGenerator) (typeDefAxis, error) {
	tda := typeDefAxis{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case byteOrderToken:
			tda.byteOrder, err = parseByteOrder(tok)
			if err != nil {
				log.Err(err).Msg("typeDefAxis byteOrder could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefAxis byteOrder successfully parsed")
		case depositToken:
			tda.deposit, err = parseDeposit(tok)
			if err != nil {
				log.Err(err).Msg("typeDefAxis deposit could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefAxis deposit successfully parsed")
		case extendedLimitsToken:
			tda.extendedLimits, err = parseExtendedLimits(tok)
			if err != nil {
				log.Err(err).Msg("typeDefAxis extendedLimits could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefAxis extendedLimits successfully parsed")
		case formatToken:
			tda.format, err = parseFormat(tok)
			if err != nil {
				log.Err(err).Msg("typeDefAxis format could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefAxis format successfully parsed")
		case monotonyToken:
			tda.monotony, err = parseMonotony(tok)
			if err != nil {
				log.Err(err).Msg("typeDefAxis monotony could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefAxis monotony successfully parsed")
		case physUnitToken:
			tda.physUnit, err = parsePhysUnit(tok)
			if err != nil {
				log.Err(err).Msg("typeDefAxis physUnit could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefAxis physUnit successfully parsed")
		case stepSizeToken:
			tda.stepSize, err = parseStepSize(tok)
			if err != nil {
				log.Err(err).Msg("typeDefAxis stepSize could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefAxis stepSize successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("typeDefAxis could not be parsed")
				break forLoop
			} else if tok.current() == endTypeDefAxisToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("typeDefAxis could not be parsed")
				break forLoop
			} else if !tda.nameSet {
				tda.name = tok.current()
				tda.nameSet = true
				log.Info().Msg("typeDefAxis name successfully parsed")
			} else if !tda.longIdentifierSet {
				tda.LongIdentifier = tok.current()
				tda.longIdentifierSet = true
				log.Info().Msg("typeDefAxis longIdentifier successfully parsed")
			} else if !tda.inputQuantitySet {
				tda.inputQuantity = tok.current()
				tda.inputQuantitySet = true
				log.Info().Msg("typeDefAxis inputQuantity successfully parsed")
			} else if !tda.recordLayoutSet {
				tda.recordLayout = tok.current()
				tda.recordLayoutSet = true
				log.Info().Msg("typeDefAxis recordLayout successfully parsed")
			} else if !tda.maxDiffSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("typeDefAxis maxDiff could not be parsed")
					break forLoop
				}
				tda.maxDiff = buf
				tda.maxDiffSet = true
				log.Info().Msg("typeDefAxis maxDiff successfully parsed")
			} else if !tda.conversionSet {
				tda.conversion = tok.current()
				tda.conversionSet = true
				log.Info().Msg("typeDefAxis conversion successfully parsed")
			} else if !tda.maxAxisPointsSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 16)
				if err != nil {
					log.Err(err).Msg("typeDefAxis maxAxisPoints could not be parsed")
					break forLoop
				}
				tda.maxAxisPoints = uint16(buf)
				tda.maxAxisPointsSet = true
				log.Info().Msg("typeDefAxis maxAxisPoints successfully parsed")
			} else if !tda.lowerLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("typeDefAxis lowerLimit could not be parsed")
					break forLoop
				}
				tda.lowerLimit = buf
				tda.lowerLimitSet = true
				log.Info().Msg("typeDefAxis lowerLimit successfully parsed")
			} else if !tda.upperLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("typeDefAxis upperLimit could not be parsed")
					break forLoop
				}
				tda.upperLimit = buf
				tda.upperLimitSet = true
				log.Info().Msg("typeDefAxis upperLimit successfully parsed")
			}
		}
	}
	return tda, err
}
