package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type srcAddr5 struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseSrcAddr5(tok *tokenGenerator) (srcAddr5, error) {
	sa5 := srcAddr5{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("srcAddr5 could not be parsed")
			break forLoop
		} else if !sa5.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("srcAddr5 position could not be parsed")
				break forLoop
			}
			sa5.position = uint16(buf)
			sa5.positionSet = true
			log.Info().Msg("srcAddr5 position successfully parsed")
		} else if !sa5.datatypeSet {
			sa5.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("srcAddr5 datatype could not be parsed")
				break forLoop
			}
			sa5.datatypeSet = true
			log.Info().Msg("srcAddr5 datatype successfully parsed")
			break forLoop
		}
	}
	return sa5, err
}
