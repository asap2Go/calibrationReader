package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type Project struct {
	Name              string
	nameSet           bool
	longIdentifier    string
	longIdentifierSet bool
	header            header
	Modules           []module
}

func parseProject(tok *tokenGenerator) (Project, error) {
	p := Project{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginHeaderToken:
			p.header, err = parseHeader(tok)
			if err != nil {
				log.Err(err).Msg("project header could not be parsed")
				break forLoop
			}
			log.Info().Msg("project header successfully parsed")
		case beginModuleToken:
			var buf module
			//decision whether to parse the module multithreaded or not
			if useMultithreading {
				buf, err = parseModuleMultithreaded(tok)
			} else {
				buf, err = parseModule(tok)
			}
			if err != nil {
				log.Err(err).Msg("project module could not be parsed")
				break forLoop
			}
			p.Modules = append(p.Modules, buf)
			log.Info().Msg("project module successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("project could not be parsed")
				break forLoop
			} else if tok.current() == endProjectToken {
				break forLoop
			} else if !p.nameSet {
				p.Name = tok.current()
				p.nameSet = true
				log.Info().Msg("project name successfully parsed")
			} else if !p.longIdentifierSet {
				p.longIdentifier = tok.current()
				p.longIdentifierSet = true
				log.Info().Msg("project longIdentifier successfully parsed")
			}
		}
	}
	return p, err
}
