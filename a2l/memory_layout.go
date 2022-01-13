package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type memoryLayout struct {
	prgType    prgTypeEnum
	prgTypeSet bool
	address    uint32
	addressSet bool
	size       uint32
	sizeSet    bool
	offset     string
	offsetSet  bool
	ifData     []IfData
}

func parseMemoryLayout(tok *tokenGenerator) (memoryLayout, error) {
	ml := memoryLayout{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginIfDataToken:
			var buf IfData
			buf, err = parseIfData(tok)
			if err != nil {
				log.Err(err).Msg("memoryLayout ifData could not be parsed")
				break forLoop
			}
			ml.ifData = append(ml.ifData, buf)
			log.Info().Msg("memoryLayout ifData successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("memoryLayout could not be parsed")
				break forLoop
			} else if tok.current() == endMemoryLayoutToken {
				break forLoop
			} else if !ml.prgTypeSet {
				ml.prgType, err = parsePrgTypeEnum(tok)
				if err != nil {
					log.Err(err).Msg("memoryLayout prgType could not be parsed")
					break forLoop
				}
				ml.prgTypeSet = true
				log.Info().Msg("memoryLayout prgType successfully parsed")
			} else if !ml.addressSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 32)
				if err != nil {
					log.Err(err).Msg("attribute address could not be parsed")
					break forLoop
				}
				ml.address = uint32(buf)
				ml.addressSet = true
			} else if !ml.sizeSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 32)
				if err != nil {
					log.Err(err).Msg("attribute size could not be parsed")
					break forLoop
				}
				ml.size = uint32(buf)
				ml.sizeSet = true
			} else if !ml.offsetSet {
				ml.offset = tok.current()
				ml.offsetSet = true
			}
		}
	}
	return ml, err
}
