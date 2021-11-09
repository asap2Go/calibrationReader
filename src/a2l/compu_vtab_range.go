package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type compuVTabRange struct {
	name                  string
	nameSet               bool
	longIdentifier        string
	longIdentifierSet     bool
	numberValueTriples    uint16
	numberValueTriplesSet bool
	inValMin              []float64
	inValMinSet           bool
	inValMax              []float64
	inValMaxSet           bool
	outVal                []string
	outValSet             bool
	defaultValue          defaultValue
}

func parseCompuVtabRange(tok *tokenGenerator) (compuVTabRange, error) {
	cvr := compuVTabRange{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case defaultValueToken:
			cvr.defaultValue, err = parseDefaultValue(tok)
			if err != nil {
					log.Err(err).Msg("compuTab defaultValue could not be parsed")
					break forLoop
			}
					log.Info().Msg("compuVtabRange defaultValue successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
					log.Err(err).Msg("compuTab could not be parsed")
					break forLoop
			} else if tok.current() == endCompuVtabRangeToken {
				break forLoop
			} else if !cvr.nameSet {
				cvr.name = tok.current()
				cvr.nameSet = true
					log.Info().Msg("compuVtabRange name successfully parsed")
			} else if !cvr.longIdentifierSet {
				cvr.longIdentifier = tok.current()
				cvr.longIdentifierSet = true
					log.Info().Msg("compuVtabRange longIdentifier successfully parsed")
			} else if !cvr.numberValueTriplesSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 16)
				if err != nil {
						log.Err(err).Msg("compuVtabRange numberValueTriples could not be parsed")
						break forLoop
				}
					cvr.numberValueTriples = uint16(buf)
					cvr.numberValueTriplesSet = true
						log.Info().Msg("compuVtabRange numberValueTriples successfully parsed")
			} else if !cvr.inValMinSet || !cvr.inValMaxSet || !cvr.outValSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
						log.Err(err).Msg("compuVtabRange inValMin could not be parsed")
						break forLoop
				}
					cvr.inValMin = append(cvr.inValMin, buf)
						log.Info().Msg("compuVtabRange inValMin successfully parsed")
					if uint16(len(cvr.inValMin)) == cvr.numberValueTriples {
						cvr.inValMinSet = true
							log.Info().Msg("compuVtabRange inValMin successfully parsed")
					}
				tok.next()
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
						log.Err(err).Msg("compuVtabRange inValMax could not be parsed")
						break forLoop
				}
					cvr.inValMax = append(cvr.inValMax, buf)
						log.Info().Msg("compuVtabRange inValMax successfully parsed")
					if uint16(len(cvr.inValMax)) == cvr.numberValueTriples {
						cvr.inValMaxSet = true
							log.Info().Msg("compuVtabRange inValMax successfully parsed")
					}
				tok.next()
				cvr.outVal = append(cvr.outVal, tok.current())
					log.Info().Msg("compuVtabRange outVal successfully parsed")
				if uint16(len(cvr.outVal)) == cvr.numberValueTriples {
					cvr.outValSet = true
				}
			}
		}
	}
	return cvr, err
}
