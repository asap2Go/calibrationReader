package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type noRescale4 struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseNoRescale4(tok *tokenGenerator) (noRescale4, error) {
	nr4 := noRescale4{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("noRescale4 could not be parsed")
			break forLoop
		} else if !nr4.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noRescale4 position could not be parsed")
				break forLoop
			}
			nr4.position = uint16(buf)
			nr4.positionSet = true
			log.Info().Msg("noRescale4 position successfully parsed")
			break forLoop
		} else if !nr4.datatypeSet {
			nr4.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noRescale4 datatype could not be parsed")
				break forLoop
			}
			nr4.datatypeSet = true
			log.Info().Msg("noRescale4 datatype successfully parsed")
		}
	}
	return nr4, err
}
