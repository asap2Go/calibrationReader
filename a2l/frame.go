package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type frame struct {
	name              string
	nameSet           bool
	longIdentifier    string
	longIdentifierSet bool
	scalingUnit       uint16
	scalingUnitSet    bool
	rate              uint32
	rateSet           bool
	frameMeasurement  frameMeasurement
	ifData            []IfData
}

func parseFrame(tok *tokenGenerator) (frame, error) {
	f := frame{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case frameMeasurementToken:
			f.frameMeasurement, err = parseFrameMeasurement(tok)
			if err != nil {
				break forLoop
			}
			log.Info().Msg("frame frameMeasurement successfully parsed")
		case beginIfDataToken:
			var buf IfData
			buf, err = parseIfData(tok)
			if err != nil {
				break forLoop
			}
			f.ifData = append(f.ifData, buf)
			log.Info().Msg("frame ifData successfully parsed")
		default:
			if tok.current() == emptyToken {
				break forLoop
			} else if tok.current() == endFrameToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("frame could not be parsed")
				break forLoop
			} else if !f.nameSet {
				f.name = tok.current()
			} else if !f.longIdentifierSet {
				f.longIdentifier = tok.current()
			} else if !f.scalingUnitSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 16)
				if err != nil {
					log.Err(err).Msg("attribute scalingUnit could not be parsed")
					break forLoop
				}
				f.scalingUnit = uint16(buf)
			} else if !f.rateSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 32)
				if err != nil {
					log.Err(err).Msg("attribute rate could not be parsed")
					break forLoop
				}
				f.rate = uint32(buf)
			}
		}
	}
	return f, err
}
