package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type overwrite struct {
	name           string
	nameSet        bool
	axisNumer      uint16
	axisNumberSet  bool
	conversion     string
	extendedLimits extendedLimits
	inputQuantity  string
	format         format
	formatSet      bool
	limits         limits
	monotony       Monotony
	physUnit       physUnit
}

func parseOverwrite(tok *tokenGenerator) (overwrite, error) {
	o := overwrite{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case conversionToken:
			o.conversion = tok.current()
			log.Info().Msg("overwrite conversion successfully parsed")
		case extendedLimitsToken:
			o.extendedLimits, err = parseExtendedLimits(tok)
			if err != nil {
				log.Err(err).Msg("overwrite extendedLimits could not be parsed")
				break forLoop
			}
			log.Info().Msg("overwrite extendedLimits successfully parsed")
		case formatToken:
			o.format, err = parseFormat(tok)
			if err != nil {
				log.Err(err).Msg("overwrite format could not be parsed")
				break forLoop
			}
			log.Info().Msg("overwrite format successfully parsed")
		case inputQuantityToken:
			o.inputQuantity = tok.current()
			log.Info().Msg("overwrite inputQuantity successfully parsed")
		case limitsToken:
			o.limits, err = parseLimits(tok)
			if err != nil {
				log.Err(err).Msg("overwrite limits could not be parsed")
				break forLoop
			}
			log.Info().Msg("overwrite limits successfully parsed")
		case monotonyToken:
			o.monotony, err = parseMonotony(tok)
			if err != nil {
				log.Err(err).Msg("overwrite monotony could not be parsed")
				break forLoop
			}
			log.Info().Msg("overwrite monotony successfully parsed")
		case physUnitToken:
			o.physUnit, err = parsePhysUnit(tok)
			if err != nil {
				log.Err(err).Msg("overwrite physUnit could not be parsed")
				break forLoop
			}
			log.Info().Msg("overwrite physUnit successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("overwrite could not be parsed")
				break forLoop
			} else if tok.current() == endOverwriteToken {
				break forLoop
			} else if !o.nameSet {
				o.name = tok.current()
				o.nameSet = true
				log.Info().Msg("overwrite name successfully parsed")
			} else if !o.axisNumberSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 16)
				if err != nil {
					log.Err(err).Msg("overwrite axisNumber could not be parsed")
					break forLoop
				}
				o.axisNumer = uint16(buf)
				o.axisNumberSet = true
				log.Info().Msg("overwrite axisNumber successfully parsed")
			}
		}
	}
	return o, err
}
