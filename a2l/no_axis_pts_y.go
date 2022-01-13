package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type noAxisPtsY struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseNoAxisPtsY(tok *tokenGenerator) (noAxisPtsY, error) {
	napy := noAxisPtsY{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("noAxisPtsy could not be parsed")
			break forLoop
		} else if !napy.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noAxisPtsy position could not be parsed")
				break forLoop
			}
			napy.position = uint16(buf)
			napy.positionSet = true
			log.Info().Msg("noAxisPtsy position successfully parsed")
		} else if !napy.datatypeSet {
			napy.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noAxisPtsy datatype could not be parsed")
				break forLoop
			}
			napy.datatypeSet = true
			log.Info().Msg("noAxisPtsy datatype successfully parsed")
			break forLoop
		}
	}
	return napy, err
}
