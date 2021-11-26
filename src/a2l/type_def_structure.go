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
	addressType        AddrTypeEnum
	consistentExchange consistentExchangeKeyword
	structureComponent structureComponent
	symbolTypeLink     string
}

func parseTypeDefStructure(tok *tokenGenerator) (typeDefStructure, error) {
	sc := typeDefStructure{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case addressTypeToken:
			sc.addressType, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("typeDefStructure addressType could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefStructure addressType successfully parsed")
		case consistentExchangeToken:
			sc.consistentExchange, err = parseConsistentExchangeKeyword(tok)
			if err != nil {
				log.Err(err).Msg("typeDefStructure consistentExchange could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefStructure consistentExchange successfully parsed")
		case beginStructureComponentToken:
			sc.structureComponent, err = parseStructureComponent(tok)
			if err != nil {
				log.Err(err).Msg("typeDefStructure structureComponent could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefStructure typeDefStructure successfully parsed")
		case symbolTypeLinkToken:
			sc.symbolTypeLink = tok.current()
			log.Info().Msg("typeDefStructure symbolTypeLink successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("typeDefStructure could not be parsed")
				break forLoop
			} else if tok.current() == endTypeDefStructureToken {
				break forLoop
			} else if !sc.nameSet {
				sc.name = tok.current()
				sc.nameSet = true
				log.Info().Msg("typeDefStructure name successfully parsed")
			} else if !sc.longIdentifierSet {
				sc.longIdentifier = tok.current()
				sc.longIdentifierSet = true
				log.Info().Msg("typeDefStructure typeDefName successfully parsed")
			} else if !sc.sizeSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 32)
				if err != nil {
					log.Err(err).Msg("typeDefStructure address could not be parsed")
					break forLoop
				}
				sc.size = uint32(buf)
				sc.sizeSet = true
				log.Info().Msg("typeDefStructure address successfully parsed")
			}
		}
	}
	return sc, err
}
