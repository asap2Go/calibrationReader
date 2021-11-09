package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type layout struct {
	indexMode    IndexModeEnum
	indexModeSet bool
}

func parseLayout(tok *tokenGenerator) (layout, error) {
	l := layout{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
			log.Err(err).Msg("layout could not be parsed")
	} else if !l.indexModeSet {
		var buf IndexModeEnum
		buf, err = parseIndexModeEnum(tok)
		if err != nil {
				log.Err(err).Msg("layout indexMode could not be parsed")
		}
		l.indexMode = buf
		l.indexModeSet = true
			log.Info().Msg("layout indexMode successfully parsed")
	}
	return l, err
}
