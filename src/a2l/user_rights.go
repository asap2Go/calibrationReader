package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type userRights struct {
	userLevelId    string
	userLevelIdSet bool
	readOnly       readOnly
	refGroup       []refGroup
}

func parseUserRights(tok *tokenGenerator) (userRights, error) {
	ur := userRights{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case readOnlyToken:
			ur.readOnly, err = parseReadOnly(tok)
			if err != nil {
				log.Err(err).Msg("userRights readOnly could not be parsed")
				break forLoop
			}
			log.Info().Msg("userRights readOnly successfully parsed")
		case beginRefGroupToken:
			var buf refGroup
			buf, err = parseRefGroup(tok)
			if err != nil {
				log.Err(err).Msg("userRights refGroup could not be parsed")
				break forLoop
			}
			ur.refGroup = append(ur.refGroup, buf)
			log.Info().Msg("userRights refGroup successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("userRights could not be parsed")
				break forLoop
			} else if tok.current() == endUserRightsToken {
				break forLoop
			} else if !ur.userLevelIdSet {
				ur.userLevelId = tok.current()
				ur.userLevelIdSet = true
				log.Info().Msg("userRights userLevelId successfully parsed")
			}
		}
	}
	return ur, err
}
