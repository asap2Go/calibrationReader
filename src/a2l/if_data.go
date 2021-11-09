package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type IfData struct {
	name    string
	nameSet bool
	data    []string
	dataSet bool
}

func parseIfData(tok *tokenGenerator) (IfData, error) {
	id := IfData{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("ifData: could not be parsed")
			break forLoop
		} else if tok.current() == endIfDataToken {
			id.dataSet = true
				log.Info().Msg("ifData data successfully parsed")
			break forLoop
		} else if !id.nameSet {
			id.name = tok.current()
			id.nameSet = true
				log.Info().Msg("ifData name successfully parsed")
		} else if !id.dataSet {
			id.data = append(id.data, tok.current())
		}
	}
	return id, err
}
