package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
Reserved can be used to skip specific elements in an adjustable object whose
meaning must not be interpreted by the measurement and calibration system
(e.g. for extensions: new parameters in the adjustable objects).
*/
type Reserved struct {
	Position    uint16
	PositionSet bool
	DataSize    DataSizeEnum
	DataSizeSet bool
}

func parseReserved(tok *tokenGenerator) (Reserved, error) {
	r := Reserved{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("reserved could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("reserved could not be parsed")
			break forLoop
		} else if !r.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("reserved position could not be parsed")
				break forLoop
			}
			r.Position = uint16(buf)
			r.PositionSet = true
			log.Info().Msg("reserved position successfully parsed")
		} else if !r.DataSizeSet {
			r.DataSize, err = parseDataSizeEnum(tok)
			if err != nil {
				log.Err(err).Msg("reserved dataSize could not be parsed")
				break forLoop
			}
			r.DataSizeSet = true
			log.Info().Msg("reserved dataSize successfully parsed")
			break forLoop
		}
	}
	return r, err
}
