package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type structureComponent struct {
	name             string
	nameSet          bool
	typeDefName      string
	typeDefNameSet   bool
	addressOffset    string //uint32
	addressOffsetSet bool
	addressType      AddrTypeEnum
	layout           layout
	matrixDim        MatrixDim
	symbolTypeLink   symbolTypeLink
}

func parseStructureComponent(tok *tokenGenerator) (structureComponent, error) {
	sc := structureComponent{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case addressTypeToken:
			sc.addressType, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("structureComponent addressType could not be parsed")
				break forLoop
			}
			log.Info().Msg("structureComponent addressType successfully parsed")
		case layoutToken:
			sc.layout, err = parseLayout(tok)
			if err != nil {
				log.Err(err).Msg("structureComponent layout could not be parsed")
				break forLoop
			}
			log.Info().Msg("structureComponent layout successfully parsed")
		case matrixDimToken:
			sc.matrixDim, err = parseMatrixDim(tok)
			if err != nil {
				log.Err(err).Msg("structureComponent matrixDim could not be parsed")
				break forLoop
			}
			log.Info().Msg("structureComponent matrixDim successfully parsed")
		case symbolTypeLinkToken:
			sc.symbolTypeLink, err = parseSymbolTypeLink(tok)
			if err != nil {
				log.Err(err).Msg("typeDefStructure symbolTypeLink could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefStructure symbolTypeLink successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("structureComponent could not be parsed")
				break forLoop
			} else if tok.current() == endStructureComponentToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("structureComponent could not be parsed")
				break forLoop
			} else if !sc.nameSet {
				sc.name = tok.current()
				sc.nameSet = true
				log.Info().Msg("structureComponent name successfully parsed")
			} else if !sc.typeDefNameSet {
				sc.typeDefName = tok.current()
				sc.typeDefNameSet = true
				log.Info().Msg("structureComponent typeDefName successfully parsed")
			} else if !sc.addressOffsetSet {
				sc.addressOffset = tok.current()
				sc.addressOffsetSet = true
				log.Info().Msg("structureComponent address successfully parsed")
			}
		}
	}
	return sc, err
}
