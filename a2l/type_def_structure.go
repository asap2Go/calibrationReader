package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type typeDefStructure struct {
	name               string
	nameSet            bool
	longIdentifier     string
	longIdentifierSet  bool
	size               uint32
	sizeSet            bool
	addressType        addrTypeEnum
	consistentExchange consistentExchangeKeyword
	structureComponent []structureComponent
	symbolTypeLink     symbolTypeLink
}

func parseTypeDefStructure(tok *tokenGenerator) (typeDefStructure, error) {
	tds := typeDefStructure{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case addressTypeToken:
			tds.addressType, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("typeDefStructure addressType could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefStructure addressType successfully parsed")
		case consistentExchangeToken:
			tds.consistentExchange, err = parseConsistentExchangeKeyword(tok)
			if err != nil {
				log.Err(err).Msg("typeDefStructure consistentExchange could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefStructure consistentExchange successfully parsed")
		case beginStructureComponentToken:
			var buf structureComponent
			buf, err = parseStructureComponent(tok)
			if err != nil {
				log.Err(err).Msg("typeDefStructure structureComponent could not be parsed")
				break forLoop
			}
			tds.structureComponent = append(tds.structureComponent, buf)
			log.Info().Msg("typeDefStructure typeDefStructure successfully parsed")
		case symbolTypeLinkToken:
			tds.symbolTypeLink, err = parseSymbolTypeLink(tok)
			if err != nil {
				log.Err(err).Msg("typeDefStructure symbolTypeLink could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefStructure symbolTypeLink successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("typeDefStructure could not be parsed")
				break forLoop
			} else if tok.current() == endTypeDefStructureToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("typeDefStructure could not be parsed")
				break forLoop
			} else if !tds.nameSet {
				tds.name = tok.current()
				tds.nameSet = true
				log.Info().Msg("typeDefStructure name successfully parsed")
			} else if !tds.longIdentifierSet {
				tds.longIdentifier = tok.current()
				tds.longIdentifierSet = true
				log.Info().Msg("typeDefStructure typeDefName successfully parsed")
			} else if !tds.sizeSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 32)
				if err != nil {
					log.Err(err).Msg("typeDefStructure address could not be parsed")
					break forLoop
				}
				tds.size = uint32(buf)
				tds.sizeSet = true
				log.Info().Msg("typeDefStructure address successfully parsed")
			}
		}
	}
	return tds, err
}
