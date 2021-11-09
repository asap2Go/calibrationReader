package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type noAxisPts4 struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseNoAxisPts4(tok *tokenGenerator) (noAxisPts4, error) {
	nap4 := noAxisPts4{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("noAxisPts4 could not be parsed")
			break forLoop
		} else if !nap4.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
					log.Err(err).Msg("noAxisPts4 position could not be parsed")
				break forLoop
			}
			nap4.position = uint16(buf)
			nap4.positionSet = true
				log.Info().Msg("noAxisPts4 position successfully parsed")
		} else if !nap4.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
					log.Err(err).Msg("noAxisPts4 datatype could not be parsed")
				break forLoop
			}
			nap4.datatype = buf
			nap4.datatypeSet = true
				log.Info().Msg("noAxisPts4 datatype successfully parsed")
			break forLoop
		}
	}
	return nap4, err
}
