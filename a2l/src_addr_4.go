package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

// srcAddr4 is the description of the address of the input quantity in an adjustable object
type srcAddr4 struct {
	position    uint16
	positionSet bool
	datatype    DataTypeEnum
	datatypeSet bool
}

func parseSrcAddr4(tok *tokenGenerator) (srcAddr4, error) {
	sa4 := srcAddr4{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("srcAddr4 could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("srcAddr4 could not be parsed")
			break forLoop
		} else if !sa4.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("srcAddr4 position could not be parsed")
				break forLoop
			}
			sa4.position = uint16(buf)
			sa4.positionSet = true
			log.Info().Msg("srcAddr4 position successfully parsed")
		} else if !sa4.datatypeSet {
			sa4.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("srcAddr4 datatype could not be parsed")
				break forLoop
			}
			sa4.datatypeSet = true
			log.Info().Msg("srcAddr4 datatype successfully parsed")
			break forLoop
		}
	}
	return sa4, err
}
