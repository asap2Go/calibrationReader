package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type noRescale5 struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseNoRescale5(tok *tokenGenerator) (noRescale5, error) {
	nr5 := noRescale5{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("noRescale5 could not be parsed")
			break forLoop
		} else if !nr5.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("noRescale5 position could not be parsed")
				break forLoop
			}
			nr5.position = uint16(buf)
			nr5.positionSet = true
			log.Info().Msg("noRescale5 position successfully parsed")
			break forLoop
		} else if !nr5.datatypeSet {
			nr5.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("noRescale5 datatype could not be parsed")
				break forLoop
			}
			nr5.datatypeSet = true
			log.Info().Msg("noRescale5 datatype successfully parsed")
		}
	}
	return nr5, err
}
