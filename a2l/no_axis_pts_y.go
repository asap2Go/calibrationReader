package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type NoAxisPtsY struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseNoAxisPtsY(tok *tokenGenerator) (NoAxisPtsY, error) {
	napy := NoAxisPtsY{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("noAxisPtsy could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("noAxisPtsY could not be parsed")
			break forLoop
		} else if !napy.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noAxisPtsy position could not be parsed")
				break forLoop
			}
			napy.Position = uint16(buf)
			napy.PositionSet = true
			log.Info().Msg("noAxisPtsy position successfully parsed")
		} else if !napy.DatatypeSet {
			napy.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noAxisPtsy datatype could not be parsed")
				break forLoop
			}
			napy.DatatypeSet = true
			log.Info().Msg("noAxisPtsy datatype successfully parsed")
			break forLoop
		}
	}
	return napy, err
}
