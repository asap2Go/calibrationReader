package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type varNaming struct {
	tag    tagEnum
	tagSet bool
}

func parseVarNaming(tok *tokenGenerator) (varNaming, error) {
	vn := varNaming{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("varNaming could not be parsed")
	} else if !vn.tagSet {
		var buf tagEnum
		buf, err = parseTagEnum(tok)
		if err != nil {
			log.Err(err).Msg("varNaming tag could not be parsed")
		}
		vn.tag = buf
		log.Info().Msg("varNaming tag successfully parsed")
		vn.tagSet = true
		log.Info().Msg("varNaming tag successfully parsed")
	}
	return vn, err
}
