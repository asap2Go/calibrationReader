package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type matrixDim struct {
	xDim    uint16
	xDimSet bool
	yDim    uint16
	yDimSet bool
	zDim    uint16
	zDimSet bool
}

func parseMatrixDim(tok *tokenGenerator) (matrixDim, error) {
	md := matrixDim{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("matrixDim could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			log.Info().Str("current token", tok.current()).Msg("matrixDim detected keyword:")
			tok.previous()
			log.Info().Str("current token", tok.current()).Msg("matrixDim rolled back to:")
			break forLoop
		} else if !md.xDimSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("matrixDim xDim could not be parsed")
				break forLoop
			}
			md.xDim = uint16(buf)
			md.xDimSet = true
			log.Info().Msg("matrixDim xDim successfully parsed")
		} else if !md.yDimSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("matrixDim yDim could not be parsed")
				break forLoop
			}
			md.yDim = uint16(buf)
			log.Info().Msg("matrixDim yDim successfully parsed")
		} else if !md.zDimSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("matrixDim zDim could not be parsed")
				break forLoop
			}
			md.zDim = uint16(buf)
			log.Info().Msg("matrixDim zDim successfully parsed")
		}
	}
	return md, err
}
