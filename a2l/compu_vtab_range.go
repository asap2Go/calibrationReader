package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type CompuVTabRange struct {
	Name                  string
	NameSet               bool
	LongIdentifier        string
	LongIdentifierSet     bool
	NumberValueTriples    uint16
	NumberValueTriplesSet bool
	InValMin              []float64
	InValMinSet           bool
	InValMax              []float64
	InValMaxSet           bool
	OutVal                []string
	OutValSet             bool
	DefaultValue          DefaultValue
}

func parseCompuVtabRange(tok *tokenGenerator) (CompuVTabRange, error) {
	cvr := CompuVTabRange{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case defaultValueToken:
			cvr.DefaultValue, err = parseDefaultValue(tok)
			if err != nil {
				log.Err(err).Msg("compuTab defaultValue could not be parsed")
				break forLoop
			}
			log.Info().Msg("compuVtabRange defaultValue successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("compuVTabRange could not be parsed")
				break forLoop
			} else if tok.current() == endCompuVtabRangeToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("compuVTabRange could not be parsed")
				break forLoop
			} else if !cvr.NameSet {
				cvr.Name = tok.current()
				cvr.NameSet = true
				log.Info().Msg("compuVtabRange name successfully parsed")
			} else if !cvr.LongIdentifierSet {
				cvr.LongIdentifier = tok.current()
				cvr.LongIdentifierSet = true
				log.Info().Msg("compuVtabRange longIdentifier successfully parsed")
			} else if !cvr.NumberValueTriplesSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 16)
				if err != nil {
					log.Err(err).Msg("compuVtabRange numberValueTriples could not be parsed")
					break forLoop
				}
				cvr.NumberValueTriples = uint16(buf)
				cvr.NumberValueTriplesSet = true
				log.Info().Msg("compuVtabRange numberValueTriples successfully parsed")
			} else if !cvr.InValMinSet || !cvr.InValMaxSet || !cvr.OutValSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("compuVtabRange inValMin could not be parsed")
					break forLoop
				}
				cvr.InValMin = append(cvr.InValMin, buf)
				log.Info().Msg("compuVtabRange inValMin successfully parsed")
				if uint16(len(cvr.InValMin)) == cvr.NumberValueTriples {
					cvr.InValMinSet = true
					log.Info().Msg("compuVtabRange inValMin successfully parsed")
				}
				tok.next()
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("compuVtabRange inValMax could not be parsed")
					break forLoop
				}
				cvr.InValMax = append(cvr.InValMax, buf)
				log.Info().Msg("compuVtabRange inValMax successfully parsed")
				if uint16(len(cvr.InValMax)) == cvr.NumberValueTriples {
					cvr.InValMaxSet = true
					log.Info().Msg("compuVtabRange inValMax successfully parsed")
				}
				tok.next()
				cvr.OutVal = append(cvr.OutVal, tok.current())
				log.Info().Msg("compuVtabRange outVal successfully parsed")
				if uint16(len(cvr.OutVal)) == cvr.NumberValueTriples {
					cvr.OutValSet = true
				}
			}
		}
	}
	return cvr, err
}
