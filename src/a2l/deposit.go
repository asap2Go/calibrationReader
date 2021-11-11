package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type Deposit struct {
	mode    ModeEnum
	modeSet bool
}

func parseDeposit(tok *tokenGenerator) (Deposit, error) {
	d := Deposit{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("deposit could not be parsed")
	} else if !d.modeSet {
		d.mode, err = parseModeEnum(tok)
		if err != nil {
			log.Err(err).Msg("deposit could not be parsed")
		}
		d.modeSet = true
		log.Info().Msg("deposit mode successfully parsed")
	}
	return d, err
}
