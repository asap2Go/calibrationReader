package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

//annotationText may represent a multi-line ASCII description text (voluminous description).
//Its purpose is to be an application note which explains the function of an identifier for the calibration engineer.
type annotationText struct {
	annotationText    string
	annotationTextSet bool
}

func parseAnnotationText(tok *tokenGenerator) (annotationText, error) {
	at := annotationText{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("annotationText could not be parsed")
			break forLoop
		} else if tok.current() == endAnnotationTextToken {
			at.annotationTextSet = true
			break forLoop
		} else if !at.annotationTextSet {
			at.annotationText = at.annotationText + spaceToken + tok.current()
		}
	}
	return at, err
}
