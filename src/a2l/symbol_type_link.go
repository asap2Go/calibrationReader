package a2l

import (
	"github.com/rs/zerolog/log"
)

type symbolTypeLink struct {
	symbolName    string
	symbolNameSet bool
}

func parseSymbolTypeLink(tok *tokenGenerator) symbolTypeLink {
	stl := symbolTypeLink{}
	if !stl.symbolNameSet {
		stl.symbolName = tok.current()
		stl.symbolNameSet = true
		log.Info().Msg("symbolTypeLink symbolName successfully parsed")
	}
	return stl
}
