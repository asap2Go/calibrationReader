package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

//arPrototypeOf usage shall not lead to circular dependencies between functions and prototypes.
//arPrototypeOf shall not reference to the FUNCTION where it is stated.
type arPrototypeOf struct {
	//Name of the FUNCTION which describes the "SwComponentType" from which "SwComponentPrototype" is derived.
	name    string
	nameSet bool
}

func parseArPrototypeOf(tok *tokenGenerator) (arPrototypeOf, error) {
	apo := arPrototypeOf{}
	var err error
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("arComponent could not be parsed")
	} else if !apo.nameSet {
		apo.name = tok.current()
		apo.nameSet = true
		log.Info().Msg("arPrototypeOf name successfully parsed")
	}
	return apo, err
}
