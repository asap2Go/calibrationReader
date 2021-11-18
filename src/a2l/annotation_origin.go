package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

//annotationOrigin specifies the creator or creating system of the annotation
type annotationOrigin struct {
	origin    string
	originSet bool
}

func parseAnnotationOrigin(tok *tokenGenerator) (annotationOrigin, error) {
	ao := annotationOrigin{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("annotationOrigin could not be parsed")
	} else if !ao.originSet {
		ao.origin = tok.current()
		ao.originSet = true
		log.Info().Msg("annotationOrigin origin successfully parsed")
	}
	return ao, err
}
