package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

//annotationLabel is the label or title of the annotation
type annotationLabel struct {
	label    string
	labelSet bool
}

func parseAnnotationLabel(tok *tokenGenerator) (annotationLabel, error) {
	al := annotationLabel{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("annotationLabel could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("annotationLabel could not be parsed")
	} else if !al.labelSet {
		al.label = tok.current()
		al.labelSet = true
		log.Info().Msg("annotationLabel label successfully parsed")
	}
	return al, err
}
