package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type noAxisPts5 struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseNoAxisPts5(tok *tokenGenerator) (noAxisPts5, error) {
	nap5 := noAxisPts5{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("noAxisPts5 could not be parsed")
			break forLoop
		} else if !nap5.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noAxisPts5 position could not be parsed")
				break forLoop
			}
			nap5.position = uint16(buf)
			nap5.positionSet = true
			log.Info().Msg("noAxisPts5 position successfully parsed")
		} else if !nap5.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noAxisPts5 datatype could not be parsed")
				break forLoop
			}
			nap5.datatype = buf
			nap5.datatypeSet = true
			log.Info().Msg("noAxisPts5 datatype successfully parsed")
			break forLoop
		}
	}
	return nap5, err
}
