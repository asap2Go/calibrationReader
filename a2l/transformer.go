package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

//transformer may represent a voluminous description.
//Its purpose is to be e.g. an application note which explains the function of an identifier for the calibration engineer.
type transformer struct {
	name                  string
	nameSet               bool
	version               string
	versionSet            bool
	executable32          string
	executable32Set       bool
	executable64          string
	executable64Set       bool
	timeOut               uint32
	timeOutSet            bool
	trigger               triggerEnum
	triggerSet            bool
	inverseTransformer    string
	inverseTransformerSet bool
	transformerInObjects  transformerInObject
	transformerOutObjects transformerOutObject
}

func parseTransformer(tok *tokenGenerator) (transformer, error) {
	t := transformer{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginTransformerInObjectsToken:
			t.transformerInObjects, err = parseTransformerInObject(tok)
			if err != nil {
				log.Err(err).Msg("transformer transformerInObjects could not be parsed")
				break forLoop
			}
			log.Info().Msg("transformer transformerInObjects successfully parsed")
		case beginTransformerOutObjectsToken:
			t.transformerOutObjects, err = parseTransformerOutObject(tok)
			if err != nil {
				log.Err(err).Msg("transformer transformerOutObjects could not be parsed")
				break forLoop
			}
			log.Info().Msg("transformer transformerOutObjects successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("transformer could not be parsed")
				break forLoop
			} else if tok.current() == endTransformerToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("transformer could not be parsed")
				break forLoop
			} else if !t.nameSet {
				t.name = tok.current()
				t.nameSet = true
				log.Info().Msg("transformer name successfully parsed")
			} else if !t.versionSet {
				t.version = tok.current()
				t.versionSet = true
				log.Info().Msg("transformer version successfully parsed")
			} else if !t.executable32Set {
				t.executable32 = tok.current()
				t.executable32Set = true
				log.Info().Msg("transformer executable32 successfully parsed")
			} else if !t.executable64Set {
				t.executable64 = tok.current()
				t.executable64Set = true
				log.Info().Msg("transformer executable64 successfully parsed")
			} else if !t.timeOutSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 32)
				if err != nil {
					log.Err(err).Msg("transformer timeOut could not be parsed")
					break forLoop
				}
				t.timeOut = uint32(buf)
				t.timeOutSet = true
				log.Info().Msg("transformer timeOut successfully parsed")
			} else if !t.triggerSet {
				t.trigger, err = parseTriggerEnum(tok)
				if err != nil {
					log.Err(err).Msg("transformer trigger could not be parsed")
					break forLoop
				}
				t.triggerSet = true
				log.Info().Msg("transformer trigger successfully parsed")
			} else if !t.inverseTransformerSet {
				t.inverseTransformer = tok.current()
				t.inverseTransformerSet = true
				log.Info().Msg("transformer inverseTransformer successfully parsed")
			}
		}
	}
	return t, err
}
