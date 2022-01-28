package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type header struct {
	comment    string
	commentSet bool
	projectNo  projectNo
	version    version
}

func parseHeader(tok *tokenGenerator) (header, error) {
	h := header{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case projectNoToken:
			h.projectNo, err = parseProjectNo(tok)
			if err != nil {
				log.Err(err).Msg("header projectNo could not be parsed")
				break forLoop
			}
			log.Info().Msg("header projectNo successfully parsed")
		case versionToken:
			h.version, err = parseVersion(tok)
			if err != nil {
				log.Err(err).Msg("header version could not be parsed")
				break forLoop
			}
			log.Info().Msg("header version successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("header could not be parsed: unexpected end of file")
				break forLoop
			} else if tok.current() == endHeaderToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("header could not be parsed")
				break forLoop
			} else if !h.commentSet {
				h.comment = tok.current()
				h.commentSet = true
				log.Info().Msg("header comment successfully parsed")
			}
		}
	}
	return h, err
}
