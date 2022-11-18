package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type NoAxisPts4 struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseNoAxisPts4(tok *tokenGenerator) (NoAxisPts4, error) {
	nap4 := NoAxisPts4{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("noAxisPts4 could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("noAxisPts4 could not be parsed")
			break forLoop
		} else if !nap4.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noAxisPts4 position could not be parsed")
				break forLoop
			}
			nap4.Position = uint16(buf)
			nap4.PositionSet = true
			log.Info().Msg("noAxisPts4 position successfully parsed")
		} else if !nap4.DatatypeSet {
			nap4.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noAxisPts4 datatype could not be parsed")
				break forLoop
			}
			nap4.DatatypeSet = true
			log.Info().Msg("noAxisPts4 datatype successfully parsed")
			break forLoop
		}
	}
	return nap4, err
}
