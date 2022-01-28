package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type refMemorySegment struct {
	name    string
	nameSet bool
}

func parseRefMemorySegment(tok *tokenGenerator) (refMemorySegment, error) {
	rms := refMemorySegment{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("refMemorySegment: could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("refMemorySegment could not be parsed")
	} else if !rms.nameSet {
		rms.name = tok.current()
		rms.nameSet = true
		log.Info().Msg("refMemorySegment name successfully parsed")
	}
	return rms, err
}
