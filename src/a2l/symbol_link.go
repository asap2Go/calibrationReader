package a2l

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type symbolLink struct {
	symbolName    string
	symbolNameSet bool
	offset        int32
	offsetSet     bool
}

func parseSymbolLink(tok *tokenGenerator) (symbolLink, error) {
	sl := symbolLink{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("symbolLink could not be parsed")
			break forLoop
		} else if !sl.symbolNameSet {
			sl.symbolName = tok.current()
			sl.symbolNameSet = true
			log.Info().Msg("symbolLink symbolName successfully parsed")
		} else if !sl.offsetSet {
			var buf int64
			buf, err = strconv.ParseInt(tok.current(), 10, 32)
			if err != nil {
				log.Err(err).Msg("symbolLink offset could not be parsed")
				break forLoop
			}
			sl.offset = int32(buf)
			sl.offsetSet = true
			log.Info().Msg("symbolLink offset successfully parsed")
			break forLoop
		}

	}
	return sl, err
}
