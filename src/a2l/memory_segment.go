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
			} else if !ms.nameSet {
				ms.name = tok.current()
				ms.nameSet = true
				log.Info().Msg("memorySegment name successfully parsed")
			} else if !ms.longIdentifierSet {
				ms.longIdentifier = tok.current()
				ms.longIdentifierSet = true
				log.Info().Msg("memorySegment longIdentifier successfully parsed")
			} else if !ms.prgTypeSet {
				var buf prgTypeEnum
				buf, err = parsePrgTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("memorySegment prgType could not be parsed")
					break forLoop
				}
				ms.prgType = buf
				ms.prgTypeSet = true
				log.Info().Msg("memorySegment prgType successfully parsed")
			} else if !ms.memoryTypeSet {
				var buf memoryTypeEnum
				buf, err = parseMemoryTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("memorySegment memoryType could not be parsed")
					break forLoop
				}
				ms.memoryType = buf
				ms.memoryTypeSet = true
				log.Info().Msg("memorySegment memoryType successfully parsed")
			} else if !ms.attributeSet {
				var buf attributeEnum
				buf, err = parseAttributeEnum(tok)
				if err != nil {
					log.Err(err).Msg("memorySegment attribute could not be parsed")
					break forLoop
				}
				ms.attribute = buf
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
