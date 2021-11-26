package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type typeDefBlob struct {
	name              string
	nameSet           bool
	longIdentifier    string
	longIdentifierSet bool
	size              uint32
	sizeSet           bool
	addressType       AddrTypeEnum
}

func parseTypeDefBlob(tok *tokenGenerator) (typeDefBlob, error) {
	tdb := typeDefBlob{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case addressTypeToken:
			tdb.addressType, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("typeDefBlob addressType could not be parsed")
				break forLoop
			}
			log.Info().Msg("typeDefBlob addressType successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("typeDefBlob could not be parsed")
				break forLoop
			} else if tok.current() == endTypeDefBlobToken {
				break forLoop
			} else if !tdb.nameSet {
				tdb.name = tok.current()
				tdb.nameSet = true
				log.Info().Msg("typeDefBlob name successfully parsed")
			} else if !tdb.longIdentifierSet {
				tdb.longIdentifier = tok.current()
				tdb.longIdentifierSet = true
				log.Info().Msg("typeDefBlob longIdentifier successfully parsed")
			} else if !tdb.sizeSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 32)
				if err != nil {
					log.Err(err).Msg("typeDefBlob size could not be parsed")
					break forLoop
				}
				tdb.size = uint32(buf)
				tdb.sizeSet = true
				log.Info().Msg("typeDefBlob size successfully parsed")
			}
		}
	}
	return tdb, err
}
