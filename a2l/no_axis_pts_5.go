package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type NoAxisPts5 struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseNoAxisPts5(tok *tokenGenerator) (NoAxisPts5, error) {
	nap5 := NoAxisPts5{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("noAxisPts5 could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("noAxisPts5 could not be parsed")
			break forLoop
		} else if !nap5.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noAxisPts5 position could not be parsed")
				break forLoop
			}
			nap5.Position = uint16(buf)
			nap5.PositionSet = true
			log.Info().Msg("noAxisPts5 position successfully parsed")
		} else if !nap5.DatatypeSet {
			nap5.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noAxisPts5 datatype could not be parsed")
				break forLoop
			}
			nap5.DatatypeSet = true
			log.Info().Msg("noAxisPts5 datatype successfully parsed")
			break forLoop
		}
	}
	return nap5, err
}
