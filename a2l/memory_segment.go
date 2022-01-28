package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type memorySegment struct {
	name              string
	nameSet           bool
	longIdentifier    string
	longIdentifierSet bool
	prgType           prgTypeEnum
	prgTypeSet        bool
	memoryType        memoryTypeEnum
	memoryTypeSet     bool
	attribute         attributeEnum
	attributeSet      bool
	address           string
	addressSet        bool
	size              string
	sizeSet           bool
	offset            [5]string
	offsetSet         bool
	ifData            []IfData
}

func parseMemorySegment(tok *tokenGenerator) (memorySegment, error) {
	ms := memorySegment{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginIfDataToken:
			var buf IfData
			buf, err = parseIfData(tok)
			if err != nil {
				log.Err(err).Msg("memorySegment ifData could not be parsed")
				break forLoop
			}
			ms.ifData = append(ms.ifData, buf)
			log.Info().Msg("memorySegment ifData successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("memorySegment could not be parsed")
				break forLoop
			} else if tok.current() == endMemorySegmentToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("memorySegment could not be parsed")
				break forLoop
			} else if !ms.nameSet {
				ms.name = tok.current()
				ms.nameSet = true
				log.Info().Msg("memorySegment name successfully parsed")
			} else if !ms.longIdentifierSet {
				ms.longIdentifier = tok.current()
				ms.longIdentifierSet = true
				log.Info().Msg("memorySegment longIdentifier successfully parsed")
			} else if !ms.prgTypeSet {
				ms.prgType, err = parsePrgTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("memorySegment prgType could not be parsed")
					break forLoop
				}
				ms.prgTypeSet = true
				log.Info().Msg("memorySegment prgType successfully parsed")
			} else if !ms.memoryTypeSet {
				ms.memoryType, err = parseMemoryTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("memorySegment memoryType could not be parsed")
					break forLoop
				}
				ms.memoryTypeSet = true
				log.Info().Msg("memorySegment memoryType successfully parsed")
			} else if !ms.attributeSet {
				ms.attribute, err = parseAttributeEnum(tok)
				if err != nil {
					log.Err(err).Msg("memorySegment attribute could not be parsed")
					break forLoop
				}
				ms.attributeSet = true
				log.Info().Msg("memorySegment address successfully parsed")
			} else if !ms.addressSet {
				ms.address = tok.current()
				ms.addressSet = true
			} else if !ms.sizeSet {
				ms.size = tok.current()
				ms.sizeSet = true
			} else if !ms.offsetSet {
				for i := 0; i < 5; i++ {
					ms.offset[i] = tok.current()
				}
				ms.offsetSet = true
			}
		}
	}
	return ms, err
}
