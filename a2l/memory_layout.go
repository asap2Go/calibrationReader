package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type memoryLayout struct {
	prgType    prgTypeEnum
	prgTypeSet bool
	address    string
	addressSet bool
	size       string
	sizeSet    bool
	offset     []int32
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
				ml.address = tok.current()
				ml.addressSet = true
			} else if !ml.sizeSet {
				ml.size = tok.current()
				ml.sizeSet = true
			} else if !ml.offsetSet {
				var buf int64
				buf, err = strconv.ParseInt(tok.current(), 10, 32)
				if err != nil {
					log.Err(err).Msg("shiftOp4 position could not be parsed")
					break forLoop
				}
				ml.offset = append(ml.offset, int32(buf))
				if len(ml.offset) == 5 {
					ml.offsetSet = true
				}
				log.Info().Msg("shiftOp4 position successfully parsed")
			}
		}
	}
	return ml, err
}
