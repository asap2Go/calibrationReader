package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type user struct {
	userName    string
	userNameSet bool
}

func parseUser(tok *tokenGenerator) (user, error) {
	u := user{}
	var err error
	tok.next()
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("user could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("user could not be parsed")
	} else if !u.userNameSet {
		u.userName = tok.current()
		u.userNameSet = true
		log.Info().Msg("user userName successfully parsed")
	}

	return u, err
}
