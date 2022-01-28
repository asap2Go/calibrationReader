package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

//annotation may represent a voluminous description.
//Its purpose is to be e.g. an application note which explains the function of an identifier for the calibration engineer.
type annotation struct {
	annotationLabel  annotationLabel
	annotationOrigin annotationOrigin
	annotationText   []annotationText
}

func parseAnnotation(tok *tokenGenerator) (annotation, error) {
	an := annotation{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case annotationLabelToken:
			an.annotationLabel, err = parseAnnotationLabel(tok)
			if err != nil {
				log.Err(err).Msg("annotation annotationLabel could not be parsed")
				break forLoop
			}
			log.Info().Msg("annotation annotationLabel successfully parsed")
		case annotationOriginToken:
			an.annotationOrigin, err = parseAnnotationOrigin(tok)
			if err != nil {
				log.Err(err).Msg("annotation annotationOrigin could not be parsed")
				break forLoop
			}
			log.Info().Msg("annotation annotationOrigin successfully parsed")
		case beginAnnotationTextToken:
			var buf annotationText
			buf, err = parseAnnotationText(tok)
			if err != nil {
				log.Err(err).Msg("annotation annotationText could not be parsed")
				break forLoop
			}
			an.annotationText = append(an.annotationText, buf)
			log.Info().Msg("annotation annotationText successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("annotation could not be parsed")
				break forLoop
			} else if tok.current() == endAnnotationToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("annotation could not be parsed")
				break forLoop
			}
		}
	}
	return an, err
}
